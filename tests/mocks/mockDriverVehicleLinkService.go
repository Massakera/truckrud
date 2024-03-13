package mocks

import (
	"truckrud/models"

	"github.com/stretchr/testify/mock"
)

type MockDriverVehicleLinkService struct {
	mock.Mock
}

func (_m *MockDriverVehicleLinkService) LinkDriverToVehicle(driverID, vehicleID uint) (models.DriverVehicleLink, error) {
	ret := _m.Called(driverID, vehicleID)
	var r0 models.DriverVehicleLink
	if rf, ok := ret.Get(0).(func(uint, uint) models.DriverVehicleLink); ok {
		r0 = rf(driverID, vehicleID)
	} else {
		r0 = ret.Get(0).(models.DriverVehicleLink)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(driverID, vehicleID)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

func (_m *MockDriverVehicleLinkService) UnlinkDriverFromVehicle(linkID uint) error {
	ret := _m.Called(linkID)
	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(linkID)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}
