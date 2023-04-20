// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package alarms

import (
	"context"
	"sync"
)

// Ensure, that AlarmRepositoryMock does implement AlarmRepository.
// If this is not the case, regenerate this file with moq.
var _ AlarmRepository = &AlarmRepositoryMock{}

// AlarmRepositoryMock is a mock implementation of AlarmRepository.
//
//	func TestSomethingThatUsesAlarmRepository(t *testing.T) {
//
//		// make and configure a mocked AlarmRepository
//		mockedAlarmRepository := &AlarmRepositoryMock{
//			AddFunc: func(ctx context.Context, alarm Alarm) (int, error) {
//				panic("mock out the Add method")
//			},
//			CloseFunc: func(ctx context.Context, alarmID int) error {
//				panic("mock out the Close method")
//			},
//			GetAllFunc: func(ctx context.Context, tenants ...string) ([]Alarm, error) {
//				panic("mock out the GetAll method")
//			},
//			GetByIDFunc: func(ctx context.Context, alarmID int) (Alarm, error) {
//				panic("mock out the GetByID method")
//			},
//			GetByRefIDFunc: func(ctx context.Context, refID string) ([]Alarm, error) {
//				panic("mock out the GetByRefID method")
//			},
//		}
//
//		// use mockedAlarmRepository in code that requires AlarmRepository
//		// and then make assertions.
//
//	}
type AlarmRepositoryMock struct {
	// AddFunc mocks the Add method.
	AddFunc func(ctx context.Context, alarm Alarm) (int, error)

	// CloseFunc mocks the Close method.
	CloseFunc func(ctx context.Context, alarmID int) error

	// GetAllFunc mocks the GetAll method.
	GetAllFunc func(ctx context.Context, tenants ...string) ([]Alarm, error)

	// GetByIDFunc mocks the GetByID method.
	GetByIDFunc func(ctx context.Context, alarmID int) (Alarm, error)

	// GetByRefIDFunc mocks the GetByRefID method.
	GetByRefIDFunc func(ctx context.Context, refID string) ([]Alarm, error)

	// calls tracks calls to the methods.
	calls struct {
		// Add holds details about calls to the Add method.
		Add []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Alarm is the alarm argument value.
			Alarm Alarm
		}
		// Close holds details about calls to the Close method.
		Close []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// AlarmID is the alarmID argument value.
			AlarmID int
		}
		// GetAll holds details about calls to the GetAll method.
		GetAll []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenants is the tenants argument value.
			Tenants []string
		}
		// GetByID holds details about calls to the GetByID method.
		GetByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// AlarmID is the alarmID argument value.
			AlarmID int
		}
		// GetByRefID holds details about calls to the GetByRefID method.
		GetByRefID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// RefID is the refID argument value.
			RefID string
		}
	}
	lockAdd        sync.RWMutex
	lockClose      sync.RWMutex
	lockGetAll     sync.RWMutex
	lockGetByID    sync.RWMutex
	lockGetByRefID sync.RWMutex
}

// Add calls AddFunc.
func (mock *AlarmRepositoryMock) Add(ctx context.Context, alarm Alarm) (int, error) {
	if mock.AddFunc == nil {
		panic("AlarmRepositoryMock.AddFunc: method is nil but AlarmRepository.Add was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Alarm Alarm
	}{
		Ctx:   ctx,
		Alarm: alarm,
	}
	mock.lockAdd.Lock()
	mock.calls.Add = append(mock.calls.Add, callInfo)
	mock.lockAdd.Unlock()
	return mock.AddFunc(ctx, alarm)
}

// AddCalls gets all the calls that were made to Add.
// Check the length with:
//
//	len(mockedAlarmRepository.AddCalls())
func (mock *AlarmRepositoryMock) AddCalls() []struct {
	Ctx   context.Context
	Alarm Alarm
} {
	var calls []struct {
		Ctx   context.Context
		Alarm Alarm
	}
	mock.lockAdd.RLock()
	calls = mock.calls.Add
	mock.lockAdd.RUnlock()
	return calls
}

// Close calls CloseFunc.
func (mock *AlarmRepositoryMock) Close(ctx context.Context, alarmID int) error {
	if mock.CloseFunc == nil {
		panic("AlarmRepositoryMock.CloseFunc: method is nil but AlarmRepository.Close was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		AlarmID int
	}{
		Ctx:     ctx,
		AlarmID: alarmID,
	}
	mock.lockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	mock.lockClose.Unlock()
	return mock.CloseFunc(ctx, alarmID)
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//
//	len(mockedAlarmRepository.CloseCalls())
func (mock *AlarmRepositoryMock) CloseCalls() []struct {
	Ctx     context.Context
	AlarmID int
} {
	var calls []struct {
		Ctx     context.Context
		AlarmID int
	}
	mock.lockClose.RLock()
	calls = mock.calls.Close
	mock.lockClose.RUnlock()
	return calls
}

// GetAll calls GetAllFunc.
func (mock *AlarmRepositoryMock) GetAll(ctx context.Context, tenants ...string) ([]Alarm, error) {
	if mock.GetAllFunc == nil {
		panic("AlarmRepositoryMock.GetAllFunc: method is nil but AlarmRepository.GetAll was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Tenants []string
	}{
		Ctx:     ctx,
		Tenants: tenants,
	}
	mock.lockGetAll.Lock()
	mock.calls.GetAll = append(mock.calls.GetAll, callInfo)
	mock.lockGetAll.Unlock()
	return mock.GetAllFunc(ctx, tenants...)
}

// GetAllCalls gets all the calls that were made to GetAll.
// Check the length with:
//
//	len(mockedAlarmRepository.GetAllCalls())
func (mock *AlarmRepositoryMock) GetAllCalls() []struct {
	Ctx     context.Context
	Tenants []string
} {
	var calls []struct {
		Ctx     context.Context
		Tenants []string
	}
	mock.lockGetAll.RLock()
	calls = mock.calls.GetAll
	mock.lockGetAll.RUnlock()
	return calls
}

// GetByID calls GetByIDFunc.
func (mock *AlarmRepositoryMock) GetByID(ctx context.Context, alarmID int) (Alarm, error) {
	if mock.GetByIDFunc == nil {
		panic("AlarmRepositoryMock.GetByIDFunc: method is nil but AlarmRepository.GetByID was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		AlarmID int
	}{
		Ctx:     ctx,
		AlarmID: alarmID,
	}
	mock.lockGetByID.Lock()
	mock.calls.GetByID = append(mock.calls.GetByID, callInfo)
	mock.lockGetByID.Unlock()
	return mock.GetByIDFunc(ctx, alarmID)
}

// GetByIDCalls gets all the calls that were made to GetByID.
// Check the length with:
//
//	len(mockedAlarmRepository.GetByIDCalls())
func (mock *AlarmRepositoryMock) GetByIDCalls() []struct {
	Ctx     context.Context
	AlarmID int
} {
	var calls []struct {
		Ctx     context.Context
		AlarmID int
	}
	mock.lockGetByID.RLock()
	calls = mock.calls.GetByID
	mock.lockGetByID.RUnlock()
	return calls
}

// GetByRefID calls GetByRefIDFunc.
func (mock *AlarmRepositoryMock) GetByRefID(ctx context.Context, refID string) ([]Alarm, error) {
	if mock.GetByRefIDFunc == nil {
		panic("AlarmRepositoryMock.GetByRefIDFunc: method is nil but AlarmRepository.GetByRefID was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		RefID string
	}{
		Ctx:   ctx,
		RefID: refID,
	}
	mock.lockGetByRefID.Lock()
	mock.calls.GetByRefID = append(mock.calls.GetByRefID, callInfo)
	mock.lockGetByRefID.Unlock()
	return mock.GetByRefIDFunc(ctx, refID)
}

// GetByRefIDCalls gets all the calls that were made to GetByRefID.
// Check the length with:
//
//	len(mockedAlarmRepository.GetByRefIDCalls())
func (mock *AlarmRepositoryMock) GetByRefIDCalls() []struct {
	Ctx   context.Context
	RefID string
} {
	var calls []struct {
		Ctx   context.Context
		RefID string
	}
	mock.lockGetByRefID.RLock()
	calls = mock.calls.GetByRefID
	mock.lockGetByRefID.RUnlock()
	return calls
}
