package bds20

import (
	"encoding/hex"
	"github.com/twuillemin/modes/pkg/bds/register"
	"testing"
)

func TestReadAircraftIdentificationValid(t *testing.T) {

	id1, _ := hex.DecodeString("20508673E19820")
	message, err := ReadAircraftIdentification(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS20 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS20)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.Identification != "THY38Y  " {
		t.Errorf("for Identification: got %s, want %s", message.Identification, "THY38Y  ")
	}
}

func TestReadAircraftIdentificationIncoherent(t *testing.T) {

	id1, _ := hex.DecodeString("20508673E1982F")
	message, err := ReadAircraftIdentification(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.CheckCoherency() == nil {
		t.Errorf("for Coherency: got nil, want error")
	}

	if message.Identification != "THY38Y #" {
		t.Errorf("for Identification: got %s, want %s", message.Identification, "THY38Y #")
	}
}

func TestReadAircraftIdentificationErroneous(t *testing.T) {
	tooShortMessage, _ := hex.DecodeString("20508673E198")
	_, err := ReadAircraftIdentification(tooShortMessage)
	if err == nil {
		t.Errorf("for ReadAircraftIdentification: got nil, want error")
	}

	wrongHeaderMessage, _ := hex.DecodeString("30508673E19820")
	_, err = ReadAircraftIdentification(wrongHeaderMessage)
	if err == nil {
		t.Errorf("for ReadAircraftIdentification: got nil, want error")
	}
}
