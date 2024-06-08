package bds09

import (
	"encoding/hex"
	"github.com/twuillemin/modes/pkg/bds/bds09/fields"
	"testing"

	"github.com/twuillemin/modes/pkg/bds/register"
)

func TestReadAirborneVelocityValid1(t *testing.T) {

	id1, _ := hex.DecodeString("99440994083817")
	message, err := ReadAirborneVelocity(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS09 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS09)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.GetSubtype() != fields.SubtypeGroundSpeedNormal {
		t.Errorf("for GetSubtype: got %v, want %v", message.GetSubtype(), fields.SubtypeGroundSpeedNormal)
	}
}

func TestReadAirborneVelocityValid2(t *testing.T) {

	id1, _ := hex.DecodeString("9B06B6AF189400")
	message, err := ReadAirborneVelocity(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS09 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS09)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.GetSubtype() != fields.SubtypeAirspeedNormal {
		t.Errorf("for GetSubtype: got %v, want %v", message.GetSubtype(), fields.SubtypeAirspeedNormal)
	}
}
