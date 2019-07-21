package fields

import "github.com/twuillemin/modes/pkg/common"

// ContainmentRadius is the general definition for the Containment Radius in ADSB V0
type ContainmentRadius interface {
	common.Printable

	// ToContainmentRadius returns the ContainmentRadius
	ToContainmentRadius() ContainmentRadius
}
