package bds44

import (
	"encoding/hex"
	"testing"

	"github.com/twuillemin/modes/pkg/bds/register"
)

func TestReadMeteorologicalRoutineAirReportV1Valid(t *testing.T) {

	//a0000f3c18bba9cdc000004c6477
	id1, _ := hex.DecodeString("08BBA9CDC00000")
	message, err := ReadMeteorologicalRoutineAirReportV1(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS44 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS44)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}
}
