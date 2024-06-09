package bds44

import (
	"encoding/hex"
	"github.com/twuillemin/modes/pkg/bds/bds44/fields"
	"testing"

	"github.com/twuillemin/modes/pkg/bds/register"
)

func TestReadMeteorologicalRoutineAirReportValid(t *testing.T) {

	//a0000f3c18bba9cdc000004c6477
	id1, _ := hex.DecodeString("18BBA9CDC00000")
	message, err := ReadMeteorologicalRoutineAirReport(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS44 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS44)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	// Should report Invalid as 0 => Message V1
	if message.GetSource() != fields.SourceInvalid {
		t.Errorf("for field GetSource: got %v, want %v", message.GetSource(), 0)
	}
}
