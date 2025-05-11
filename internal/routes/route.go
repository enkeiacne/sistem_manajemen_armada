package routes

import (
	"github.com/gin-gonic/gin"
	moduleFrontOfficeVehicleLocationRoute "transakarta_BE_test/internal/modules/front-office/vehicle-location/route"
)

func SetupRoutes(r *gin.Engine) {
	moduleFrontOfficeVehicleLocationRoute.RouteVehicleLocation(r)
}
