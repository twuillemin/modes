package fields

import "fmt"

// HPLAirborneGNSS is the HPL (Horizontal Protection Limit) for Airborne aircraft format definition
// when altitude type is GNSS
//
// Specified in Doc 9871 / A.2.3.1
type HPLAirborneGNSS byte

const (
	// HPLGLowerThan7Dot5M denotes HPL < 7.5 m
	HPLGLowerThan7Dot5M HPLAirborneGNSS = 0
	// HPLGLowerThan25M denotes HPL < 25 m
	HPLGLowerThan25M HPLAirborneGNSS = 1
	// HPLGGreaterThan25M denotes 25 m <= HPL
	HPLGGreaterThan25M HPLAirborneGNSS = 2
)

// GetHPLLowerBound returns the lower bound of the HPL. If the given HPL does not have a lower bound returns -1.
// Notes:
//    - The bound is returned in meters.
//    - The bound is inclusive to the value, meaning that bound <= value.
func (hpl HPLAirborneGNSS) GetHPLLowerBound() float64 {
	switch hpl {
	case HPLGGreaterThan25M:
		return 25
	default:
		return -1
	}
}

// GetHPLHigherBound returns the higher bound of the HPL. If the given HPL does not have a higher bound returns -1.
// Notes:
//    - The bound is returned in meters.
//    - The bound is exclusive to the value, meaning that value < bound.
func (hpl HPLAirborneGNSS) GetHPLHigherBound() float64 {
	switch hpl {
	case HPLGLowerThan7Dot5M:
		return 7.5
	case HPLGLowerThan25M:
		return 25
	default:
		return -1
	}
}

// ToString returns a basic, but readable, representation of the field
func (hpl HPLAirborneGNSS) ToString() string {
	switch hpl {
	case HPLGLowerThan7Dot5M:
		return "HPL < 7.5 m"
	case HPLGLowerThan25M:
		return "HPL < 25 m"
	case HPLGGreaterThan25M:
		return "25 m <= HPL"
	default:
		return fmt.Sprintf("%v - Unknown code", hpl)
	}
}
