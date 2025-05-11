package moduleFrontOfficeVehicleLocationRoute

import (
	"github.com/gin-gonic/gin"
	modulesFrontOfficeVehicleLocationController "transakarta_BE_test/internal/modules/front-office/vehicle-location/controller"
	moduleFrontOfficeVehicleLocationRepository "transakarta_BE_test/internal/modules/front-office/vehicle-location/repository"
)

func RouteVehicleLocation(rg *gin.Engine) {
	repo := moduleFrontOfficeVehicleLocationRepository.NewVehicleLocationRepository()
	controller := modulesFrontOfficeVehicleLocationController.NewVehicleLocationController(repo)

	r := rg.Group("/vehicles")
	r.GET("/:id/location", controller.FindOneLatestLocationByVehicleIdController)
	r.GET("/:id/location/history", controller.PaginateHistoryVehicleLocationByVehicleIdController)
}
