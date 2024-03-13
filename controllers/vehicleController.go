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

func (vc *VehicleController) GetAllVehicles(c *gin.Context) {
	vehicles, err := vc.Service.GetAllVehicles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vehicles)
}

func (vc *VehicleController) GetVehicleByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	vehicle, err := vc.Service.GetVehicleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}
	c.JSON(http.StatusOK, vehicle)
}

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

func (vc *VehicleController) DeleteVehicle(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	err := vc.Service.DeleteVehicle(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Vehicle deleted successfully"})
}
