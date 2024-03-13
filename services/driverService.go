package services

import (
	"truckrud/models"

	"gorm.io/gorm"
)

type DriverService struct {
	DB *gorm.DB
}

func NewDriverService(db *gorm.DB) *DriverService {
	return &DriverService{DB: db}
}

func (s *DriverService) GetAllDrivers() ([]models.Driver, error) {
	var drivers []models.Driver
	result := s.DB.Preload("Vehicles").Find(&drivers)
	return drivers, result.Error
}

func (s *DriverService) GetDriverByID(id uint) (models.Driver, error) {
	var driver models.Driver
	result := s.DB.Preload("Vehicles").First(&driver, id)
	return driver, result.Error
}

func (s *DriverService) AddDriver(driver models.Driver) (models.Driver, error) {
	result := s.DB.Create(&driver)
	return driver, result.Error
}

func (s *DriverService) UpdateDriver(id uint, driverData models.Driver) (models.Driver, error) {
	var driver models.Driver
	result := s.DB.First(&driver, id)
	if result.Error != nil {
		return models.Driver{}, result.Error
	}

	result = s.DB.Model(&driver).Updates(driverData)
	return driver, result.Error
}

func (s *DriverService) DeleteDriver(id uint) error {
	result := s.DB.Delete(&models.Driver{}, id)
	return result.Error
}
