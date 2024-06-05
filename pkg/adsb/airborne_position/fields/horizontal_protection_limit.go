package fields

import "fmt"

// HorizontalProtectionLimitBarometric is the HorizontalProtectionLimit (Horizontal Protection Limit) for Airborne aircraft format definition
// when altitude type is barometric in ADSB V0
//
// Specified in Doc 9871 / A.2.3.1
type HorizontalProtectionLimit byte

const (
	// HPLBaroLowerThan7Dot5M denotes HorizontalProtectionLimit < 7.5 m
	HPLBaroLowerThan7Dot5M HorizontalProtectionLimit = 9
	// HPLBaroBetween7Dot5MAnd25M denotes 7.5 m <= HorizontalProtectionLimit < 25 m
	HPLBaroBetween7Dot5MAnd25M HorizontalProtectionLimit = 10
	// HPLBaroBetween25MAnd185Dot2M denotes 25 m <= HorizontalProtectionLimit < 185.2 m (0.1NM)
	HPLBaroBetween25MAnd185Dot2M HorizontalProtectionLimit = 11
	// HPLBaroBetween185Dot2MAnd370Dot4M denotes 185.2 m <= HorizontalProtectionLimit < 370.4 m (0.1 NM <= HorizontalProtectionLimit < 0.2 NM)
	HPLBaroBetween185Dot2MAnd370Dot4M HorizontalProtectionLimit = 12
	// HPLBaroBetween370Dot4MAnd926M denotes 370.4 m <= HorizontalProtectionLimit < 926 m (0.2 NM <= HorizontalProtectionLimit < 0.5 NM)
	HPLBaroBetween370Dot4MAnd926M HorizontalProtectionLimit = 13
	// HPLBaroBetween926MAnd1852M denotes 926 m <= HorizontalProtectionLimit < 1 852 m (0.5 NM <= HorizontalProtectionLimit < 1 NM)
	HPLBaroBetween926MAnd1852M HorizontalProtectionLimit = 14
	// HPLBaroBetween1852MAnd3704M denotes 1 852 m <= HorizontalProtectionLimit < 3 704 m (1 NM <= HorizontalProtectionLimit < 2 NM)
	HPLBaroBetween1852MAnd3704M HorizontalProtectionLimit = 15
	// HPLBaroBetween3704MAnd18Point52Km denotes 3.704 km <= HorizontalProtectionLimit < 18.52 km (2 NM <= HorizontalProtectionLimit < 10 NM)
	HPLBaroBetween3704MAnd18Point52Km HorizontalProtectionLimit = 16
	// HPLBaroBetween18Point52KmAnd37Point04Km denotes 18.52 km <= HorizontalProtectionLimit < 37.04 km (10 NM <= HorizontalProtectionLimit < 20 NM)
	HPLBaroBetween18Point52KmAnd37Point04Km HorizontalProtectionLimit = 17
	// HPLBaroLargerThan37Point04Km denotes HorizontalProtectionLimit >= 37.04 km (HorizontalProtectionLimit >= 20 NM)
	HPLBaroLargerThan37Point04Km HorizontalProtectionLimit = 18
	// HPLGNSSLowerThan7Dot5M denotes HorizontalProtectionLimit < 7.5 m
	HPLGNSSLowerThan7Dot5M HorizontalProtectionLimit = 20
	// HPLGNSSLowerThan25M denotes HorizontalProtectionLimit < 25 m
	HPLGNSSLowerThan25M HorizontalProtectionLimit = 21
	// HPLGNSSGreaterThan25M denotes 25 m <= HorizontalProtectionLimit
	HPLGNSSGreaterThan25M HorizontalProtectionLimit = 22
)

// ToString returns a basic, but readable, representation of the field
func (hpl HorizontalProtectionLimit) ToString() string {
	switch hpl {
	case HPLBaroLowerThan7Dot5M:
		return "(Baro) HorizontalProtectionLimit < 7.5 m"
	case HPLBaroBetween7Dot5MAnd25M:
		return "(Baro) 7.5 m <= HorizontalProtectionLimit < 25 m"
	case HPLBaroBetween25MAnd185Dot2M:
		return "(Baro) 25 m <= HorizontalProtectionLimit < 185.2 m (0.1NM)"
	case HPLBaroBetween185Dot2MAnd370Dot4M:
		return "(Baro) 185.2 m <= HorizontalProtectionLimit < 370.4 m (0.1 NM <= HorizontalProtectionLimit < 0.2 NM)"
	case HPLBaroBetween370Dot4MAnd926M:
		return "(Baro) 370.4 m <= HorizontalProtectionLimit < 926 m (0.2 NM <= HorizontalProtectionLimit < 0.5 NM)"
	case HPLBaroBetween926MAnd1852M:
		return "(Baro) 926 m <= HorizontalProtectionLimit < 1 852 m (0.5 NM <= HorizontalProtectionLimit < 1.0 NM)"
	case HPLBaroBetween1852MAnd3704M:
		return "(Baro) 1 852 m <= HorizontalProtectionLimit < 3 704 m (1.0 NM <= HorizontalProtectionLimit < 2.0 NM)"
	case HPLBaroBetween3704MAnd18Point52Km:
		return "(Baro) 3.704 km <= HorizontalProtectionLimit < 18.52 km (2.0 NM <= HorizontalProtectionLimit < 10 NM)"
	case HPLBaroBetween18Point52KmAnd37Point04Km:
		return "(Baro) 18.52 km <= HorizontalProtectionLimit < 37.04 km (10 NM <= HorizontalProtectionLimit < 20 NM)"
	case HPLBaroLargerThan37Point04Km:
		return "(Baro) HorizontalProtectionLimit >= 37.04 km (HorizontalProtectionLimit >= 20 NM)"
	case HPLGNSSLowerThan7Dot5M:
		return "(GNSS) HorizontalProtectionLimit < 7.5 m"
	case HPLGNSSLowerThan25M:
		return "(GNSS) HorizontalProtectionLimit < 25 m"
	case HPLGNSSGreaterThan25M:
		return "(GNSS) 25 m <= HorizontalProtectionLimit"
	default:
		return fmt.Sprintf("%v - Unknown code", hpl)
	}
}
