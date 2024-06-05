package fields

import "fmt"

// ContainmentRadius is the 95% Containment radius,μ and v, on horizontal and vertical position error for
// Airborne aircraft format definition in ADSB V0
//
// Specified in Doc 9871 / A.2.3.1
type ContainmentRadius byte

const (
	// CRBaroLowerThan3M denotes μ < 3m
	CRBaroLowerThan3M ContainmentRadius = 9
	// CRBaroBetween3MAnd10M denotes 3 m <= μ < 10 m
	CRBaroBetween3MAnd10M ContainmentRadius = 10
	// CRBaroBetween10MAnd92Dot6M denotes 10 m <= μ < 92.6 m (0.05 NM)
	CRBaroBetween10MAnd92Dot6M ContainmentRadius = 11
	// CRBaroBetween92Dot6MAnd185Dot2M denotes 92.6 m <= μ < 185.2 m (0.05 NM <= μ < 0.1 NM)
	CRBaroBetween92Dot6MAnd185Dot2M ContainmentRadius = 12
	// CRBaroBetween185Dot2MAnd463M denotes 185.2 m <= μ < 463 m (0.1 NM <= μ < 0.25 NM)
	CRBaroBetween185Dot2MAnd463M ContainmentRadius = 13
	// CRBaroBetween463MAnd926M denotes 463 m <= μ < 926 m (0.25 NM <= μ < 0.5 NM)
	CRBaroBetween463MAnd926M ContainmentRadius = 14
	// CRBaroBetween926MAnd1852M denotes 926 m <= μ < 1852 m (0.5 NM <= μ < 1.0 NM)
	CRBaroBetween926MAnd1852M ContainmentRadius = 15
	// CRBaroBetween1Point852KmAnd9Point26Km denotes 1.852 km <= μ < 9.26 km (1.0 NM <= μ < 5.0 NM)
	CRBaroBetween1Point852KmAnd9Point26Km ContainmentRadius = 16
	// CRBaroBetween9Point26KmAnd18Point52Km denotes 9.26 km <= μ < 18.52 km (5.0 NM <= μ < 10.0 NM)
	CRBaroBetween9Point26KmAnd18Point52Km ContainmentRadius = 17
	// CRBaroLargerThan18Point52Km denotes 18.52 km <= μ (10.0 NM <= μ)
	CRBaroLargerThan18Point52Km ContainmentRadius = 18
	// CRGNSSHorizontalLowerThan3MAndVerticalLowerThan4M denotes μ < 3 m and v < 4 m
	CRGNSSHorizontalLowerThan3MAndVerticalLowerThan4M ContainmentRadius = 20
	// CRGNSSHorizontalLowerThan10MAndVerticalLowerThan15M denotes μ < 10 m and v < 15 m
	CRGNSSHorizontalLowerThan10MAndVerticalLowerThan15M ContainmentRadius = 21
	// CRGNSSHorizontalGreaterThan10MOrVerticalGreaterThan15M denotes μ > 10 m or v >= 15 m
	CRGNSSHorizontalGreaterThan10MOrVerticalGreaterThan15M ContainmentRadius = 22
)

// ToString returns a basic, but readable, representation of the field
func (cr ContainmentRadius) ToString() string {
	switch cr {
	case CRBaroLowerThan3M:
		return "(Baro) 0 μ < 3m"
	case CRBaroBetween3MAnd10M:
		return "(Baro) 3 m <= μ < 10 m"
	case CRBaroBetween10MAnd92Dot6M:
		return "(Baro) 10 m <= μ < 92.6 m (0.05 NM)"
	case CRBaroBetween92Dot6MAnd185Dot2M:
		return "(Baro) 92.6 m <= μ < 185.2 m (0.05 NM <= μ < 0.1 NM)"
	case CRBaroBetween185Dot2MAnd463M:
		return "(Baro) 185.2 m <= μ < 463 m (0.1 NM <= μ < 0.25 NM)"
	case CRBaroBetween463MAnd926M:
		return "(Baro) 463 m <= μ < 926 m (0.25 NM <= μ < 0.5 NM)"
	case CRBaroBetween926MAnd1852M:
		return "(Baro) 926 m <= μ < 1852 m (0.5 NM <= μ < 1.0 NM)"
	case CRBaroBetween1Point852KmAnd9Point26Km:
		return "(Baro) 1.852 km <= μ < 9.26 km (1.0 NM <= μ < 5.0 NM)"
	case CRBaroBetween9Point26KmAnd18Point52Km:
		return "(Baro) 9.26 km <= μ < 18.52 km (5.0 NM <= μ < 10.0 NM)"
	case CRBaroLargerThan18Point52Km:
		return "(Baro) 18.52 km <= μ (10.0 NM <= μ)"
	case CRGNSSHorizontalLowerThan3MAndVerticalLowerThan4M:
		return "(GNSS) μ < 3 m and v < 4 m"
	case CRGNSSHorizontalLowerThan10MAndVerticalLowerThan15M:
		return "(GNSS) μ < 10 m and v < 15 m"
	case CRGNSSHorizontalGreaterThan10MOrVerticalGreaterThan15M:
		return "(GNSS) μ > 10 m or v >= 15 m"
	default:
		return fmt.Sprintf("%v - Unknown code", cr)
	}
}
