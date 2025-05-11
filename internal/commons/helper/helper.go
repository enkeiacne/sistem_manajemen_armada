package commonsHelper

import (
	"math"
	"time"
)

func CheckUnixTimestamp(timestamp int64) bool {
	now := time.Now().Unix()
	m := now + 10*365*24*60*60
	return timestamp > 0 && timestamp < m
}

func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000

	lat1Rad := lat1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	deltaLat := (lat2 - lat1) * math.Pi / 180
	deltaLon := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}
