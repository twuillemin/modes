package fields

import "fmt"

// HorizontalProtectionLimit is the HPL (Horizontal Protection Limit) definition
//
// Specified in Doc 9871 / A.2.3.1
type HorizontalProtectionLimit byte

const (
	// HPLLowerThan7Point5M denotes HPL < 7.5 m
	HPLLowerThan7Point5M HorizontalProtectionLimit = 5
	// HPLLowerThan25M denotes HPL < 25 m
	HPLLowerThan25M HorizontalProtectionLimit = 6
	// HPLLowerThan185Point2M denotes HPL < 185.2 m
	HPLLowerThan185Point2M HorizontalProtectionLimit = 7
	// HPLGreaterThan185Point2M denotes HPL >= 185.2 m
	HPLGreaterThan185Point2M HorizontalProtectionLimit = 8
)

// GetHPLLowerBound returns the lower bound of the HPL. If the given HPL does not have a lower bound returns -1.
// Notes:
//   - The bound is returned in meters.
//   - The bound is inclusive to the value, meaning that bound <= value.
func (hpl HorizontalProtectionLimit) GetHPLLowerBound() float64 {
	switch hpl {
	case HPLGreaterThan185Point2M:
		return 185.2
	default:
		return -1
	}
}

// GetHPLHigherBound returns the higher bound of the HPL. If the given HPL does not have a higher bound returns -1.
// Notes:
//   - The bound is returned in meters.
//   - The bound is exclusive to the value, meaning that value < bound.
func (hpl HorizontalProtectionLimit) GetHPLHigherBound() float64 {
	switch hpl {
	case HPLLowerThan7Point5M:
		return 7.5
	case HPLLowerThan25M:
		return 25
	case HPLLowerThan185Point2M:
		return 185.2
	default:
		return -1
	}
}

// ToString returns a basic, but readable, representation of the field
func (hpl HorizontalProtectionLimit) ToString() string {
	switch hpl {
	case HPLLowerThan7Point5M:
		return "HPL < 7.5 m"
	case HPLLowerThan25M:
		return "HPL < 25 m"
	case HPLLowerThan185Point2M:
		return "HPL < 185.2 m (0.1 NM)"
	case HPLGreaterThan185Point2M:
		return "185.2 m (0.1 NM) <= HPL"
	default:
		return fmt.Sprintf("%v - Unknown code", hpl)
	}
}
