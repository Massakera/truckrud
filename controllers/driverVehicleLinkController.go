package controllers

import (
	"net/http"
	"strconv"
	"truckrud/models"
	"truckrud/services"

	"github.com/gin-gonic/gin"
)

type DriverVehicleLinkController struct {
	Service services.IDriverVehicleLinkService
}

func NewDriverVehicleLinkController(service services.IDriverVehicleLinkService) *DriverVehicleLinkController {
	return &DriverVehicleLinkController{Service: service}
}

// @Summary Link a driver to a vehicle
// @Description Create a link between a driver and a vehicle
// @Tags link
// @Accept json
// @Produce json
// @Param request body models.DriverVehicleLink true "Link request"
// @Success 201 {object} models.DriverVehicleLink
// @Failure 400 {object} object{error=string}
// @Failure 500 {object} object{error=string}
// @Router /link/driver/vehicle [post]
func (c *DriverVehicleLinkController) LinkDriverToVehicle(ctx *gin.Context) {
	var request models.DriverVehicleLink
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	link, err := c.Service.LinkDriverToVehicle(request.DriverID, request.VehicleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, link)
}

// @Summary Unlink a driver from a vehicle
// @Description Remove the link between a driver and a vehicle using the link ID
// @Tags link
// @Accept json
// @Produce json
// @Param id path int true "Link ID"
// @Success 200 {object} object{message=string}
// @Failure 400 {object} object{error=string}
// @Failure 500 {object} object{error=string}
// @Router /unlink/driver/vehicle/{id} [delete]
func (c *DriverVehicleLinkController) UnlinkDriverFromVehicle(ctx *gin.Context) {
	linkID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid link ID"})
		return
	}

	if err := c.Service.UnlinkDriverFromVehicle(uint(linkID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Link removed successfully"})
}
