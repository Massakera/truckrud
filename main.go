// file: main.go

package main

import (
	"truckrud/controllers"
	"truckrud/database"
	docs "truckrud/docs"
	"truckrud/routes"
	"truckrud/services"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.SetupRoutes(r, driverController, vehicleController, linkController)

	r.Run()
}
