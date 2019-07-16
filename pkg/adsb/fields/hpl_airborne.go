package fields

import "fmt"

// HPLAirborne is the HPL (Horizontal Protection Limit) for Airborne aircraft format definition
//
// Specified in Doc 9871 / A.2.3.1
type HPLAirborne byte

const (
	// HPLALowerThan7Dot5M denotes HPL < 7.5 m
	HPLALowerThan7Dot5M HPLAirborne = 0
	// HPLABetween7Dot5MAnd25M denotes 7.5 m ≤ HPL < 25 m
	HPLABetween7Dot5MAnd25M HPLAirborne = 1
	// HPLABetween25MAnd185Dot2M denotes 25 m ≤ HPL < 185.2 m (0.1NM)
	HPLABetween25MAnd185Dot2M HPLAirborne = 2
	// HPLABetween185Dot2MAnd370Dot4M denotes 185.2 m ≤ HPL < 370.4 m (0.1 NM ≤ HPL < 0.2 NM)
	HPLABetween185Dot2MAnd370Dot4M HPLAirborne = 3
	// HPLABetween370Dot4MAnd926M denotes 370.4 m ≤ HPL < 926 m (0.2 NM ≤ HPL < 0.5 NM)
	HPLABetween370Dot4MAnd926M HPLAirborne = 4
	// HPLABetween926MAnd1852M denotes 926 m ≤ HPL < 1 852 m (0.5 NM ≤ HPL < 1 NM)
	HPLABetween926MAnd1852M HPLAirborne = 5
	// HPLABetween1852MAnd3704M denotes 1 852 m ≤ HPL < 3 704 m (1 NM ≤ HPL < 2 NM)
	HPLABetween1852MAnd3704M HPLAirborne = 6
	// HPLABetween3704MAnd18Point52Km denotes 3.704 km ≤ HPL < 18.52 km (2 NM ≤ HPL < 10 NM)
	HPLABetween3704MAnd18Point52Km HPLAirborne = 7
	// HPLABetween18Point52KmAnd37Point04Km denotes 18.52 km ≤ HPL < 37.04 km (10 NM ≤ HPL < 20 NM)
	HPLABetween18Point52KmAnd37Point04Km HPLAirborne = 8
	// HPLALargerThan37Point04Km denotes HPL ≥ 37.04 km (HPL ≥ 20 NM)
	HPLALargerThan37Point04Km HPLAirborne = 9
)

// ToString returns a basic, but readable, representation of the field
func (hpl HPLAirborne) ToString() string {
	switch hpl {
	case HPLALowerThan7Dot5M:
		return "HPL < 7.5 m"
	case HPLABetween7Dot5MAnd25M:
		return "7.5 m ≤ HPL < 25 m"
	case HPLABetween25MAnd185Dot2M:
		return "25 m ≤ HPL < 185.2 m (0.1NM)"
	case HPLABetween185Dot2MAnd370Dot4M:
		return "185.2 m ≤ HPL < 370.4 m (0.1 NM ≤ HPL < 0.2 NM)"
	case HPLABetween370Dot4MAnd926M:
		return "370.4 m ≤ HPL < 926 m (0.2 NM ≤ HPL < 0.5 NM)"
	case HPLABetween926MAnd1852M:
		return "926 m ≤ HPL < 1 852 m (0.5 NM ≤ HPL < 1.0 NM)"
	case HPLABetween1852MAnd3704M:
		return "1 852 m ≤ HPL < 3 704 m (1.0 NM ≤ HPL < 2.0 NM)"
	case HPLABetween3704MAnd18Point52Km:
		return "3.704 km ≤ HPL < 18.52 km (2.0 NM ≤ HPL < 10 NM)"
	case HPLABetween18Point52KmAnd37Point04Km:
		return "18.52 km ≤ HPL < 37.04 km (10 NM ≤ HPL < 20 NM)"
	case HPLALargerThan37Point04Km:
		return "HPL ≥ 37.04 km (HPL ≥ 20 NM)"
	default:
		return fmt.Sprintf("%v - Unknown code", hpl)
	}
}
