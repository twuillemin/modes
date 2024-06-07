package bds08

import (
	"encoding/hex"
	"github.com/twuillemin/modes/pkg/bds/bds08/fields"
	"testing"

	"github.com/twuillemin/modes/pkg/bds/register"
)

func TestReadAircraftIdentificationAndCategoryValid(t *testing.T) {

	id1, _ := hex.DecodeString("11508673E19820")
	message, err := ReadAircraftIdentificationAndCategory(id1)
	if err != nil {
		t.Fatal(err)
	}

	if message.GetRegister() != register.BDS08 {
		t.Errorf("for Register: got %v, want %v", message.GetRegister(), register.BDS08)
	}

	if message.CheckCoherency() != nil {
		t.Errorf("for Coherency: got %v, want nil", message.CheckCoherency())
	}

	if message.Category != fields.ACSCSurfaceEmergency {
		t.Errorf("for Category: got %v, want %v", message.Category, fields.ACSCSurfaceEmergency)
	}

	if message.Identification != "THY38Y  " {
		t.Errorf("for Identification: got %s, want %s", message.Identification, "THY38Y  ")
	}
}

func TestReadAircraftIdentificationAndCategoryIncoherent(t *testing.T) {

	id1, _ := hex.DecodeString("11508673E1982F")
	message, err := ReadAircraftIdentificationAndCategory(id1)
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

func TestReadAircraftIdentificationAndCategoryErroneous(t *testing.T) {
	tooShortMessage, _ := hex.DecodeString("11508673E198")
	_, err := ReadAircraftIdentificationAndCategory(tooShortMessage)
	if err == nil {
		t.Errorf("for ReadAircraftIdentificationAndCategory: got nil, want error")
	}

	wrongHeaderMessage, _ := hex.DecodeString("F0508673E19820")
	_, err = ReadAircraftIdentificationAndCategory(wrongHeaderMessage)
	if err == nil {
		t.Errorf("for ReadAircraftIdentificationAndCategory: got nil, want error")
	}
}
