package bds09

import (
	"encoding/hex"
	"github.com/twuillemin/modes/pkg/bds/bds09/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
	"testing"
)

func TestReadAirborneVelocityAirSpeedNormalValid(t *testing.T) {

	id1, _ := hex.DecodeString("9B06B6AF189400")
	message, err := ReadAirborneVelocityAirSpeedNormal(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS09 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS09)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.AirspeedStatus != fields.NVSRegular {
		t.Errorf("for AirspeedStatus: got %v, want %v", message.AirspeedStatus, fields.NVSRegular)
	}

	// Field value is 376, so Velocity is 376 - 1 => 8
	if message.Airspeed != 375 {
		t.Errorf("for Airspeed: got %v, want %v", message.Airspeed, 375)
	}

	if message.MagneticHeadingStatus != fields.MHSAvailable {
		t.Errorf("for MagneticHeadingStatus: got %v, want %v", message.MagneticHeadingStatus, fields.MHSAvailable)
	}

	if message.MagneticHeading < 243.0 || message.MagneticHeading > 244.0 {
		t.Errorf("for MagneticHeading: got %v, want ~ %v", message.MagneticHeading, 243.98)
	}

	if message.VerticalRate != -2304 {
		t.Errorf("for VerticalRate: got %v, want %v", message.VerticalRate, -832)
	}
}
