package geo

import (
	"math"
)

// ComputeGroundDistance computes the metric distance in meters with the defined reference point
//
// Params:
//   - lat1: latitude of the first point
//   - long1: longitude of the first point
//   - lat2: latitude of the second point
//   - long2: longitude of the second point
//
// Returns the distance in meter
func ComputeGroundDistance(lat1, long1, lat2, long2 float64) float64 {

	earthRadius := 6371000.0
	phi1 := lat1 * math.Pi / 180.0
	phi2 := lat2 * math.Pi / 180.0

	deltaPhi := (lat2 - lat1) * math.Pi / 180.0
	deltaLambda := (long2 - long1) * math.Pi / 180.0

	a := math.Pow(math.Sin(deltaPhi/2.0), 2) + math.Cos(phi1)*math.Cos(phi2)*math.Pow(math.Sin(deltaLambda/2.0), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	metricDistance := earthRadius * c

	return metricDistance
}
