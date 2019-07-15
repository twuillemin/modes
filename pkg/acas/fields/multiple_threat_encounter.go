package fields

import (
	"fmt"
)

// MultipleThreatEncounter indicates whether two or more simultaneous threats are currently being processed by the ACAS
// threat resolution logic.
type MultipleThreatEncounter int

const (
	// MTEOneOrZero signifies that one or no threat is being processed
	MTEOneOrZero MultipleThreatEncounter = 0
	// MTETwoOrMore signifies that two or more simultaneous threats are being processed
	MTETwoOrMore MultipleThreatEncounter = 1
)

// ToString returns a basic, but readable, representation of the field
func (multipleThreatEncounter MultipleThreatEncounter) ToString() string {
	switch multipleThreatEncounter {
	case MTEOneOrZero:
		return "One or no threat is currently being processed by the ACAS"
	case MTETwoOrMore:
		return "Two or more threats are currently being processed by the ACAS"
	default:
		return fmt.Sprintf("%v - Unknown code", multipleThreatEncounter)
	}
}

// ReadMultipleThreatEncounter reads the bit data that constitutes the MultipleThreatEncounter field (MTE)
//
// Params:
//    - value: the single bit of the MTE field. value is right packed in a 8 bits int.
//
// Returns an MultipleThreatEncounter properly filled
func ReadMultipleThreatEncounter(value byte) MultipleThreatEncounter {

	mte := MTEOneOrZero
	if (value & 0x01) != 0 {
		mte = MTETwoOrMore
	}

	return mte
}
