package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model
	LicensePlate string   `json:"license_plate" gorm:"unique"`
	VehicleModel string   `json:"vehiclemodel"`
	Year         int      `json:"year"`
	Drivers      []Driver `gorm:"many2many:driver_vehicle_links;"`
}
