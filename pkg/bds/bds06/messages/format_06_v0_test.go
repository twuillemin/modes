// Package messages holds the definition of the messages
package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"testing"
)

func TestReadFormat06V0Valid(t *testing.T) {

	data := buildValidBDS06Message()
	data[0] = data[0] | 0x30

	msg, err := readFormat06V0(data)
	if err != nil {
		t.Fatal(err)
	}

	if msg.GetMessageFormat() != adsb.Format06V0 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format06V0.ToString(),
			msg.GetMessageFormat().ToString())
	}

	isBDS06Valid(t, msg)
}

func TestReadFormat06V0TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS06Message()[:6]
	data[0] = data[0] | 0x30

	_, err := readFormat06V0(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat06V0BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS06Message()
	data[0] = data[0] | 0x01

	_, err := readFormat06V0(data)
	if err == nil {
		t.Error(err)
	}
}
