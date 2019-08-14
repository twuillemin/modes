// Package messages holds the definition of the messages
package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"testing"
)

func TestReadFormat08V0Valid(t *testing.T) {

	data := buildValidBDS06Message()
	data[0] = data[0] | 0x40

	msg, err := readFormat08V0(data)
	if err != nil {
		t.Fatal(err)
	}

	if msg.GetMessageFormat() != adsb.Format08V0 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format08V0.ToString(),
			msg.GetMessageFormat().ToString())
	}

	isBDS06Valid(t, msg)
}

func TestReadFormat08V0TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS06Message()[:6]
	data[0] = data[0] | 0x40

	_, err := readFormat08V0(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat08V0BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS06Message()
	data[0] = data[0] | 0x01

	_, err := readFormat08V0(data)
	if err == nil {
		t.Error(err)
	}
}
