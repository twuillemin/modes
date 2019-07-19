package fields

import (
	"fmt"
)

// ActiveRAOneThreatOrSameSeparation is one of the possible resolution advisory contained in a MV field
type ActiveRAOneThreatOrSameSeparation struct {
	PreventiveCorrective         ActiveRAPreventiveCorrective
	Sense                        ActiveRASense
	IsIncreasedRate              bool
	IsSenseReversal              bool
	IsAltitudeCrossing           bool
	VerticalSpeedLimitOrPositive ActiveRAVerticalSpeedLimitOrPositive
}

// GetType returns the type of Resolution Advisory
func (oneThreatOrSameSeparation ActiveRAOneThreatOrSameSeparation) GetType() RAType {
	return OneThreatOrSameSeparation
}

// ToString returns a basic, but readable, representation of the field
func (oneThreatOrSameSeparation ActiveRAOneThreatOrSameSeparation) ToString() string {
	return fmt.Sprintf("Type: One threat or multiple threats with separation in same direction\n"+
		"Preventive/Corrective: %v\n"+
		"Sense generated: %v\n"+
		"Is increased rate: %v\n"+
		"Is sense reversal: %v\n"+
		"Is altitude crossing: %v\n"+
		"Vertical Speed Limit/Positive: %v",
		oneThreatOrSameSeparation.PreventiveCorrective.ToString(),
		oneThreatOrSameSeparation.Sense.ToString(),
		oneThreatOrSameSeparation.IsIncreasedRate,
		oneThreatOrSameSeparation.IsSenseReversal,
		oneThreatOrSameSeparation.IsAltitudeCrossing,
		oneThreatOrSameSeparation.VerticalSpeedLimitOrPositive.ToString())
}

// ReadARAOneThreatOrSameSeparation reads the 14 bits data that constitutes the Active Resolution field (ARA)
func ReadARAOneThreatOrSameSeparation(data []byte) *ActiveRAOneThreatOrSameSeparation {

	preventiveCorrective := ActiveRAPreventive
	if (data[0] & 0x40) != 0 {
		preventiveCorrective = ActiveRACorrective
	}

	sense := ActiveRASenseUpward
	if (data[0] & 0x20) != 0 {
		sense = ActiveRASenseDownward
	}

	isIncreasedRate := false
	if (data[0] & 0x10) != 0 {
		isIncreasedRate = true
	}

	isSenseReversal := false
	if (data[0] & 0x08) != 0 {
		isSenseReversal = true
	}

	isAltitudeCrossing := false
	if (data[0] & 0x04) != 0 {
		isAltitudeCrossing = true
	}

	verticalSpeedLimitOrPositive := ActiveRAVerticalSpeedLimit
	if (data[0] & 0x02) != 0 {
		verticalSpeedLimitOrPositive = ActiveRAPositive
	}

	return &ActiveRAOneThreatOrSameSeparation{
		PreventiveCorrective:         preventiveCorrective,
		Sense:                        sense,
		IsIncreasedRate:              isIncreasedRate,
		IsSenseReversal:              isSenseReversal,
		IsAltitudeCrossing:           isAltitudeCrossing,
		VerticalSpeedLimitOrPositive: verticalSpeedLimitOrPositive,
	}
}

// ---------------------------------------------------------------------------------------------
//
//                                    Subfields
//
// ---------------------------------------------------------------------------------------------

// ActiveRAPreventiveCorrective is a subfield of the ActiveRA
type ActiveRAPreventiveCorrective int

const (
	// ActiveRAPreventive signifies the RA is preventive
	ActiveRAPreventive ActiveRAPreventiveCorrective = 0
	// ActiveRACorrective signifies the RA is corrective
	ActiveRACorrective ActiveRAPreventiveCorrective = 1
)

// ToString returns a basic, but readable, representation of the field
func (preventiveCorrective ActiveRAPreventiveCorrective) ToString() string {
	switch preventiveCorrective {
	case ActiveRAPreventive:
		return "Preventive"
	case ActiveRACorrective:
		return "Corrective"
	default:
		return fmt.Sprintf("%v - Unknown code", preventiveCorrective)
	}
}

// ActiveRASense is a subfield of the ActiveRA
type ActiveRASense int

const (
	// ActiveRASenseUpward signifies that the upward sense RA has been generated
	ActiveRASenseUpward ActiveRASense = 0
	// ActiveRASenseDownward signifies that the downward sense RA has been generated
	ActiveRASenseDownward ActiveRASense = 1
)

// ToString returns a basic, but readable, representation of the field
func (sense ActiveRASense) ToString() string {
	switch ActiveRASenseUpward {
	case ActiveRASenseUpward:
		return "Upward"
	case ActiveRASenseDownward:
		return "Downward"
	default:
		return fmt.Sprintf("%v - Unknown code", sense)
	}
}

// ActiveRAVerticalSpeedLimitOrPositive is a subfield of the ActiveRA
type ActiveRAVerticalSpeedLimitOrPositive int

const (
	// ActiveRAVerticalSpeedLimit signifies that the RA is VerticalSpeedLimit
	ActiveRAVerticalSpeedLimit ActiveRAVerticalSpeedLimitOrPositive = 0
	// ActiveRAPositive signifies that the RA is positive
	ActiveRAPositive ActiveRAVerticalSpeedLimitOrPositive = 1
)

// ToString returns a basic, but readable, representation of the field
func (verticalSpeedLimitOrPositive ActiveRAVerticalSpeedLimitOrPositive) ToString() string {
	switch verticalSpeedLimitOrPositive {
	case ActiveRAVerticalSpeedLimit:
		return "Vertical Speed Limit"
	case ActiveRAPositive:
		return "Positive"
	default:
		return fmt.Sprintf("%v - Unknown code", verticalSpeedLimitOrPositive)
	}
}
