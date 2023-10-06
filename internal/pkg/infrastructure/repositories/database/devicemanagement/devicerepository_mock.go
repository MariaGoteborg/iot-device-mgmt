// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package devicemanagement

import (
	"context"
	"io"
	"sync"
	"time"
)

// Ensure, that DeviceRepositoryMock does implement DeviceRepository.
// If this is not the case, regenerate this file with moq.
var _ DeviceRepository = &DeviceRepositoryMock{}

// DeviceRepositoryMock is a mock implementation of DeviceRepository.
//
//	func TestSomethingThatUsesDeviceRepository(t *testing.T) {
//
//		// make and configure a mocked DeviceRepository
//		mockedDeviceRepository := &DeviceRepositoryMock{
//			AddAlarmFunc: func(ctx context.Context, deviceID string, alarmID int, severity int, observedAt time.Time) error {
//				panic("mock out the AddAlarm method")
//			},
//			GetDeviceByDeviceIDFunc: func(ctx context.Context, deviceID string, tenants ...string) (Device, error) {
//				panic("mock out the GetDeviceByDeviceID method")
//			},
//			GetDeviceBySensorIDFunc: func(ctx context.Context, sensorID string, tenants ...string) (Device, error) {
//				panic("mock out the GetDeviceBySensorID method")
//			},
//			GetDevicesFunc: func(ctx context.Context, tenants ...string) ([]Device, error) {
//				panic("mock out the GetDevices method")
//			},
//			GetOnlineDevicesFunc: func(ctx context.Context, tenants ...string) ([]Device, error) {
//				panic("mock out the GetOnlineDevices method")
//			},
//			RemoveAlarmByIDFunc: func(ctx context.Context, alarmID int) (string, error) {
//				panic("mock out the RemoveAlarmByID method")
//			},
//			SaveFunc: func(ctx context.Context, device *Device) error {
//				panic("mock out the Save method")
//			},
//			SeedFunc: func(contextMoqParam context.Context, reader io.Reader, strings ...string) error {
//				panic("mock out the Seed method")
//			},
//			UpdateDeviceStateFunc: func(ctx context.Context, deviceID string, deviceState DeviceState) error {
//				panic("mock out the UpdateDeviceState method")
//			},
//			UpdateDeviceStatusFunc: func(ctx context.Context, deviceID string, deviceStatus DeviceStatus) error {
//				panic("mock out the UpdateDeviceStatus method")
//			},
//		}
//
//		// use mockedDeviceRepository in code that requires DeviceRepository
//		// and then make assertions.
//
//	}
type DeviceRepositoryMock struct {
	// AddAlarmFunc mocks the AddAlarm method.
	AddAlarmFunc func(ctx context.Context, deviceID string, alarmID int, severity int, observedAt time.Time) error

	// GetDeviceByDeviceIDFunc mocks the GetDeviceByDeviceID method.
	GetDeviceByDeviceIDFunc func(ctx context.Context, deviceID string, tenants ...string) (Device, error)

	// GetDeviceBySensorIDFunc mocks the GetDeviceBySensorID method.
	GetDeviceBySensorIDFunc func(ctx context.Context, sensorID string, tenants ...string) (Device, error)

	// GetDevicesFunc mocks the GetDevices method.
	GetDevicesFunc func(ctx context.Context, tenants ...string) ([]Device, error)

	// GetOnlineDevicesFunc mocks the GetOnlineDevices method.
	GetOnlineDevicesFunc func(ctx context.Context, tenants ...string) ([]Device, error)

	// RemoveAlarmByIDFunc mocks the RemoveAlarmByID method.
	RemoveAlarmByIDFunc func(ctx context.Context, alarmID int) (string, error)

	// SaveFunc mocks the Save method.
	SaveFunc func(ctx context.Context, device *Device) error

	// SeedFunc mocks the Seed method.
	SeedFunc func(contextMoqParam context.Context, reader io.Reader, strings ...string) error

	// UpdateDeviceStateFunc mocks the UpdateDeviceState method.
	UpdateDeviceStateFunc func(ctx context.Context, deviceID string, deviceState DeviceState) error

	// UpdateDeviceStatusFunc mocks the UpdateDeviceStatus method.
	UpdateDeviceStatusFunc func(ctx context.Context, deviceID string, deviceStatus DeviceStatus) error

	// calls tracks calls to the methods.
	calls struct {
		// AddAlarm holds details about calls to the AddAlarm method.
		AddAlarm []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// DeviceID is the deviceID argument value.
			DeviceID string
			// AlarmID is the alarmID argument value.
			AlarmID int
			// Severity is the severity argument value.
			Severity int
			// ObservedAt is the observedAt argument value.
			ObservedAt time.Time
		}
		// GetDeviceByDeviceID holds details about calls to the GetDeviceByDeviceID method.
		GetDeviceByDeviceID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// DeviceID is the deviceID argument value.
			DeviceID string
			// Tenants is the tenants argument value.
			Tenants []string
		}
		// GetDeviceBySensorID holds details about calls to the GetDeviceBySensorID method.
		GetDeviceBySensorID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// SensorID is the sensorID argument value.
			SensorID string
			// Tenants is the tenants argument value.
			Tenants []string
		}
		// GetDevices holds details about calls to the GetDevices method.
		GetDevices []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenants is the tenants argument value.
			Tenants []string
		}
		// GetOnlineDevices holds details about calls to the GetOnlineDevices method.
		GetOnlineDevices []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenants is the tenants argument value.
			Tenants []string
		}
		// RemoveAlarmByID holds details about calls to the RemoveAlarmByID method.
		RemoveAlarmByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// AlarmID is the alarmID argument value.
			AlarmID int
		}
		// Save holds details about calls to the Save method.
		Save []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Device is the device argument value.
			Device *Device
		}
		// Seed holds details about calls to the Seed method.
		Seed []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Reader is the reader argument value.
			Reader io.Reader
			// Strings is the strings argument value.
			Strings []string
		}
		// UpdateDeviceState holds details about calls to the UpdateDeviceState method.
		UpdateDeviceState []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// DeviceID is the deviceID argument value.
			DeviceID string
			// DeviceState is the deviceState argument value.
			DeviceState DeviceState
		}
		// UpdateDeviceStatus holds details about calls to the UpdateDeviceStatus method.
		UpdateDeviceStatus []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// DeviceID is the deviceID argument value.
			DeviceID string
			// DeviceStatus is the deviceStatus argument value.
			DeviceStatus DeviceStatus
		}
	}
	lockAddAlarm            sync.RWMutex
	lockGetDeviceByDeviceID sync.RWMutex
	lockGetDeviceBySensorID sync.RWMutex
	lockGetDevices          sync.RWMutex
	lockGetOnlineDevices    sync.RWMutex
	lockRemoveAlarmByID     sync.RWMutex
	lockSave                sync.RWMutex
	lockSeed                sync.RWMutex
	lockUpdateDeviceState   sync.RWMutex
	lockUpdateDeviceStatus  sync.RWMutex
}

// AddAlarm calls AddAlarmFunc.
func (mock *DeviceRepositoryMock) AddAlarm(ctx context.Context, deviceID string, alarmID int, severity int, observedAt time.Time) error {
	if mock.AddAlarmFunc == nil {
		panic("DeviceRepositoryMock.AddAlarmFunc: method is nil but DeviceRepository.AddAlarm was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		DeviceID   string
		AlarmID    int
		Severity   int
		ObservedAt time.Time
	}{
		Ctx:        ctx,
		DeviceID:   deviceID,
		AlarmID:    alarmID,
		Severity:   severity,
		ObservedAt: observedAt,
	}
	mock.lockAddAlarm.Lock()
	mock.calls.AddAlarm = append(mock.calls.AddAlarm, callInfo)
	mock.lockAddAlarm.Unlock()
	return mock.AddAlarmFunc(ctx, deviceID, alarmID, severity, observedAt)
}

// AddAlarmCalls gets all the calls that were made to AddAlarm.
// Check the length with:
//
//	len(mockedDeviceRepository.AddAlarmCalls())
func (mock *DeviceRepositoryMock) AddAlarmCalls() []struct {
	Ctx        context.Context
	DeviceID   string
	AlarmID    int
	Severity   int
	ObservedAt time.Time
} {
	var calls []struct {
		Ctx        context.Context
		DeviceID   string
		AlarmID    int
		Severity   int
		ObservedAt time.Time
	}
	mock.lockAddAlarm.RLock()
	calls = mock.calls.AddAlarm
	mock.lockAddAlarm.RUnlock()
	return calls
}

// GetDeviceByDeviceID calls GetDeviceByDeviceIDFunc.
func (mock *DeviceRepositoryMock) GetDeviceByDeviceID(ctx context.Context, deviceID string, tenants ...string) (Device, error) {
	if mock.GetDeviceByDeviceIDFunc == nil {
		panic("DeviceRepositoryMock.GetDeviceByDeviceIDFunc: method is nil but DeviceRepository.GetDeviceByDeviceID was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		DeviceID string
		Tenants  []string
	}{
		Ctx:      ctx,
		DeviceID: deviceID,
		Tenants:  tenants,
	}
	mock.lockGetDeviceByDeviceID.Lock()
	mock.calls.GetDeviceByDeviceID = append(mock.calls.GetDeviceByDeviceID, callInfo)
	mock.lockGetDeviceByDeviceID.Unlock()
	return mock.GetDeviceByDeviceIDFunc(ctx, deviceID, tenants...)
}

// GetDeviceByDeviceIDCalls gets all the calls that were made to GetDeviceByDeviceID.
// Check the length with:
//
//	len(mockedDeviceRepository.GetDeviceByDeviceIDCalls())
func (mock *DeviceRepositoryMock) GetDeviceByDeviceIDCalls() []struct {
	Ctx      context.Context
	DeviceID string
	Tenants  []string
} {
	var calls []struct {
		Ctx      context.Context
		DeviceID string
		Tenants  []string
	}
	mock.lockGetDeviceByDeviceID.RLock()
	calls = mock.calls.GetDeviceByDeviceID
	mock.lockGetDeviceByDeviceID.RUnlock()
	return calls
}

// GetDeviceBySensorID calls GetDeviceBySensorIDFunc.
func (mock *DeviceRepositoryMock) GetDeviceBySensorID(ctx context.Context, sensorID string, tenants ...string) (Device, error) {
	if mock.GetDeviceBySensorIDFunc == nil {
		panic("DeviceRepositoryMock.GetDeviceBySensorIDFunc: method is nil but DeviceRepository.GetDeviceBySensorID was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		SensorID string
		Tenants  []string
	}{
		Ctx:      ctx,
		SensorID: sensorID,
		Tenants:  tenants,
	}
	mock.lockGetDeviceBySensorID.Lock()
	mock.calls.GetDeviceBySensorID = append(mock.calls.GetDeviceBySensorID, callInfo)
	mock.lockGetDeviceBySensorID.Unlock()
	return mock.GetDeviceBySensorIDFunc(ctx, sensorID, tenants...)
}

// GetDeviceBySensorIDCalls gets all the calls that were made to GetDeviceBySensorID.
// Check the length with:
//
//	len(mockedDeviceRepository.GetDeviceBySensorIDCalls())
func (mock *DeviceRepositoryMock) GetDeviceBySensorIDCalls() []struct {
	Ctx      context.Context
	SensorID string
	Tenants  []string
} {
	var calls []struct {
		Ctx      context.Context
		SensorID string
		Tenants  []string
	}
	mock.lockGetDeviceBySensorID.RLock()
	calls = mock.calls.GetDeviceBySensorID
	mock.lockGetDeviceBySensorID.RUnlock()
	return calls
}

// GetDevices calls GetDevicesFunc.
func (mock *DeviceRepositoryMock) GetDevices(ctx context.Context, tenants ...string) ([]Device, error) {
	if mock.GetDevicesFunc == nil {
		panic("DeviceRepositoryMock.GetDevicesFunc: method is nil but DeviceRepository.GetDevices was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Tenants []string
	}{
		Ctx:     ctx,
		Tenants: tenants,
	}
	mock.lockGetDevices.Lock()
	mock.calls.GetDevices = append(mock.calls.GetDevices, callInfo)
	mock.lockGetDevices.Unlock()
	return mock.GetDevicesFunc(ctx, tenants...)
}

// GetDevicesCalls gets all the calls that were made to GetDevices.
// Check the length with:
//
//	len(mockedDeviceRepository.GetDevicesCalls())
func (mock *DeviceRepositoryMock) GetDevicesCalls() []struct {
	Ctx     context.Context
	Tenants []string
} {
	var calls []struct {
		Ctx     context.Context
		Tenants []string
	}
	mock.lockGetDevices.RLock()
	calls = mock.calls.GetDevices
	mock.lockGetDevices.RUnlock()
	return calls
}

// GetOnlineDevices calls GetOnlineDevicesFunc.
func (mock *DeviceRepositoryMock) GetOnlineDevices(ctx context.Context, tenants ...string) ([]Device, error) {
	if mock.GetOnlineDevicesFunc == nil {
		panic("DeviceRepositoryMock.GetOnlineDevicesFunc: method is nil but DeviceRepository.GetOnlineDevices was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Tenants []string
	}{
		Ctx:     ctx,
		Tenants: tenants,
	}
	mock.lockGetOnlineDevices.Lock()
	mock.calls.GetOnlineDevices = append(mock.calls.GetOnlineDevices, callInfo)
	mock.lockGetOnlineDevices.Unlock()
	return mock.GetOnlineDevicesFunc(ctx, tenants...)
}

// GetOnlineDevicesCalls gets all the calls that were made to GetOnlineDevices.
// Check the length with:
//
//	len(mockedDeviceRepository.GetOnlineDevicesCalls())
func (mock *DeviceRepositoryMock) GetOnlineDevicesCalls() []struct {
	Ctx     context.Context
	Tenants []string
} {
	var calls []struct {
		Ctx     context.Context
		Tenants []string
	}
	mock.lockGetOnlineDevices.RLock()
	calls = mock.calls.GetOnlineDevices
	mock.lockGetOnlineDevices.RUnlock()
	return calls
}

// RemoveAlarmByID calls RemoveAlarmByIDFunc.
func (mock *DeviceRepositoryMock) RemoveAlarmByID(ctx context.Context, alarmID int) (string, error) {
	if mock.RemoveAlarmByIDFunc == nil {
		panic("DeviceRepositoryMock.RemoveAlarmByIDFunc: method is nil but DeviceRepository.RemoveAlarmByID was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		AlarmID int
	}{
		Ctx:     ctx,
		AlarmID: alarmID,
	}
	mock.lockRemoveAlarmByID.Lock()
	mock.calls.RemoveAlarmByID = append(mock.calls.RemoveAlarmByID, callInfo)
	mock.lockRemoveAlarmByID.Unlock()
	return mock.RemoveAlarmByIDFunc(ctx, alarmID)
}

// RemoveAlarmByIDCalls gets all the calls that were made to RemoveAlarmByID.
// Check the length with:
//
//	len(mockedDeviceRepository.RemoveAlarmByIDCalls())
func (mock *DeviceRepositoryMock) RemoveAlarmByIDCalls() []struct {
	Ctx     context.Context
	AlarmID int
} {
	var calls []struct {
		Ctx     context.Context
		AlarmID int
	}
	mock.lockRemoveAlarmByID.RLock()
	calls = mock.calls.RemoveAlarmByID
	mock.lockRemoveAlarmByID.RUnlock()
	return calls
}

// Save calls SaveFunc.
func (mock *DeviceRepositoryMock) Save(ctx context.Context, device *Device) error {
	if mock.SaveFunc == nil {
		panic("DeviceRepositoryMock.SaveFunc: method is nil but DeviceRepository.Save was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Device *Device
	}{
		Ctx:    ctx,
		Device: device,
	}
	mock.lockSave.Lock()
	mock.calls.Save = append(mock.calls.Save, callInfo)
	mock.lockSave.Unlock()
	return mock.SaveFunc(ctx, device)
}

// SaveCalls gets all the calls that were made to Save.
// Check the length with:
//
//	len(mockedDeviceRepository.SaveCalls())
func (mock *DeviceRepositoryMock) SaveCalls() []struct {
	Ctx    context.Context
	Device *Device
} {
	var calls []struct {
		Ctx    context.Context
		Device *Device
	}
	mock.lockSave.RLock()
	calls = mock.calls.Save
	mock.lockSave.RUnlock()
	return calls
}

// Seed calls SeedFunc.
func (mock *DeviceRepositoryMock) Seed(contextMoqParam context.Context, reader io.Reader, strings ...string) error {
	if mock.SeedFunc == nil {
		panic("DeviceRepositoryMock.SeedFunc: method is nil but DeviceRepository.Seed was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		Reader          io.Reader
		Strings         []string
	}{
		ContextMoqParam: contextMoqParam,
		Reader:          reader,
		Strings:         strings,
	}
	mock.lockSeed.Lock()
	mock.calls.Seed = append(mock.calls.Seed, callInfo)
	mock.lockSeed.Unlock()
	return mock.SeedFunc(contextMoqParam, reader, strings...)
}

// SeedCalls gets all the calls that were made to Seed.
// Check the length with:
//
//	len(mockedDeviceRepository.SeedCalls())
func (mock *DeviceRepositoryMock) SeedCalls() []struct {
	ContextMoqParam context.Context
	Reader          io.Reader
	Strings         []string
} {
	var calls []struct {
		ContextMoqParam context.Context
		Reader          io.Reader
		Strings         []string
	}
	mock.lockSeed.RLock()
	calls = mock.calls.Seed
	mock.lockSeed.RUnlock()
	return calls
}

// UpdateDeviceState calls UpdateDeviceStateFunc.
func (mock *DeviceRepositoryMock) UpdateDeviceState(ctx context.Context, deviceID string, deviceState DeviceState) error {
	if mock.UpdateDeviceStateFunc == nil {
		panic("DeviceRepositoryMock.UpdateDeviceStateFunc: method is nil but DeviceRepository.UpdateDeviceState was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		DeviceID    string
		DeviceState DeviceState
	}{
		Ctx:         ctx,
		DeviceID:    deviceID,
		DeviceState: deviceState,
	}
	mock.lockUpdateDeviceState.Lock()
	mock.calls.UpdateDeviceState = append(mock.calls.UpdateDeviceState, callInfo)
	mock.lockUpdateDeviceState.Unlock()
	return mock.UpdateDeviceStateFunc(ctx, deviceID, deviceState)
}

// UpdateDeviceStateCalls gets all the calls that were made to UpdateDeviceState.
// Check the length with:
//
//	len(mockedDeviceRepository.UpdateDeviceStateCalls())
func (mock *DeviceRepositoryMock) UpdateDeviceStateCalls() []struct {
	Ctx         context.Context
	DeviceID    string
	DeviceState DeviceState
} {
	var calls []struct {
		Ctx         context.Context
		DeviceID    string
		DeviceState DeviceState
	}
	mock.lockUpdateDeviceState.RLock()
	calls = mock.calls.UpdateDeviceState
	mock.lockUpdateDeviceState.RUnlock()
	return calls
}

// UpdateDeviceStatus calls UpdateDeviceStatusFunc.
func (mock *DeviceRepositoryMock) UpdateDeviceStatus(ctx context.Context, deviceID string, deviceStatus DeviceStatus) error {
	if mock.UpdateDeviceStatusFunc == nil {
		panic("DeviceRepositoryMock.UpdateDeviceStatusFunc: method is nil but DeviceRepository.UpdateDeviceStatus was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		DeviceID     string
		DeviceStatus DeviceStatus
	}{
		Ctx:          ctx,
		DeviceID:     deviceID,
		DeviceStatus: deviceStatus,
	}
	mock.lockUpdateDeviceStatus.Lock()
	mock.calls.UpdateDeviceStatus = append(mock.calls.UpdateDeviceStatus, callInfo)
	mock.lockUpdateDeviceStatus.Unlock()
	return mock.UpdateDeviceStatusFunc(ctx, deviceID, deviceStatus)
}

// UpdateDeviceStatusCalls gets all the calls that were made to UpdateDeviceStatus.
// Check the length with:
//
//	len(mockedDeviceRepository.UpdateDeviceStatusCalls())
func (mock *DeviceRepositoryMock) UpdateDeviceStatusCalls() []struct {
	Ctx          context.Context
	DeviceID     string
	DeviceStatus DeviceStatus
} {
	var calls []struct {
		Ctx          context.Context
		DeviceID     string
		DeviceStatus DeviceStatus
	}
	mock.lockUpdateDeviceStatus.RLock()
	calls = mock.calls.UpdateDeviceStatus
	mock.lockUpdateDeviceStatus.RUnlock()
	return calls
}
