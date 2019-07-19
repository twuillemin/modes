package fields

import "fmt"

// ContainmentRadius is the 95% Containment radius, μ and v, on horizontal and vertical position error for
// Airborne aircraft format definition
//
// Specified in Doc 9871 / A.2.3.1
type ContainmentRadius byte

const (
	// CRLowerThan3M denotes μ < 3m
	CRLowerThan3M ContainmentRadius = 0
	// CRBetween3MAnd10M denotes 3 m <= μ < 10 m
	CRBetween3MAnd10M ContainmentRadius = 1
	// CRBetween10MAnd92Point6M denotes 10 m <= μ < 92.6 m (0.05 NM)
	CRBetween10MAnd92Point6M ContainmentRadius = 2
	// CRGreaterThan92Point6M denotes 92.6 m (0.05 NM) <= μ
	CRGreaterThan92Point6M ContainmentRadius = 3
)

// GetHorizontalContainmentLowerBound returns the lower bound of the CR. If the given CR does not have a lower bound returns -1.
// Notes:
//    - The bound is returned in meters.
//    - The bound is inclusive to the value, meaning that bound <= value.
func (cr ContainmentRadius) GetHorizontalContainmentLowerBound() float64 {
	switch cr {
	case CRBetween3MAnd10M:
		return 3
	case CRBetween10MAnd92Point6M:
		return 10
	case CRGreaterThan92Point6M:
		return 92.6
	default:
		return -1
	}
}

// GetHorizontalContainmentHigherBound returns the higher bound of the CR. If the given CR does not have a higher bound returns -1.
// Notes:
//    - The bound is returned in meters.
//    - The bound is exclusive to the value, meaning that value < bound.
func (cr ContainmentRadius) GetHorizontalContainmentHigherBound() float64 {
	switch cr {
	case CRLowerThan3M:
		return 3
	case CRBetween3MAnd10M:
		return 10
	case CRBetween10MAnd92Point6M:
		return 92.6
	default:
		return -1
	}
}

// ToString returns a basic, but readable, representation of the field
func (cr ContainmentRadius) ToString() string {
	switch cr {
	case CRLowerThan3M:
		return "μ < 3 m"
	case CRBetween3MAnd10M:
		return "3 m <= μ < 10 m"
	case CRBetween10MAnd92Point6M:
		return "10 m <= μ < 92.6 m (0.05 NM)"
	case CRGreaterThan92Point6M:
		return "92.6 m (0.05 NM) <= μ"
	default:
		return fmt.Sprintf("%v - Unknown code", cr)
	}
}
