package alarms

import (
	"context"
	"fmt"
	"time"

	. "github.com/diwise/iot-device-mgmt/internal/pkg/application/alarms/events"
	db "github.com/diwise/iot-device-mgmt/internal/pkg/infrastructure/repositories/database/alarms"
	"github.com/diwise/messaging-golang/pkg/messaging"
	"github.com/diwise/service-chassis/pkg/infrastructure/o11y/logging"
)

//go:generate moq -rm -out alarmservice_mock.go . AlarmService
type AlarmService interface {
	Start()
	Stop()

	GetAlarms(ctx context.Context, tenants ...string) ([]db.Alarm, error)
	GetAlarmsByID(ctx context.Context, id ...int) ([]db.Alarm, error)
	AddAlarm(ctx context.Context, alarm db.Alarm) error
	CloseAlarm(ctx context.Context, alarmID int) error

	GetConfiguration() Configuration
}

type alarmService struct {
	alarmRepository db.AlarmRepository
	messenger       messaging.MsgContext
	config          *Configuration
}

func New(d db.AlarmRepository, m messaging.MsgContext, cfg *Configuration) AlarmService {
	as := &alarmService{
		alarmRepository: d,
		messenger:       m,
		config:          cfg,
	}

	as.messenger.RegisterTopicMessageHandler("device-status", DeviceStatusHandler(m, as))
	as.messenger.RegisterTopicMessageHandler("watchdog.batteryLevelChanged", BatteryLevelChangedHandler(m, as))
	as.messenger.RegisterTopicMessageHandler("watchdog.deviceNotObserved", DeviceNotObservedHandler(m, as))
	as.messenger.RegisterTopicMessageHandler("function.updated", FunctionUpdatedHandler(m, as))

	return as
}

func (a *alarmService) Start() {}
func (a *alarmService) Stop()  {}

func (a *alarmService) GetAlarms(ctx context.Context, tenants ...string) ([]db.Alarm, error) {
	alarms, err := a.alarmRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return alarms, nil
}

func (a *alarmService) GetAlarmsByID(ctx context.Context, id ...int) ([]db.Alarm, error) {
	alarms, err := a.alarmRepository.GetByID(ctx, id...)
	if err != nil {
		return nil, err
	}

	return alarms, nil
}

func (a *alarmService) AddAlarm(ctx context.Context, alarm db.Alarm) error {
	id, err := a.alarmRepository.Add(ctx, alarm)
	if err != nil {
		return err
	}

	alarms, err := a.alarmRepository.GetByID(ctx, id)
	if err != nil || len(alarms) != 1 {
		return fmt.Errorf("failed to add alarm, could not fetch newly created alarm")
	}

	return a.messenger.PublishOnTopic(ctx, &AlarmCreated{
		Alarm:     alarms[0],
		Tenant:    alarms[0].Tenant,
		Timestamp: alarms[0].ObservedAt,
	})
}

func (a *alarmService) CloseAlarm(ctx context.Context, alarmID int) error {
	logger := logging.GetFromContext(ctx)

	alarm, err := a.alarmRepository.GetByID(ctx, alarmID)
	if len(alarm) == 0 || err != nil {
		logger.Debug().Msgf("alarm %d could not be fetched by ID", alarmID)
		return err
	}

	err = a.alarmRepository.Close(ctx, alarmID)
	if err != nil {
		return err
	}

	return a.messenger.PublishOnTopic(ctx, &AlarmClosed{ID: alarmID, Tenant: alarm[0].Tenant, Timestamp: time.Now().UTC()})
}

func (a *alarmService) GetConfiguration() Configuration {
	return *a.config
}
