package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"testing"
)

func TestDetectAndReadReadFormat29Subtype0Valid(t *testing.T) {

	msg, adsbResult, err := ReadBDS62(adsb.ReaderLevel0OrMore, buildValidFormat29Subtype0Message())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format29Subtype0); !ok {
		t.Errorf("expected a message of type Format29Subtype0, but got %v instead", msg.GetMessageFormat().ToString())
	}
	if adsbResult != adsb.ReaderLevel1OrMore {
		t.Errorf("expected adsbLevel to be returned as ReaderLevel1OrMore, but got %v instead", adsbResult.ToString())
	}
}

func TestDetectAndReadReadFormat29Subtype1Valid(t *testing.T) {

	msg, adsbResult, err := ReadBDS62(adsb.ReaderLevel0OrMore, buildValidFormat29Subtype1Message())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format29Subtype1); !ok {
		t.Errorf("expected a message of type Format29Subtype1, but got %v instead", msg.GetMessageFormat().ToString())
	}
	if adsbResult != adsb.ReaderLevel2 {
		t.Errorf("expected adsbLevel to be returned as ReaderLevel2, but got %v instead", adsbResult.ToString())
	}
}

func TestDetectBadFormat(t *testing.T) {

	// make a Format 05 message
	data := buildValidFormat29Subtype1Message()
	data[0] = (data[0] & 0x07) | 0x28

	_, _, err := ReadBDS62(adsb.ReaderLevel0OrMore, data)
	if err == nil {
		t.Fatal("Expected an error while reading a message with format 5, but message was read")
	}
}
