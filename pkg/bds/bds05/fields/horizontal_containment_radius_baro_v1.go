package fields

import "fmt"

// HorizontalContainmentRadiusBarometricV1 is the Horizontal Containment Radius Barometric definition for ADBS V1
// messages when altitude type is barometric
//
// Specified in Doc 9871 / B.2.3.1
type HorizontalContainmentRadiusBarometricV1 byte

const (
	// HCRBV1RcLowerThan7Point5MAndVPLLowerThan11M indicates Rc < 7.5 m and VPL < 11 m
	HCRBV1RcLowerThan7Point5MAndVPLLowerThan11M HorizontalContainmentRadiusBarometricV1 = 0

	// HCRBV1RcLowerThan25MAndVPLLowerThan37Point5M indicates Rc < 25 m and VPL < 37.5 m
	HCRBV1RcLowerThan25MAndVPLLowerThan37Point5M HorizontalContainmentRadiusBarometricV1 = 1

	// HCRBV1RcLowerThan75MAndVPLLowerThan112M indicates Rc < 75 m and VPL < 112 m
	HCRBV1RcLowerThan75MAndVPLLowerThan112M HorizontalContainmentRadiusBarometricV1 = 2
	// HCRBV1RcLowerThan0Point1NM indicates Rc < 0.1 NM (185.2 m)
	HCRBV1RcLowerThan0Point1NM HorizontalContainmentRadiusBarometricV1 = 3

	// HCRBV1RcLowerThan0Point2NM indicates Rc < 0.2 NM (370.4 m)
	HCRBV1RcLowerThan0Point2NM HorizontalContainmentRadiusBarometricV1 = 4

	// HCRBV1RcLowerThan0Point6NM indicates Rc < 0.6 NM (1111.2 m)
	HCRBV1RcLowerThan0Point6NM HorizontalContainmentRadiusBarometricV1 = 5
	// HCRBV1RcLowerThan0Point5NM indicates Rc < 0.5 NM (926 m)
	HCRBV1RcLowerThan0Point5NM HorizontalContainmentRadiusBarometricV1 = 6

	// HCRBV1RcLowerThan1Point0NM indicates Rc < 1.0 NM (1852 m)
	HCRBV1RcLowerThan1Point0NM HorizontalContainmentRadiusBarometricV1 = 7

	// HCRBV1RcLowerThan2NM indicates Rc < 2 NM (3.704 km)
	HCRBV1RcLowerThan2NM HorizontalContainmentRadiusBarometricV1 = 8

	// HCRBV1RcLowerThan4NM indicates Rc < 4 NM (7.408 km)
	HCRBV1RcLowerThan4NM HorizontalContainmentRadiusBarometricV1 = 9
	// HCRBV1RcLowerThan8NM indicates Rc < 8 NM (14.816 km)
	HCRBV1RcLowerThan8NM HorizontalContainmentRadiusBarometricV1 = 10

	// HCRBV1RcLowerThan20NM indicates Rc < 20 NM (37.04 km)
	HCRBV1RcLowerThan20NM HorizontalContainmentRadiusBarometricV1 = 11

	// HCRBV1RcGreaterThan20NM indicates Rc >= 20 NM (37.04 km) or unknown
	HCRBV1RcGreaterThan20NM HorizontalContainmentRadiusBarometricV1 = 12
)

// ToHorizontalContainmentRadius returns the HorizontalContainmentRadius
func (hcr HorizontalContainmentRadiusBarometricV1) ToHorizontalContainmentRadius() HorizontalContainmentRadius {
	return hcr
}

// ToString returns a basic, but readable, representation of the field
func (hcr HorizontalContainmentRadiusBarometricV1) ToString() string {
	switch hcr {
	case HCRBV1RcLowerThan7Point5MAndVPLLowerThan11M:
		return "Rc < 7.5 m and VPL < 11 m"
	case HCRBV1RcLowerThan25MAndVPLLowerThan37Point5M:
		return "Rc < 25 m and VPL < 37.5 m"
	case HCRBV1RcLowerThan75MAndVPLLowerThan112M:
		return "Rc < 75 m and VPL < 112 m"
	case HCRBV1RcLowerThan0Point1NM:
		return "Rc < 0.1 NM (185.2 m)"
	case HCRBV1RcLowerThan0Point2NM:
		return "Rc < 0.2 NM (370.4 m)"
	case HCRBV1RcLowerThan0Point6NM:
		return "Rc < 0.6 NM (1111.2 m)"
	case HCRBV1RcLowerThan0Point5NM:
		return "Rc < 0.5 NM (926 m)"
	case HCRBV1RcLowerThan1Point0NM:
		return "Rc < 1.0 NM (1852 m)"
	case HCRBV1RcLowerThan2NM:
		return "Rc < 2 NM (3.704 km)"
	case HCRBV1RcLowerThan4NM:
		return "Rc < 4 NM (7.408 km)"
	case HCRBV1RcLowerThan8NM:
		return "Rc < 8 NM (14.816 km)"
	case HCRBV1RcLowerThan20NM:
		return "Rc < 20 NM (37.04 km)"
	case HCRBV1RcGreaterThan20NM:
		return "Rc >= 20 NM (37.04 km) or unknown"
	default:
		return fmt.Sprintf("%v - Unknown code", hcr)
	}
}
