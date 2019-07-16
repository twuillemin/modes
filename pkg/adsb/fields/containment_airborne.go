package fields

import "fmt"

// ContainmentRadiusAirborne is the 95% Containment radius,μ and v, on horizontal and vertical position error for
// Airborne aircraft format definition
//
// Specified in Doc 9871 / A.2.3.1
type ContainmentRadiusAirborne byte

const (
	// CRALowerThan3M denotes μ < 3m
	CRALowerThan3M ContainmentRadiusAirborne = 0
	// CRABetween3MAnd10M denotes 3 m ≤ μ < 10 m
	CRABetween3MAnd10M ContainmentRadiusAirborne = 1
	// CRABetween10MAnd92Dot6M denotes 10 m ≤ μ < 92.6 m (0.05 NM)
	CRABetween10MAnd92Dot6M ContainmentRadiusAirborne = 2
	// CRABetween92Dot6MAnd185Dot2M denotes 92.6 m ≤ μ < 185.2 m (0.05 NM ≤ μ < 0.1 NM)
	CRABetween92Dot6MAnd185Dot2M ContainmentRadiusAirborne = 3
	// CRABetween185Dot2MAnd463M denotes 185.2 m ≤ μ < 463 m (0.1 NM ≤ μ < 0.25 NM)
	CRABetween185Dot2MAnd463M ContainmentRadiusAirborne = 4
	// CRABetween463MAnd926M denotes 463 m ≤ μ < 926 m (0.25 NM ≤ μ < 0.5 NM)
	CRABetween463MAnd926M ContainmentRadiusAirborne = 5
	// CRABetween926MAnd1852M denotes 926 m ≤ μ < 1852 m (0.5 NM ≤ μ < 1.0 NM)
	CRABetween926MAnd1852M ContainmentRadiusAirborne = 6
	// CRABetween1Point852KmAnd9Point26Km denotes 1.852 km ≤ μ < 9.26 km (1.0 NM ≤ μ < 5.0 NM)
	CRABetween1Point852KmAnd9Point26Km ContainmentRadiusAirborne = 7
	// CRABetween9Point26KmAnd18Point52 denotes 9.26 km ≤ μ < 18.52 km (5.0 NM ≤ μ < 10.0 NM)
	CRABetween9Point26KmAnd18Point52 ContainmentRadiusAirborne = 8
	// CRALargerThan18Point52 denotes 18.52 km ≤ μ (10.0 NM ≤ μ)
	CRALargerThan18Point52 ContainmentRadiusAirborne = 9
)

// ToString returns a basic, but readable, representation of the field
func (cr ContainmentRadiusAirborne) ToString() string {
	switch cr {
	case CRALowerThan3M:
		return "μ < 3m"
	case CRABetween3MAnd10M:
		return "3 m ≤ μ < 10 m"
	case CRABetween10MAnd92Dot6M:
		return "10 m ≤ μ < 92.6 m (0.05 NM)"
	case CRABetween92Dot6MAnd185Dot2M:
		return "92.6 m ≤ μ < 185.2 m (0.05 NM ≤ μ < 0.1 NM)"
	case CRABetween185Dot2MAnd463M:
		return "185.2 m ≤ μ < 463 m (0.1 NM ≤ μ < 0.25 NM)"
	case CRABetween463MAnd926M:
		return "463 m ≤ μ < 926 m (0.25 NM ≤ μ < 0.5 NM)"
	case CRABetween926MAnd1852M:
		return "926 m ≤ μ < 1852 m (0.5 NM ≤ μ < 1.0 NM)"
	case CRABetween1Point852KmAnd9Point26Km:
		return "1.852 km ≤ μ < 9.26 km (1.0 NM ≤ μ < 5.0 NM)"
	case CRABetween9Point26KmAnd18Point52:
		return "9.26 km ≤ μ < 18.52 km (5.0 NM ≤ μ < 10.0 NM)"
	case CRALargerThan18Point52:
		return "18.52 km ≤ μ (10.0 NM ≤ μ)"
	default:
		return fmt.Sprintf("%v - Unknown code", cr)
	}
}
