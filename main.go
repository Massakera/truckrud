// file: main.go

package main

import (
	"truckrud/controllers"
	"truckrud/database"
	"truckrud/routes"
	"truckrud/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// initialize the database
	database.InitDatabase()

	// create service instances
	driverService := services.NewDriverService(database.DB)
	vehicleService := services.NewVehicleService(database.DB)
	linkService := services.NewDriverVehicleLinkService(database.DB)

	// create controller instances
	driverController := controllers.NewDriverController(driverService)
	vehicleController := controllers.NewVehicleController(vehicleService)
	linkController := controllers.NewDriverVehicleLinkController(linkService)

	r := gin.Default()

	routes.SetupRoutes(r, driverController, vehicleController, linkController)

	r.Run()
}
