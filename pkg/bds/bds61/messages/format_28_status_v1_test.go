package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
	"testing"
)

func TestReadFormat28StatusV1Valid(t *testing.T) {

	msg, err := ReadFormat28StatusV1(buildValidFormat28StatusV1Message())
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

	if msg.GetSubtype() != fields.SubtypeEmergencyPriorityStatus {
		t.Errorf("Expected Subtype \"%v\", got \"%v\"",
			fields.SubtypeEmergencyPriorityStatus.ToString(),
			msg.GetSubtype().ToString())
	}

	if msg.GetEmergencyPriorityStatus() != fields.EPSUnlawfulInterference {
		t.Errorf("Expected Emergency \"%v\", got \"%v\"",
			fields.EPSUnlawfulInterference.ToString(),
			msg.GetEmergencyPriorityStatus().ToString())
	}

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func TestReadFormat28StatusV1TooShort(t *testing.T) {

	// Shorten the data
	data := buildValidFormat28StatusV1Message()[:6]

	_, err := ReadFormat28StatusV1(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat28StatusV1BadFormat(t *testing.T) {

	// Change code to 2
	data := buildValidFormat28StatusV1Message()
	data[0] = (data[0] & 0x07) | 0x10

	_, err := ReadFormat28StatusV1(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat28StatusV1BadSubtype(t *testing.T) {

	// Change subtype to 2
	data := buildValidFormat28StatusV1Message()
	data[0] = (data[0] & 0xF8) | 0x02

	_, err := ReadFormat28StatusV1(data)
	if err == nil {
		t.Error(err)
	}
}

func buildValidFormat28StatusV1Message() []byte {
	data := make([]byte, 7)

	// 1110 0001: code 28 (11100) + subtype 1 (001)
	data[0] = 0xE1

	// 1010 0000: Emergency State: Unlawful interference (101) + Reserved (0 0000)
	data[1] = 0xA0

	// 0000 0000: Reserved (0000 0000)
	data[2] = 0x00

	// 0000 0000: Reserved (0000 0000)
	data[3] = 0x00

	// 0000 0000: Reserved (0000 0000)
	data[4] = 0x00

	// 0000 0000: Reserved (0000 0000)
	data[5] = 0x00

	// 0000 0000: Reserved (0000 0000)
	data[6] = 0x00

	return data
}
