package controllers

import (
	"net/http"
	"strconv"
	"truckrud/models"
	"truckrud/services"

	"github.com/gin-gonic/gin"
)

type VehicleController struct {
	Service services.IVehicleService
}

func NewVehicleController(service services.IVehicleService) *VehicleController {
	return &VehicleController{Service: service}
}

// @Summary List all vehicles
// @Description Get a list of all vehicles
// @Tags vehicles
// @Accept json
// @Produce json
// @Success 200 {array} models.Vehicle
// @Failure 500 {object} object{error=string}
// @Router /vehicles [get]
func (vc *VehicleController) GetAllVehicles(c *gin.Context) {
	vehicles, err := vc.Service.GetAllVehicles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vehicles)
}

// @Summary Get a vehicle by ID
// @Description Get details of a vehicle with the given ID
// @Tags vehicles
// @Accept json
// @Produce json
// @Param id path int true "Vehicle ID"
// @Success 200 {object} models.Vehicle
// @Failure 404 {object} object{error=string}
// @Router /vehicles/{id} [get]
func (vc *VehicleController) GetVehicleByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	vehicle, err := vc.Service.GetVehicleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}
	c.JSON(http.StatusOK, vehicle)
}

// @Summary Add a new vehicle
// @Description Add a new vehicle with the provided information
// @Tags vehicles
// @Accept json
// @Produce json
// @Param vehicle body models.Vehicle true "Vehicle data"
// @Success 201 {object} models.Vehicle
// @Failure 400 {object} object{error=string}
// @Router /vehicles [post]
func (vc *VehicleController) AddVehicle(c *gin.Context) {
	var vehicle models.Vehicle
	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdVehicle, err := vc.Service.AddVehicle(vehicle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdVehicle)
}

// @Summary Update a vehicle
// @Description Update a vehicle with the given ID
// @Tags vehicles
// @Accept json
// @Produce json
// @Param id path int true "Vehicle ID"
// @Param vehicle body models.Vehicle true "Vehicle data"
// @Success 200 {object} models.Vehicle
// @Failure 400 {object} object{error=string}
// @Failure 500 {object} object{error=string}
// @Router /vehicles/{id} [put]
func (vc *VehicleController) UpdateVehicle(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var vehicleData models.Vehicle
	if err := c.ShouldBindJSON(&vehicleData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedVehicle, err := vc.Service.UpdateVehicle(uint(id), vehicleData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedVehicle)
}

// @Summary Delete a vehicle
// @Description Delete the vehicle with the given ID
// @Tags vehicles
// @Accept json
// @Produce json
// @Param id path int true "Vehicle ID"
// @Success 200 {object} object{message=string}
// @Failure 404 {object} object{error=string}
// @Router /vehicles/{id} [delete]
func (vc *VehicleController) DeleteVehicle(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	err := vc.Service.DeleteVehicle(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Vehicle deleted successfully"})
}
