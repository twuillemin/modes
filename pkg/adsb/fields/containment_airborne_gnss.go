package fields

import "fmt"

// ContainmentRadiusAirborneGNSS is the 95% Containment radius,μ and v, on horizontal and vertical position error for
// Airborne aircraft format definition
//
// Specified in Doc 9871 / A.2.3.1
type ContainmentRadiusAirborneGNSS byte

const (
	// CRAGHorizontalLowerThan3MAndVerticalLowerThan4M denotes μ < 3 m and v < 4 m
	CRAGHorizontalLowerThan3MAndVerticalLowerThan4M ContainmentRadiusAirborneGNSS = 0
	// CRAGHorizontalLowerThan10MAndVerticalLowerThan15M denotes μ < 10 m and v < 15 m
	CRAGHorizontalLowerThan10MAndVerticalLowerThan15M ContainmentRadiusAirborneGNSS = 1
	// CRAGHorizontalGreaterThan10MOrVerticalGreaterThan15M denotes μ > 10 m or v ≥ 15 m
	CRAGHorizontalGreaterThan10MOrVerticalGreaterThan15M ContainmentRadiusAirborneGNSS = 2
)

// GetHorizontalContainmentLowerBound returns the lower bound of the CR. If the given CR does not have a lower bound returns -1.
// Notes:
//    - The bound is returned in meters.
//    - The bound is inclusive to the value, meaning that bound <= value.
func (cr ContainmentRadiusAirborneGNSS) GetHorizontalContainmentLowerBound() float64 {
	switch cr {
	case CRAGHorizontalGreaterThan10MOrVerticalGreaterThan15M:
		return 10
	default:
		return -1
	}
}

// GetHorizontalContainmentHigherBound returns the higher bound of the CR. If the given CR does not have a higher bound returns -1.
// Notes:
//    - The bound is returned in meters.
//    - The bound is exclusive to the value, meaning that value < bound.
func (cr ContainmentRadiusAirborneGNSS) GetHorizontalContainmentHigherBound() float64 {
	switch cr {
	case CRAGHorizontalLowerThan3MAndVerticalLowerThan4M:
		return 3
	case CRAGHorizontalLowerThan10MAndVerticalLowerThan15M:
		return 10
	default:
		return -1
	}
}

// ToString returns a basic, but readable, representation of the field
func (cr ContainmentRadiusAirborneGNSS) ToString() string {
	switch cr {
	case CRAGHorizontalLowerThan3MAndVerticalLowerThan4M:
		return "μ < 3 m and v < 4 m"
	case CRAGHorizontalLowerThan10MAndVerticalLowerThan15M:
		return "μ < 10 m and v < 15 m"
	case CRAGHorizontalGreaterThan10MOrVerticalGreaterThan15M:
		return "μ > 10 m or v ≥ 15 m"
	default:
		return fmt.Sprintf("%v - Unknown code", cr)
	}
}
