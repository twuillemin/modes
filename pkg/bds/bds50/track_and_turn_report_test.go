package bds50

import (
	"encoding/hex"
	"testing"

	"github.com/twuillemin/modes/pkg/bds/register"
)

func TestSelectedReadTrackAndTurnReportValid1(t *testing.T) {

	id1, _ := hex.DecodeString("81951536E024D4")
	message, err := ReadTrackAndTurnReport(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS50 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS50)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.RollAngleStatus != true {
		t.Errorf("for field RollAngleStatus: got %v, want %v", message.RollAngleStatus, true)
	}

	if message.RollAngle < 2.05 || message.RollAngle > 2.15 {
		t.Errorf("for field RollAngle: got %v, want ~ %v", message.RollAngle, 2.1)
	}

	if message.TrueTrackAngleStatus != true {
		t.Errorf("for field TrueTrackAngleStatus: got %v, want %v", message.TrueTrackAngleStatus, true)
	}

	if message.TrueTrackAngle < 114.2 || message.TrueTrackAngle > 114.3 {
		t.Errorf("for field TrueTrackAngle: got %v, want ~ %v", message.TrueTrackAngle, 114.258)
	}

	if message.GroundSpeedStatus != true {
		t.Errorf("for field GroundSpeedStatus: got %v, want %v", message.GroundSpeedStatus, true)
	}

	if message.GroundSpeed != 438 {
		t.Errorf("for field GroundSpeed: got %v, want %v", message.GroundSpeed, 438)
	}

	if message.TrackAngleRateStatus != true {
		t.Errorf("for field TrackAngleRateStatus: got %v, want %v", message.TrackAngleRateStatus, true)
	}

	if message.TrackAngleRate < 0.1 || message.TrackAngleRate > 0.2 {
		t.Errorf("for field TrackAngleRate: got %v, want ~ %v", message.TrackAngleRate, 0.125)
	}

	if message.TrueAirSpeedStatus != true {
		t.Errorf("for field TrueAirSpeedStatus: got %v, want %v", message.TrueAirSpeedStatus, true)
	}

	if message.TrueAirSpeed != 424 {
		t.Errorf("for field TrueAirSpeed: got %v, want ~ %v", message.TrueAirSpeed, 424)
	}
}

func TestSelectedReadTrackAndTurnReportValid2(t *testing.T) {

	id1, _ := hex.DecodeString("FFD263377FFCE0")
	message, err := ReadTrackAndTurnReport(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS50 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS50)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.RollAngleStatus != true {
		t.Errorf("for field RollAngleStatus: got %v, want %v", message.RollAngleStatus, true)
	}

	if message.RollAngle < -0.36 || message.RollAngle > -0.35 {
		t.Errorf("for field RollAngle: got %v, want ~ %v", message.RollAngle, -0.351)
	}
}

func TestSelectedReadTrackAndTurnReportIncoherent(t *testing.T) {
	allStatusFalse, _ := hex.DecodeString("00000000000000")
	messageAllStatusFalse, err := ReadTrackAndTurnReport(allStatusFalse)
	if err != nil {
		t.Fatal(err)
	}

	if messageAllStatusFalse.CheckCoherency() == nil {
		t.Errorf("for Coherency: got nil, want error")
	}

	rollAngleTooHigh, _ := hex.DecodeString("BF951536E024D4")
	messageRollAngleTooHigh, err := ReadTrackAndTurnReport(rollAngleTooHigh)
	if err != nil {
		t.Fatal(err)
	}

	if messageRollAngleTooHigh.CheckCoherency() == nil {
		t.Errorf("for Coherency: got nil, want error")
	}

	// Ground speed 1974
	groundSpeedTooHigh, _ := hex.DecodeString("819515F6E024D4")
	messageGroundSpeedTooHigh, err := ReadTrackAndTurnReport(groundSpeedTooHigh)
	if err != nil {
		t.Fatal(err)
	}

	if messageGroundSpeedTooHigh.CheckCoherency() == nil {
		t.Errorf("for Coherency: got nil, want error")
	}

	// Ground speed 1960
	airSpeedTooHigh, _ := hex.DecodeString("81951536E027D4")
	messageAirSpeedTooHigh, err := ReadTrackAndTurnReport(airSpeedTooHigh)
	if err != nil {
		t.Fatal(err)
	}

	if messageAirSpeedTooHigh.CheckCoherency() == nil {
		t.Errorf("for Coherency: got nil, want error")
	}

	// Ground speed 438 Air speed 0
	differenceTooHigh, _ := hex.DecodeString("81951536E02400")
	messageDifferenceTooHigh, err := ReadTrackAndTurnReport(differenceTooHigh)
	if err != nil {
		t.Fatal(err)
	}

	if messageDifferenceTooHigh.CheckCoherency() == nil {
		t.Errorf("for Coherency: got nil, want error")
	}
}

func TestSelectedReadTrackAndTurnReportErroneous(t *testing.T) {
	tooShortMessage, _ := hex.DecodeString("85E42F31300")
	_, err := ReadTrackAndTurnReport(tooShortMessage)
	if err == nil {
		t.Errorf("for ReadSelectedVerticalIntention: got nil, want error")
	}
}
