package messages

import (
	"github.com/twuillemin/modes/pkg/acas/ra/fields"
	"testing"
)

func TestReadResolutionAdvisoryMultipleThreatSameSeparationValid(t *testing.T) {

	msg, err := ParseResolutionAdvisory(buildValidResolutionAdvisoryMultipleThreatSameSeparationMessage())
	if err != nil {
		t.Fatal(err)
	}

	if msg.ActiveRA.GetType() != fields.OneThreatOrSameSeparation {
		t.Fatalf("Expected Type \"OneThreatOrSameSeparation\", got \"%v\"",
			msg.ActiveRA.GetType())
	}

	resolutionAdvisory, ok := msg.ActiveRA.(*fields.ActiveRAOneThreatOrSameSeparation)
	if ok != true {
		t.Errorf("Expected format to be ActiveRAOneThreatOrSameSeparation, but it was not possible to convert")
	}

	if resolutionAdvisory.PreventiveCorrective != fields.ActiveRACorrective {
		t.Errorf("Expected PreventiveCorrective \"%v\", got \"%v\"",
			fields.ActiveRACorrective.ToString(),
			resolutionAdvisory.PreventiveCorrective.ToString())
	}

	if resolutionAdvisory.Sense != fields.ActiveRASenseDownward {
		t.Errorf("Expected Sense \"%v\", got \"%v\"",
			fields.ActiveRASenseDownward.ToString(),
			resolutionAdvisory.Sense.ToString())
	}

	if resolutionAdvisory.IsIncreasedRate != true {
		t.Errorf("Expected IsIncreasedRate \"true\", got \"%v\"", resolutionAdvisory.IsIncreasedRate)
	}

	if resolutionAdvisory.IsSenseReversal != true {
		t.Errorf("Expected IsSenseReversal \"true\", got \"%v\"", resolutionAdvisory.IsSenseReversal)
	}

	if resolutionAdvisory.IsAltitudeCrossing != true {
		t.Errorf("Expected IsAltitudeCrossing \"true\", got \"%v\"", resolutionAdvisory.IsAltitudeCrossing)
	}

	if resolutionAdvisory.VerticalSpeedLimitOrPositive != fields.ActiveRAPositive {
		t.Errorf("Expected IsAltitudeCrossing \"%v\", got \"%v\"",
			fields.ActiveRAPositive,
			resolutionAdvisory.VerticalSpeedLimitOrPositive.ToString())
	}

	if msg.RAComplement.DoNotPassBelow != true {
		t.Errorf("Expected DoNotPassBelow \"true\", got \"%v\"", msg.RAComplement.DoNotPassBelow)
	}

	if msg.RAComplement.DoNotPassAbove != true {
		t.Errorf("Expected DoNotPassAbove \"true\", got \"%v\"", msg.RAComplement.DoNotPassAbove)
	}

	if msg.RAComplement.DoNotTurnLeft != true {
		t.Errorf("Expected DoNotTurnLeft \"true\", got \"%v\"", msg.RAComplement.DoNotTurnLeft)
	}

	if msg.RAComplement.DoNotTurnRight != true {
		t.Errorf("Expected DoNotTurnRight \"true\", got \"%v\"", msg.RAComplement.DoNotTurnRight)
	}

	if msg.MultipleThreatEncounter != fields.MTETwoOrMore {
		t.Errorf("Expected MultipleThreatEncounter \"%v\", got \"%v\"",
			fields.MTETwoOrMore.ToString(),
			msg.MultipleThreatEncounter.ToString())
	}

	if msg.ThreatTypeIndicator != fields.ThreatTypeModeS {
		t.Errorf("Expected ThreatTypeIndicator \"%v\", got \"%v\"",
			fields.ThreatTypeModeS.ToString(),
			msg.ThreatTypeIndicator.ToString())
	}

	address := msg.ThreatIdentityAddress
	if address == nil {
		t.Errorf("Expected ThreatIdentityAddress \"not nil\", got \"nil\"")
	} else if *address != 0x005A697D {
		t.Errorf("Expected ThreatIdentityAddress to be \"0x005A697D = 5925245\", got \"%v\"", *address)
	}

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func buildValidResolutionAdvisoryMultipleThreatSameSeparationMessage() []byte {

	// Format of the message is as follow:
	//        0                 1                 2                 3                 4                 5
	//                 |             RAC |  R  R M  T  TID |       TID       |       TID       |       TID       |
	//       ARA       |   ARA       RAC |  A  A T  T  d d | d d d d d d d d | d d d d d d d d | d d d d d d _ _ |
	// ^                 ^           ^         ^ ^ ^   ^
	// 41  	            49          55       59 60 61  63

	data := make([]byte, 6)

	// First bit 1 OneThreatOrSameSeparation
	data[0] |= 0x80
	// RA Corrective
	data[0] |= 0x40
	// Downward sense
	data[0] |= 0x20
	// Increasing rate
	data[0] |= 0x10
	// Sense Reversal
	data[0] |= 0x08
	// Altitude crossing
	data[0] |= 0x04
	// Positive
	data[0] |= 0x02

	// RAC do not pass below, above, left or right
	data[1] |= 0x02
	data[1] |= 0x01
	data[2] |= 0x80
	data[2] |= 0x40

	// RAT : Terminated
	data[2] |= 0x20

	// MTE: Yes
	data[2] |= 0x10

	// TTI Mode-S
	data[2] |= 0x04

	// 01011010: 5A 01101001: 69 01111101: 7D
	// Mode-S: 01 | 0110 1001 | 1010 0101 | 1111 01

	data[2] |= 0x01
	data[3] |= 0x69
	data[4] |= 0xA5
	data[5] |= 0xF4

	return data
}

func TestReadResolutionAdvisoryMultipleThreatDifferentSeparationValid(t *testing.T) {

	msg, err := ParseResolutionAdvisory(buildValidResolutionAdvisoryMultipleThreatDifferentSeparation())
	if err != nil {
		t.Fatal(err)
	}

	if msg.ActiveRA.GetType() != fields.MultipleThreatDifferentSeparation {
		t.Fatalf("Expected Type \"MultipleThreatDifferentSeparation\", got \"%v\"",
			msg.ActiveRA.GetType())
	}

	resolutionAdvisory, ok := msg.ActiveRA.(*fields.ActiveRAMultipleThreatsDifferentSeparation)
	if ok != true {
		t.Errorf("Expected format to be ActiveRAMultipleThreatsDifferentSeparation, but it was not possible to convert")
	}

	if resolutionAdvisory.RequiresCorrectionUpwardSense != true {
		t.Errorf("Expected RequiresCorrectionUpwardSense \"true\", got \"%v\"", resolutionAdvisory.RequiresCorrectionUpwardSense)
	}

	if resolutionAdvisory.RequiresPositiveClimb != true {
		t.Errorf("Expected RequiresPositiveClimb \"true\", got \"%v\"", resolutionAdvisory.RequiresPositiveClimb)
	}

	if resolutionAdvisory.RequiresCorrectionDownwardSense != true {
		t.Errorf("Expected RequiresCorrectionDownwardSense \"true\", got \"%v\"", resolutionAdvisory.RequiresCorrectionDownwardSense)
	}

	if resolutionAdvisory.RequiresPositiveDescend != true {
		t.Errorf("Expected RequiresPositiveDescend \"true\", got \"%v\"", resolutionAdvisory.RequiresPositiveDescend)
	}

	if resolutionAdvisory.RequiresCrossing != true {
		t.Errorf("Expected RequiresCrossing \"true\", got \"%v\"", resolutionAdvisory.RequiresCrossing)
	}

	if resolutionAdvisory.IsSenseReversal != true {
		t.Errorf("Expected IsSenseReversal \"true\", got \"%v\"", resolutionAdvisory.IsSenseReversal)
	}

	if msg.RAComplement.DoNotPassBelow != false {
		t.Errorf("Expected DoNotPassBelow \"true\", got \"%v\"", msg.RAComplement.DoNotPassBelow)
	}

	if msg.RAComplement.DoNotPassAbove != true {
		t.Errorf("Expected DoNotPassAbove \"true\", got \"%v\"", msg.RAComplement.DoNotPassAbove)
	}

	if msg.RAComplement.DoNotTurnLeft != false {
		t.Errorf("Expected DoNotTurnLeft \"true\", got \"%v\"", msg.RAComplement.DoNotTurnLeft)
	}

	if msg.RAComplement.DoNotTurnRight != true {
		t.Errorf("Expected DoNotTurnRight \"true\", got \"%v\"", msg.RAComplement.DoNotTurnRight)
	}

	if msg.MultipleThreatEncounter != fields.MTETwoOrMore {
		t.Errorf("Expected MultipleThreatEncounter \"%v\", got \"%v\"",
			fields.MTETwoOrMore.ToString(),
			msg.MultipleThreatEncounter.ToString())
	}

	if msg.ThreatTypeIndicator != fields.ThreatTypeAltitudeRangeBearing {
		t.Errorf("Expected ThreatTypeIndicator \"%v\", got \"%v\"",
			fields.ThreatTypeAltitudeRangeBearing.ToString(),
			msg.ThreatTypeIndicator.ToString())
	}

	threatAltitude := msg.ThreatIdentityAltitude
	if threatAltitude == nil {
		t.Errorf("Expected ThreatIdentityAltitude \"not nil\", got \"nil\"")
	} else if (*threatAltitude).AltitudeInFeet != 0 {
		t.Errorf("Expected ThreatIdentityAltitude to be \"0\", got \"%v\"", (*threatAltitude).AltitudeInFeet)
	}

	threatRange := msg.ThreatIdentityRange
	if threatRange == nil {
		t.Errorf("Expected ThreatIdentityRange \"not nil\", got \"nil\"")
	} else if threatRange.GetRange() != 8.5 {
		t.Errorf("Expected ThreatIdentityRange to be \"8.5\", got \"%v\"", threatRange.GetRange())
	}

	threatBearing := msg.ThreatIdentityBearing
	if threatBearing == nil {
		t.Errorf("Expected ThreatIdentityBearing \"not nil\", got \"nil\"")
	} else if threatBearing.GetBearing() != 249 {
		t.Errorf("Expected ThreatIdentityBearing to be \"249\", got \"%v\"", threatBearing.GetBearing())
	}

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func buildValidResolutionAdvisoryMultipleThreatDifferentSeparation() []byte {

	// Format of the message is as follow:
	//        0                 1                 2                 3                 4                 5
	//                 |             RAC |  C  T E  I  TIDA|       TIDA      | TIDA    TIDR    |TIDR     TIDB    |
	// a a a a a a a a | a a a a a a c c | c c t m i i a a | a a a a a a a a | a a a r r r r r | r r b b b b b b |
	// ^                 ^           ^         ^ ^ ^   ^
	// 41  	            49          55       59 60 61  63

	data := make([]byte, 6)

	// First bit 1 OneThreatOrSameSeparation: false
	//--data[0] |= 0x80

	// RequiresCorrectionUpwardSense: true
	data[0] |= 0x40
	// RequiresPositiveClimb: true
	data[0] |= 0x20
	// RequiresCorrectionDownwardSense: true
	data[0] |= 0x10
	// RequiresPositiveDescend: true
	data[0] |= 0x08
	// RequiresCrossing: true
	data[0] |= 0x04
	// IsSenseReversal: true
	data[0] |= 0x02

	// RAC do not pass below: No, above: Yes, left: No or right: Yes
	//--data[1] |= 0x02
	data[1] |= 0x01
	//--data[2] |= 0x80
	data[2] |= 0x40

	// RAT : Terminated
	data[2] |= 0x20

	// MTE: true
	data[2] |= 0x10

	// TTI 2 Altitude Range Bearing
	data[2] |= 0x08

	// Altitude :0M => __00 | 1000 0001 | 010_ ___
	// Set C2
	data[3] = data[3] | 0x80
	// Set B2
	data[3] = data[3] | 0x01
	// Set B4
	data[4] = data[4] | 0x40

	// Range
	// 8.5NM => value: 86 => | ___1 0101 | 10__
	data[4] = data[4] | 0x10
	data[4] = data[4] | 0x04
	data[4] = data[4] | 0x01
	data[5] = data[5] | 0x80

	// Bearing
	// 249 deg => value 42 => | __10 1010 |
	data[5] = data[5] | 0x20
	data[5] = data[5] | 0x08
	data[5] = data[5] | 0x02

	return data
}
