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

// @Summary List all drivers
// @Description Get a list of all drivers
// @Tags drivers
// @Accept json
// @Produce json
// @Success 200 {array} models.Driver
// @Router /drivers [get]
func (dc *DriverController) GetAllDrivers(c *gin.Context) {
	drivers, err := dc.Service.GetAllDrivers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, drivers)
}

// @Summary Get a driver by ID
// @Description Get details of a driver with the given ID
// @Tags drivers
// @Accept json
// @Produce json
// @Param id path int true "Driver ID"
// @Success 200 {object} models.Driver
// @Failure 404 {object} object{error=string}
// @Router /drivers/{id} [get]
func (dc *DriverController) GetDriverByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	driver, err := dc.Service.GetDriverByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Driver not found"})
		return
	}
	c.JSON(http.StatusOK, driver)
}

// @Summary Add a new driver
// @Description Add a new driver with the provided information
// @Tags drivers
// @Accept json
// @Produce json
// @Param driver body models.Driver true "Driver data"
// @Success 201 {object} models.Driver
// @Failure 400 {object} object{error=string}
// @Router /drivers [post]
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

// @Summary Update a driver
// @Description Update a driver with the given ID
// @Tags drivers
// @Accept json
// @Produce json
// @Param id path int true "Driver ID"
// @Param driver body models.Driver true "Driver data"
// @Success 200 {object} models.Driver
// @Failure 400 {object} object{error=string}
// @Failure 404 {object} object{error=string}
// @Router /drivers/{id} [put]
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

// @Summary Delete a driver
// @Description Delete the driver with the given ID
// @Tags drivers
// @Accept json
// @Produce json
// @Param id path int true "Driver ID"
// @Success 200 {object} object{message=string}
// @Failure 404 {object} object{error=string}
// @Router /drivers/{id} [delete]
func (dc *DriverController) DeleteDriver(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	err := dc.Service.DeleteDriver(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Driver not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Driver deleted successfully"})
}
