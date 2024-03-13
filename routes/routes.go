// file: routes/routes.go

package routes

import (
	"truckrud/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine, driverController *controllers.DriverController,
	vehicleController *controllers.VehicleController,
	linkController *controllers.DriverVehicleLinkController,
) {
	// routes for Drivers
	router.GET("/drivers", driverController.GetAllDrivers)
	router.GET("/drivers/:id", driverController.GetDriverByID)
	router.POST("/drivers", driverController.AddDriver)
	router.PUT("/drivers/:id", driverController.UpdateDriver)
	router.DELETE("/drivers/:id", driverController.DeleteDriver)

	// routes for Vehicles
	router.GET("/vehicles", vehicleController.GetAllVehicles)
	router.GET("/vehicles/:id", vehicleController.GetVehicleByID)
	router.POST("/vehicles", vehicleController.AddVehicle)
	router.PUT("/vehicles/:id", vehicleController.UpdateVehicle)
	router.DELETE("/vehicles/:id", vehicleController.DeleteVehicle)

	// routes for linking drivers to vehicles
	router.POST("/link/driver-to-vehicle", linkController.LinkDriverToVehicle)
	router.DELETE("/unlink/driver-from-vehicle/:id", linkController.UnlinkDriverFromVehicle)
}
