package bds45

import (
	"encoding/hex"
	"github.com/twuillemin/modes/pkg/bds/bds45/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
	"testing"
)

func TestReadMeteorologicalRoutineAirReportV2Valid(t *testing.T) {

	id1, _ := hex.DecodeString("F2E4ECB6FD3C9C")
	message, err := ReadMeteorologicalHazardReportV1(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS45 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS45)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.WindShearStatus != true {
		t.Errorf("for WindShearStatus: got %v, want %v", message.WindShearStatus, true)
	}

	if message.WindShear != fields.HLSevere {
		t.Errorf("for WindShear: got %v, want %v", message.WindShear, fields.HLSevere)
	}

	if message.MicroBurstStatus != true {
		t.Errorf("for MicroBurstStatus: got %v, want %v", message.MicroBurstStatus, true)
	}

	if message.MicroBurst != fields.HLNil {
		t.Errorf("for MicroBurst: got %v, want %v", message.MicroBurst, fields.HLNil)
	}

	if message.IcingStatus != true {
		t.Errorf("for IcingStatus: got %v, want %v", message.IcingStatus, true)
	}

	if message.Icing != fields.HLLight {
		t.Errorf("for Icing: got %v, want %v", message.Icing, fields.HLLight)
	}

	if message.WakeVortexStatus != true {
		t.Errorf("for WakeVortexStatus: got %v, want %v", message.WakeVortexStatus, true)
	}

	if message.WakeVortex != fields.HLModerate {
		t.Errorf("for WakeVortex: got %v, want %v", message.WakeVortex, fields.HLModerate)
	}

	if message.StaticAirTemperatureStatus != false {
		t.Errorf("for StaticAirTemperatureStatus: got %v, want %v", message.StaticAirTemperatureStatus, false)
	}

	if message.StaticAirTemperature != -98.5 {
		t.Errorf("for StaticAirTemperature: got %v, want %v", message.StaticAirTemperature, -98.5)
	}

	if message.RadioHeightStatus != true {
		t.Errorf("for RadioHeightStatus: got %v, want %v", message.RadioHeightStatus, true)
	}

	if message.RadioHeight != 3518 {
		t.Errorf("for RadioHeight: got %v, want %v", message.RadioHeight, 3518)
	}

	if message.TurbulenceStatus != true {
		t.Errorf("for TurbulenceStatus: got %v, want %v", message.TurbulenceStatus, true)
	}

	if message.TurbulenceAverageEDR != 0.38 {
		t.Errorf("for TurbulenceAverageEDR: got %v, want %v", message.TurbulenceAverageEDR, 0.38)
	}

	if message.TurbulencePeakEDR != 1.0 {
		t.Errorf("for TurbulencePeakEDR: got %v, want %v", message.TurbulencePeakEDR, 1.0)
	}

	if message.TurbulencePeakDelayInterval != 7 {
		t.Errorf("for TurbulencePeakDelayInterval: got %v, want %v", message.TurbulencePeakDelayInterval, 7)
	}
}
