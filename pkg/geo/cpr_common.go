package geo

import "math"

const Nz = 15.0

// CPR coordinates are on 17 bits, so max value 2^17 => 131072.0
const maxBinaryValue = 131072.0

// mod is the modulo function to be used in CPR computation
// Please note, that this function is not the same as the standard Go function.
func mod(x float64, y float64) float64 {
	return x - y*math.Floor(x/y)
}

// getNumberOfLongitude computes the number of longitude zones at a latitude
//
// Params:
//   - lat: The latitude
//
// Return the number of longitude zones at a latitude
func getNumberOfLongitude(lat float64) float64 {

	// Deal with extreme coordinates
	if lat == 0 {
		return 59
	} else if math.Abs(lat) == 87 {
		return 2
	} else if math.Abs(lat) > 87 {
		return 1
	}

	a := 1 - math.Cos(math.Pi/(2*Nz))
	b := math.Pow(math.Cos(math.Pi/180.0*math.Abs(lat)), 2.0)
	nl := 2 * math.Pi / (math.Acos(1 - a/b))

	return math.Floor(nl)
}
