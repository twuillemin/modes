package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"testing"
)

func TestDetectAndReadReadFormat01Valid(t *testing.T) {

	msg, adsbResult, err := ReadBDS08(adsb.Level0OrMore, buildValidFormat01Message())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format01); !ok {
		t.Errorf("expected a message of type Format01, but got %v instead", msg.GetMessageFormat().ToString())
	}
	if adsbResult != adsb.Level0OrMore {
		t.Errorf("expected adsbLevel to be returned as Level0OrMore, but got %v instead", adsbResult.ToString())
	}
}

func TestDetectAndReadReadFormat02Valid(t *testing.T) {

	msg, adsbResult, err := ReadBDS08(adsb.Level1OrMore, buildValidFormat02Message())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format02); !ok {
		t.Errorf("expected a message of type Format02, but got %v instead", msg.GetMessageFormat().ToString())
	}
	if adsbResult != adsb.Level1OrMore {
		t.Errorf("expected adsbLevel to be returned as Level1OrMore, but got %v instead", adsbResult.ToString())
	}
}

func TestDetectAndReadReadFormat03Valid(t *testing.T) {

	msg, adsbResult, err := ReadBDS08(adsb.Level2, buildValidFormat03Message())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format03); !ok {
		t.Errorf("expected a message of type Format03, but got %v instead", msg.GetMessageFormat().ToString())
	}
	if adsbResult != adsb.Level2 {
		t.Errorf("expected adsbLevel to be returned as Level2, but got %v instead", adsbResult.ToString())
	}
}

func TestDetectAndReadReadFormat04Valid(t *testing.T) {

	msg, adsbResult, err := ReadBDS08(adsb.Level0OrMore, buildValidFormat04Message())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format04); !ok {
		t.Errorf("expected a message of type Format04, but got %v instead", msg.GetMessageFormat().ToString())
	}
	if adsbResult != adsb.Level0OrMore {
		t.Errorf("expected adsbLevel to be returned as Level0OrMore, but got %v instead", adsbResult.ToString())
	}
}

func TestDetectBadFormat(t *testing.T) {

	// make a Format 05 message
	data := buildValidFormat04Message()
	data[0] = (data[0] & 0x07) | 0x28

	_, _, err := ReadBDS08(adsb.Level0OrMore, data)
	if err == nil {
		t.Fatal("Expected an error while reading a message with format 5, but message was read")
	}
}
