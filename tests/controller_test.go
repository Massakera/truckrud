package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"truckrud/controllers"
	"truckrud/models"
	"truckrud/tests/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func TestGetAllDrivers(t *testing.T) {
	mockService := new(mocks.MockDriverService)
	controller := controllers.NewDriverController(mockService)

	expectedDrivers := []models.Driver{{Name: "John Doe"}}
	mockService.On("GetAllDrivers").Return(expectedDrivers, nil)

	router := setupRouter()
	router.GET("/drivers", controller.GetAllDrivers)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/drivers", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var actualDrivers []models.Driver
	json.Unmarshal(w.Body.Bytes(), &actualDrivers)
	assert.Equal(t, expectedDrivers, actualDrivers)
	mockService.AssertExpectations(t)
}

func TestGetDriverByID(t *testing.T) {
	mockService := new(mocks.MockDriverService)
	controller := controllers.NewDriverController(mockService)
	expectedDriver := models.Driver{Model: gorm.Model{ID: 1}, Name: "John Doe"}

	mockService.On("GetDriverByID", uint(1)).Return(expectedDriver, nil)

	router := setupRouter()
	router.GET("/drivers/:id", controller.GetDriverByID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/drivers/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var actualDriver models.Driver
	err := json.Unmarshal(w.Body.Bytes(), &actualDriver)
	assert.NoError(t, err)
	assert.Equal(t, expectedDriver, actualDriver)

	mockService.AssertExpectations(t)
}

func TestAddDriver(t *testing.T) {
	mockService := new(mocks.MockDriverService)
	controller := controllers.NewDriverController(mockService)
	driverToAdd := models.Driver{Name: "New Driver", CNH: "XYZ", BirthDate: "1990-01-01"}
	expectedDriver := models.Driver{Model: gorm.Model{ID: 1}, Name: "New Driver", CNH: "XYZ", BirthDate: "1990-01-01"}

	mockService.On("AddDriver", mock.AnythingOfType("models.Driver")).Return(expectedDriver, nil)

	router := setupRouter()
	router.POST("/drivers", controller.AddDriver)

	requestBody, _ := json.Marshal(driverToAdd)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/drivers", bytes.NewBuffer(requestBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var actualDriver models.Driver
	err := json.Unmarshal(w.Body.Bytes(), &actualDriver)
	assert.NoError(t, err)
	assert.Equal(t, expectedDriver, actualDriver)

	mockService.AssertExpectations(t)
}

func TestUpdateDriver(t *testing.T) {
	mockService := new(mocks.MockDriverService)
	controller := controllers.NewDriverController(mockService)
	updatedDriver := models.Driver{Model: gorm.Model{ID: 1}, Name: "Updated Name", CNH: "ABC", BirthDate: "1989-12-12"}

	mockService.On("UpdateDriver", uint(1), mock.AnythingOfType("models.Driver")).Return(updatedDriver, nil)

	router := setupRouter()
	router.PUT("/drivers/:id", controller.UpdateDriver)

	requestBody, _ := json.Marshal(updatedDriver)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/drivers/1", bytes.NewBuffer(requestBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var actualDriver models.Driver
	err := json.Unmarshal(w.Body.Bytes(), &actualDriver)
	assert.NoError(t, err)
	assert.Equal(t, updatedDriver, actualDriver)

	mockService.AssertExpectations(t)
}

func TestDeleteDriver(t *testing.T) {
	mockService := new(mocks.MockDriverService)
	controller := controllers.NewDriverController(mockService)

	mockService.On("DeleteDriver", uint(1)).Return(nil)

	router := setupRouter()
	router.DELETE("/drivers/:id", controller.DeleteDriver)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/drivers/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	mockService.AssertExpectations(t)
}

func TestGetAllVehicles(t *testing.T) {
	mockService := new(mocks.MockVehicleService)
	controller := controllers.NewVehicleController(mockService)

	expectedVehicles := []models.Vehicle{{Model: gorm.Model{ID: 1}, LicensePlate: "XYZ123", VehicleModel: "Toyota", Year: 2020}}
	mockService.On("GetAllVehicles").Return(expectedVehicles, nil)

	router := setupRouter()
	router.GET("/vehicles", controller.GetAllVehicles)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/vehicles", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var actualVehicles []models.Vehicle
	json.Unmarshal(w.Body.Bytes(), &actualVehicles)
	assert.Equal(t, expectedVehicles, actualVehicles)

	mockService.AssertExpectations(t)
}

func TestGetVehicleByID(t *testing.T) {
	mockService := new(mocks.MockVehicleService)
	controller := controllers.NewVehicleController(mockService)

	expectedVehicle := models.Vehicle{Model: gorm.Model{ID: 1}, LicensePlate: "XYZ123", VehicleModel: "Toyota", Year: 2020}
	mockService.On("GetVehicleByID", uint(1)).Return(expectedVehicle, nil)

	router := setupRouter()
	router.GET("/vehicles/:id", controller.GetVehicleByID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/vehicles/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var actualVehicle models.Vehicle
	json.Unmarshal(w.Body.Bytes(), &actualVehicle)
	assert.Equal(t, expectedVehicle, actualVehicle)

	mockService.AssertExpectations(t)
}

func TestAddVehicle(t *testing.T) {
	mockService := new(mocks.MockVehicleService)
	controller := controllers.NewVehicleController(mockService)

	vehicleToAdd := models.Vehicle{LicensePlate: "XYZ456", VehicleModel: "Honda", Year: 2021}
	expectedVehicle := models.Vehicle{Model: gorm.Model{ID: 2}, LicensePlate: "XYZ456", VehicleModel: "Honda", Year: 2021}
	mockService.On("AddVehicle", mock.AnythingOfType("models.Vehicle")).Return(expectedVehicle, nil)

	router := setupRouter()
	router.POST("/vehicles", controller.AddVehicle)

	requestBody, _ := json.Marshal(vehicleToAdd)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/vehicles", bytes.NewBuffer(requestBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var actualVehicle models.Vehicle
	json.Unmarshal(w.Body.Bytes(), &actualVehicle)
	assert.Equal(t, expectedVehicle, actualVehicle)

	mockService.AssertExpectations(t)
}

func TestUpdateVehicle(t *testing.T) {
	mockService := new(mocks.MockVehicleService)
	controller := controllers.NewVehicleController(mockService)

	updatedVehicleInput := models.Vehicle{LicensePlate: "ABC123", VehicleModel: "Updated Model", Year: 2021}
	updatedVehicle := models.Vehicle{Model: gorm.Model{ID: 1}, LicensePlate: "ABC123", VehicleModel: "Updated Model", Year: 2021}

	mockService.On("UpdateVehicle", uint(1), mock.AnythingOfType("models.Vehicle")).Return(updatedVehicle, nil)

	router := setupRouter()
	router.PUT("/vehicles/:id", controller.UpdateVehicle)

	requestBody, _ := json.Marshal(updatedVehicleInput)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/vehicles/1", bytes.NewBuffer(requestBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var actualVehicle models.Vehicle
	err := json.Unmarshal(w.Body.Bytes(), &actualVehicle)
	assert.NoError(t, err)
	assert.Equal(t, updatedVehicle, actualVehicle)

	mockService.AssertExpectations(t)
}

func TestDeleteVehicle(t *testing.T) {
	mockService := new(mocks.MockVehicleService)
	controller := controllers.NewVehicleController(mockService)

	mockService.On("DeleteVehicle", uint(1)).Return(nil) // Assume successful deletion

	router := setupRouter()
	router.DELETE("/vehicles/:id", controller.DeleteVehicle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/vehicles/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"message\":\"Vehicle deleted successfully\"}", w.Body.String())

	mockService.AssertExpectations(t)
}
