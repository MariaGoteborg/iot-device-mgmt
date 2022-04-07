package application

import (
	"context"

	"github.com/diwise/iot-device-mgmt/internal/pkg/infrastructure/repositories/database"
)

type DeviceManagement interface {
	GetDevice(context.Context, string) (database.Device, error)
}

func New(db database.Datastore) DeviceManagement {
	a := &app{
		db: db,
	}

	return a
}

type app struct {
	db database.Datastore
}

func (a *app) GetDevice(ctx context.Context, externalID string) (database.Device, error) {

	device, err := a.db.GetDeviceFromDevEUI(externalID)
	if err != nil {
		return nil, err
	}

	return device, nil
}
