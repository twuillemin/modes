package fields

import "fmt"

// HPLSurface is the HPL (Horizontal Protection Limit) for Surface aircraft format definition
//
// Specified in Doc 9871 / A.2.3.1
type HPLSurface byte

const (
	// HPLSLowerThan7Dot5M denotes HPL < 7.5 m
	HPLSLowerThan7Dot5M HPLSurface = 0
	// HPLSLowerThan25M denotes HPL < 25 m
	HPLSLowerThan25M HPLSurface = 1
	// HPLSLowerThan185Point2M denotes HPL < 25 m
	HPLSLowerThan185Point2M HPLSurface = 2
	// HPLSGreaterThan185Point2M denotes 25 m ≤ HPL
	HPLSGreaterThan185Point2M HPLSurface = 3
)

// GetHPLLowerBound returns the lower bound of the HPL. If the given HPL does not have a lower bound returns -1.
// Notes:
//    - The bound is returned in meters.
//    - The bound is inclusive to the value, meaning that bound <= value.
func (hpl HPLSurface) GetHPLLowerBound() float64 {
	switch hpl {
	case HPLSGreaterThan185Point2M:
		return 185.2
	default:
		return -1
	}
}

// GetHPLHigherBound returns the higher bound of the HPL. If the given HPL does not have a higher bound returns -1.
// Notes:
//    - The bound is returned in meters.
//    - The bound is exclusive to the value, meaning that value < bound.
func (hpl HPLSurface) GetHPLHigherBound() float64 {
	switch hpl {
	case HPLSLowerThan7Dot5M:
		return 7.5
	case HPLSLowerThan25M:
		return 25
	case HPLSLowerThan185Point2M:
		return 185.2
	default:
		return -1
	}
}

// ToString returns a basic, but readable, representation of the field
func (hpl HPLSurface) ToString() string {
	switch hpl {
	case HPLSLowerThan7Dot5M:
		return "HPL < 7.5 m"
	case HPLSLowerThan25M:
		return "HPL < 25 m"
	case HPLSLowerThan185Point2M:
		return "HPL < 185.2 m (0.1 NM)"
	case HPLSGreaterThan185Point2M:
		return "185.2 m (0.1 NM) ≤ HPL"
	default:
		return fmt.Sprintf("%v - Unknown code", hpl)
	}
}
