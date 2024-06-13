package bds44

import (
	"encoding/hex"
	"github.com/twuillemin/modes/pkg/bds/bds44/fields"
	"testing"

	"github.com/twuillemin/modes/pkg/bds/register"
)

func TestReadMeteorologicalRoutineAirReportV2Valid(t *testing.T) {

	id1, _ := hex.DecodeString("18BBA9CDC00000")
	message, err := ReadMeteorologicalRoutineAirReportV2(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS44 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS44)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.Source != fields.SourceINS {
		t.Errorf("for Source: got %v, want %v", message.Source, fields.SourceINS)
	}

	if message.WindSpeedStatus != true {
		t.Errorf("for WindSpeedStatus: got %v, want %v", message.WindSpeedStatus, true)
	}

	if message.WindSpeed != 46 {
		t.Errorf("for WindSpeed: got %v, want %v", message.WindSpeed, 46)
	}

	if message.WindDirectionStatus != true {
		t.Errorf("for WindDirectionStatus: got %v, want %v", message.WindDirectionStatus, true)
	}

	if message.WindDirection < 298 || message.WindDirection > 298.5 {
		t.Errorf("for WindDirection: got %v, want ~ %v", message.WindDirection, 298.125)
	}

	if message.StaticAirTemperatureStatus != true {
		t.Errorf("for StaticAirTemperatureStatus: got %v, want %v", message.StaticAirTemperatureStatus, true)
	}

	if message.StaticAirTemperature < -50.5 || message.StaticAirTemperature > -50.0 {
		t.Errorf("for StaticAirTemperature: got %v, want ~ %v", message.StaticAirTemperature, -50.25)
	}
}
