package fields

import "fmt"

// ContainmentRadiusAirborneBarometric is the 95% Containment radius,μ and v, on horizontal and vertical position error for
// Airborne aircraft format definition
//
// Specified in Doc 9871 / A.2.3.1
type ContainmentRadiusAirborneBarometric byte

const (
	// CRABLowerThan3M denotes μ < 3m
	CRABLowerThan3M ContainmentRadiusAirborneBarometric = 0
	// CRABBetween3MAnd10M denotes 3 m ≤ μ < 10 m
	CRABBetween3MAnd10M ContainmentRadiusAirborneBarometric = 1
	// CRABBetween10MAnd92Dot6M denotes 10 m ≤ μ < 92.6 m (0.05 NM)
	CRABBetween10MAnd92Dot6M ContainmentRadiusAirborneBarometric = 2
	// CRABBetween92Dot6MAnd185Dot2M denotes 92.6 m ≤ μ < 185.2 m (0.05 NM ≤ μ < 0.1 NM)
	CRABBetween92Dot6MAnd185Dot2M ContainmentRadiusAirborneBarometric = 3
	// CRABBetween185Dot2MAnd463M denotes 185.2 m ≤ μ < 463 m (0.1 NM ≤ μ < 0.25 NM)
	CRABBetween185Dot2MAnd463M ContainmentRadiusAirborneBarometric = 4
	// CRABBetween463MAnd926M denotes 463 m ≤ μ < 926 m (0.25 NM ≤ μ < 0.5 NM)
	CRABBetween463MAnd926M ContainmentRadiusAirborneBarometric = 5
	// CRABBetween926MAnd1852M denotes 926 m ≤ μ < 1852 m (0.5 NM ≤ μ < 1.0 NM)
	CRABBetween926MAnd1852M ContainmentRadiusAirborneBarometric = 6
	// CRABBetween1Point852KmAnd9Point26Km denotes 1.852 km ≤ μ < 9.26 km (1.0 NM ≤ μ < 5.0 NM)
	CRABBetween1Point852KmAnd9Point26Km ContainmentRadiusAirborneBarometric = 7
	// CRABBetween9Point26KmAnd18Point52Km denotes 9.26 km ≤ μ < 18.52 km (5.0 NM ≤ μ < 10.0 NM)
	CRABBetween9Point26KmAnd18Point52Km ContainmentRadiusAirborneBarometric = 8
	// CRABLargerThan18Point52Km denotes 18.52 km ≤ μ (10.0 NM ≤ μ)
	CRABLargerThan18Point52Km ContainmentRadiusAirborneBarometric = 9
)

// GetHorizontalContainmentLowerBound returns the lower bound of the CR. If the given CR does not have a lower bound returns -1.
// Notes:
//    - The bound is returned in meters.
//    - The bound is inclusive to the value, meaning that bound <= value.
func (cr ContainmentRadiusAirborneBarometric) GetHorizontalContainmentLowerBound() float64 {
	switch cr {
	case CRABBetween3MAnd10M:
		return 3
	case CRABBetween10MAnd92Dot6M:
		return 10
	case CRABBetween92Dot6MAnd185Dot2M:
		return 92.6
	case CRABBetween185Dot2MAnd463M:
		return 185.2
	case CRABBetween463MAnd926M:
		return 463
	case CRABBetween926MAnd1852M:
		return 926
	case CRABBetween1Point852KmAnd9Point26Km:
		return 1852
	case CRABBetween9Point26KmAnd18Point52Km:
		return 9260
	case CRABLargerThan18Point52Km:
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
	case CRABLowerThan3M:
		return 3
	case CRABBetween3MAnd10M:
		return 10
	case CRABBetween10MAnd92Dot6M:
		return 92.6
	case CRABBetween92Dot6MAnd185Dot2M:
		return 185.2
	case CRABBetween185Dot2MAnd463M:
		return 463
	case CRABBetween463MAnd926M:
		return 926
	case CRABBetween926MAnd1852M:
		return 1852
	case CRABBetween1Point852KmAnd9Point26Km:
		return 9260
	case CRABBetween9Point26KmAnd18Point52Km:
		return 18520
	default:
		return -1
	}
}

// ToString returns a basic, but readable, representation of the field
func (cr ContainmentRadiusAirborneBarometric) ToString() string {
	switch cr {
	case CRABLowerThan3M:
		return "μ < 3m"
	case CRABBetween3MAnd10M:
		return "3 m ≤ μ < 10 m"
	case CRABBetween10MAnd92Dot6M:
		return "10 m ≤ μ < 92.6 m (0.05 NM)"
	case CRABBetween92Dot6MAnd185Dot2M:
		return "92.6 m ≤ μ < 185.2 m (0.05 NM ≤ μ < 0.1 NM)"
	case CRABBetween185Dot2MAnd463M:
		return "185.2 m ≤ μ < 463 m (0.1 NM ≤ μ < 0.25 NM)"
	case CRABBetween463MAnd926M:
		return "463 m ≤ μ < 926 m (0.25 NM ≤ μ < 0.5 NM)"
	case CRABBetween926MAnd1852M:
		return "926 m ≤ μ < 1852 m (0.5 NM ≤ μ < 1.0 NM)"
	case CRABBetween1Point852KmAnd9Point26Km:
		return "1.852 km ≤ μ < 9.26 km (1.0 NM ≤ μ < 5.0 NM)"
	case CRABBetween9Point26KmAnd18Point52Km:
		return "9.26 km ≤ μ < 18.52 km (5.0 NM ≤ μ < 10.0 NM)"
	case CRABLargerThan18Point52Km:
		return "18.52 km ≤ μ (10.0 NM ≤ μ)"
	default:
		return fmt.Sprintf("%v - Unknown code", cr)
	}
}
