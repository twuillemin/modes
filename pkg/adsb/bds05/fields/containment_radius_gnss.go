package fields

import "fmt"

// ContainmentRadiusGNSS is the 95% Containment radius,μ and v, on horizontal and vertical position error for
// Airborne aircraft format definition in ADSB V0
//
// Specified in Doc 9871 / A.2.3.1
type ContainmentRadiusGNSS byte

const (
	// CRGHorizontalLowerThan3MAndVerticalLowerThan4M denotes μ < 3 m and v < 4 m
	CRGHorizontalLowerThan3MAndVerticalLowerThan4M ContainmentRadiusGNSS = 0
	// CRGHorizontalLowerThan10MAndVerticalLowerThan15M denotes μ < 10 m and v < 15 m
	CRGHorizontalLowerThan10MAndVerticalLowerThan15M ContainmentRadiusGNSS = 1
	// CRGHorizontalGreaterThan10MOrVerticalGreaterThan15M denotes μ > 10 m or v >= 15 m
	CRGHorizontalGreaterThan10MOrVerticalGreaterThan15M ContainmentRadiusGNSS = 2
)

// ToContainmentRadius returns the ContainmentRadius
func (cr ContainmentRadiusGNSS) ToContainmentRadius() ContainmentRadius {
	return cr
}

// ToString returns a basic, but readable, representation of the field
func (cr ContainmentRadiusGNSS) ToString() string {
	switch cr {
	case CRGHorizontalLowerThan3MAndVerticalLowerThan4M:
		return "μ < 3 m and v < 4 m"
	case CRGHorizontalLowerThan10MAndVerticalLowerThan15M:
		return "μ < 10 m and v < 15 m"
	case CRGHorizontalGreaterThan10MOrVerticalGreaterThan15M:
		return "μ > 10 m or v >= 15 m"
	default:
		return fmt.Sprintf("%v - Unknown code", cr)
	}
}
