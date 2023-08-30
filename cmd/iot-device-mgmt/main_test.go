package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/diwise/iot-device-mgmt/internal/pkg/application/alarms"
	"github.com/diwise/iot-device-mgmt/internal/pkg/application/devicemanagement"
	db "github.com/diwise/iot-device-mgmt/internal/pkg/infrastructure/repositories/database"
	dmDb "github.com/diwise/iot-device-mgmt/internal/pkg/infrastructure/repositories/database/devicemanagement"
	"github.com/diwise/iot-device-mgmt/internal/pkg/infrastructure/router"
	"github.com/diwise/iot-device-mgmt/internal/pkg/presentation/api"
	"github.com/diwise/iot-device-mgmt/pkg/types"
	"github.com/diwise/messaging-golang/pkg/messaging"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/matryer/is"
	"github.com/rs/zerolog"
)

const noToken string = ""

func TestThatHealthEndpointReturns204NoContent(t *testing.T) {
	r, is := setupTest(t)
	server := httptest.NewServer(r)
	defer server.Close()

	resp, _ := testRequest(is, server, http.MethodGet, "/health", noToken, nil)

	is.Equal(resp.StatusCode, http.StatusNoContent)
}

func TestThatGetUnknownDeviceReturns404(t *testing.T) {
	r, is := setupTest(t)
	server := httptest.NewServer(r)
	defer server.Close()

	token := createJWTWithTenants([]string{"default"})
	resp, _ := testRequest(is, server, http.MethodGet, "/api/v0/devices/nosuchdevice", token, nil)

	is.Equal(resp.StatusCode, http.StatusNotFound)
}

func TestThatGetKnownDeviceByEUIReturns200(t *testing.T) {
	r, is := setupTest(t)
	server := httptest.NewServer(r)
	defer server.Close()

	token := createJWTWithTenants([]string{"default"})
	resp, body := testRequest(is, server, http.MethodGet, "/api/v0/devices?devEUI=a81758fffe06bfa3", token, nil)

	is.Equal(resp.StatusCode, http.StatusOK)
	is.Equal(body, `{"meta":{"totalRecords":1,"count":1},"data":[{"active":true,"sensorID":"a81758fffe06bfa3","deviceID":"intern-a81758fffe06bfa3","tenant":{"name":"default"},"name":"name-a81758fffe06bfa3","description":"desc-a81758fffe06bfa3","location":{"latitude":62.3916,"longitude":17.30723,"altitude":0},"environment":"water","source":"source","types":[{"urn":"urn:oma:lwm2m:ext:3303"},{"urn":"urn:oma:lwm2m:ext:3302"},{"urn":"urn:oma:lwm2m:ext:3301"}],"tags":[],"deviceProfile":{"name":"elsys_codec","decoder":"elsys_codec","interval":60},"deviceStatus":{"batteryLevel":-1,"lastObservedAt":"0001-01-01T00:00:00Z"},"deviceState":{"online":false,"state":-1,"observedAt":"0001-01-01T00:00:00Z"}}],"links":{"self":"https://diwise.io/api/v0/devices?devEUI=a81758fffe06bfa3"}}`)
}

func TestThatGetKnownDeviceReturns200(t *testing.T) {
	r, is := setupTest(t)
	server := httptest.NewServer(r)
	defer server.Close()

	token := createJWTWithTenants([]string{"default"})
	resp, body := testRequest(is, server, http.MethodGet, "/api/v0/devices/intern-a81758fffe06bfa3", token, nil)

	d := struct {
		DevEui string `json:"sensorID"`
	}{}
	json.Unmarshal([]byte(body), &d)

	is.Equal(resp.StatusCode, http.StatusOK)
	is.Equal("a81758fffe06bfa3", d.DevEui)
	//is.Equal(body, `{"devEUI":"a81758fffe06bfa3","deviceID":"intern-a81758fffe06bfa3","name":"name-a81758fffe06bfa3","description":"desc-a81758fffe06bfa3","location":{"latitude":62.3916,"longitude":17.30723,"altitude":0},"environment":"water","types":["urn:oma:lwm2m:ext:3303","urn:oma:lwm2m:ext:3302","urn:oma:lwm2m:ext:3301"],"sensorType":{"id":1,"name":"elsys","description":"","interval":3600},"lastObserved":"0001-01-01T00:00:00Z","active":true,"tenant":"default","status":{"batteryLevel":0,"statusCode":0,"timestamp":""},"interval":60}`)
}

func TestThatGetKnownDeviceMarshalToType(t *testing.T) {
	r, is := setupTest(t)
	server := httptest.NewServer(r)
	defer server.Close()

	token := createJWTWithTenants([]string{"default"})
	resp, body := testRequest(is, server, http.MethodGet, "/api/v0/devices/intern-a81758fffe06bfa3", token, nil)

	d := types.Device{}
	json.Unmarshal([]byte(body), &d)

	is.Equal(resp.StatusCode, http.StatusOK)
	is.Equal("a81758fffe06bfa3", d.SensorID)
	is.Equal("default", d.Tenant.Name)
}

func TestThatGetKnownDeviceByEUIFromNonAllowedTenantReturns404(t *testing.T) {
	r, is := setupTest(t)
	server := httptest.NewServer(r)
	defer server.Close()

	token := createJWTWithTenants([]string{"wrongtenant"})
	resp, _ := testRequest(is, server, http.MethodGet, "/api/v0/devices?devEUI=a81758fffe06bfa3", token, nil)

	is.Equal(resp.StatusCode, http.StatusNotFound)
}

func TestThatGetKnownDeviceFromNonAllowedTenantReturns404(t *testing.T) {
	r, is := setupTest(t)
	server := httptest.NewServer(r)
	defer server.Close()

	token := createJWTWithTenants([]string{"wrongtenant"})
	resp, _ := testRequest(is, server, http.MethodGet, "/api/v0/devices/intern-a81758fffe06bfa3", token, nil)

	is.Equal(resp.StatusCode, http.StatusNotFound)
}

func setupTest(t *testing.T) (*chi.Mux, *is.I) {
	is := is.New(t)

	db, err := dmDb.NewDeviceRepository(db.NewSQLiteConnector(zerolog.Logger{}))
	is.NoErr(err)

	err = db.Seed(context.Background(), bytes.NewBuffer([]byte(csvMock)))
	is.NoErr(err)

	app := devicemanagement.New(db, &messaging.MsgContextMock{})
	router := router.New("testService")

	policies := bytes.NewBufferString(opaModule)
	api.RegisterHandlers(context.Background(), router, policies, app, &alarms.AlarmServiceMock{})

	return router, is
}

func testRequest(is *is.I, ts *httptest.Server, method, path string, token string, body io.Reader) (*http.Response, string) {
	req, _ := http.NewRequest(method, ts.URL+path, body)

	if len(token) > 0 {
		req.Header.Add("Authorization", "Bearer "+token)
	}

	req.Header.Add("X-Forwarded-Host", "diwise.io")
	req.Header.Add("X-Forwarded-Proto", "https")

	resp, _ := http.DefaultClient.Do(req)
	respBody, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	return resp, string(respBody)
}

func createJWTWithTenants(tenants []string) string {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	_, tokenString, _ := tokenAuth.Encode(map[string]any{"user_id": 123, "azp": "diwise-frontend", "tenants": tenants})
	return tokenString
}

const csvMock string = `devEUI;internalID;lat;lon;where;types;sensorType;name;description;active;tenant;interval;source
a81758fffe06bfa3;intern-a81758fffe06bfa3;62.39160;17.30723;water;urn:oma:lwm2m:ext:3303,urn:oma:lwm2m:ext:3302,urn:oma:lwm2m:ext:3301;Elsys_Codec;name-a81758fffe06bfa3;desc-a81758fffe06bfa3;true;default;60;source
a81758fffe051d00;intern-a81758fffe051d00;0.0;0.0;air;urn:oma:lwm2m:ext:3303;Elsys_Codec;name-a81758fffe051d00;desc-a81758fffe051d00;true;default;60;
a81758fffe04d83f;intern-a81758fffe04d83f;0.0;0.0;air;urn:oma:lwm2m:ext:3303;Elsys_Codec;name-a81758fffe04d83f;desc-a81758fffe04d83f;true;default;60;`

const opaModule string = `
#
# Use https://play.openpolicyagent.org for easier editing/validation of this policy file
#

package example.authz

default allow := false

allow = response {
    is_valid_token

    input.method == "GET"
    pathstart := array.slice(input.path, 0, 3)
    pathstart == ["api", "v0", "devices"]

    token.payload.azp == "diwise-frontend"

    response := {
        "tenants": token.payload.tenants
    }
}

is_valid_token {
    1 == 1
}

token := {"payload": payload} {
    [_, payload, _] := io.jwt.decode(input.token)
}
`
