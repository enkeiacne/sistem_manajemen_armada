package databaseEntities

import "time"

type VehicleLocation struct {
	ID        string    `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4();not null"`
	VehicleID string    `gorm:"column:vehicle_id;type:varchar(50);not null; index;"`
	Latitude  float64   `gorm:"column:latitude;type:double precision;not null"`
	Longitude float64   `gorm:"column:longitude;type:double precision;not null"`
	Timestamp time.Time `gorm:"column:timestamp;type:timestamptz;not null"`
}
