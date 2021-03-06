package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"testing"
)

func TestReadFormat31ReservedValid(t *testing.T) {

	msg, err := ReadFormat31Reserved(buildValidFormat31ReservedMessage())
	if err != nil {
		t.Fatal(err)
	}

	if msg.GetMessageFormat() != adsb.Format31 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format31.ToString(),
			msg.GetMessageFormat().ToString())
	}

	if msg.GetRegister().GetId() != bds.BDS65.GetId() {
		t.Errorf("Expected Register \"%v\", got \"%v\"",
			bds.BDS65.GetId(),
			msg.GetRegister().GetId())
	}

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func TestReadFormat31ReservedTooShort(t *testing.T) {

	// Get too short data
	data := buildValidFormat31ReservedMessage()[:6]

	_, err := ReadFormat31Reserved(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat31ReservedBadCode(t *testing.T) {

	// Change code to 2
	data := buildValidFormat31ReservedMessage()
	data[0] = (data[0] & 0x07) | 0x10

	_, err := ReadFormat31Reserved(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat31ReservedBadADSBLevel(t *testing.T) {

	// Get data at ADSB level 1
	data := buildValidFormat31ReservedMessage()
	data[5] = data[5] | 0x20

	_, err := ReadFormat31Reserved(data)
	if err == nil {
		t.Error(err)
	}
}

func buildValidFormat31ReservedMessage() []byte {
	data := make([]byte, 7)

	// 1111 1000: code 31 (11111) + reserved (000)
	data[0] = 0xF8

	// 0000 0000: En Route Operational Capabilities Reserved (0000) + Terminal Area Capabilities Reserved (0000)
	data[1] = 0x00

	// 0000 0000: Approach Operational Capabilities Reserved (0000) + Surface Capabilities Reserved (0000)
	data[2] = 0x00

	// 0000 0000: En Route Operational Status Reserved (0000) + Terminal Area Status Reserved (0000)
	data[3] = 0x08

	// 0000 0000:  Approach Operational Status Reserved (0000) + Surface Status Reserved (0000)
	data[4] = 0x00

	// 0000 0000: ADSB Version (000) + Reserved (0 0000)
	data[5] = 0x00

	// 0000 0000: Reserved (0000 0000)
	data[6] = 0x00

	return data
}
