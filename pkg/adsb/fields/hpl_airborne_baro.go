package fields

import "fmt"

// HPLAirborneBarometric is the HPL (Horizontal Protection Limit) for Airborne aircraft format definition
// when altitude type is barometric
//
// Specified in Doc 9871 / A.2.3.1
type HPLAirborneBarometric byte

const (
	// HPLABLowerThan7Dot5M denotes HPL < 7.5 m
	HPLABLowerThan7Dot5M HPLAirborneBarometric = 0
	// HPLABBetween7Dot5MAnd25M denotes 7.5 m ≤ HPL < 25 m
	HPLABBetween7Dot5MAnd25M HPLAirborneBarometric = 1
	// HPLABBetween25MAnd185Dot2M denotes 25 m ≤ HPL < 185.2 m (0.1NM)
	HPLABBetween25MAnd185Dot2M HPLAirborneBarometric = 2
	// HPLABBetween185Dot2MAnd370Dot4M denotes 185.2 m ≤ HPL < 370.4 m (0.1 NM ≤ HPL < 0.2 NM)
	HPLABBetween185Dot2MAnd370Dot4M HPLAirborneBarometric = 3
	// HPLABBetween370Dot4MAnd926M denotes 370.4 m ≤ HPL < 926 m (0.2 NM ≤ HPL < 0.5 NM)
	HPLABBetween370Dot4MAnd926M HPLAirborneBarometric = 4
	// HPLABBetween926MAnd1852M denotes 926 m ≤ HPL < 1 852 m (0.5 NM ≤ HPL < 1 NM)
	HPLABBetween926MAnd1852M HPLAirborneBarometric = 5
	// HPLABBetween1852MAnd3704M denotes 1 852 m ≤ HPL < 3 704 m (1 NM ≤ HPL < 2 NM)
	HPLABBetween1852MAnd3704M HPLAirborneBarometric = 6
	// HPLABBetween3704MAnd18Point52Km denotes 3.704 km ≤ HPL < 18.52 km (2 NM ≤ HPL < 10 NM)
	HPLABBetween3704MAnd18Point52Km HPLAirborneBarometric = 7
	// HPLABBetween18Point52KmAnd37Point04Km denotes 18.52 km ≤ HPL < 37.04 km (10 NM ≤ HPL < 20 NM)
	HPLABBetween18Point52KmAnd37Point04Km HPLAirborneBarometric = 8
	// HPLABLargerThan37Point04Km denotes HPL ≥ 37.04 km (HPL ≥ 20 NM)
	HPLABLargerThan37Point04Km HPLAirborneBarometric = 9
)

// GetHPLLowerBound returns the lower bound of the HPL. If the given HPL does not have a lower bound returns -1.
// Notes:
//    - The bound is returned in meters.
//    - The bound is inclusive to the value, meaning that bound <= value.
func (hpl HPLAirborneBarometric) GetHPLLowerBound() float64 {
	switch hpl {
	case HPLABBetween7Dot5MAnd25M:
		return 7.5
	case HPLABBetween25MAnd185Dot2M:
		return 25
	case HPLABBetween185Dot2MAnd370Dot4M:
		return 185.2
	case HPLABBetween370Dot4MAnd926M:
		return 370.4
	case HPLABBetween926MAnd1852M:
		return 926
	case HPLABBetween1852MAnd3704M:
		return 1852
	case HPLABBetween3704MAnd18Point52Km:
		return 3704
	case HPLABBetween18Point52KmAnd37Point04Km:
		return 18520
	case HPLABLargerThan37Point04Km:
		return 37040
	default:
		return -1
	}
}

// GetHPLHigherBound returns the higher bound of the HPL. If the given HPL does not have a higher bound returns -1.
// Notes:
//    - The bound is returned in meters.
//    - The bound is exclusive to the value, meaning that value < bound.
func (hpl HPLAirborneBarometric) GetHPLHigherBound() float64 {
	switch hpl {
	case HPLABLowerThan7Dot5M:
		return 7.5
	case HPLABBetween7Dot5MAnd25M:
		return 25
	case HPLABBetween25MAnd185Dot2M:
		return 185.2
	case HPLABBetween185Dot2MAnd370Dot4M:
		return 370.4
	case HPLABBetween370Dot4MAnd926M:
		return 926
	case HPLABBetween926MAnd1852M:
		return 1852
	case HPLABBetween1852MAnd3704M:
		return 3704
	case HPLABBetween3704MAnd18Point52Km:
		return 18520
	case HPLABBetween18Point52KmAnd37Point04Km:
		return 37040
	default:
		return -1
	}
}

// ToString returns a basic, but readable, representation of the field
func (hpl HPLAirborneBarometric) ToString() string {
	switch hpl {
	case HPLABLowerThan7Dot5M:
		return "HPL < 7.5 m"
	case HPLABBetween7Dot5MAnd25M:
		return "7.5 m ≤ HPL < 25 m"
	case HPLABBetween25MAnd185Dot2M:
		return "25 m ≤ HPL < 185.2 m (0.1NM)"
	case HPLABBetween185Dot2MAnd370Dot4M:
		return "185.2 m ≤ HPL < 370.4 m (0.1 NM ≤ HPL < 0.2 NM)"
	case HPLABBetween370Dot4MAnd926M:
		return "370.4 m ≤ HPL < 926 m (0.2 NM ≤ HPL < 0.5 NM)"
	case HPLABBetween926MAnd1852M:
		return "926 m ≤ HPL < 1 852 m (0.5 NM ≤ HPL < 1.0 NM)"
	case HPLABBetween1852MAnd3704M:
		return "1 852 m ≤ HPL < 3 704 m (1.0 NM ≤ HPL < 2.0 NM)"
	case HPLABBetween3704MAnd18Point52Km:
		return "3.704 km ≤ HPL < 18.52 km (2.0 NM ≤ HPL < 10 NM)"
	case HPLABBetween18Point52KmAnd37Point04Km:
		return "18.52 km ≤ HPL < 37.04 km (10 NM ≤ HPL < 20 NM)"
	case HPLABLargerThan37Point04Km:
		return "HPL ≥ 37.04 km (HPL ≥ 20 NM)"
	default:
		return fmt.Sprintf("%v - Unknown code", hpl)
	}
}
