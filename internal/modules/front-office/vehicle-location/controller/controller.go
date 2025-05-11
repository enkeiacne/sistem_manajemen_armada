package modulesFrontOfficeVehicleLocationController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	commonsHelper "transakarta_BE_test/internal/commons/helper"
	moduleFrontOfficeVehicleLocationDtoRequest "transakarta_BE_test/internal/modules/front-office/vehicle-location/dto/request"
	moduleFrontOfficeVehicleLocationDtoResponse "transakarta_BE_test/internal/modules/front-office/vehicle-location/dto/response"
	moduleFrontOfficeVehicleLocationRepository "transakarta_BE_test/internal/modules/front-office/vehicle-location/repository"
)

type vehicleLocationController struct {
	repository moduleFrontOfficeVehicleLocationRepository.VehicleLocationRepository
}

func NewVehicleLocationController(repository moduleFrontOfficeVehicleLocationRepository.VehicleLocationRepository) *vehicleLocationController {
	return &vehicleLocationController{repository}
}

func (v *vehicleLocationController) FindOneLatestLocationByVehicleIdController(c *gin.Context) {
	vehicleId := c.Param("id")
	if vehicleId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "VehicleId is required"})
		return
	}
	data, err := v.repository.FindOneLatestLocationByVehicleId(vehicleId)
	if err != nil || data == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := moduleFrontOfficeVehicleLocationDtoResponse.VehicleLocationResponse{
		VehicleID: data.VehicleID,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		Timestamp: data.Timestamp.Unix(),
	}
	c.JSON(http.StatusOK, response)
}
func (v *vehicleLocationController) PaginateHistoryVehicleLocationByVehicleIdController(c *gin.Context) {
	vehicleId := c.Param("id")
	if vehicleId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "VehicleId is required"})
		return
	}

	var queryParams moduleFrontOfficeVehicleLocationDtoRequest.FrontOfficeVehicleLocationHistoryQueryParamRequest
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	if queryParams.End != nil {
		end := *queryParams.End
		check := commonsHelper.CheckUnixTimestamp(end)
		if !check {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query end parameters"})
		}
	}
	if queryParams.Start != nil {
		start := *queryParams.Start
		check := commonsHelper.CheckUnixTimestamp(start)
		if !check {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query start parameters"})
		}
	}
	fmt.Print(queryParams)
	pagination, err := v.repository.PaginateHistoryLocationByVehicleId(queryParams, vehicleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, pagination)
}
