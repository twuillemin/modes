package fields

import "fmt"

// ContainmentRadiusBarometric is the 95% Containment radius,μ and v, on horizontal and vertical position error for
// Airborne aircraft format definition in ADSB V0
//
// Specified in Doc 9871 / A.2.3.1
type ContainmentRadiusBarometric byte

const (
	// CRBLowerThan3M denotes μ < 3m
	CRBLowerThan3M ContainmentRadiusBarometric = 0
	// CRBBetween3MAnd10M denotes 3 m <= μ < 10 m
	CRBBetween3MAnd10M ContainmentRadiusBarometric = 1
	// CRBBetween10MAnd92Dot6M denotes 10 m <= μ < 92.6 m (0.05 NM)
	CRBBetween10MAnd92Dot6M ContainmentRadiusBarometric = 2
	// CRBBetween92Dot6MAnd185Dot2M denotes 92.6 m <= μ < 185.2 m (0.05 NM <= μ < 0.1 NM)
	CRBBetween92Dot6MAnd185Dot2M ContainmentRadiusBarometric = 3
	// CRBBetween185Dot2MAnd463M denotes 185.2 m <= μ < 463 m (0.1 NM <= μ < 0.25 NM)
	CRBBetween185Dot2MAnd463M ContainmentRadiusBarometric = 4
	// CRBBetween463MAnd926M denotes 463 m <= μ < 926 m (0.25 NM <= μ < 0.5 NM)
	CRBBetween463MAnd926M ContainmentRadiusBarometric = 5
	// CRBBetween926MAnd1852M denotes 926 m <= μ < 1852 m (0.5 NM <= μ < 1.0 NM)
	CRBBetween926MAnd1852M ContainmentRadiusBarometric = 6
	// CRBBetween1Point852KmAnd9Point26Km denotes 1.852 km <= μ < 9.26 km (1.0 NM <= μ < 5.0 NM)
	CRBBetween1Point852KmAnd9Point26Km ContainmentRadiusBarometric = 7
	// CRBBetween9Point26KmAnd18Point52Km denotes 9.26 km <= μ < 18.52 km (5.0 NM <= μ < 10.0 NM)
	CRBBetween9Point26KmAnd18Point52Km ContainmentRadiusBarometric = 8
	// CRBLargerThan18Point52Km denotes 18.52 km <= μ (10.0 NM <= μ)
	CRBLargerThan18Point52Km ContainmentRadiusBarometric = 9
)

// ToContainmentRadius returns the ContainmentRadius
func (cr ContainmentRadiusBarometric) ToContainmentRadius() ContainmentRadius {
	return cr
}

// ToString returns a basic, but readable, representation of the field
func (cr ContainmentRadiusBarometric) ToString() string {
	switch cr {
	case CRBLowerThan3M:
		return "μ < 3m"
	case CRBBetween3MAnd10M:
		return "3 m <= μ < 10 m"
	case CRBBetween10MAnd92Dot6M:
		return "10 m <= μ < 92.6 m (0.05 NM)"
	case CRBBetween92Dot6MAnd185Dot2M:
		return "92.6 m <= μ < 185.2 m (0.05 NM <= μ < 0.1 NM)"
	case CRBBetween185Dot2MAnd463M:
		return "185.2 m <= μ < 463 m (0.1 NM <= μ < 0.25 NM)"
	case CRBBetween463MAnd926M:
		return "463 m <= μ < 926 m (0.25 NM <= μ < 0.5 NM)"
	case CRBBetween926MAnd1852M:
		return "926 m <= μ < 1852 m (0.5 NM <= μ < 1.0 NM)"
	case CRBBetween1Point852KmAnd9Point26Km:
		return "1.852 km <= μ < 9.26 km (1.0 NM <= μ < 5.0 NM)"
	case CRBBetween9Point26KmAnd18Point52Km:
		return "9.26 km <= μ < 18.52 km (5.0 NM <= μ < 10.0 NM)"
	case CRBLargerThan18Point52Km:
		return "18.52 km <= μ (10.0 NM <= μ)"
	default:
		return fmt.Sprintf("%v - Unknown code", cr)
	}
}
