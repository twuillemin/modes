package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds08/fields"
	"testing"
)

func TestReadFormat02Valid(t *testing.T) {

	msg, err := ReadFormat02(buildValidFormat02Message())
	if err != nil {
		t.Error(err)
	}

	if msg.GetMessageFormat() != adsb.Format02 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format02.ToString(),
			msg.GetMessageFormat().ToString())
	}

	if msg.GetRegister().GetId() != bds.BDS08.GetId() {
		t.Errorf("Expected Register \"%v\", got \"%v\"",
			bds.BDS08.GetId(),
			msg.GetRegister().GetId())
	}

	if msg.AircraftCategory != fields.ACSCSurfaceService {
		t.Errorf("Expected category \"%v\", got \"%v\"",
			fields.ACSCSurfaceService,
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

func TestReadFormat02TooShort(t *testing.T) {

	// Get too short data
	data := buildValidFormat02Message()[:6]

	_, err := ReadFormat02(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat02BadCode(t *testing.T) {

	// Change code to 10
	data := buildValidFormat02Message()
	data[0] = (data[0] & 0x07) | 0x80

	_, err := ReadFormat02(data)
	if err == nil {
		t.Error(err)
	}
}

func buildValidFormat02Message() []byte {
	data := make([]byte, 7)

	// 0001 0010: code 1 (00010) + category C / SubtypeSurface Service Vehicle (010)
	data[0] = 0x12

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
