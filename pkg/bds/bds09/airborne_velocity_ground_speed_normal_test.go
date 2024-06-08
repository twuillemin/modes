package bds09

import (
	"encoding/hex"
	"testing"

	"github.com/twuillemin/modes/pkg/bds/bds09/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

func TestReadAirborneVelocityGroundSpeedNormalValid(t *testing.T) {

	id1, _ := hex.DecodeString("99440994083817")
	message, err := ReadAirborneVelocityGroundSpeedNormal(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS09 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS09)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.DirectionEastWest != fields.DEWWest {
		t.Errorf("for DirectionEastWest: got %v, want %v", message.DirectionEastWest, fields.DEWWest)
	}

	// Field value is 9, so Velocity is 9 - 1 => 8
	if message.VelocityEW != 8 {
		t.Errorf("for VelocityEW: got %v, want %v", message.VelocityEW, 8)
	}

	if message.DirectionNorthSouth != fields.DNSSouth {
		t.Errorf("for DirectionNorthSouth: got %v, want %v", message.DirectionNorthSouth, fields.DNSSouth)
	}

	// Field value is 160, so Velocity is 160 - 1 => 159
	if message.VelocityNS != 159 {
		t.Errorf("for VelocityNS: got %v, want %v", message.VelocityNS, 159)
	}

	if message.VerticalRateStatus != fields.NVSRegular {
		t.Errorf("for VerticalRateStatus: got %v, want %v", message.VerticalRateStatus, fields.NVSRegular)
	}

	// Field value is 14, so rate is (14 - 1) * 64 => 832
	if message.VerticalRate != -832 {
		t.Errorf("for VerticalRate: got %v, want %v", message.VerticalRate, -832)
	}

	if message.DifferenceAltitudeGNSSBaroStatus != fields.NVSRegular {
		t.Errorf("for DifferenceAltitudeGNSSBaroStatus: got %v, want %v", message.DifferenceAltitudeGNSSBaroStatus, fields.NVSRegular)
	}

	if message.DifferenceAltitudeGNSSBaro != 550 {
		t.Errorf("for DifferenceAltitudeGNSSBaro: got %v, want %v", message.DifferenceAltitudeGNSSBaro, 550)
	}

	speed, err := message.GetSpeed()
	if err != nil {
		t.Errorf("for GetSpeed(): got %v, want nil", err)
	}

	if speed != 159 {
		t.Errorf("for GetSpeed(): got %v, want %v", speed, 159)
	}

	track, err := message.GetTrack()
	if err != nil {
		t.Errorf("for GetTrack(): got %v, want nil", err)
	}

	if track < 182.0 || track > 183.0 {
		t.Errorf("for GetTrack(): got %v, want ~ %v", track, 182.88)
	}
}
