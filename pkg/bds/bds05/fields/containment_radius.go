package fields

import "github.com/twuillemin/modes/pkg/common"

// ContainmentRadius is the general definition for the ContainmentRadius
type ContainmentRadius interface {
	common.Printable

	// GetHorizontalContainmentLowerBound returns the lower bound of the CR. If the given CR does not have a lower bound returns -1.
	// Notes:
	//    - The bound is returned in meters.
	//    - The bound is inclusive to the value, meaning that bound <= value.
	GetHorizontalContainmentLowerBound() float64

	// GetHorizontalContainmentHigherBound returns the higher bound of the CR. If the given CR does not have a higher bound returns -1.
	// Notes:
	//    - The bound is returned in meters.
	//    - The bound is exclusive to the value, meaning that value < bound.
	GetHorizontalContainmentHigherBound() float64
}
