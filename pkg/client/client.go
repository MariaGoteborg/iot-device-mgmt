package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/diwise/iot-device-mgmt/pkg/types"
	"github.com/diwise/service-chassis/pkg/infrastructure/o11y/logging"
	"github.com/diwise/service-chassis/pkg/infrastructure/o11y/tracing"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"golang.org/x/oauth2/clientcredentials"
)

type DeviceManagementClient interface {
	FindDeviceFromDevEUI(ctx context.Context, devEUI string) (Device, error)
	FindDeviceFromInternalID(ctx context.Context, deviceID string) (Device, error)
	Close(ctx context.Context)
}

type deviceState int

const (
	Refreshing deviceState = iota
	Ready
	Error
)

type devEUIState struct {
	state      deviceState
	err        error
	internalID string
}

type lookupResult struct {
	state  deviceState
	device Device
	err    error
	when   time.Time
}

type devManagementClient struct {
	url               string
	clientCredentials *clientcredentials.Config

	cacheByInternalID map[string]lookupResult
	knownDevEUI       map[string]devEUIState
	queue             chan (func())

	keepRunning *atomic.Bool
	wg          sync.WaitGroup
}

var tracer = otel.Tracer("device-mgmt-client")

func New(ctx context.Context, devMgmtUrl, oauthTokenURL, oauthClientID, oauthClientSecret string) (DeviceManagementClient, error) {
	oauthConfig := &clientcredentials.Config{
		ClientID:     oauthClientID,
		ClientSecret: oauthClientSecret,
		TokenURL:     oauthTokenURL,
	}

	token, err := oauthConfig.Token(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get client credentials from %s: %w", oauthConfig.TokenURL, err)
	}

	if !token.Valid() {
		return nil, fmt.Errorf("an invalid token was returned from %s", oauthTokenURL)
	}

	dmc := &devManagementClient{
		url:               devMgmtUrl,
		clientCredentials: oauthConfig,

		cacheByInternalID: make(map[string]lookupResult, 100),
		knownDevEUI:       make(map[string]devEUIState, 100),
		queue:             make(chan func()),
		keepRunning:       &atomic.Bool{},
	}

	go dmc.run(ctx)

	return dmc, nil
}

func (dmc *devManagementClient) run(ctx context.Context) {
	dmc.wg.Add(1)
	defer dmc.wg.Done()

	logger := logging.GetFromContext(ctx)
	logger.Info().Msg("starting up device management client")

	// use atomic swap to avoid startup races
	alreadyStarted := dmc.keepRunning.Swap(true)
	if alreadyStarted {
		logger.Error().Msg("attempt to start the device management client multiple times")
		return
	}

	for dmc.keepRunning.Load() {
		fn := <-dmc.queue
		fn()
	}

	logger.Info().Msg("device management client exiting")
}

func (dmc *devManagementClient) Close(ctx context.Context) {
	dmc.queue <- func() {
		dmc.keepRunning.Store(false)
	}

	dmc.wg.Wait()
}

var ErrDeviceNotFound error = errors.New("not found")

var errInternal error = errors.New("internal error")
var errRetry error = errors.New("retry")

func (dmc *devManagementClient) FindDeviceFromDevEUI(ctx context.Context, devEUI string) (Device, error) {

	resultchan := make(chan Device)
	errchan := make(chan error)

	dmc.queue <- func() {
		device, ok := dmc.knownDevEUI[devEUI]

		if ok {
			switch device.state {
			case Ready:
				go func() {
					deviceByID, err := dmc.FindDeviceFromInternalID(ctx, device.internalID)
					if err != nil {
						errchan <- err
					} else {
						resultchan <- deviceByID
					}
				}()
			case Error:
				errchan <- device.err
			case Refreshing:
				errchan <- errRetry
			default:
				errchan <- errInternal
			}

			return
		}

		dmc.knownDevEUI[devEUI] = devEUIState{state: Refreshing}
		go func() {
			dmc.updateDeviceCacheFromDevEUI(ctx, devEUI)
		}()
		errchan <- errRetry
	}

	select {
	case d := <-resultchan:
		return d, nil
	case e := <-errchan:
		if errors.Is(e, errRetry) {
			time.Sleep(10 * time.Millisecond)
			return dmc.FindDeviceFromDevEUI(ctx, devEUI)
		}
		return nil, e
	}
}

func (dmc *devManagementClient) updateDeviceCacheFromDevEUI(ctx context.Context, devEUI string) error {
	device, err := dmc.findDeviceFromDevEUI(ctx, devEUI)

	dmc.queue <- func() {
		if err != nil {
			log := logging.GetFromContext(ctx)
			log.Error().Err(err).Msg("failed to update device cache")

			dmc.knownDevEUI[devEUI] = devEUIState{state: Error, err: err}
		} else {
			dmc.knownDevEUI[devEUI] = devEUIState{state: Ready, internalID: device.ID()}
			dmc.cacheByInternalID[device.ID()] = lookupResult{state: Ready, device: device, when: time.Now()}
		}
	}

	return err
}

func (dmc *devManagementClient) findDeviceFromDevEUI(ctx context.Context, devEUI string) (Device, error) {
	var err error
	ctx, span := tracer.Start(ctx, "find-device-from-deveui")
	defer func() { tracing.RecordAnyErrorAndEndSpan(err, span) }()

	log := logging.GetFromContext(ctx)
	log.Info().Msgf("looking up internal id and types for devEUI %s", devEUI)

	httpClient := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	url := dmc.url + "/api/v0/devices?devEUI=" + devEUI

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		err = fmt.Errorf("failed to create http request: %w", err)
		return nil, err
	}

	if dmc.clientCredentials != nil {
		token, err := dmc.clientCredentials.Token(ctx)
		if err != nil {
			err = fmt.Errorf("failed to get client credentials from %s: %w", dmc.clientCredentials.TokenURL, err)
			return nil, err
		}

		req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		err = fmt.Errorf("failed to retrieve device information from devEUI: %w", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		err = fmt.Errorf("request failed, not authorized")
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrDeviceNotFound
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("request failed with status code %d", resp.StatusCode)
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("failed to read response body: %w", err)
		return nil, err
	}

	impls := struct {
		Devices []types.Device `json:"data"`
	}{}

	err = json.Unmarshal(respBody, &impls)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal response body: %w", err)
		return nil, err
	}

	if len(impls.Devices) == 0 {
		err = fmt.Errorf("device management returned an empty list of devices")
		return nil, err
	}

	device := impls.Devices[0]
	return &deviceWrapper{&device}, nil
}

func (dmc *devManagementClient) FindDeviceFromInternalID(ctx context.Context, deviceID string) (Device, error) {

	resultchan := make(chan Device)
	errchan := make(chan error)

	dmc.queue <- func() {
		r, ok := dmc.cacheByInternalID[deviceID]

		if ok {
			switch r.state {
			case Ready:
				resultchan <- r.device
			case Error:
				errchan <- r.err
			case Refreshing:
				errchan <- errRetry
			default:
				errchan <- errInternal
			}

			return
		}

		// if not in cache, send request to device management
		r = lookupResult{state: Refreshing, when: time.Now()}
		dmc.cacheByInternalID[deviceID] = r

		go func() {
			dmc.updateDeviceCacheFromInternalID(ctx, deviceID)
		}()

		errchan <- errRetry
	}

	select {
	case d := <-resultchan:
		return d, nil
	case e := <-errchan:
		if errors.Is(e, errRetry) {
			time.Sleep(10 * time.Millisecond)
			return dmc.FindDeviceFromInternalID(ctx, deviceID)
		}
		return nil, e
	}
}

func (dmc *devManagementClient) updateDeviceCacheFromInternalID(ctx context.Context, deviceID string) error {
	device, err := dmc.findDeviceFromInternalID(ctx, deviceID)

	dmc.queue <- func() {
		if err != nil {
			log := logging.GetFromContext(ctx)
			log.Error().Err(err).Msg("failed to update device cache")

			dmc.cacheByInternalID[deviceID] = lookupResult{state: Error, err: err, when: time.Now()}
		} else {
			dmc.cacheByInternalID[deviceID] = lookupResult{state: Ready, device: device, when: time.Now()}
		}
	}

	return err
}

func (dmc *devManagementClient) findDeviceFromInternalID(ctx context.Context, deviceID string) (Device, error) {
	var err error
	ctx, span := tracer.Start(ctx, "find-device-from-id")
	defer func() { tracing.RecordAnyErrorAndEndSpan(err, span) }()

	log := logging.GetFromContext(ctx)
	log.Info().Msgf("looking up properties for device %s", deviceID)

	httpClient := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	url := dmc.url + "/api/v0/devices/" + deviceID

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		err = fmt.Errorf("failed to create http request: %w", err)
		return nil, err
	}

	if dmc.clientCredentials != nil {
		token, err := dmc.clientCredentials.Token(ctx)
		if err != nil {
			err = fmt.Errorf("failed to get client credentials from %s: %w", dmc.clientCredentials.TokenURL, err)
			return nil, err
		}

		req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		err = fmt.Errorf("failed to retrieve information for device: %w", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		err = fmt.Errorf("request failed, not authorized")
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrDeviceNotFound
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("request failed with status code %d", resp.StatusCode)
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("failed to read response body: %w", err)
		return nil, err
	}

	impl := &types.Device{}

	err = json.Unmarshal(respBody, impl)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal response body: %w", err)
		return nil, err
	}

	return &deviceWrapper{impl}, nil
}

//go:generate moq -rm -out ../test/device_mock.go . Device

// you need to modify the generated ../test/device_mock.go
// import "github.com/diwise/iot-device-mgmt/pkg/client"
// change "var _ Device = &DeviceMock{}" to "var _ client.Device = &DeviceMock{}"
type Device interface {
	ID() string
	Environment() string
	IsActive() bool
	Latitude() float64
	Longitude() float64
	SensorType() string
	Source() string
	Tenant() string
	Types() []string
}

type deviceWrapper struct {
	impl *types.Device
}

func (d *deviceWrapper) ID() string {
	return d.impl.DeviceID
}

func (d *deviceWrapper) Latitude() float64 {
	return d.impl.Location.Latitude
}

func (d *deviceWrapper) Longitude() float64 {
	return d.impl.Location.Longitude
}

func (d *deviceWrapper) Environment() string {
	return d.impl.Environment
}

func (d *deviceWrapper) SensorType() string {
	return d.impl.DeviceProfile.Decoder
}

func (d *deviceWrapper) Types() []string {
	types := []string{}
	for _, t := range d.impl.Lwm2mTypes {
		types = append(types, t.Urn)
	}
	return types
}

func (d *deviceWrapper) IsActive() bool {
	return d.impl.Active
}

func (d *deviceWrapper) Tenant() string {
	return d.impl.Tenant.Name
}

func (d *deviceWrapper) Source() string {
	return d.impl.Source
}
