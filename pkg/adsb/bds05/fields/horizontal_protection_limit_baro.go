package fields

import "fmt"

// HorizontalProtectionLimitBarometric is the HorizontalProtectionLimit (Horizontal Protection Limit) for Airborne aircraft format definition
// when altitude type is barometric in ADSB V0
//
// Specified in Doc 9871 / A.2.3.1
type HorizontalProtectionLimitBarometric byte

const (
	// HPLBLowerThan7Dot5M denotes HorizontalProtectionLimit < 7.5 m
	HPLBLowerThan7Dot5M HorizontalProtectionLimitBarometric = 0
	// HPLBBetween7Dot5MAnd25M denotes 7.5 m <= HorizontalProtectionLimit < 25 m
	HPLBBetween7Dot5MAnd25M HorizontalProtectionLimitBarometric = 1
	// HPLBBetween25MAnd185Dot2M denotes 25 m <= HorizontalProtectionLimit < 185.2 m (0.1NM)
	HPLBBetween25MAnd185Dot2M HorizontalProtectionLimitBarometric = 2
	// HPLBBetween185Dot2MAnd370Dot4M denotes 185.2 m <= HorizontalProtectionLimit < 370.4 m (0.1 NM <= HorizontalProtectionLimit < 0.2 NM)
	HPLBBetween185Dot2MAnd370Dot4M HorizontalProtectionLimitBarometric = 3
	// HPLBBetween370Dot4MAnd926M denotes 370.4 m <= HorizontalProtectionLimit < 926 m (0.2 NM <= HorizontalProtectionLimit < 0.5 NM)
	HPLBBetween370Dot4MAnd926M HorizontalProtectionLimitBarometric = 4
	// HPLBBetween926MAnd1852M denotes 926 m <= HorizontalProtectionLimit < 1 852 m (0.5 NM <= HorizontalProtectionLimit < 1 NM)
	HPLBBetween926MAnd1852M HorizontalProtectionLimitBarometric = 5
	// HPLBBetween1852MAnd3704M denotes 1 852 m <= HorizontalProtectionLimit < 3 704 m (1 NM <= HorizontalProtectionLimit < 2 NM)
	HPLBBetween1852MAnd3704M HorizontalProtectionLimitBarometric = 6
	// HPLBBetween3704MAnd18Point52Km denotes 3.704 km <= HorizontalProtectionLimit < 18.52 km (2 NM <= HorizontalProtectionLimit < 10 NM)
	HPLBBetween3704MAnd18Point52Km HorizontalProtectionLimitBarometric = 7
	// HPLBBetween18Point52KmAnd37Point04Km denotes 18.52 km <= HorizontalProtectionLimit < 37.04 km (10 NM <= HorizontalProtectionLimit < 20 NM)
	HPLBBetween18Point52KmAnd37Point04Km HorizontalProtectionLimitBarometric = 8
	// HPLBLargerThan37Point04Km denotes HorizontalProtectionLimit >= 37.04 km (HorizontalProtectionLimit >= 20 NM)
	HPLBLargerThan37Point04Km HorizontalProtectionLimitBarometric = 9
)

// ToHorizontalProtectionLimit returns the HorizontalProtectionLimit
func (hpl HorizontalProtectionLimitBarometric) ToHorizontalProtectionLimit() HorizontalProtectionLimit {
	return hpl
}

// ToString returns a basic, but readable, representation of the field
func (hpl HorizontalProtectionLimitBarometric) ToString() string {
	switch hpl {
	case HPLBLowerThan7Dot5M:
		return "HorizontalProtectionLimit < 7.5 m"
	case HPLBBetween7Dot5MAnd25M:
		return "7.5 m <= HorizontalProtectionLimit < 25 m"
	case HPLBBetween25MAnd185Dot2M:
		return "25 m <= HorizontalProtectionLimit < 185.2 m (0.1NM)"
	case HPLBBetween185Dot2MAnd370Dot4M:
		return "185.2 m <= HorizontalProtectionLimit < 370.4 m (0.1 NM <= HorizontalProtectionLimit < 0.2 NM)"
	case HPLBBetween370Dot4MAnd926M:
		return "370.4 m <= HorizontalProtectionLimit < 926 m (0.2 NM <= HorizontalProtectionLimit < 0.5 NM)"
	case HPLBBetween926MAnd1852M:
		return "926 m <= HorizontalProtectionLimit < 1 852 m (0.5 NM <= HorizontalProtectionLimit < 1.0 NM)"
	case HPLBBetween1852MAnd3704M:
		return "1 852 m <= HorizontalProtectionLimit < 3 704 m (1.0 NM <= HorizontalProtectionLimit < 2.0 NM)"
	case HPLBBetween3704MAnd18Point52Km:
		return "3.704 km <= HorizontalProtectionLimit < 18.52 km (2.0 NM <= HorizontalProtectionLimit < 10 NM)"
	case HPLBBetween18Point52KmAnd37Point04Km:
		return "18.52 km <= HorizontalProtectionLimit < 37.04 km (10 NM <= HorizontalProtectionLimit < 20 NM)"
	case HPLBLargerThan37Point04Km:
		return "HorizontalProtectionLimit >= 37.04 km (HorizontalProtectionLimit >= 20 NM)"
	default:
		return fmt.Sprintf("%v - Unknown code", hpl)
	}
}
