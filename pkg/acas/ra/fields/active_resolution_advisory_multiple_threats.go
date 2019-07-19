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
	return fmt.Sprintf("Type: Multiple threats with separation in different directions\n"+
		"Require a correction in the upward sense: %v\n"+
		"Requires a positive climb: %v\n"+
		"Requires a correction in the downward sense: %v\n"+
		"Requires a positive descend: %v\n"+
		"Requires a crossing: %v\n"+
		"Is a sense reversal: %v",
		multipleThreatsDifferentSeparation.RequiresCorrectionUpwardSense,
		multipleThreatsDifferentSeparation.RequiresPositiveClimb,
		multipleThreatsDifferentSeparation.RequiresCorrectionDownwardSense,
		multipleThreatsDifferentSeparation.RequiresPositiveDescend,
		multipleThreatsDifferentSeparation.RequiresCrossing,
		multipleThreatsDifferentSeparation.IsSenseReversal)
}

// ReadARAMultipleThreatsDifferentSeparation reads the 14 bits data that constitutes the Active Resolution field (ARA)
func ReadARAMultipleThreatsDifferentSeparation(data []byte) *ActiveRAMultipleThreatsDifferentSeparation {

	requiresCorrectionUpwardSense := false
	if (data[0] & 0x40) != 0 {
		requiresCorrectionUpwardSense = true
	}

	requiresPositiveClimb := false
	if (data[0] & 0x20) != 0 {
		requiresPositiveClimb = true
	}

	requiresCorrectionDownwardSense := false
	if (data[0] & 0x10) != 0 {
		requiresCorrectionDownwardSense = true
	}

	requiresPositiveDescend := false
	if (data[0] & 0x08) != 0 {
		requiresPositiveDescend = true
	}

	requiresCrossing := false
	if (data[0] & 0x04) != 0 {
		requiresCrossing = true
	}

	isSenseReversal := false
	if (data[0] & 0x02) != 0 {
		isSenseReversal = true
	}

	return &ActiveRAMultipleThreatsDifferentSeparation{
		RequiresCorrectionUpwardSense:   requiresCorrectionUpwardSense,
		RequiresPositiveClimb:           requiresPositiveClimb,
		RequiresCorrectionDownwardSense: requiresCorrectionDownwardSense,
		RequiresPositiveDescend:         requiresPositiveDescend,
		RequiresCrossing:                requiresCrossing,
		IsSenseReversal:                 isSenseReversal,
	}
}
