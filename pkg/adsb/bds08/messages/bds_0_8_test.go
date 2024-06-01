package messages

import (
	"github.com/twuillemin/modes/pkg/adsb"
	"testing"
)

func TestDetectAndReadReadFormat01Valid(t *testing.T) {

	msg, err := ReadBDS08(adsb.ReaderLevel0, buildValidFormat01Message())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format01); !ok {
		t.Errorf("expected a message of type Format01, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadReadFormat02Valid(t *testing.T) {

	msg, err := ReadBDS08(adsb.ReaderLevel1, buildValidFormat02Message())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format02); !ok {
		t.Errorf("expected a message of type Format02, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadReadFormat03Valid(t *testing.T) {

	msg, err := ReadBDS08(adsb.ReaderLevel2, buildValidFormat03Message())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format03); !ok {
		t.Errorf("expected a message of type Format03, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadReadFormat04Valid(t *testing.T) {

	msg, err := ReadBDS08(adsb.ReaderLevel0, buildValidFormat04Message())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format04); !ok {
		t.Errorf("expected a message of type Format04, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectBadFormat(t *testing.T) {

	// make a Format 05 message
	data := buildValidFormat04Message()
	data[0] = (data[0] & 0x07) | 0x28

	_, err := ReadBDS08(adsb.ReaderLevel0, data)
	if err == nil {
		t.Fatal("Expected an error while reading a message with format 5, but message was read")
	}
}
