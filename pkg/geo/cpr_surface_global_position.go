package geo

import (
	"errors"
	"math"
)

// GetCPRSurfaceGlobalPosition returns the exact position (based on two measure) of a plane position
//
// Params:
//   - evenCPRLatitude: The latitude (CPR encoded) coming from the even message
//   - evenCPRLongitude: The longitude (CPR encoded) coming from the even message
//   - oddCPRLatitude: The latitude (CPR encoded) coming from the even message
//   - oddCPRLongitude: The longitude (CPR encoded) coming from the even message
//   - evenDataReceivedAfterOddData: indicates that the even data was received after the odd data. This can boil down
//     to something like even.Timestamp > odd.Timestamp.
//   - referenceLatitude: The latitude coming from the reference point
//   - referenceLongitude: The longitude coming from the reference point
//
// Return the latitude and longitude (in this order) or an error.
func GetCPRSurfaceGlobalPosition(
	evenCPRLatitude uint32,
	evenCPRLongitude uint32,
	oddCPRLatitude uint32,
	oddCPRLongitude uint32,
	evenDataReceivedAfterOddData bool,
	referenceLatitude float64,
	referenceLongitude float64,
) (float64, float64, error) {

	// it is possible to extract the encoded CPR latitude and quadrantLongitude binary and convert them to decimal format.
	// Then, they are divided by the 2^17, representing the fractions of the positions within the latitude and
	// quadrantLongitude grids
	cprLatitudeEven := float64(evenCPRLatitude) / maxBinaryValue
	cprLongitudeEven := float64(evenCPRLongitude) / maxBinaryValue
	cprLatitudeOdd := float64(oddCPRLatitude) / maxBinaryValue
	cprLongitudeOdd := float64(oddCPRLongitude) / maxBinaryValue

	// We can calculate the latitude index
	latitudeIndex := math.Floor((4*Nz-1)*cprLatitudeEven - (4*Nz)*cprLatitudeOdd + 0.5)

	// Then, we can decode the latitudes from both even and odd messages
	latitudeEven := 90 / (4 * Nz) * (mod(latitudeIndex, 4*Nz) + cprLatitudeEven)
	latitudeOdd := 90 / (4*Nz - 1) * (mod(latitudeIndex, 4*Nz-1) + cprLatitudeOdd)

	// If the reference in the southern hemisphere latitude solution
	latitudeEvenFixed := latitudeEven
	if referenceLatitude < 0 {
		latitudeEvenFixed = latitudeEvenFixed - 90
	}

	latitudeOddFixed := latitudeOdd
	if referenceLatitude < 0 {
		latitudeOddFixed = latitudeOddFixed - 90
	}

	// Before proceeding to the quadrantLongitude calculation, we need to compute NL(lat_even) and NL(lat_odd) to check if both
	//values are the same. If not, this means the pair of messages are from different quadrantLongitude zones,
	//and it is not possible to compute the correct global position.
	numberLongitudeEven := getNumberOfLongitude(latitudeEvenFixed)
	numberLongitudeOdd := getNumberOfLongitude(latitudeOddFixed)

	// In this case, decoding should be stopped, and it is necessary to wait for a pair of messages that are from the
	//same latitude zone. This situation happens when aircraft are flying across the boundaries of quadrantLongitude zones.
	if numberLongitudeEven != numberLongitudeOdd {
		return 0, 0, errors.New("both latitude are not in the same quadrantLongitude zone")
	}
	numberLongitude := numberLongitudeEven

	// We can continue to calculate the global position. Since the even message is the most recent message, the
	// northern hemisphere latitude solution
	latitude := 0.0
	if evenDataReceivedAfterOddData {
		latitude = latitudeEvenFixed
	} else {
		latitude = latitudeOddFixed
	}

	// The quadrantLongitude solution in the first quadrant of 0 to 90 degrees is calculated based on the most
	// recent message
	m := math.Floor(cprLongitudeEven*(numberLongitude-1) - cprLongitudeOdd*numberLongitude + 0.5)
	nEven := math.Max(numberLongitude-0, 1)
	nOdd := math.Max(numberLongitude-1, 1)

	// Then, the quadrantLongitude is calculated
	lonEven := 90 / nEven * (mod(m, nEven) + cprLongitudeEven)
	lonOdd := 90 / nOdd * (mod(m, nOdd) + cprLongitudeOdd)

	quadrantLongitude := 0.0
	if evenDataReceivedAfterOddData {
		quadrantLongitude = lonEven
	} else {
		quadrantLongitude = lonOdd
	}

	// All the solutions are The three other possible solutions are
	longitudes := [4]float64{
		quadrantLongitude,
		quadrantLongitude + 90,
		quadrantLongitude + 180,
		quadrantLongitude + 270,
	}

	longitudesCorrected := [4]float64{}
	for i, longitude := range longitudes {
		longitudesCorrected[i] = mod((longitude+180), 360) - 180
	}

	distancesToRef := [4]float64{}
	for i, longitudeCorrected := range longitudesCorrected {
		distancesToRef[i] = math.Abs(referenceLongitude - longitudeCorrected)
	}

	smallestDistance := 10000000.0
	minIndex := -1
	for i, distance := range distancesToRef {
		if distance < smallestDistance {
			smallestDistance = distance
			minIndex = i
		}
	}

	longitude := longitudesCorrected[minIndex]

	return latitude, longitude, nil
}
