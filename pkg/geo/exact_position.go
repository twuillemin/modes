package geo

import (
	"errors"
	"math"
)

// GetCPRExactPosition returns the exact position (based on two measure) of a plane position
//
// Params:
//    - evenCPRLatitude: The latitude (CPR encoded) coming from the even message
//    - evenCPRLongitude: The longitude (CPR encoded) coming from the even message
//    - oddCPRLatitude: The latitude (CPR encoded) coming from the even message
//    - oddCPRLongitude: The longitude (CPR encoded) coming from the even message
//    - evenDataReceivedAfterOddData: indicates that the even data was received after the odd data. This can boil down
//        to something like even.Timestamp > odd.Timestamp.
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

	dLatitudeEven := 360.0 / 60.0
	dLatitudeOdd := 360.0 / 59.0

	// Compute latitude index
	j := math.Floor(59*cprLatitudeEven - 60*cprLatitudeOdd + 0.5)

	latitudeEven := dLatitudeEven * (math.Mod(j, 60.0) + cprLatitudeEven)
	latitudeOdd := dLatitudeOdd * (math.Mod(j, 59.0) + cprLatitudeOdd)

	if latitudeEven >= 270 {
		latitudeEven -= 360
	}

	if latitudeOdd >= 270 {
		latitudeOdd -= 360
	}

	nlLatitudeEven := getNumberOfEvenLongitude(latitudeEven)
	nlLatitudeOdd := getNumberOfEvenLongitude(latitudeOdd)

	if nlLatitudeEven != nlLatitudeOdd {
		return 0, 0, errors.New("both latitude are not in the same latitude zone")
	}

	lon := 0.0
	lat := 0.0

	if evenDataReceivedAfterOddData {
		ni := nlLatitudeEven
		if ni < 1 {
			ni = 1
		}

		m := math.Floor(cprLongitudeEven*(nlLatitudeEven-1) - cprLongitudeOdd*nlLatitudeEven + 0.5)

		lon = (360.0 / ni) * (math.Mod(m, ni) + cprLongitudeEven)
		lat = latitudeEven
	} else {
		ni := nlLatitudeEven - 1
		if ni < 1 {
			ni = 1
		}

		m := math.Floor(cprLongitudeEven*(nlLatitudeOdd-1) - cprLongitudeOdd*nlLatitudeOdd + 0.5)

		lon = (360.0 / ni) * (math.Mod(m, ni) + cprLongitudeOdd)
		lat = latitudeOdd
	}

	if lon > 180 {
		lon -= 360
	}

	return lat, lon, nil
}

// getNumberOfEvenLongitude computes the number of even longitude zones at a latitude
//
// Params:
//    - lat: The latitude
//
// Return the number of even longitude zones at a latitude
func getNumberOfEvenLongitude(lat float64) float64 {

	// Deal with extreme coordinates
	if lat == 0 {
		return 59
	} else if math.Abs(lat) == 87 {
		return 2
	} else if math.Abs(lat) > 87 {
		return 1
	}

	nz := 15.0
	a := 1 - math.Cos(math.Pi/(2*nz))
	b := math.Pow(math.Cos(math.Pi/180.0*math.Abs(lat)), 2.0)
	nl := 2 * math.Pi / (math.Acos(1 - a/b))

	return math.Floor(nl)
}
