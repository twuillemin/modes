package bds40

import (
	"encoding/hex"
	"testing"

	"github.com/twuillemin/modes/pkg/bds/bds40/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

func TestReadSelectedVerticalIntentionValid(t *testing.T) {

	id1, _ := hex.DecodeString("85E42F313001E7")
	message, err := ReadSelectedVerticalIntention(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS40 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS40)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.FMSSelectedAltitudeAvailable != true {
		t.Errorf("for field FMSSelectedAltitudeAvailable: got %v, want %v", message.FMSSelectedAltitudeAvailable, true)
	}

	if message.FMSSelectedAltitude != 3008 {
		t.Errorf("for field FMSSelectedAltitude: got %v, want %v", message.FMSSelectedAltitude, 3008)
	}

	if message.MCPFCUSelectedAltitudeAvailable != true {
		t.Errorf("for field MCPFCUSelectedAltitudeAvailable: got %v, want %v", message.MCPFCUSelectedAltitudeAvailable, true)
	}

	if message.MCPFCUSelectedAltitude != 3008 {
		t.Errorf("for field MCPFCUSelectedAltitude: got %v, want %v", message.MCPFCUSelectedAltitude, 3008)
	}

	if message.BarometricPressureSettingAvailable != true {
		t.Errorf("for field BarometricPressureSettingAvailable: got %v, want %v", message.BarometricPressureSettingAvailable, true)
	}

	if message.BarometricPressureSetting != 1020.0 {
		t.Errorf("for field BarometricPressureSetting: got %v, want %v", message.BarometricPressureSetting, 1020.0)
	}

	if message.StatusOfModeBits != fields.SMBInformationProvided {
		t.Errorf("for field StatusOfModeBits: got %v, want %v", message.StatusOfModeBits, fields.SMBInformationProvided)
	}

	if message.VNAVMode != true {
		t.Errorf("for field VNAVMode: got %v, want %v", message.VNAVMode, true)
	}

	if message.AltitudeHoldMode != true {
		t.Errorf("for field AltitudeHoldMode: got %v, want %v", message.AltitudeHoldMode, true)
	}

	if message.ApproachMode != true {
		t.Errorf("for field ApproachMode: got %v, want %v", message.ApproachMode, true)
	}

	if message.StatusOfTargetAltitudeSource != fields.STASInformationProvided {
		t.Errorf("for field StatusOfTargetAltitudeSource: got %v, want %v", message.StatusOfTargetAltitudeSource, fields.STASInformationProvided)
	}

	if message.TargetAltitudeSource != fields.TASFMSSelectedAltitude {
		t.Errorf("for field TargetAltitudeSource: got %v, want %v", message.TargetAltitudeSource, true)
	}
}

func TestReadSelectedVerticalIntentionIncoherent(t *testing.T) {
	allStatusFalse, _ := hex.DecodeString("00000000000000")
	messageAllStatusFalse, err := ReadSelectedVerticalIntention(allStatusFalse)
	if err != nil {
		t.Fatal(err)
	}

	if messageAllStatusFalse.CheckCoherency() == nil {
		t.Errorf("for Coherency: got nil, want error")
	}
}

func TestReadSelectedVerticalIntentionErroneous(t *testing.T) {
	tooShortMessage, _ := hex.DecodeString("85E42F31300")
	_, err := ReadSelectedVerticalIntention(tooShortMessage)
	if err == nil {
		t.Errorf("for ReadSelectedVerticalIntention: got nil, want error")
	}

	// Bit 40 is reserved
	bit40SetMessage, _ := hex.DecodeString("85E42F31310000")
	_, err = ReadSelectedVerticalIntention(bit40SetMessage)
	if err == nil {
		t.Errorf("for ReadSelectedVerticalIntention: got nil, want error")
	}

	// Bit 41 - 47 are reserved
	bit41SetMessage, _ := hex.DecodeString("85E42F31308000")
	_, err = ReadSelectedVerticalIntention(bit41SetMessage)
	if err == nil {
		t.Errorf("for ReadSelectedVerticalIntention: got nil, want error")
	}

	// Bit 52 - 53 are reserved
	bit52SetMessage, _ := hex.DecodeString("85E42F31300010")
	_, err = ReadSelectedVerticalIntention(bit52SetMessage)
	if err == nil {
		t.Errorf("for ReadSelectedVerticalIntention: got nil, want error")
	}
}
