package messages

import (
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/bds62/fields"
	"testing"
)

func TestReadFormat29Subtype0Valid(t *testing.T) {

	msg, err := ReadFormat29Subtype0(buildValidFormat29Subtype0Message())
	if err != nil {
		t.Fatal(err)
	}

	if msg.GetMessageFormat() != adsb.Format29 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format29.ToString(),
			msg.GetMessageFormat().ToString())
	}

	if msg.GetSubtype() != fields.Subtype0 {
		t.Errorf("Expected Subtype \"%v\", got \"%v\"",
			fields.Subtype0.ToString(),
			msg.GetSubtype().ToString())
	}

	if msg.VerticalDataAvailableSourceIndicator != fields.VDAFMS {
		t.Errorf("Expected Vertical Data Available / Source Indicator \"%v\", got \"%v\"",
			fields.VDAFMS.ToString(),
			msg.VerticalDataAvailableSourceIndicator.ToString())
	}

	if msg.TargetAltitudeType != fields.TATReferencedToBarometricAltitude {
		t.Errorf("Expected Target Altitude Type \"%v\", got \"%v\"",
			fields.TATReferencedToBarometricAltitude.ToString(),
			msg.TargetAltitudeType.ToString())
	}

	if msg.TargetAltitudeCapability != fields.TACAltitudeOrAutopilotOrFMS {
		t.Errorf("Expected Target Altitude Capability \"%v\", got \"%v\"",
			fields.TACAltitudeOrAutopilotOrFMS.ToString(),
			msg.TargetAltitudeCapability.ToString())
	}

	if msg.VerticalModeIndicator != fields.VMICapturingMode {
		t.Errorf("Expected Vertical Mode Indicator \"%v\", got \"%v\"",
			fields.VMICapturingMode.ToString(),
			msg.VerticalModeIndicator.ToString())
	}

	if msg.TargetAltitude.GetTargetAltitude() != 67300 {
		t.Errorf("Expected Target Altitude to be 67300, got \"%v\"",
			msg.VerticalModeIndicator.ToString())
	}

	if msg.HorizontalDataAvailableSourceIndicator != fields.HDAFMS {
		t.Errorf("Expected Horizontal Data Available Source Indicator \"%v\", got \"%v\"",
			fields.HDAFMS.ToString(),
			msg.HorizontalDataAvailableSourceIndicator.ToString())
	}

	if msg.TargetHeadingTrackAngle.GetTargetHeadingTrackAngle() != 359 {
		t.Errorf("Expected Target Heading Track Angle to be 359, got \"%v\"",
			msg.TargetHeadingTrackAngle)
	}

	if msg.TargetHeadingTrackIndicator != fields.THTITrack {
		t.Errorf("Expected Target Heading Track Indicator \"%v\", got \"%v\"",
			fields.THTITrack.ToString(),
			msg.TargetHeadingTrackIndicator.ToString())
	}

	if msg.HorizontalModeIndicator != fields.HMICapturingMode {
		t.Errorf("Expected Horizontal Mode Indicator \"%v\", got \"%v\"",
			fields.HMICapturingMode.ToString(),
			msg.HorizontalModeIndicator.ToString())
	}

	if msg.NavigationalAccuracyCategoryPosition != fields.NACPV1EPULowerThan10MAndVEPULowerThan15M {
		t.Errorf("Expected Navigational Accuracy Category Position \"%v\", got \"%v\"",
			fields.NACPV1EPULowerThan10MAndVEPULowerThan15M.ToString(),
			msg.NavigationalAccuracyCategoryPosition.ToString())
	}

	if msg.NICBaro != fields.NICBGilhamCrossCheckedOrNonGilham {
		t.Errorf("Expected NICBaro \"%v\", got \"%v\"",
			fields.NICBGilhamCrossCheckedOrNonGilham.ToString(),
			msg.NICBaro.ToString())
	}

	if msg.SurveillanceIntegrityLevel != fields.SUILLevel3 {
		t.Errorf("Expected Surveillance Integrity ReaderLevel \"%v\", got \"%v\"",
			fields.SUILLevel3.ToString(),
			msg.SurveillanceIntegrityLevel.ToString())
	}

	if msg.CapabilityModeCode != fields.CMCNotOperationalRAActive {
		t.Errorf("Expected Capability Mode Code \"%v\", got \"%v\"",
			fields.CMCNotOperationalRAActive.ToString(),
			msg.CapabilityModeCode.ToString())
	}

	if msg.EmergencyPriorityStatus != fields.EPSDownedAircraft {
		t.Errorf("Expected EmergencyPriorityStatus \"%v\", got \"%v\"",
			fields.EPSDownedAircraft.ToString(),
			msg.EmergencyPriorityStatus.ToString())
	}

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func TestReadFormat29Subtype0TooShort(t *testing.T) {

	// Shorten the data
	data := buildValidFormat29Subtype0Message()[:6]

	_, err := ReadFormat29Subtype0(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat29Subtype0BadFormat(t *testing.T) {

	// Change code to 2
	data := buildValidFormat29Subtype0Message()
	data[0] = (data[0] & 0x07) | 0x10

	_, err := ReadFormat29Subtype0(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat29Subtype0BadSubtype(t *testing.T) {

	// Change subtype to 2
	data := buildValidFormat29Subtype0Message()
	data[0] = data[0] | 0x04

	_, err := ReadFormat29Subtype0(data)
	if err == nil {
		t.Error(err)
	}
}

func buildValidFormat29Subtype0Message() []byte {
	data := make([]byte, 7)

	// 1110 1001: code 29 (11101) + subtype 0 (00) + VerticalData FMS (1[1])
	data[0] = 0xE9

	// 1101 0101: VerticalData: FMS([1]1) + Target Altitude Type: mean sea level (1) + Reserved (0) + Target Altitude
	// Capability: FMS (10) + Vertical Mode: Capturing (10) + Target altitude: 67300 (1 [0101 0101 1])
	data[1] = 0xD5

	// 0101 0101: Target altitude ([1] 0101 0101 [1])
	data[2] = 0x55

	// 1111 0110: Target altitude ([1 0101 0101] 1) + HorizontalData: FMS (11) + TargetHeading: 359(1 0110 [0111])
	data[3] = 0xF6

	// 0111 1101:  TargetHeading: 359([1 0110] 0111) + TrackIndicator: angle (1) + Horizontal Indicator:Capturing (10)
	// + NACp: EPU < 10 m and VEPU < 15 m (1 [010])
	data[4] = 0x7D

	// 0101 1100: NACp: EPU < 10 m and VEPU < 15 m ([1] 010) + NicBaro: Crosschecked (1) + SIL <= 10-7 (11)
	// + Reserved (00)
	data[5] = 0x5C

	// 0001 1110: Reserved (000) + CapabilityMode ACAS not operational + Resolution Advisory active (11) +
	// Emergency: Downed (110)
	data[6] = 0x1E

	return data
}
