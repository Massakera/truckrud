package services

import (
	"truckrud/models"

	"gorm.io/gorm"
)

type DriverVehicleLinkService struct {
	DB *gorm.DB
}

func NewDriverVehicleLinkService(db *gorm.DB) *DriverVehicleLinkService {
	return &DriverVehicleLinkService{DB: db}
}

func (svc *DriverVehicleLinkService) LinkDriverToVehicle(driverID, vehicleID uint) (models.DriverVehicleLink, error) {
	link := models.DriverVehicleLink{DriverID: driverID, VehicleID: vehicleID}
	result := svc.DB.Create(&link)
	return link, result.Error
}

func (svc *DriverVehicleLinkService) UnlinkDriverFromVehicle(linkID uint) error {
	result := svc.DB.Delete(&models.DriverVehicleLink{}, linkID)
	return result.Error
}
