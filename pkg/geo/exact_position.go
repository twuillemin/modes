package geo

import (
	"errors"
	"math"
)

const Nz = 15.0

// GetCPRExactPosition returns the exact position (based on two measure) of a plane position
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
func GetCPRExactPosition(
	evenCPRLatitude uint32,
	evenCPRLongitude uint32,
	oddCPRLatitude uint32,
	oddCPRLongitude uint32,
	evenDataReceivedAfterOddData bool) (float64, float64, error) {

	cprLatitudeEven := float64(evenCPRLatitude) / 131072.0
	cprLongitudeEven := float64(evenCPRLongitude) / 131072.0
	cprLatitudeOdd := float64(oddCPRLatitude) / 131072.0
	cprLongitudeOdd := float64(oddCPRLongitude) / 131072.0

	// Compute the latitude zone sizes
	dLatitudeEven := 360.0 / (4 * Nz)
	dLatitudeOdd := 360.0 / (4*Nz - 1)

	// Compute latitude index
	j := math.Floor(59*cprLatitudeEven - 60*cprLatitudeOdd + 0.5)

	// Compute the two latitudes from odd and even frames
	latitudeEven := dLatitudeEven * (mod(j, 60.0) + cprLatitudeEven)
	latitudeOdd := dLatitudeOdd * (mod(j, 59.0) + cprLatitudeOdd)

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
	nlLatitudeEven := getNumberOfLongitude(latitudeEven)
	nlLatitudeOdd := getNumberOfLongitude(latitudeOdd)

	// In this case, decoding should be stopped, and it is necessary to wait for a pair of messages that are from the
	//same latitude zone. This situation happens when aircraft are flying across the boundaries of longitude zones.
	if nlLatitudeEven != nlLatitudeOdd {
		return 0, 0, errors.New("both latitude are not in the same latitude zone")
	}

	latitude := 0.0
	if evenDataReceivedAfterOddData {
		latitude = latitudeEven
	} else {
		latitude = latitudeOdd
	}

	nlLatitude := getNumberOfLongitude(latitude)

	// m is the longitude index
	m := math.Floor(cprLongitudeEven*(nlLatitude-1) - cprLongitudeOdd*nlLatitude + 0.5)

	// Calculate the longitude zone size, which is dependent on the latitude. For even and odd messages, the number of
	// longitude zones
	nEven := math.Max(nlLatitude, 1)
	nOdd := math.Max(nlLatitude-1, 1)

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
