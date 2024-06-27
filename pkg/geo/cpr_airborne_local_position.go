package geo

import "math"

// GetCPRAirborneLocalPosition returns the position (based on a single measure and a reference point) of a plane position
//
// Params:
//   - cprLatitude: The latitude (CPR encoded) coming from the message
//   - cprLongitude: The longitude (CPR encoded) coming from the message
//   - referenceLatitude: The latitude coming from the reference point
//   - referenceLongitude: The longitude coming from the reference point
//   - isOdd: indicates that the message is an odd message.
//
// Return the latitude and longitude (in this order).
func GetCPRAirborneLocalPosition(
	cprLatitude uint32,
	cprLongitude uint32,
	referenceLatitude float64,
	referenceLongitude float64,
	isOdd bool) (float64, float64) {

	i := 0.0
	if isOdd {
		i = 1.0
	}

	latCPR := float64(cprLatitude) / maxBinaryValue
	lonCPR := float64(cprLongitude) / maxBinaryValue

	// The latitude zone size is different depending on the message type
	dLat := 360.0 / (4.0*Nz - i)

	// Then, the latitude zone index, j, is calculated as
	j := math.Floor(referenceLatitude/dLat) + math.Floor(mod(referenceLatitude, dLat)/dLat-latCPR+0.5)

	// Knowing the latitude zone index, the latitude of the new position is:
	latitude := dLat * (j + latCPR)

	// Next, we can calculate the increment of the longitude per zone based on the decoded latitude, which is
	// dependent on both message type and latitude
	dLon := 360.0 / (math.Max(getNumberOfLongitude(latitude)-i, 1.0))

	// Then, the longitude zone index, m, is calculated as
	m := math.Floor(referenceLongitude/dLon) + math.Floor(mod(referenceLongitude, dLon)/dLon-lonCPR+0.5)

	// Knowing the longitude zone index, the longitude of the new position is
	longitude := dLon * (m + lonCPR)

	return latitude, longitude
}
