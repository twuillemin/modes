package fields

import "fmt"

// HorizontalProtectionLimitGNSS is the HorizontalProtectionLimit (Horizontal Protection Limit) for Airborne aircraft format definition
// when altitude type is GNSS in ADSB V0
//
// Specified in Doc 9871 / A.2.3.1
type HorizontalProtectionLimitGNSS byte

const (
	// HPLGLowerThan7Dot5M denotes HorizontalProtectionLimit < 7.5 m
	HPLGLowerThan7Dot5M HorizontalProtectionLimitGNSS = 0
	// HPLGLowerThan25M denotes HorizontalProtectionLimit < 25 m
	HPLGLowerThan25M HorizontalProtectionLimitGNSS = 1
	// HPLGGreaterThan25M denotes 25 m <= HorizontalProtectionLimit
	HPLGGreaterThan25M HorizontalProtectionLimitGNSS = 2
)

// ToHorizontalProtectionLimit returns the HorizontalProtectionLimit
func (hpl HorizontalProtectionLimitGNSS) ToHorizontalProtectionLimit() HorizontalProtectionLimit {
	return hpl
}

// ToString returns a basic, but readable, representation of the field
func (hpl HorizontalProtectionLimitGNSS) ToString() string {
	switch hpl {
	case HPLGLowerThan7Dot5M:
		return "HorizontalProtectionLimit < 7.5 m"
	case HPLGLowerThan25M:
		return "HorizontalProtectionLimit < 25 m"
	case HPLGGreaterThan25M:
		return "25 m <= HorizontalProtectionLimit"
	default:
		return fmt.Sprintf("%v - Unknown code", hpl)
	}
}
