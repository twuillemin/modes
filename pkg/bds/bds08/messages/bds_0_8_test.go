package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"testing"
)

func TestDetectAndReadReadFormat01Valid(t *testing.T) {

	msg, adsbResult, err := ReadBDS08(adsb.ReaderLevel0, buildValidFormat01Message())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format01); !ok {
		t.Errorf("expected a message of type Format01, but got %v instead", msg.GetMessageFormat().ToString())
	}
	if adsbResult != adsb.ReaderLevel0 {
		t.Errorf("expected adsbLevel to be returned as ReaderLevel0, but got %v instead", adsbResult.ToString())
	}
}

func TestDetectAndReadReadFormat02Valid(t *testing.T) {

	msg, adsbResult, err := ReadBDS08(adsb.ReaderLevel1, buildValidFormat02Message())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format02); !ok {
		t.Errorf("expected a message of type Format02, but got %v instead", msg.GetMessageFormat().ToString())
	}
	if adsbResult != adsb.ReaderLevel1 {
		t.Errorf("expected adsbLevel to be returned as ReaderLevel1, but got %v instead", adsbResult.ToString())
	}
}

func TestDetectAndReadReadFormat03Valid(t *testing.T) {

	msg, adsbResult, err := ReadBDS08(adsb.ReaderLevel2, buildValidFormat03Message())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format03); !ok {
		t.Errorf("expected a message of type Format03, but got %v instead", msg.GetMessageFormat().ToString())
	}
	if adsbResult != adsb.ReaderLevel2 {
		t.Errorf("expected adsbLevel to be returned as ReaderLevel2, but got %v instead", adsbResult.ToString())
	}
}

func TestDetectAndReadReadFormat04Valid(t *testing.T) {

	msg, adsbResult, err := ReadBDS08(adsb.ReaderLevel0, buildValidFormat04Message())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format04); !ok {
		t.Errorf("expected a message of type Format04, but got %v instead", msg.GetMessageFormat().ToString())
	}
	if adsbResult != adsb.ReaderLevel0 {
		t.Errorf("expected adsbLevel to be returned as ReaderLevel0, but got %v instead", adsbResult.ToString())
	}
}

func TestDetectBadFormat(t *testing.T) {

	// make a Format 05 message
	data := buildValidFormat04Message()
	data[0] = (data[0] & 0x07) | 0x28

	_, _, err := ReadBDS08(adsb.ReaderLevel0, data)
	if err == nil {
		t.Fatal("Expected an error while reading a message with format 5, but message was read")
	}
}
