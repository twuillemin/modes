package fields

import "fmt"

// HorizontalContainmentRadiusGNSSV1 is the Horizontal Containment Radius Barometric definition for ADBS V1
// messages when altitude type is barometric
//
// Specified in Doc 9871 / B.2.3.1
type HorizontalContainmentRadiusGNSSV1 byte

const (
	// HCRGV1RcLowerThan7Point5MAndVPLLowerThan11M indicates Rc < 7.5 m and VPL < 11 m
	HCRGV1RcLowerThan7Point5MAndVPLLowerThan11M HorizontalContainmentRadiusGNSSV1 = 0

	// HCRGV1RcLowerThan25MAndVPLLowerThan37Point5M indicates Rc < 25 m and VPL < 37.5 m
	HCRGV1RcLowerThan25MAndVPLLowerThan37Point5M HorizontalContainmentRadiusGNSSV1 = 1

	// HCRGV1RcGreaterThan25MOrVPLGreaterThan37Point5MOrUnknown indicates Rc >= 25 m or Pc >= 37.5 m or unknown
	HCRGV1RcGreaterThan25MOrVPLGreaterThan37Point5MOrUnknown HorizontalContainmentRadiusGNSSV1 = 2
)

// ToHorizontalContainmentRadius returns the HorizontalContainmentRadius
func (hcr HorizontalContainmentRadiusGNSSV1) ToHorizontalContainmentRadius() HorizontalContainmentRadius {
	return hcr
}

// ToString returns a basic, but readable, representation of the field
func (hcr HorizontalContainmentRadiusGNSSV1) ToString() string {
	switch hcr {
	case HCRGV1RcLowerThan7Point5MAndVPLLowerThan11M:
		return "Rc < 7.5 m and VPL < 11 m"
	case HCRGV1RcLowerThan25MAndVPLLowerThan37Point5M:
		return "Rc < 25 m and VPL < 37.5 m"
	case HCRGV1RcGreaterThan25MOrVPLGreaterThan37Point5MOrUnknown:
		return "Rc >= 25 m or VPL >= 37.5 m or unknown"
	default:
		return fmt.Sprintf("%v - Unknown code", hcr)
	}
}
