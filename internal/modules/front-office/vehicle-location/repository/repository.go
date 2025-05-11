package moduleFrontOfficeVehicleLocationRepository

import (
	"time"
	commonsHelperPagination "transakarta_BE_test/internal/commons/helper/pagination"
	"transakarta_BE_test/internal/database"
	databaseEntities "transakarta_BE_test/internal/database/entities"
	moduleFrontOfficeVehicleLocationDtoRequest "transakarta_BE_test/internal/modules/front-office/vehicle-location/dto/request"
	moduleFrontOfficeVehicleLocationDtoResponse "transakarta_BE_test/internal/modules/front-office/vehicle-location/dto/response"
)

type VehicleLocationRepository interface {
	FindOneLatestLocationByVehicleId(vehicleId string) (*databaseEntities.VehicleLocation, error)
	PaginateHistoryLocationByVehicleId(request moduleFrontOfficeVehicleLocationDtoRequest.FrontOfficeVehicleLocationHistoryQueryParamRequest, vehicleId string) (*commonsHelperPagination.PaginationResult, error)
	Create(databaseEntities.VehicleLocation) (*databaseEntities.VehicleLocation, error)
}
type vehicleLocationRepositoryImpl struct {
}

func NewVehicleLocationRepository() VehicleLocationRepository {
	return &vehicleLocationRepositoryImpl{}
}

// ============ logic ===================
func (v *vehicleLocationRepositoryImpl) FindOneLatestLocationByVehicleId(vehicleId string) (*databaseEntities.VehicleLocation, error) {
	var data databaseEntities.VehicleLocation
	if err := database.DB.First(&data, "vehicle_id = ?", vehicleId).Order("timestamp DESC").Error; err != nil {
		return nil, err
	}
	return &data, nil
}
func (v *vehicleLocationRepositoryImpl) PaginateHistoryLocationByVehicleId(queryParam moduleFrontOfficeVehicleLocationDtoRequest.FrontOfficeVehicleLocationHistoryQueryParamRequest, vehicleId string) (*commonsHelperPagination.PaginationResult, error) {
	var dataVehicleLocations []databaseEntities.VehicleLocation
	limit := 10
	page := 1

	if queryParam.Page != nil {
		page = int(*queryParam.Page)
	}
	if queryParam.Limit != nil {
		limit = int(*queryParam.Limit)
	}

	sql := database.DB.Where("vehicle_id = ?", vehicleId).Order("timestamp DESC")

	//	==== filter =====
	if queryParam.Start != nil {
		sql = sql.Where("timestamp >= ?", time.Unix(*queryParam.Start, 0))
	}
	if queryParam.End != nil {
		sql = sql.Where("timestamp <= ?", time.Unix(*queryParam.End, 0))
	}

	pagination, err := commonsHelperPagination.Paginate(sql, page, limit, &databaseEntities.VehicleLocation{}, &dataVehicleLocations)
	if err != nil {
		return nil, err
	}
	mapped := make([]moduleFrontOfficeVehicleLocationDtoResponse.VehicleLocationResponse, 0, len(dataVehicleLocations))
	for _, item := range dataVehicleLocations {
		mapped = append(mapped, moduleFrontOfficeVehicleLocationDtoResponse.VehicleLocationResponse{
			VehicleID: item.VehicleID,
			Latitude:  item.Latitude,
			Longitude: item.Longitude,
			Timestamp: item.Timestamp.Unix(),
		})
	}
	pagination.Data = mapped

	return pagination, nil
}
func (v *vehicleLocationRepositoryImpl) Create(data databaseEntities.VehicleLocation) (*databaseEntities.VehicleLocation, error) {
	var newVehicleLocation databaseEntities.VehicleLocation
	newVehicleLocation.VehicleID = data.VehicleID
	newVehicleLocation.Latitude = data.Latitude
	newVehicleLocation.Longitude = data.Longitude
	newVehicleLocation.Timestamp = data.Timestamp

	if err := database.DB.Create(&newVehicleLocation).Error; err != nil {
		return nil, err
	}
	return &newVehicleLocation, nil
}
