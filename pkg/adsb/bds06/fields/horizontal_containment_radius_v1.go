package fields

import "fmt"

// HorizontalContainmentRadiusV1 is the Horizontal Containment Radius Barometric definition for ADSB V1 messages
//
// Specified in Doc 9871 / C.2.3.1
type HorizontalContainmentRadiusV1 byte

const (
	// HCRV1RcLowerThan7Point5M indicates Rc < 7.5 m
	HCRV1RcLowerThan7Point5M HorizontalContainmentRadiusV1 = 0

	// HCRV1RcLowerThan25M indicates Rc < 25 m
	HCRV1RcLowerThan25M HorizontalContainmentRadiusV1 = 1

	// HCRV1RcLowerThan75M indicates Rc < 75 m
	HCRV1RcLowerThan75M HorizontalContainmentRadiusV1 = 2

	// HCRV1RcLowerThan0Point1NM indicates Rc < 0.1 NM (185.2 m)
	HCRV1RcLowerThan0Point1NM HorizontalContainmentRadiusV1 = 3

	// HCRV1RcGreaterThan0Point1NM indicates Rc >= 0.1 NM (185.2 m) or unknown
	HCRV1RcGreaterThan0Point1NM HorizontalContainmentRadiusV1 = 4
)

// ToHorizontalContainmentRadius returns the HorizontalContainmentRadius
func (hcr HorizontalContainmentRadiusV1) ToHorizontalContainmentRadius() HorizontalContainmentRadius {
	return hcr
}

// ToString returns a basic, but readable, representation of the field
func (hcr HorizontalContainmentRadiusV1) ToString() string {
	switch hcr {
	case HCRV1RcLowerThan7Point5M:
		return "Rc < 7.5 m"
	case HCRV1RcLowerThan25M:
		return "Rc < 25 m"
	case HCRV1RcLowerThan75M:
		return "Rc < 75 m"
	case HCRV1RcLowerThan0Point1NM:
		return "Rc < 0.1 NM (185.2 m)"
	case HCRV1RcGreaterThan0Point1NM:
		return "Rc >= 0.1 NM (185.2 m) or unknown"
	default:
		return fmt.Sprintf("%v - Unknown code", hcr)
	}
}
