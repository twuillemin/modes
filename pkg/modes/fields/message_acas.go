package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
)

// -----------------------------------------------------------------------------------------
//
//                                 Message ACAS (MV)
//
// -----------------------------------------------------------------------------------------

// MessageACAS (MV) field shall contain the aircraft address which provides the long ACAS message
//
// Defined at 3.1.2.8.3.1
type MessageACAS struct {
	VDS1                    byte
	VDS2                    byte
	ActiveRA                ResolutionAdvisory
	RAComplement            RAComplement
	RATerminatedIndicator   RATerminatedIndicator
	MultipleThreatEncounter MultipleThreatEncounter
}

// ReadMessageACAS reads the MV field from a message
func ReadMessageACAS(message common.MessageData) MessageACAS {

	// Format of the message is as follow:
	//
	//  VDS1     VDS2  |       ARA       |   ARA       RAC | RAC RAT MTE Res | Reserved
	// v v v v v v v v | a a a a a a a a | a a a a a a c c | c c t m _ _ _ _ | 18 bits

	// Extract the easy values first
	vds1 := message.Payload[4] >> 4
	vds2 := message.Payload[4] & 0x0F

	rac := (message.Payload[6]&0x03)<<2 + (message.Payload[7])>>6
	doNotPassBelow := (rac & 0x80) != 0
	doNotPassAbove := (rac & 0x40) != 0
	doNotTurnLeft := (rac & 0x20) != 0
	doNotTurnRight := (rac & 0x10) != 0

	rat := RATCurrent
	if (message.Payload[7] & 0x20) != 0 {
		rat = RATTerminated
	}

	mte := MTEOneOrZero
	if (message.Payload[7] & 0x10) != 0 {
		mte = MTETwoOrMore
	}

	araFirstBit := (message.Payload[4] & 0x80) >> 7

	var activeRA ResolutionAdvisory
	if araFirstBit == 1 {

		preventiveCorrective := ActiveRAPreventive
		if (message.Payload[4] & 0x40) != 0 {
			preventiveCorrective = ActiveRACorrective
		}

		sense := ActiveRASenseUpward
		if (message.Payload[4] & 0x20) != 0 {
			sense = ActiveRASenseDownward
		}

		isIncreasedRate := false
		if (message.Payload[4] & 0x10) != 0 {
			isIncreasedRate = true
		}

		isSenseReversal := false
		if (message.Payload[4] & 0x08) != 0 {
			isSenseReversal = true
		}

		isAltitudeCrossing := false
		if (message.Payload[4] & 0x04) != 0 {
			isAltitudeCrossing = true
		}

		verticalSpeedLimitOrPositive := ActiveRAVerticalSpeedLimit
		if (message.Payload[4] & 0x20) != 0 {
			verticalSpeedLimitOrPositive = ActiveRAPositive
		}

		activeRA = ActiveRAOneThreatOrSameSeparation{
			PreventiveCorrective:         preventiveCorrective,
			Sense:                        sense,
			IsIncreasedRate:              isIncreasedRate,
			IsSenseReversal:              isSenseReversal,
			IsAltitudeCrossing:           isAltitudeCrossing,
			VerticalSpeedLimitOrPositive: verticalSpeedLimitOrPositive,
		}
	} else if mte == MTETwoOrMore {

		requiresCorrectionUpwardSense := false
		if (message.Payload[4] & 0x40) != 0 {
			requiresCorrectionUpwardSense = true
		}

		requiresPositiveClimb := false
		if (message.Payload[4] & 0x20) != 0 {
			requiresPositiveClimb = true
		}

		requiresCorrectionDownwardSense := false
		if (message.Payload[4] & 0x10) != 0 {
			requiresCorrectionDownwardSense = true
		}

		requiresPositiveDescend := false
		if (message.Payload[4] & 0x08) != 0 {
			requiresPositiveDescend = true
		}

		requiresCrossing := false
		if (message.Payload[4] & 0x04) != 0 {
			requiresCrossing = true
		}

		isSenseReversal := false
		if (message.Payload[4] & 0x02) != 0 {
			isSenseReversal = true
		}

		activeRA = ActiveRAMultipleThreatsDifferentSeparation{
			RequiresCorrectionUpwardSense:   requiresCorrectionUpwardSense,
			RequiresPositiveClimb:           requiresPositiveClimb,
			RequiresCorrectionDownwardSense: requiresCorrectionDownwardSense,
			RequiresPositiveDescend:         requiresPositiveDescend,
			RequiresCrossing:                requiresCrossing,
			IsSenseReversal:                 isSenseReversal,
		}
	} else {
		activeRA = ActiveRANoVerticalRAGenerated{}
	}

	return MessageACAS{
		VDS1:     vds1,
		VDS2:     vds2,
		ActiveRA: activeRA,
		RAComplement: RAComplement{
			DoNotPassBelow: doNotPassBelow,
			DoNotPassAbove: doNotPassAbove,
			DoNotTurnLeft:  doNotTurnLeft,
			DoNotTurnRight: doNotTurnRight,
		},
		RATerminatedIndicator:   rat,
		MultipleThreatEncounter: mte,
	}
}

// ToString returns a basic, but readable, representation of the field
func (messageACAS MessageACAS) ToString() string {
	return fmt.Sprintf("Message ACAS VDS 1: %X\n"+
		"Message ACAS VDS 2: %X\n"+
		"Message ACAS Active Resolution Advisory :\n %v\n"+
		"Message ACAS Active Resolution Advisory Complement: %v\n"+
		"Message ACAS Active Resolution Advisory Terminated Indicator: %v\n"+
		"Message ACAS Active Multiple Threat Encounter: %v",
		messageACAS.VDS1,
		messageACAS.VDS2,
		messageACAS.ActiveRA.ToString(),
		messageACAS.RAComplement.ToString(),
		messageACAS.RATerminatedIndicator.ToString(),
		messageACAS.MultipleThreatEncounter.ToString())
}

// ActiveRAPreventiveCorrective is part of the MV fields, subfield ActiveRA
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

// ActiveRASense is part of the MV fields, subfield ActiveRA
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

// ActiveRAVerticalSpeedLimitOrPositive is part of the MV fields, subfield ActiveRA
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

// RAType defines the type of the Resolution Advisory
type RAType int

const (
	// OneThreatOrSameSeparation indicates that the Resolution Advisory is for single threat or multiple threats with
	// the same separation
	OneThreatOrSameSeparation RAType = 0
	// MultipleThreatDifferentSeparation indicates that the Resolution Advisory is for multiple threats with
	// different separation
	MultipleThreatDifferentSeparation RAType = 1
	// NoVerticalRAGenerated indicates that no Resolution Advisory was generated
	NoVerticalRAGenerated RAType = 2
)

// ResolutionAdvisory is the base type that all Resolution Advisory should implement
type ResolutionAdvisory interface {
	// GetType returns the type of Resolution Advisory
	GetType() RAType
	// ToString returns a basic, but readable, representation of the field
	ToString() string
}

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
	return fmt.Sprintf("Resolution Advisory type: One threat or multiple threats with separation in same direction\n"+
		"Resolution Advisory Preventive/Corrective: %v\n"+
		"Resolution Advisory sense generated: %v\n"+
		"Resolution Advisory is increased rate: %v\n"+
		"Resolution Advisory is sense reversal: %v\n"+
		"Resolution Advisory is altitude crossing: %v\n"+
		"Resolution Advisory Vertical Speed Limit/Positive: %v",
		oneThreatOrSameSeparation.PreventiveCorrective,
		oneThreatOrSameSeparation.Sense,
		oneThreatOrSameSeparation.IsIncreasedRate,
		oneThreatOrSameSeparation.IsSenseReversal,
		oneThreatOrSameSeparation.IsAltitudeCrossing,
		oneThreatOrSameSeparation.VerticalSpeedLimitOrPositive)
}

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
	return fmt.Sprintf("Resolution Advisory type: Multiple threats with separation in different directions\n"+
		"Resolution Advisory require a correction in the upward sense: %v\n"+
		"Resolution Advisory requires a positive climb: %v\n"+
		"Resolution Advisory requires a correction in the downward sense: %v\n"+
		"Resolution Advisory requires a positive descend: %v\n"+
		"Resolution Advisory requires a crossing: %v\n"+
		"Resolution Advisory is a sense reversal: %v",
		multipleThreatsDifferentSeparation.RequiresCorrectionUpwardSense,
		multipleThreatsDifferentSeparation.RequiresPositiveClimb,
		multipleThreatsDifferentSeparation.RequiresCorrectionDownwardSense,
		multipleThreatsDifferentSeparation.RequiresPositiveDescend,
		multipleThreatsDifferentSeparation.RequiresCrossing,
		multipleThreatsDifferentSeparation.IsSenseReversal)
}

// ActiveRANoVerticalRAGenerated is one of the possible resolution advisory contained in a MV field
type ActiveRANoVerticalRAGenerated struct{}

// GetType returns the type of Resolution Advisory
func (noVerticalRAGenerated ActiveRANoVerticalRAGenerated) GetType() RAType {
	return NoVerticalRAGenerated
}

// ToString returns a basic, but readable, representation of the field
func (noVerticalRAGenerated ActiveRANoVerticalRAGenerated) ToString() string {
	return fmt.Sprintf("Resolution Advisory type: no Resolution Advisory has been generated")
}

// RAComplement indicates all the currently active RACs, if any, received from other ACAS aircraft.
type RAComplement struct {
	// DoNotPassBelow signifies that the other aircraft ask to not pass below
	DoNotPassBelow bool
	// DoNotPassAbove signifies that the other aircraft ask to not pass above
	DoNotPassAbove bool
	// DoNotTurnLeft signifies that the other aircraft ask to not turn left
	DoNotTurnLeft bool
	// DoNotTurnRight signifies that the other aircraft ask to not turn right
	DoNotTurnRight bool
}

// ToString returns a basic, but readable, representation of the field
func (complement RAComplement) ToString() string {
	return fmt.Sprintf("Resolution Advisory Complement: Do not pass below: %v\n"+
		"Resolution Advisory Complement: Do not pass above: %v\n"+
		"Resolution Advisory Complement: Do not turn left: %v\n"+
		"Resolution Advisory Complement: Do not turn right: %v",
		complement.DoNotPassBelow,
		complement.DoNotPassAbove,
		complement.DoNotTurnLeft,
		complement.DoNotTurnRight)
}

// RATerminatedIndicator indicates when an RA previously generated by ACAS has ceased being generated.
type RATerminatedIndicator int

const (
	// RATCurrent signifies that ACAS is currently generating the RA indicated in the ARA subfield
	RATCurrent RATerminatedIndicator = 0
	// RATTerminated signifies that the RA indicated by the ARA subfield has been terminated
	RATTerminated RATerminatedIndicator = 1
)

// ToString returns a basic, but readable, representation of the field
func (terminatedIndicator RATerminatedIndicator) ToString() string {
	switch terminatedIndicator {
	case RATCurrent:
		return "ACAS is currently generating the RA"
	case RATTerminated:
		return "the RA indicated has been terminated"
	default:
		return fmt.Sprintf("%v - Unknown code", terminatedIndicator)
	}
}

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
		return "ACAS is currently generating the RA"
	case MTETwoOrMore:
		return "the RA indicated has been terminated"
	default:
		return fmt.Sprintf("%v - Unknown code", multipleThreatEncounter)
	}
}
