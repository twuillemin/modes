package fields

import "fmt"

// HorizontalContainmentRadiusBarometricV2 is the Horizontal Containment Radius Barometric definition for ADSB V2
// messages when altitude type is barometric
//
// Specified in Doc 9871 / C.2.3.1
type HorizontalContainmentRadiusBarometricV2 byte

const (
	// HCRBV2RcLowerThan7Point5M indicates Rc < 7.5 m
	HCRBV2RcLowerThan7Point5M HorizontalContainmentRadiusBarometricV2 = 0

	// HCRBV2RcLowerThan25M indicates Rc < 25 m
	HCRBV2RcLowerThan25M HorizontalContainmentRadiusBarometricV2 = 1

	// HCRBV2RcLowerThan75M indicates Rc < 75 m
	HCRBV2RcLowerThan75M HorizontalContainmentRadiusBarometricV2 = 2
	// HCRBV2RcLowerThan0Point1NM indicates Rc < 0.1 NM (185.2 m)
	HCRBV2RcLowerThan0Point1NM HorizontalContainmentRadiusBarometricV2 = 3

	// HCRBV2RcLowerThan0Point2NM indicates Rc < 0.2 NM (370.4 m)
	HCRBV2RcLowerThan0Point2NM HorizontalContainmentRadiusBarometricV2 = 4

	// HCRBV2RcLowerThan0Point3NM indicates Rc < 0.3 NM (555.6 m)
	HCRBV2RcLowerThan0Point3NM HorizontalContainmentRadiusBarometricV2 = 5
	// HCRBV2RcLowerThan0Point5NM indicates Rc < 0.5 NM (926 m)
	HCRBV2RcLowerThan0Point5NM HorizontalContainmentRadiusBarometricV2 = 6
	// HCRBV2RcLowerThan0Point6NM indicates Rc < 0.6 NM (1111.2 m)
	HCRBV2RcLowerThan0Point6NM HorizontalContainmentRadiusBarometricV2 = 7

	// HCRBV2RcLowerThan1Point0NM indicates Rc < 1.0 NM (1852 m)
	HCRBV2RcLowerThan1Point0NM HorizontalContainmentRadiusBarometricV2 = 8

	// HCRBV2RcLowerThan2NM indicates Rc < 2 NM (3.704 km)
	HCRBV2RcLowerThan2NM HorizontalContainmentRadiusBarometricV2 = 9

	// HCRBV2RcLowerThan4NM indicates Rc < 4 NM (7.408 km)
	HCRBV2RcLowerThan4NM HorizontalContainmentRadiusBarometricV2 = 10
	// HCRBV2RcLowerThan8NM indicates Rc < 8 NM (14.816 km)
	HCRBV2RcLowerThan8NM HorizontalContainmentRadiusBarometricV2 = 11

	// HCRBV2RcLowerThan20NM indicates Rc < 20 NM (37.04 km)
	HCRBV2RcLowerThan20NM HorizontalContainmentRadiusBarometricV2 = 12

	// HCRBV2RcGreaterThan20NM indicates Rc >= 20 NM (37.04 km) or unknown
	HCRBV2RcGreaterThan20NM HorizontalContainmentRadiusBarometricV2 = 13
)

// ToHorizontalContainmentRadius returns the HorizontalContainmentRadius
func (hcr HorizontalContainmentRadiusBarometricV2) ToHorizontalContainmentRadius() HorizontalContainmentRadius {
	return hcr
}

// ToString returns a basic, but readable, representation of the field
func (hcr HorizontalContainmentRadiusBarometricV2) ToString() string {
	switch hcr {
	case HCRBV2RcLowerThan7Point5M:
		return "Rc < 7.5 m"
	case HCRBV2RcLowerThan25M:
		return "Rc < 25 m"
	case HCRBV2RcLowerThan75M:
		return "Rc < 75 m"
	case HCRBV2RcLowerThan0Point1NM:
		return "Rc < 0.1 NM (185.2 m)"
	case HCRBV2RcLowerThan0Point2NM:
		return "Rc < 0.2 NM (370.4 m)"
	case HCRBV2RcLowerThan0Point3NM:
		return "Rc < 0.3 NM (555.6 m)"
	case HCRBV2RcLowerThan0Point5NM:
		return "Rc < 0.5 NM (926 m)"
	case HCRBV2RcLowerThan0Point6NM:
		return "Rc < 0.6 NM (1111.2 m)"
	case HCRBV2RcLowerThan1Point0NM:
		return "Rc < 1.0 NM (1852 m)"
	case HCRBV2RcLowerThan2NM:
		return "Rc < 2 NM (3.704 km)"
	case HCRBV2RcLowerThan4NM:
		return "Rc < 4 NM (7.408 km)"
	case HCRBV2RcLowerThan8NM:
		return "Rc < 8 NM (14.816 km)"
	case HCRBV2RcLowerThan20NM:
		return "Rc < 20 NM (37.04 km)"
	case HCRBV2RcGreaterThan20NM:
		return "Rc >= 20 NM (37.04 km) or unknown"
	default:
		return fmt.Sprintf("%v - Unknown code", hcr)
	}
}
