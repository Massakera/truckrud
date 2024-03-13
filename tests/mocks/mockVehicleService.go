package mocks

import (
	"truckrud/models"

	"github.com/stretchr/testify/mock"
)

type MockVehicleService struct {
	mock.Mock
}

func (m *MockVehicleService) GetAllVehicles() ([]models.Vehicle, error) {
	args := m.Called()
	return args.Get(0).([]models.Vehicle), args.Error(1)
}

func (m *MockVehicleService) GetVehicleByID(id uint) (models.Vehicle, error) {
	args := m.Called(id)
	return args.Get(0).(models.Vehicle), args.Error(1)
}

func (m *MockVehicleService) AddVehicle(vehicle models.Vehicle) (models.Vehicle, error) {
	args := m.Called(vehicle)
	return args.Get(0).(models.Vehicle), args.Error(1)
}

func (m *MockVehicleService) UpdateVehicle(id uint, vehicle models.Vehicle) (models.Vehicle, error) {
	args := m.Called(id, vehicle)
	return args.Get(0).(models.Vehicle), args.Error(1)
}

func (m *MockVehicleService) DeleteVehicle(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
