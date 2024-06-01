package messages

import (
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/bds62/fields"
	"github.com/twuillemin/modes/pkg/bds"
	"testing"
)

func TestReadFormat29Subtype1Valid(t *testing.T) {

	msg, err := ReadFormat29Subtype1(buildValidFormat29Subtype1Message())
	if err != nil {
		t.Fatal(err)
	}

	if msg.GetMessageFormat() != adsb.Format29 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format29.ToString(),
			msg.GetMessageFormat().ToString())
	}

	if msg.GetRegister().GetId() != bds.BDS62.GetId() {
		t.Errorf("Expected Register \"%v\", got \"%v\"",
			bds.BDS62.GetId(),
			msg.GetRegister().GetId())
	}

	if msg.GetSubtype() != fields.Subtype1 {
		t.Errorf("Expected Subtype \"%v\", got \"%v\"",
			fields.Subtype1.ToString(),
			msg.GetSubtype().ToString())
	}

	if msg.SourceIntegrityLevelSupplement != fields.SILSBySample {
		t.Errorf("Expected Source Integrity ReaderLevel Supplement \"%v\", got \"%v\"",
			fields.SILSBySample.ToString(),
			msg.SourceIntegrityLevelSupplement.ToString())
	}

	if msg.SelectedAltitudeType != fields.SATByFMS {
		t.Errorf("Expected Selected Altitude Type \"%v\", got \"%v\"",
			fields.SATByFMS.ToString(),
			msg.SelectedAltitudeType.ToString())
	}

	if msg.SelectedAltitude.GetSelectedAltitude() != 21792 {
		t.Errorf("Expected Selected Altitude to be 21792, got \"%v\"",
			msg.SelectedAltitude.GetSelectedAltitude())
	}

	if msg.BarometricPressureSetting.GetBarometricPressureSetting() != 272.8 {
		t.Errorf("Expected GetBarometric Pressure Setting to be 272.8, got \"%v\"",
			msg.BarometricPressureSetting.GetBarometricPressureSetting())
	}

	if msg.SelectedHeadingStatus != fields.SHSValid {
		t.Errorf("Expected Selected Heading Status \"%v\", got \"%v\"",
			fields.SHSValid.ToString(),
			msg.SelectedHeadingStatus.ToString())
	}

	if msg.SelectedHeadingSign != fields.SHSNegative {
		t.Errorf("Expected Selected Heading Sign \"%v\", got \"%v\"",
			fields.SHSNegative.ToString(),
			msg.SelectedHeadingSign.ToString())
	}

	if selectedHeading := msg.SelectedHeading.GetSelectedHeading(msg.SelectedHeadingStatus, msg.SelectedHeadingSign); selectedHeading > -59.5 || selectedHeading < -60 {
		t.Errorf("Expected Selected Heading to be -59.78, got \"%v\"",
			selectedHeading)
	}

	if msg.NavigationalAccuracyCategoryPosition != fields.NACV2PEPULowerThan10M {
		t.Errorf("Expected Navigational Accuracy Category Position \"%v\", got \"%v\"",
			fields.NACV2PEPULowerThan10M.ToString(),
			msg.NavigationalAccuracyCategoryPosition.ToString())
	}

	if msg.NICBaro != fields.NICBGilhamCrossCheckedOrNonGilham {
		t.Errorf("Expected NICBaro \"%v\", got \"%v\"",
			fields.NICBGilhamCrossCheckedOrNonGilham.ToString(),
			msg.NICBaro.ToString())
	}

	if msg.SourceIntegrityLevel != fields.SILLevel3 {
		t.Errorf("Expected Source Integrity ReaderLevel \"%v\", got \"%v\"",
			fields.SILLevel3.ToString(),
			msg.SourceIntegrityLevel.ToString())
	}

	if msg.StatusMCPFCUBits != fields.SMFBInformationProvided {
		t.Errorf("Expected Status MCP FCU Bits \"%v\", got \"%v\"",
			fields.SMFBInformationProvided.ToString(),
			msg.StatusMCPFCUBits.ToString())
	}

	if msg.AutopilotEngaged != fields.AEEngaged {
		t.Errorf("Expected Autopilot Engaged \"%v\", got \"%v\"",
			fields.AEEngaged.ToString(),
			msg.AutopilotEngaged.ToString())
	}

	if msg.VNAVModeEngaged != fields.VMEngaged {
		t.Errorf("Expected VNAV Mode Engaged \"%v\", got \"%v\"",
			fields.VMEngaged.ToString(),
			msg.VNAVModeEngaged.ToString())
	}

	if msg.AltitudeHoldModeEngaged != fields.AHEngaged {
		t.Errorf("Expected Altitude Hold Mode Engaged \"%v\", got \"%v\"",
			fields.AHEngaged.ToString(),
			msg.AltitudeHoldModeEngaged.ToString())
	}

	if msg.ApproachModeEngaged != fields.AMEngaged {
		t.Errorf("Expected Approach Mode Engaged \"%v\", got \"%v\"",
			fields.AMEngaged.ToString(),
			msg.ApproachModeEngaged.ToString())
	}

	if msg.ACASOperational != fields.AOOperational {
		t.Errorf("Expected ACAS Operational \"%v\", got \"%v\"",
			fields.AOOperational.ToString(),
			msg.ACASOperational.ToString())
	}

	if msg.LNAVModeEngaged != fields.LMEngaged {
		t.Errorf("Expected LNAV Mode Engaged \"%v\", got \"%v\"",
			fields.LMEngaged.ToString(),
			msg.LNAVModeEngaged.ToString())
	}

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func TestReadFormat29Subtype1TooShort(t *testing.T) {

	// Shorten the data
	data := buildValidFormat29Subtype1Message()[:6]

	_, err := ReadFormat29Subtype1(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat29Subtype1BadFormat(t *testing.T) {

	// Change code to 2
	data := buildValidFormat29Subtype1Message()
	data[0] = (data[0] & 0x07) | 0x10

	_, err := ReadFormat29Subtype1(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat29Subtype1BadSubtype(t *testing.T) {

	// Change subtype to 5
	data := buildValidFormat29Subtype1Message()
	data[0] = data[0] | 0x04

	_, err := ReadFormat29Subtype1(data)
	if err == nil {
		t.Error(err)
	}
}

func buildValidFormat29Subtype1Message() []byte {
	data := make([]byte, 7)

	// 1110 1011: code 29 (11101) + subtype 0 (01) + SIL Supplement: By Sample (1)
	data[0] = 0xEB

	// 1010 1010: Selected Altitude Type FMS (1) + Target Altitude 21792 -> 682 (010 1010 [1010])
	data[1] = 0xAA

	// 1010 1010: Target Altitude 21792 -> 682 ([010 1010] 1010) + Baro Pressure 272.8 -> 342 (1010 [10110])
	data[2] = 0xAA

	// 1011 0110: Baro Pressure 272.8 -> 342 ([1010] 10110) + Status Valid (1) + Sign Negative (1) + Heading 59.77 (0[101 0101])
	data[3] = 0xB6

	// 1010 1011:  Heading 59.77 ([0] 101 0101) + NACp: EPU < 10 m (1 [010])
	data[4] = 0xAB

	// 0101 1111: NACp: EPU < 10 m ([1] 010) + NicBaro: Crosschecked (1) + SIL <= 10-7 (11) +
	// MCP status: Valid (1) + autopilot engaged (1)
	data[5] = 0x5F

	// 1101 1100: VNAV engaged (1) + Altitude hold engaged (1) + Reserved (0) + Approach engaged (1)
	// + ACAS Operational (1) + LNAV engaged (1) + Reserved (00)
	data[6] = 0xDC

	return data
}
