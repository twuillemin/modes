package fields

import "fmt"

// HPLAirborneGNSS is the HPL (Horizontal Protection Limit) for Airborne aircraft format definition
// when altitude type is GNSS
//
// Specified in Doc 9871 / A.2.3.1
type HPLAirborneGNSS byte

const (
	// HPLAGLowerThan7Dot5M denotes HPL < 7.5 m
	HPLAGLowerThan7Dot5M HPLAirborneGNSS = 0
	// HPLAGLowerThan25M denotes HPL < 25 m
	HPLAGLowerThan25M HPLAirborneGNSS = 1
	// HPLAGGreaterThan25M denotes 25 m ≤ HPL
	HPLAGGreaterThan25M HPLAirborneGNSS = 2
)

// GetHPLLowerBound returns the lower bound of the HPL. If the given HPL does not have a lower bound returns -1.
// Notes:
//    - The bound is returned in meters.
//    - The bound is inclusive to the value, meaning that bound <= value.
func (hpl HPLAirborneGNSS) GetHPLLowerBound() float64 {
	switch hpl {
	case HPLAGGreaterThan25M:
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
	case HPLAGLowerThan7Dot5M:
		return 7.5
	case HPLAGLowerThan25M:
		return 25
	default:
		return -1
	}
}

// ToString returns a basic, but readable, representation of the field
func (hpl HPLAirborneGNSS) ToString() string {
	switch hpl {
	case HPLAGLowerThan7Dot5M:
		return "HPL < 7.5 m"
	case HPLAGLowerThan25M:
		return "HPL < 25 m"
	case HPLAGGreaterThan25M:
		return "25 m ≤ HPL"
	default:
		return fmt.Sprintf("%v - Unknown code", hpl)
	}
}
