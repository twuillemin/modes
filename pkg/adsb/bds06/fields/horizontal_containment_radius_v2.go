package fields

import "fmt"

// HorizontalContainmentRadiusV2 is the Horizontal Containment Radius Barometric definition for ADSB V2 messages
//
// Specified in Doc 9871 / C.2.3.1
type HorizontalContainmentRadiusV2 byte

const (
	// HCRV2RcLowerThan7Point5M indicates Rc < 7.5 m
	HCRV2RcLowerThan7Point5M HorizontalContainmentRadiusV2 = 0

	// HCRV2RcLowerThan25M indicates Rc < 25 m
	HCRV2RcLowerThan25M HorizontalContainmentRadiusV2 = 1

	// HCRV2RcLowerThan75M indicates Rc < 75 m
	HCRV2RcLowerThan75M HorizontalContainmentRadiusV2 = 2

	// HCRV2RcLowerThan0Point1NM indicates Rc < 0.1 NM (185.2 m)
	HCRV2RcLowerThan0Point1NM HorizontalContainmentRadiusV2 = 3

	// HCRV2RcLowerThan0Point2NM indicates Rc < 0.2 NM (370.4 m)
	HCRV2RcLowerThan0Point2NM HorizontalContainmentRadiusV2 = 4

	// HCRV2RcLowerThan0Point3NM indicates Rc < 0.3 NM (555.6 m)
	HCRV2RcLowerThan0Point3NM HorizontalContainmentRadiusV2 = 5

	// HCRV2RcLowerThan0Point6NM indicates Rc < 0.6 NM (1111.2 m)
	HCRV2RcLowerThan0Point6NM HorizontalContainmentRadiusV2 = 6

	// HCRV2RcGreaterThan0Point6NM indicates Rc >= 0.6 NM (1111.2 m) or unknown
	HCRV2RcGreaterThan0Point6NM HorizontalContainmentRadiusV2 = 7
)

// ToHorizontalContainmentRadius returns the HorizontalContainmentRadius
func (hcr HorizontalContainmentRadiusV2) ToHorizontalContainmentRadius() HorizontalContainmentRadius {
	return hcr
}

// ToString returns a basic, but readable, representation of the field
func (hcr HorizontalContainmentRadiusV2) ToString() string {
	switch hcr {
	case HCRV2RcLowerThan7Point5M:
		return "Rc < 7.5 m"
	case HCRV2RcLowerThan25M:
		return "Rc < 25 m"
	case HCRV2RcLowerThan75M:
		return "Rc < 75 m"
	case HCRV2RcLowerThan0Point1NM:
		return "Rc < 0.1 NM (185.2 m)"
	case HCRV2RcLowerThan0Point2NM:
		return "Rc < 0.2 NM (370.4 m)"
	case HCRV2RcLowerThan0Point3NM:
		return "Rc < 0.3 NM (555.6 m)"
	case HCRV2RcLowerThan0Point6NM:
		return "Rc < 0.6 NM (1111.2 m)"
	case HCRV2RcGreaterThan0Point6NM:
		return "Rc >= 0.6 NM (1111.2 m) or unknown"
	default:
		return fmt.Sprintf("%v - Unknown code", hcr)
	}
}
