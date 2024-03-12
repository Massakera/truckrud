package models

import "gorm.io/gorm"

type DriverVehicleLink struct {
	gorm.Model
	DriverID  uint `json:"driver_id"`
	VehicleID uint `json:"vehicle_id"`
}
