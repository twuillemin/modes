package messages

import (
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/bds08/fields"
	"github.com/twuillemin/modes/pkg/bds"
	"testing"
)

func TestReadFormat04Valid(t *testing.T) {

	msg, err := ReadFormat04(buildValidFormat04Message())
	if err != nil {
		t.Error(err)
	}

	if msg.GetMessageFormat() != adsb.Format04 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format04.ToString(),
			msg.GetMessageFormat().ToString())
	}

	if msg.GetRegister().GetId() != bds.BDS08.GetId() {
		t.Errorf("Expected Register \"%v\", got \"%v\"",
			bds.BDS08.GetId(),
			msg.GetRegister().GetId())
	}

	if msg.AircraftCategory != fields.ACSAHighVortex {
		t.Errorf("Expected category \"%v\", got \"%v\"",
			fields.ACSAHighVortex,
			msg.AircraftCategory.ToString())
	}

	if msg.AircraftIdentification != "ABOPZ09 " {
		t.Errorf("Expected identification \"%v\", got \"%v\"",
			"ABOPZ09 ",
			msg.AircraftIdentification)
	}

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func TestReadFormat04TooShort(t *testing.T) {

	// Get too short data
	data := buildValidFormat04Message()[:6]

	_, err := ReadFormat04(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat04BadCode(t *testing.T) {

	// Change code to 11
	data := buildValidFormat04Message()
	data[0] = (data[0] & 0x07) | 0x80

	_, err := ReadFormat04(data)
	if err == nil {
		t.Error(err)
	}
}

func buildValidFormat04Message() []byte {
	data := make([]byte, 7)

	// 0010 0100: code 1 (00100) + set A / High vortex (100)
	data[0] = 0x24

	// 0000 0100: A (000001) + B (00[0010])
	data[1] = 0x04

	// 0010 0011: B ([00]0010) + O (0011[11])
	data[2] = 0x23

	// 1101 0000: O ([0011]11) + P (010000)
	data[3] = 0xD0

	// 0110 1011:  Z (011010) + 0 (11[0000])
	data[4] = 0x6B

	// 0000 1110: 0 ([11]0000) + 9 (1110[01])
	data[5] = 0x0E

	// 0110 0000: 9 ([1110]01) + Space (100000)
	data[6] = 0x60

	return data
}
