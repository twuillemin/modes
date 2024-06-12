package bds61

import (
	"encoding/hex"
	"testing"

	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

func TestReadHeadingAndSpeedReportValid(t *testing.T) {

	id1, _ := hex.DecodeString("E1191D00000000")
	message, err := ReadAircraftStatusEmergency(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS61 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS60)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.EmergencyPriorityStatus != fields.EPSNoEmergency {
		t.Errorf("for field EmergencyPriorityStatus: got %v, want %v", message.EmergencyPriorityStatus, fields.EPSNoEmergency)
	}
}
