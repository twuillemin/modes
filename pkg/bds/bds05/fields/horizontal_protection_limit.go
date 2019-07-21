package fields

import "github.com/twuillemin/modes/pkg/common"

// HorizontalProtectionLimit is the general definition for the Horizontal Protection Limit in ADSB V0
type HorizontalProtectionLimit interface {
	common.Printable

	// ToHorizontalProtectionLimit returns the HorizontalProtectionLimit
	ToHorizontalProtectionLimit() HorizontalProtectionLimit
}
