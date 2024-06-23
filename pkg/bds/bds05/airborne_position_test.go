package bds05

import (
	"encoding/hex"
	"testing"

	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

func TestReadAirbornePositionAltitudeBaro25FeetIncrement(t *testing.T) {

	rawMessage, _ := hex.DecodeString("58C382D690C8AC")
	message, err := ReadAirbornePosition(rawMessage)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS05 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS05)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.AltitudeSource != fields.AltitudeBarometric {
		t.Errorf("for AltitudeSource: got %v, want %v", message.AltitudeSource, fields.AltitudeBarometric)
	}

	if message.AltitudeReportMethod != fields.AltitudeReport25FootIncrements {
		t.Errorf("for AltitudeReportMethod: got %v, want %v", message.AltitudeReportMethod, fields.AltitudeReport25FootIncrements)
	}

	if message.AltitudeInFeet != 38000 {
		t.Errorf("for AltitudeInFeet: got %v, want %v", message.AltitudeInFeet, 38000)
	}
}

func TestReadAirbornePositionAltitudeGNSS25FeetIncrement(t *testing.T) {

	rawMessage, _ := hex.DecodeString("58C901375147EF")
	message, err := ReadAirbornePosition(rawMessage)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS05 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS05)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.AltitudeSource != fields.AltitudeBarometric {
		t.Errorf("for AltitudeSource: got %v, want %v", message.AltitudeSource, fields.AltitudeBarometric)
	}

	if message.AltitudeReportMethod != fields.AltitudeReport25FootIncrements {
		t.Errorf("for AltitudeReportMethod: got %v, want %v", message.AltitudeReportMethod, fields.AltitudeReport25FootIncrements)
	}

	if message.AltitudeInFeet != 39000 {
		t.Errorf("for AltitudeInFeet: got %v, want %v", message.AltitudeInFeet, 39000)
	}
}
