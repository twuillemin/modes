package bds62

import (
	"encoding/hex"
	"testing"

	"github.com/twuillemin/modes/pkg/bds/bds62/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

func TestReadTargetStateAndStatusValid1(t *testing.T) {

	id1, _ := hex.DecodeString("EA21485CBF3F8C")
	message, err := ReadTargetStateAndStatus(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS62 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS62)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.GetSubtype() != fields.Subtype1 {
		t.Errorf("for GetSubtype: got %v, want %v", message.GetSubtype(), fields.Subtype0)
	}
}
