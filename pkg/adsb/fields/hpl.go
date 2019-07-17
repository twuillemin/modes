package fields

import "github.com/twuillemin/modes/pkg/common"

// HPL is the general definition for the HPL (Horizontal Protection Limit)
type HPL interface {
	common.Printable

	// GetHPLLowerBound returns the lower bound of the HPL. If the given HPL does not have a lower bound returns -1.
	// Notes:
	//    - The bound is returned in meters.
	//    - The bound is inclusive to the value, meaning that bound <= value.
	GetHPLLowerBound() float64

	// GetHPLHigherBound returns the higher bound of the HPL. If the given HPL does not have a higher bound returns -1.
	// Notes:
	//    - The bound is returned in meters.
	//    - The bound is exclusive to the value, meaning that value < bound.
	GetHPLHigherBound() float64
}
