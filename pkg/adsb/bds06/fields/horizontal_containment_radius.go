package fields

import "github.com/twuillemin/modes/pkg/common"

// HorizontalContainmentRadius is the general definition for the HorizontalContainmentRadius
type HorizontalContainmentRadius interface {
	common.Printable

	// ToHorizontalContainmentRadius returns the HorizontalContainmentRadius
	ToHorizontalContainmentRadius() HorizontalContainmentRadius
}
