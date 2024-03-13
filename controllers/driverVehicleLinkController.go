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
