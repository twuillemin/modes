package fields

import "fmt"

// HorizontalContainmentRadiusGNSSV2 is the Horizontal Containment Radius Barometric definition for ADBS V2
// messages when altitude type is barometric
//
// Specified in Doc 9871 / C.2.3.1
type HorizontalContainmentRadiusGNSSV2 byte

const (
	// HCRGV2RcLowerThan7Point5M indicates Rc < 7.5 m
	HCRGV2RcLowerThan7Point5M HorizontalContainmentRadiusGNSSV2 = 0

	// HCRGV2RcLowerThan25M indicates Rc < 25 m
	HCRGV2RcLowerThan25M HorizontalContainmentRadiusGNSSV2 = 1

	// HCRGV2RcGreaterThan25MOrUnknown indicates Rc >= 25 m or unknown
	HCRGV2RcGreaterThan25MOrUnknown HorizontalContainmentRadiusGNSSV2 = 2
)

// ToHorizontalContainmentRadius returns the HorizontalContainmentRadius
func (hcr HorizontalContainmentRadiusGNSSV2) ToHorizontalContainmentRadius() HorizontalContainmentRadius {
	return hcr
}

// ToString returns a basic, but readable, representation of the field
func (hcr HorizontalContainmentRadiusGNSSV2) ToString() string {
	switch hcr {
	case HCRGV2RcLowerThan7Point5M:
		return "Rc < 7.5 m"
	case HCRGV2RcLowerThan25M:
		return "Rc < 25 m"
	case HCRGV2RcGreaterThan25MOrUnknown:
		return "Rc >= 25 m or unknown"
	default:
		return fmt.Sprintf("%v - Unknown code", hcr)
	}
}
