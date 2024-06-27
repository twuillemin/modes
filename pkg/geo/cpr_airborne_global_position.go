package geo

import (
	"errors"
	"math"
)

// GetCPRAirborneGlobalPosition returns the exact position (based on two measure) of a plane position
//
// Params:
//   - evenCPRLatitude: The latitude (CPR encoded) coming from the even message
//   - evenCPRLongitude: The longitude (CPR encoded) coming from the even message
//   - oddCPRLatitude: The latitude (CPR encoded) coming from the even message
//   - oddCPRLongitude: The longitude (CPR encoded) coming from the even message
//   - evenDataReceivedAfterOddData: indicates that the even data was received after the odd data. This can boil down
//     to something like even.Timestamp > odd.Timestamp.
//
// Return the latitude and longitude (in this order) or an error.
func GetCPRAirborneGlobalPosition(
	evenCPRLatitude uint32,
	evenCPRLongitude uint32,
	oddCPRLatitude uint32,
	oddCPRLongitude uint32,
	evenDataReceivedAfterOddData bool) (float64, float64, error) {

	cprLatitudeEven := float64(evenCPRLatitude) / maxBinaryValue
	cprLongitudeEven := float64(evenCPRLongitude) / maxBinaryValue
	cprLatitudeOdd := float64(oddCPRLatitude) / maxBinaryValue
	cprLongitudeOdd := float64(oddCPRLongitude) / maxBinaryValue

	// Compute the latitude zone sizes
	dLatitudeEven := 360.0 / (4 * Nz)
	dLatitudeOdd := 360.0 / (4*Nz - 1)

	// Compute latitude index
	j := math.Floor((4*Nz-1)*cprLatitudeEven - (4*Nz)*cprLatitudeOdd + 0.5)

	// Compute the two latitudes from odd and even frames
	latitudeEven := dLatitudeEven * (mod(j, 4*Nz) + cprLatitudeEven)
	latitudeOdd := dLatitudeOdd * (mod(j, 4*Nz-1) + cprLatitudeOdd)

	// For southern hemisphere values returned from previous equations range from 270 to 360 degrees.
	//Hence, we need to make sure the latitude is within the range of [-90,+90]
	if latitudeEven >= 270 {
		latitudeEven -= 360
	}

	if latitudeOdd >= 270 {
		latitudeOdd -= 360
	}

	// Before proceeding to the longitude calculation, we need to compute NL(lat_even) and NL(lat_odd) to check if both
	//values are the same. If not, this means the pair of messages are from different longitude zones,
	//and it is not possible to compute the correct global position.
	numberLongitudeEven := getNumberOfLongitude(latitudeEven)
	numberLongitudeOdd := getNumberOfLongitude(latitudeOdd)

	// In this case, decoding should be stopped, and it is necessary to wait for a pair of messages that are from the
	//same latitude zone. This situation happens when aircraft are flying across the boundaries of longitude zones.
	if numberLongitudeEven != numberLongitudeOdd {
		return 0, 0, errors.New("both latitude are not in the same latitude zone")
	}
	numberLongitude := numberLongitudeEven

	latitude := 0.0
	if evenDataReceivedAfterOddData {
		latitude = latitudeEven
	} else {
		latitude = latitudeOdd
	}

	// m is the longitude index
	m := math.Floor(cprLongitudeEven*(numberLongitude-1) - cprLongitudeOdd*numberLongitude + 0.5)

	// Calculate the longitude zone size, which is dependent on the latitude. For even and odd messages, the number of
	// longitude zones
	nEven := math.Max(numberLongitude-0, 1)
	nOdd := math.Max(numberLongitude-1, 1)

	// The longitude zone sizes are defined as follows
	dLonEven := 360.0 / nEven
	dLonOdd := 360.0 / nOdd

	// Then, the longitude is calculated as:
	lonEven := dLonEven * (mod(m, nEven) + cprLongitudeEven)
	lonOdd := dLonOdd * (mod(m, nOdd) + cprLongitudeOdd)

	longitude := 0.0
	if evenDataReceivedAfterOddData {
		longitude = lonEven
	} else {
		longitude = lonOdd
	}

	// It is worth noting that the longitudes in position messages are between 0 and 360 degrees. We often need to
	//convert them to the range between -180 and 180 degrees, which is consistent with aviation conventions.

	if longitude > 180 {
		longitude -= 360
	}

	return latitude, longitude, nil
}
