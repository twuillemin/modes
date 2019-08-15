package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
	"testing"
)

func TestReadFormat28ACASValid(t *testing.T) {

	msg, err := ReadFormat28ACAS(buildValidFormat28ACASMessage())
	if err != nil {
		t.Fatal(err)
	}

	if msg.GetMessageFormat() != adsb.Format28 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format28.ToString(),
			msg.GetMessageFormat().ToString())
	}

	if msg.GetRegister().GetId() != bds.BDS61.GetId() {
		t.Errorf("Expected Register \"%v\", got \"%v\"",
			bds.BDS61.GetId(),
			msg.GetRegister().GetId())
	}

	if msg.GetSubtype() != fields.SubtypeRABroadcast {
		t.Errorf("Expected Subtype \"%v\", got \"%v\"",
			fields.SubtypeRABroadcast.ToString(),
			msg.GetSubtype().ToString())
	}

	if msg.GetACASData()[0] != 0xAA {
		t.Errorf("Expected Data[0] to be 0xAA , got \"%v\"", msg.GetACASData()[0])
	}

	if msg.GetACASData()[1] != 0x88 {
		t.Errorf("Expected Data[1] to be 0x88 , got \"%v\"", msg.GetACASData()[1])
	}

	if msg.GetACASData()[2] != 0x66 {
		t.Errorf("Expected Data[2] to be 0x66 , got \"%v\"", msg.GetACASData()[2])
	}

	if msg.GetACASData()[3] != 0x44 {
		t.Errorf("Expected Data[3] to be 0x44 , got \"%v\"", msg.GetACASData()[3])
	}

	if msg.GetACASData()[4] != 0x22 {
		t.Errorf("Expected Data[4] to be 0x22 , got \"%v\"", msg.GetACASData()[4])
	}

	if msg.GetACASData()[5] != 0x11 {
		t.Errorf("Expected Data[5] to be 0x11 , got \"%v\"", msg.GetACASData()[5])
	}

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func TestReadFormat28ACASTooShort(t *testing.T) {

	// Shorten the data
	data := buildValidFormat28ACASMessage()[:6]

	_, err := ReadFormat28ACAS(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat28ACASBadFormat(t *testing.T) {

	// Change code to 2
	data := buildValidFormat28ACASMessage()
	data[0] = (data[0] & 0x07) | 0x10

	_, err := ReadFormat28ACAS(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat28ACASBadSubtype(t *testing.T) {

	// Change subtype to 1
	data := buildValidFormat28ACASMessage()
	data[0] = (data[0] & 0xF8) | 0x01

	_, err := ReadFormat28ACAS(data)
	if err == nil {
		t.Error(err)
	}
}

func buildValidFormat28ACASMessage() []byte {
	data := make([]byte, 7)

	// 1110 0010: code 28 (11100) + subtype 2 (010)
	data[0] = 0xE2

	// 1010 1010: ACAS Data
	data[1] = 0xAA

	// 1000 1000: ACAS Data
	data[2] = 0x88

	// 0110 0110: ACAS Data
	data[3] = 0x66

	// 0100 0100:  ACAS Data
	data[4] = 0x44

	// 0010 0010: ACAS Data
	data[5] = 0x22

	// 0001 0001: ACAS Data
	data[6] = 0x11

	return data
}
