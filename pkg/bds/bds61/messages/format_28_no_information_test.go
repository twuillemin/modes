package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
	"testing"
)

func TestReadFormat28NoInformationValid(t *testing.T) {

	msg, err := ReadFormat28NoInformation(buildValidFormat28NoInformationMessage())
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

	if msg.GetSubtype() != fields.SubtypeNoInformation {
		t.Errorf("Expected Subtype \"%v\", got \"%v\"",
			fields.SubtypeNoInformation.ToString(),
			msg.GetSubtype().ToString())
	}

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func TestReadFormat28NoInformationTooShort(t *testing.T) {

	// Shorten the data
	data := buildValidFormat28NoInformationMessage()[:6]

	_, err := ReadFormat28NoInformation(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat28NoInformationBadFormat(t *testing.T) {

	// Change code to 2
	data := buildValidFormat28NoInformationMessage()
	data[0] = (data[0] & 0x07) | 0x10

	_, err := ReadFormat28NoInformation(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat28NoInformationBadSubtype(t *testing.T) {

	// Change subtype to 1
	data := buildValidFormat28NoInformationMessage()
	data[0] = (data[0] & 0xF8) | 0x01

	_, err := ReadFormat28NoInformation(data)
	if err == nil {
		t.Error(err)
	}
}

func buildValidFormat28NoInformationMessage() []byte {
	data := make([]byte, 7)

	// 1110 0000: code 28 (11100) + subtype 0 (000)
	data[0] = 0xE0

	// 0000 0000: Reserved (0000 0000)
	data[1] = 0x00

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
