package mocks

import (
	"truckrud/models"

	"github.com/stretchr/testify/mock"
)

type MockDriverService struct {
	mock.Mock
}

func (m *MockDriverService) GetAllDrivers() ([]models.Driver, error) {
	args := m.Called()
	return args.Get(0).([]models.Driver), args.Error(1)
}

func (m *MockDriverService) GetDriverByID(id uint) (models.Driver, error) {
	args := m.Called(id)
	return args.Get(0).(models.Driver), args.Error(1)
}

func (m *MockDriverService) AddDriver(driver models.Driver) (models.Driver, error) {
	args := m.Called(driver)
	return args.Get(0).(models.Driver), args.Error(1)
}

func (m *MockDriverService) UpdateDriver(id uint, driver models.Driver) (models.Driver, error) {
	args := m.Called(id, driver)
	return args.Get(0).(models.Driver), args.Error(1)
}

func (m *MockDriverService) DeleteDriver(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
