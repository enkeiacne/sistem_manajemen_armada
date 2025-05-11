package moduleFrontOfficeVehicleLocationDtoRequest

type FrontOfficeVehicleLocationCreateRequestDto struct {
	VehicleID string  `form:"vehicle_id" binding:"required"`
	Latitude  float64 `form:"latitude" binding:"required"`
	Longitude float64 `form:"longitude" binding:"required"`
	TimeStamp int64   `form:"time_stamp" binding:"required,unix_ts"`
}
type FrontOfficeVehicleLocationHistoryQueryParamRequest struct {
	Page  *int64 `form:"page"`
	Limit *int64 `form:"limit"`
	Start *int64 `form:"start"`
	End   *int64 `form:"end"`
}
