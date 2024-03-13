package services

import "truckrud/models"

type IDriverService interface {
	GetAllDrivers() ([]models.Driver, error)
	GetDriverByID(id uint) (models.Driver, error)
	AddDriver(driver models.Driver) (models.Driver, error)
	UpdateDriver(id uint, driver models.Driver) (models.Driver, error)
	DeleteDriver(id uint) error
}

type IVehicleService interface {
	GetAllVehicles() ([]models.Vehicle, error)
	GetVehicleByID(id uint) (models.Vehicle, error)
	AddVehicle(vehicle models.Vehicle) (models.Vehicle, error)
	UpdateVehicle(id uint, vehicle models.Vehicle) (models.Vehicle, error)
	DeleteVehicle(id uint) error
}

type IDriverVehicleLinkService interface {
	LinkDriverToVehicle(driverID, vehicleID uint) (models.DriverVehicleLink, error)
	UnlinkDriverFromVehicle(linkID uint) error
}
