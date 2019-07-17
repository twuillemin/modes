package fields

import "fmt"

// ContainmentRadiusAirborneBarometric is the 95% Containment radius,μ and v, on horizontal and vertical position error for
// Airborne aircraft format definition
//
// Specified in Doc 9871 / A.2.3.1
type ContainmentRadiusSurface byte

const (
	// CRSLowerThan3M denotes μ < 3m
	CRSLowerThan3M ContainmentRadiusSurface = 0
	// CRSBetween3MAnd10M denotes 3 m ≤ μ < 10 m
	CRSBetween3MAnd10M ContainmentRadiusSurface = 1
	// CRSBetween10MAnd92Point6M denotes 10 m ≤ μ < 92.6 m (0.05 NM)
	CRSBetween10MAnd92Point6M ContainmentRadiusSurface = 2
	// CRSGreaterThan92Point6M denotes 92.6 m (0.05 NM) ≤ μ
	CRSGreaterThan92Point6M ContainmentRadiusSurface = 3
)

// GetHorizontalContainmentLowerBound returns the lower bound of the CR. If the given CR does not have a lower bound returns -1.
// Notes:
//    - The bound is returned in meters.
//    - The bound is inclusive to the value, meaning that bound <= value.
func (cr ContainmentRadiusSurface) GetHorizontalContainmentLowerBound() float64 {
	switch cr {
	case CRSBetween3MAnd10M:
		return 3
	case CRSBetween10MAnd92Point6M:
		return 10
	case CRSGreaterThan92Point6M:
		return 92.6
	default:
		return -1
	}
}

// GetHorizontalContainmentHigherBound returns the higher bound of the CR. If the given CR does not have a higher bound returns -1.
// Notes:
//    - The bound is returned in meters.
//    - The bound is exclusive to the value, meaning that value < bound.
func (cr ContainmentRadiusSurface) GetHorizontalContainmentHigherBound() float64 {
	switch cr {
	case CRSLowerThan3M:
		return 3
	case CRSBetween3MAnd10M:
		return 10
	case CRSBetween10MAnd92Point6M:
		return 92.6
	default:
		return -1
	}
}

// ToString returns a basic, but readable, representation of the field
func (cr ContainmentRadiusSurface) ToString() string {
	switch cr {
	case CRSLowerThan3M:
		return "μ < 3 m"
	case CRSBetween3MAnd10M:
		return "3 m ≤ μ < 10 m"
	case CRSBetween10MAnd92Point6M:
		return "10 m ≤ μ < 92.6 m (0.05 NM)"
	case CRSGreaterThan92Point6M:
		return "92.6 m (0.05 NM) ≤ μ"
	default:
		return fmt.Sprintf("%v - Unknown code", cr)
	}
}
