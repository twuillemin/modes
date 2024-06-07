package bds60

import (
	"encoding/hex"
	"github.com/twuillemin/modes/pkg/bds/register"
	"testing"
)

func TestReadHeadingAndSpeedReportValid(t *testing.T) {

	id1, _ := hex.DecodeString("8F39F91A7E27C4")
	message, err := ReadHeadingAndSpeedReport(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS60 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS60)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.MagneticHeadingStatus != true {
		t.Errorf("for field MagneticHeadingStatus: got %v, want %v", message.MagneticHeadingStatus, true)
	}

	if message.MagneticHeading < 42.7 || message.MagneticHeading > 42.8 {
		t.Errorf("for field MagneticHeading: got %v, want ~ %v", message.MagneticHeading, 42.71484)
	}

	if message.IndicatedAirSpeedStatus != true {
		t.Errorf("for field IndicatedAirSpeedStatus: got %v, want %v", message.IndicatedAirSpeedStatus, true)
	}

	if message.IndicatedAirSpeed != 252 {
		t.Errorf("for field IndicatedAirSpeed: got %v, want %v", message.IndicatedAirSpeed, 42.71484)
	}

	if message.MachStatus != true {
		t.Errorf("for field MachStatus: got %v, want %v", message.MachStatus, true)
	}

	if message.Mach < 0.41 || message.Mach > 0.43 {
		t.Errorf("for field Mach: got %v, want ~ %v", message.Mach, 0.42)
	}

	if message.BarometricAltitudeRateStatus != true {
		t.Errorf("for field BarometricAltitudeRateStatus: got %v, want %v", message.BarometricAltitudeRateStatus, true)
	}

	if message.BarometricAltitudeRate != -1920 {
		t.Errorf("for field BarometricAltitudeRate: got %v, want %v", message.BarometricAltitudeRate, -1920)
	}

	if message.InitialVerticalVelocityStatus != true {
		t.Errorf("for field InitialVerticalVelocityStatus: got %v, want %v", message.InitialVerticalVelocityStatus, true)
	}

	if message.InitialVerticalVelocity != -1920 {
		t.Errorf("for field InitialVerticalVelocity: got %v, want %v", message.InitialVerticalVelocity, -1920)
	}
}

func TestReadHeadingAndSpeedReportIncoherent(t *testing.T) {
	allStatusFalse, _ := hex.DecodeString("00000000000000")
	messageAllStatusFalse, err := ReadHeadingAndSpeedReport(allStatusFalse)
	if err != nil {
		t.Fatal(err)
	}

	if messageAllStatusFalse.CheckCoherency() == nil {
		t.Errorf("for Coherency: got nil, want error")
	}

	// Air speed 896
	airSpeedTooHigh, _ := hex.DecodeString("8F3F011A7E27C4")
	messageAirSpeedTooHigh, err := ReadHeadingAndSpeedReport(airSpeedTooHigh)
	if err != nil {
		t.Fatal(err)
	}

	if messageAirSpeedTooHigh.CheckCoherency() == nil {
		t.Errorf("for Coherency: got nil, want error")
	}

	// Mach 2.048
	machTooHigh, _ := hex.DecodeString("8F39F9803E27C4")
	messageMachTooHigh, err := ReadHeadingAndSpeedReport(machTooHigh)
	if err != nil {
		t.Fatal(err)
	}

	if messageMachTooHigh.CheckCoherency() == nil {
		t.Errorf("for Coherency: got nil, want error")
	}

	// Barometric -16384
	barometricTooHigh, _ := hex.DecodeString("8F39F91A7007C4")
	messageBarometricTooHigh, err := ReadHeadingAndSpeedReport(barometricTooHigh)
	if err != nil {
		t.Fatal(err)
	}

	if messageBarometricTooHigh.CheckCoherency() == nil {
		t.Errorf("for Coherency: got nil, want error")
	}

	// Inertial -16384
	InertialTooHigh, _ := hex.DecodeString("8F39F91A7E2600")
	messageInertialTooHigh, err := ReadHeadingAndSpeedReport(InertialTooHigh)
	if err != nil {
		t.Fatal(err)
	}

	if messageInertialTooHigh.CheckCoherency() == nil {
		t.Errorf("for Coherency: got nil, want error")
	}
}

func TestReadHeadingAndSpeedReportErroneous(t *testing.T) {
	tooShortMessage, _ := hex.DecodeString("8F39F91A7E27")
	_, err := ReadHeadingAndSpeedReport(tooShortMessage)
	if err == nil {
		t.Errorf("for ReadHeadingAndSpeedReport: got nil, want error")
	}
}
