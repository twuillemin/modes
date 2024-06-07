package bds00

import (
	"encoding/hex"
	"github.com/twuillemin/modes/pkg/bds/register"
	"testing"
)

func TestNoMessageAvailableValid(t *testing.T) {

	id1, _ := hex.DecodeString("00000000000000")
	message, err := ReadNoMessageAvailable(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS00 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS00)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}
}

func TestNoMessageAvailableErroneous(t *testing.T) {
	tooShortMessage, _ := hex.DecodeString("000000000000")
	_, err := ReadNoMessageAvailable(tooShortMessage)
	if err == nil {
		t.Errorf("for ReadNoMessageAvailable: got nil, want error")
	}

	wrongHeaderMessage, _ := hex.DecodeString("00000000000001")
	_, err = ReadNoMessageAvailable(wrongHeaderMessage)
	if err == nil {
		t.Errorf("for ReadNoMessageAvailable: got nil, want error")
	}
}
