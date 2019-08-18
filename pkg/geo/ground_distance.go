package geo

import (
	"math"
)

// ComputeGroundDistance computes the distance in meters with the defined reference point
//
// Params:
//    - lat1: latitude of the first point
//    - long1: longitude of the first point
//    - lat2: latitude of the second point
//    - long2: longitude of the second point
//
// Returns the distance in meter
func ComputeGroundDistance(lat1, long1, lat2, long2 float64) int {

	difLatitude := (lat1 - lat2) * math.Pi / 180
	difLongitude := (long1 - long2) * math.Pi / 180

	a := (math.Sin(difLatitude/2) * math.Sin(difLatitude/2)) + (math.Cos(lat1) * math.Cos(lat2) * math.Sin(difLongitude/2) * math.Sin(difLongitude/2))

	angularDistance := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	metricDistance := 6371000 * angularDistance

	return int(math.Floor(metricDistance))
}
