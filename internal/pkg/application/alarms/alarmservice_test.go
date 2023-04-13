package alarms

import (
	"bytes"
	"context"
	"testing"

	"github.com/diwise/iot-device-mgmt/internal/pkg/infrastructure/repositories/database/alarms"
	"github.com/diwise/messaging-golang/pkg/messaging"
	"github.com/matryer/is"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

func TestBatteryLevelChangedHandler(t *testing.T) {
	is, ctx, log := testSetup(t)

	d := amqp091.Delivery{
		RoutingKey: "watchdog.batteryLevelChanged",
		Body:       []byte(batteryLevelChangedJson),
	}

	alarmList := []alarms.Alarm{}

	m := &messaging.MsgContextMock{}
	a := &AlarmServiceMock{
		GetConfigurationFunc: func() Configuration {
			return *parseConfigFile(bytes.NewBufferString(configFileJson))
		},

		AddAlarmFunc: func(ctx context.Context, alarm alarms.Alarm) error {
			alarmList = append(alarmList, alarm)
			return nil
		},
	}

	BatteryLevelChangedHandler(m, a)(ctx, d, log)

	is.Equal(1, len(alarmList))
}

func TestFunctionUpdatedHandler_CounterOverflow(t *testing.T) {
	is, ctx, log := testSetup(t)

	d := amqp091.Delivery{
		RoutingKey: "feature.updated",
		Body:       []byte(counterOverflow1Json),
	}

	alarmList := []alarms.Alarm{}

	m := &messaging.MsgContextMock{}
	a := &AlarmServiceMock{
		GetConfigurationFunc: func() Configuration {
			return *parseConfigFile(bytes.NewBufferString(configFileJson))
		},
		AddAlarmFunc: func(ctx context.Context, alarm alarms.Alarm) error {
			alarmList = append(alarmList, alarm)
			return nil
		},
	}

	FunctionUpdatedHandler(m, a)(ctx, d, log)

	is.Equal(1, len(alarmList))
}

func TestFunctionUpdatedHandler_CounterOverflow_Between(t *testing.T) {
	is, ctx, log := testSetup(t)

	d := amqp091.Delivery{
		RoutingKey: "feature.updated",
		Body:       []byte(counterOverflow2Json),
	}

	alarmList := []alarms.Alarm{}

	m := &messaging.MsgContextMock{}
	a := &AlarmServiceMock{
		GetConfigurationFunc: func() Configuration {
			return *parseConfigFile(bytes.NewBufferString(configFileJson))
		},
		AddAlarmFunc: func(ctx context.Context, alarm alarms.Alarm) error {
			alarmList = append(alarmList, alarm)
			return nil
		},
	}

	FunctionUpdatedHandler(m, a)(ctx, d, log)

	is.Equal(1, len(alarmList))
}

func TestFunctionUpdatedHandler_LevelSand(t *testing.T) {
	is, ctx, log := testSetup(t)

	d := amqp091.Delivery{
		RoutingKey: "feature.updated",
		Body:       []byte(levelSandJson),
	}

	alarmList := []alarms.Alarm{}

	m := &messaging.MsgContextMock{}
	a := &AlarmServiceMock{
		GetConfigurationFunc: func() Configuration {
			return *parseConfigFile(bytes.NewBufferString(configFileJson))
		},
		AddAlarmFunc: func(ctx context.Context, alarm alarms.Alarm) error {
			alarmList = append(alarmList, alarm)
			return nil
		},
	}

	FunctionUpdatedHandler(m, a)(ctx, d, log)

	is.Equal(1, len(alarmList))
}

func TestDeviceStatusHandler(t *testing.T) {
	is, ctx, log := testSetup(t)

	d := amqp091.Delivery{
		RoutingKey: "device-status",
		Body:       []byte(uplinkFcntRetransmissionJson),
	}

	alarmList := []alarms.Alarm{}

	m := &messaging.MsgContextMock{}
	a := &AlarmServiceMock{
		GetConfigurationFunc: func() Configuration {
			return *parseConfigFile(bytes.NewBufferString(configFileJson))
		},
		AddAlarmFunc: func(ctx context.Context, alarm alarms.Alarm) error {
			alarmList = append(alarmList, alarm)
			return nil
		},
	}

	DeviceStatusHandler(m, a)(ctx, d, log)

	is.Equal(1, len(alarmList))
}

func TestParseConfigFile(t *testing.T) {
	is := is.New(t)
	config := parseConfigFile(bytes.NewBufferString(configFileJson))
	is.Equal(7, len(config.AlarmConfigurations))

	is.Equal("net:test:iot:a81757", config.AlarmConfigurations[0].ID)
	is.Equal("", config.AlarmConfigurations[5].ID)
	is.Equal("", config.AlarmConfigurations[6].ID)
}

func testSetup(t *testing.T) (*is.I, context.Context, zerolog.Logger) {
	is := is.New(t)
	return is, context.Background(), zerolog.Logger{}
}

const batteryLevelChangedJson = `{"deviceID":"net:test:iot:a81757","batteryLevel":10,"tenant":"default","observedAt":"2023-04-12T06:51:25.389495559Z"}`
const counterOverflow1Json string = `{"id":"a817bf9e","type":"counter","subtype":"overflow","counter":{"count":11,"state":true},"tenant":"default"}`
const counterOverflow2Json string = `{"id":"fbf9f","type":"counter","subtype":"overflow","counter":{"count":11,"state":true},"tenant":"default"}`
const levelSandJson string = `{"id":"323c6","type":"level","subtype":"sand","level":{"current":1.4,"percent":19},"tenant":"default"}`
const uplinkFcntRetransmissionJson string = `{"deviceID":"01","batteryLevel":10,"tenant":"default","statusCode":1,"statusMessages":["UPLINK_FCNT_RETRANSMISSION"],"timestamp":"2023-04-12T06:51:25.389495559Z"}`

const configFileJson string = `
deviceID;functionID;alarmName;alarmType;min;max;severity;description
;;deviceNotObserved;-;;;1;
net:test:iot:a81757;;batteryLevel;MIN;20;0;1;
;a817bf9e;counter;MAX;;10;1;
;fbf9f;counter;BETWEEN;1;10;1;Count ska vara mellan {MIN} och {MAX} men är {VALUE}
;323c6;level;MAX;;4;1;
;70t589;waterquality;BETWEEN;4;35;1;Temp ska vara mellan {MIN} och {MAX} men är {VALUE}
;;UPLINK_FCNT_RETRANSMISSION;-;;;1;
`
