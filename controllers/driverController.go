package controllers

import (
	"net/http"
	"strconv"
	"truckrud/models"
	"truckrud/services"

	"github.com/gin-gonic/gin"
)

type DriverController struct {
	Service services.IDriverService
}

func NewDriverController(service services.IDriverService) *DriverController {
	return &DriverController{Service: service}
}

func (dc *DriverController) GetAllDrivers(c *gin.Context) {
	drivers, err := dc.Service.GetAllDrivers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, drivers)
}

func (dc *DriverController) GetDriverByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	driver, err := dc.Service.GetDriverByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Driver not found"})
		return
	}
	c.JSON(http.StatusOK, driver)
}

func (dc *DriverController) AddDriver(c *gin.Context) {
	var driver models.Driver
	if err := c.ShouldBindJSON(&driver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdDriver, err := dc.Service.AddDriver(driver)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdDriver)
}

func (dc *DriverController) UpdateDriver(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var driverData models.Driver
	if err := c.ShouldBindJSON(&driverData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedDriver, err := dc.Service.UpdateDriver(uint(id), driverData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedDriver)
}

func (dc *DriverController) DeleteDriver(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	err := dc.Service.DeleteDriver(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Driver not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Driver deleted successfully"})
}
