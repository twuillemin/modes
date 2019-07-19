package fields

import "fmt"

// ContainmentRadiusAirborneBarometric is the 95% Containment radius,μ and v, on horizontal and vertical position error for
// Airborne aircraft format definition
//
// Specified in Doc 9871 / A.2.3.1
type ContainmentRadiusAirborneBarometric byte

const (
	// CRBLowerThan3M denotes μ < 3m
	CRBLowerThan3M ContainmentRadiusAirborneBarometric = 0
	// CRBBetween3MAnd10M denotes 3 m <= μ < 10 m
	CRBBetween3MAnd10M ContainmentRadiusAirborneBarometric = 1
	// CRBBetween10MAnd92Dot6M denotes 10 m <= μ < 92.6 m (0.05 NM)
	CRBBetween10MAnd92Dot6M ContainmentRadiusAirborneBarometric = 2
	// CRBBetween92Dot6MAnd185Dot2M denotes 92.6 m <= μ < 185.2 m (0.05 NM <= μ < 0.1 NM)
	CRBBetween92Dot6MAnd185Dot2M ContainmentRadiusAirborneBarometric = 3
	// CRBBetween185Dot2MAnd463M denotes 185.2 m <= μ < 463 m (0.1 NM <= μ < 0.25 NM)
	CRBBetween185Dot2MAnd463M ContainmentRadiusAirborneBarometric = 4
	// CRBBetween463MAnd926M denotes 463 m <= μ < 926 m (0.25 NM <= μ < 0.5 NM)
	CRBBetween463MAnd926M ContainmentRadiusAirborneBarometric = 5
	// CRBBetween926MAnd1852M denotes 926 m <= μ < 1852 m (0.5 NM <= μ < 1.0 NM)
	CRBBetween926MAnd1852M ContainmentRadiusAirborneBarometric = 6
	// CRBBetween1Point852KmAnd9Point26Km denotes 1.852 km <= μ < 9.26 km (1.0 NM <= μ < 5.0 NM)
	CRBBetween1Point852KmAnd9Point26Km ContainmentRadiusAirborneBarometric = 7
	// CRBBetween9Point26KmAnd18Point52Km denotes 9.26 km <= μ < 18.52 km (5.0 NM <= μ < 10.0 NM)
	CRBBetween9Point26KmAnd18Point52Km ContainmentRadiusAirborneBarometric = 8
	// CRBLargerThan18Point52Km denotes 18.52 km <= μ (10.0 NM <= μ)
	CRBLargerThan18Point52Km ContainmentRadiusAirborneBarometric = 9
)

// GetHorizontalContainmentLowerBound returns the lower bound of the CR. If the given CR does not have a lower bound returns -1.
// Notes:
//    - The bound is returned in meters.
//    - The bound is inclusive to the value, meaning that bound <= value.
func (cr ContainmentRadiusAirborneBarometric) GetHorizontalContainmentLowerBound() float64 {
	switch cr {
	case CRBBetween3MAnd10M:
		return 3
	case CRBBetween10MAnd92Dot6M:
		return 10
	case CRBBetween92Dot6MAnd185Dot2M:
		return 92.6
	case CRBBetween185Dot2MAnd463M:
		return 185.2
	case CRBBetween463MAnd926M:
		return 463
	case CRBBetween926MAnd1852M:
		return 926
	case CRBBetween1Point852KmAnd9Point26Km:
		return 1852
	case CRBBetween9Point26KmAnd18Point52Km:
		return 9260
	case CRBLargerThan18Point52Km:
		return 18520
	default:
		return -1
	}
}

// GetHorizontalContainmentHigherBound returns the higher bound of the CR. If the given CR does not have a higher bound returns -1.
// Notes:
//    - The bound is returned in meters.
//    - The bound is exclusive to the value, meaning that value < bound.
func (cr ContainmentRadiusAirborneBarometric) GetHorizontalContainmentHigherBound() float64 {
	switch cr {
	case CRBLowerThan3M:
		return 3
	case CRBBetween3MAnd10M:
		return 10
	case CRBBetween10MAnd92Dot6M:
		return 92.6
	case CRBBetween92Dot6MAnd185Dot2M:
		return 185.2
	case CRBBetween185Dot2MAnd463M:
		return 463
	case CRBBetween463MAnd926M:
		return 926
	case CRBBetween926MAnd1852M:
		return 1852
	case CRBBetween1Point852KmAnd9Point26Km:
		return 9260
	case CRBBetween9Point26KmAnd18Point52Km:
		return 18520
	default:
		return -1
	}
}

// ToString returns a basic, but readable, representation of the field
func (cr ContainmentRadiusAirborneBarometric) ToString() string {
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
