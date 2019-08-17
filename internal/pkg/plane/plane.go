package plane

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/modes/common"
	"math"
)

// Plane is the structure keeping track of the current status of a plane
type Plane struct {
	ICAOAddress        common.ICAOAddress
	ADSBLevel          adsb.ReaderLevel
	Altitude           int
	Identification     string
	FirstSeenTimestamp uint32
	LastSeenTimestamp  uint32
	EvenCPRLatitude    uint32
	EvenCPRLongitude   uint32
	EventCPRTimestamp  uint32
	OddCPRLatitude     uint32
	OddCPRLongitude    uint32
	OddCPRTimestamp    uint32
	AirSpeed           int
	AirSpeedValid      bool
	VerticalRate       int
	VerticalRateValid  bool
	NICSupplementA     bool
	NICSupplementC     bool
}

// ToString returns a very simple representation of the plane
func (plane *Plane) ToString() string {

	lat, long, err := plane.GetExactPosition()
	coord := ""
	if err == nil {

		distance := "N/A"
		if dist, errDist := plane.computeGroundDistanceDistanceWithReference(); errDist == nil {
			distance = fmt.Sprintf("%v", dist)
		}

		coord = fmt.Sprintf("lat: %8.5f, long: %8.5f, ground distance: %v km", lat, long, distance)
	} else {
		coord = err.Error()
	}

	altitude := "N/A"
	if plane.Altitude > 0 {
		altitude = fmt.Sprintf("%v", plane.Altitude)
	}

	airSpeed := "N/A"
	if plane.AirSpeedValid {
		airSpeed = fmt.Sprintf("%v", plane.AirSpeed)
	}

	verticalRate := "N/A"
	if plane.VerticalRateValid {
		verticalRate = fmt.Sprintf("%v", plane.VerticalRate)
	}

	return fmt.Sprintf("Plane: %v, Flight: %v, ADSB: %v, Position: %v, Altitude: %v feet, AirSpeed: %v knot, VerticalSpeed: %v ft/min",
		plane.ICAOAddress.ToString(),
		plane.Identification,
		plane.ADSBLevel.ToString(),
		coord,
		altitude,
		airSpeed,
		verticalRate)
}

// GetExactPosition returns the exact position (based on two measure) of the plane position
func (plane *Plane) GetExactPosition() (float64, float64, error) {

	if plane.EventCPRTimestamp == 0 || plane.OddCPRTimestamp == 0 {
		return 0, 0, errors.New("missing data to initialize")
	}

	cprLatitudeEven := float64(plane.EvenCPRLatitude) / 131072.0
	cprLongitudeEven := float64(plane.EvenCPRLongitude) / 131072.0
	cprLatitudeOdd := float64(plane.OddCPRLatitude) / 131072.0
	cprLongitudeOdd := float64(plane.OddCPRLongitude) / 131072.0

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

	if plane.EventCPRTimestamp > plane.OddCPRTimestamp {
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

/**
 * @param Rlat Even or odd Rlat value (CPR internal)
 * @return the number of even longitude zones at a latitude
 */
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

var referenceLatitude float64
var referenceLongitude float64

// SetReferenceLatitudeLongitude defines the position that can be used to determine the distance
//
// params:
//    - latitude: the reference latitude
//    - longitude: the reference longitude
func SetReferenceLatitudeLongitude(latitude float64, longitude float64) {
	referenceLatitude = latitude
	referenceLongitude = longitude
}

// computeGroundDistanceDistanceWithReference computes the distance in km with the defined reference point
func (plane *Plane) computeGroundDistanceDistanceWithReference() (int, error) {

	if referenceLatitude == 0 && referenceLongitude == 0 {
		return 0.0, errors.New("no reference defined for computing ground distance")
	}

	planeLatitude, planeLongitude, err := plane.GetExactPosition()
	if err != nil {
		return 0, err
	}

	lat1 := referenceLatitude * math.Pi / 180
	lat2 := planeLatitude * math.Pi / 180

	difLatitude := (planeLatitude - referenceLatitude) * math.Pi / 180
	difLongitude := (planeLongitude - referenceLongitude) * math.Pi / 180

	a := (math.Sin(difLatitude/2) * math.Sin(difLatitude/2)) + (math.Cos(lat1) * math.Cos(lat2) * math.Sin(difLongitude/2) * math.Sin(difLongitude/2))

	angularDistance := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	metricDistance := 6371000 * angularDistance

	return int(math.Floor(metricDistance / 1000)), nil
}
