package fields

import "github.com/twuillemin/modes/pkg/common"

// Subtype is Subtype of the message
//
// Specified in Doc 9871 / Table B-2-98
type Subtype interface {
	common.Printable

	// ToSubtype returns the subtype itself
	ToSubtype() Subtype
}
