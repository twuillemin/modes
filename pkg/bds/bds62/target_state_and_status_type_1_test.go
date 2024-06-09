package bds62

import (
	"encoding/hex"
	"testing"

	"github.com/twuillemin/modes/pkg/bds/bds62/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

func TestReadTargetStateAndStatus1Valid1(t *testing.T) {

	id1, _ := hex.DecodeString("EA21485CBF3F8C")
	message, err := ReadTargetStateAndStatus1(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS62 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS62)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.GetSubtype() != fields.Subtype1 {
		t.Errorf("for GetSubtype: got %v, want %v", message.GetSubtype(), fields.Subtype1)
	}

	if message.SourceIntegrityLevelSupplement != fields.SILSByHour {
		t.Errorf("for SourceIntegrityLevelSupplement: got %v, want %v", message.SourceIntegrityLevelSupplement, fields.SILSByHour)
	}

	if message.SelectedAltitudeType != fields.SATByMCPFCU {
		t.Errorf("for SelectedAltitudeType: got %v, want %v", message.SelectedAltitudeType, fields.SATByMCPFCU)
	}

	if message.SelectedAltitudeStatus != fields.NVSRegular {
		t.Errorf("for SelectedAltitudeStatus: got %v, want %v", message.SelectedAltitudeStatus, fields.NVSRegular)
	}

	if message.SelectedAltitude != 16992 {
		t.Errorf("for SelectedAltitude: got %v, want %v", message.SelectedAltitude, 16992)
	}

	if message.BarometricPressureSettingStatus != fields.NVSRegular {
		t.Errorf("for BarometricPressureSettingStatus: got %v, want %v", message.BarometricPressureSettingStatus, fields.NVSRegular)
	}

	if message.BarometricPressureSetting < 1012.7 || message.BarometricPressureSetting > 1012.9 {
		t.Errorf("for BarometricPressureSetting: got %v, want ~ %v", message.BarometricPressureSetting, 1012.8)
	}

	if message.SelectedHeadingStatus != fields.SHSValid {
		t.Errorf("for SelectedHeadingStatus: got %v, want %v", message.SelectedHeadingStatus, fields.SHSValid)
	}

	if message.SelectedHeading < 66.7 || message.SelectedHeading > 66.9 {
		t.Errorf("for SelectedHeading: got %v, want ~ %v", message.SelectedHeading, 66.8)
	}

	if message.AutopilotEngaged != true {
		t.Errorf("for AutopilotEngaged: got %v, want %v", message.AutopilotEngaged, true)
	}

	if message.VNAVModeEngaged != true {
		t.Errorf("for VNAVModeEngaged: got %v, want %v", message.VNAVModeEngaged, true)
	}

	if message.AltitudeHoldModeEngaged != false {
		t.Errorf("for AltitudeHoldModeEngaged: got %v, want %v", message.AltitudeHoldModeEngaged, false)
	}

	if message.ApproachModeEngaged != false {
		t.Errorf("for AltitudeHoldModeEngaged: got %v, want %v", message.AltitudeHoldModeEngaged, false)
	}

	if message.TCASACASOperational != true {
		t.Errorf("for TCASACASOperational: got %v, want %v", message.TCASACASOperational, true)
	}

	if message.LNAVModeEngaged != true {
		t.Errorf("for LNAVModeEngaged: got %v, want %v", message.LNAVModeEngaged, true)
	}
}
