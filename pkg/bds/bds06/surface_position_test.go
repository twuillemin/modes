package bds06

import (
	"encoding/hex"
	"testing"

	"github.com/twuillemin/modes/pkg/bds/register"
)

func TestReadSurfacePosition(t *testing.T) {

	rawMessage, _ := hex.DecodeString("3A9A153237AEF0")
	message, err := ReadSurfacePosition(rawMessage)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS06 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS06)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.MovementStatus != 41 {
		t.Errorf("for MovementStatus: got %v, want %v", message.MovementStatus, 41)
	}

	if message.MovementSpeed != 17 {
		t.Errorf("for MovementSpeed: got %v, want %v", message.MovementSpeed, 17)
	}

	if message.GroundTrackStatus != true {
		t.Errorf("for GroundTrackStatus: got %v, want %v", message.GroundTrackStatus, true)
	}

	if message.GroundTrack < 92.7 || message.GroundTrack > 92.9 {
		t.Errorf("for GroundTrack: got %v, want ~ %v", message.GroundTrack, 92.8)
	}
}
