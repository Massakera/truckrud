package services

import (
	"truckrud/models"

	"gorm.io/gorm"
)

type VehicleService struct {
	DB *gorm.DB
}

func NewVehicleService(db *gorm.DB) *VehicleService {
	return &VehicleService{DB: db}
}

func (s *VehicleService) GetAllVehicles() ([]models.Vehicle, error) {
	var vehicles []models.Vehicle
	result := s.DB.Find(&vehicles)
	return vehicles, result.Error
}

func (s *VehicleService) GetVehicleByID(id uint) (models.Vehicle, error) {
	var vehicle models.Vehicle
	result := s.DB.First(&vehicle, id)
	return vehicle, result.Error
}

func (s *VehicleService) AddVehicle(vehicle models.Vehicle) (models.Vehicle, error) {
	result := s.DB.Create(&vehicle)
	return vehicle, result.Error
}

func (s *VehicleService) UpdateVehicle(id uint, vehicleData models.Vehicle) (models.Vehicle, error) {
	var vehicle models.Vehicle
	result := s.DB.First(&vehicle, id)
	if result.Error != nil {
		return models.Vehicle{}, result.Error
	}

	result = s.DB.Model(&vehicle).Updates(vehicleData)
	return vehicle, result.Error
}

func (s *VehicleService) DeleteVehicle(id uint) error {
	result := s.DB.Delete(&models.Vehicle{}, id)
	return result.Error
}
