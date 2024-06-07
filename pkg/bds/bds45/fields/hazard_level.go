package fields

import (
	"fmt"
)

// HazardLevel is the level of the hazard
//
// Specified in Doc 9871 / Table A-2-69
type HazardLevel byte

const (
	// HLNil indicates Nil.
	HLNil HazardLevel = 0
	// HLLight indicates Light.
	HLLight HazardLevel = 1
	// HLModerate indicates Moderate.
	HLModerate HazardLevel = 2
	// HLSevere indicates Severe.
	HLSevere HazardLevel = 3
)

// ToString returns a basic, but readable, representation of the field
func (hl HazardLevel) ToString() string {

	switch hl {
	case HLNil:
		return "0 - Nil"
	case HLLight:
		return "1 - Light"
	case HLModerate:
		return "2 - Moderate"
	case HLSevere:
		return "3 - Severe"
	default:
		return fmt.Sprintf("%v - Unknown code", hl)
	}
}
