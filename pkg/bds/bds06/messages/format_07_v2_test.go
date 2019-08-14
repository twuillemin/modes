// Package messages holds the definition of the messages
package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"testing"
)

func TestReadFormat07V2Valid(t *testing.T) {

	data := buildValidBDS06Message()
	data[0] = data[0] | 0x38

	msg, err := readFormat07V2(false, false, data)
	if err != nil {
		t.Fatal(err)
	}

	if msg.GetMessageFormat() != adsb.Format07V2 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format07V2.ToString(),
			msg.GetMessageFormat().ToString())
	}

	isBDS06Valid(t, msg)
}

func TestReadFormat07V2TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS06Message()[:6]
	data[0] = data[0] | 0x38

	_, err := readFormat07V2(false, false, data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat07V2BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS06Message()
	data[0] = data[0] | 0x01

	_, err := readFormat07V2(false, false, data)
	if err == nil {
		t.Error(err)
	}
}
