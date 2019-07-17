package fields

import (
	"fmt"
)

// ActiveRAMultipleThreatsDifferentSeparation is one of the possible resolution advisory contained in a MV field
type ActiveRAMultipleThreatsDifferentSeparation struct {
	RequiresCorrectionUpwardSense   bool
	RequiresPositiveClimb           bool
	RequiresCorrectionDownwardSense bool
	RequiresPositiveDescend         bool
	RequiresCrossing                bool
	IsSenseReversal                 bool
}

// GetType returns the type of Resolution Advisory
func (multipleThreatsDifferentSeparation ActiveRAMultipleThreatsDifferentSeparation) GetType() RAType {
	return MultipleThreatDifferentSeparation
}

// ToString returns a basic, but readable, representation of the field
func (multipleThreatsDifferentSeparation ActiveRAMultipleThreatsDifferentSeparation) ToString() string {
	return fmt.Sprintf("    Type: Multiple threats with separation in different directions\n"+
		"    Require a correction in the upward sense: %v\n"+
		"    Requires a positive climb: %v\n"+
		"    Requires a correction in the downward sense: %v\n"+
		"    Requires a positive descend: %v\n"+
		"    Requires a crossing: %v\n"+
		"    Is a sense reversal: %v",
		multipleThreatsDifferentSeparation.RequiresCorrectionUpwardSense,
		multipleThreatsDifferentSeparation.RequiresPositiveClimb,
		multipleThreatsDifferentSeparation.RequiresCorrectionDownwardSense,
		multipleThreatsDifferentSeparation.RequiresPositiveDescend,
		multipleThreatsDifferentSeparation.RequiresCrossing,
		multipleThreatsDifferentSeparation.IsSenseReversal)
}

// ReadARAMultipleThreatsDifferentSeparation reads the 14 bits data that constitutes the Active Resolution field (ARA)
//
// Params:
//    - value: the 14 bits of the ARA field. value is right packed in a 16 bit int
//
// Returns an ActiveRAMultipleThreatsDifferentSeparation properly filled
func ReadARAMultipleThreatsDifferentSeparation(value uint16) ActiveResolutionAdvisory {

	requiresCorrectionUpwardSense := false
	if (value & 0x1000) != 0 {
		requiresCorrectionUpwardSense = true
	}

	requiresPositiveClimb := false
	if (value & 0x0800) != 0 {
		requiresPositiveClimb = true
	}

	requiresCorrectionDownwardSense := false
	if (value & 0x0400) != 0 {
		requiresCorrectionDownwardSense = true
	}

	requiresPositiveDescend := false
	if (value & 0x0200) != 0 {
		requiresPositiveDescend = true
	}

	requiresCrossing := false
	if (value & 0x0100) != 0 {
		requiresCrossing = true
	}

	isSenseReversal := false
	if (value & 0x0080) != 0 {
		isSenseReversal = true
	}

	return ActiveRAMultipleThreatsDifferentSeparation{
		RequiresCorrectionUpwardSense:   requiresCorrectionUpwardSense,
		RequiresPositiveClimb:           requiresPositiveClimb,
		RequiresCorrectionDownwardSense: requiresCorrectionDownwardSense,
		RequiresPositiveDescend:         requiresPositiveDescend,
		RequiresCrossing:                requiresCrossing,
		IsSenseReversal:                 isSenseReversal,
	}
}
