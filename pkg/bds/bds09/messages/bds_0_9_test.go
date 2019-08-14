package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"testing"
)

func TestDetectAndReadFormat19AirspeedNormalValid(t *testing.T) {

	msg, adsbResult, err := ReadBDS09(adsb.Level0OrMore, buildValidFormat19AirspeedNormalMessage())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format19AirspeedNormal); !ok {
		t.Errorf("expected a message of type Format19AirspeedNormal, but got %v instead", msg.GetMessageFormat().ToString())
	}
	if adsbResult != adsb.Level0OrMore {
		t.Errorf("expected adsbLevel to be returned as Level0OrMore, but got %v instead", adsbResult.ToString())
	}
}

func TestDetectAndReadFormat19AirspeedSupersonicValid(t *testing.T) {

	msg, adsbResult, err := ReadBDS09(adsb.Level0OrMore, buildValidFormat19AirspeedSupersonicMessage())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format19AirspeedSupersonic); !ok {
		t.Errorf("expected a message of type Format19AirspeedSupersonic, but got %v instead", msg.GetMessageFormat().ToString())
	}
	if adsbResult != adsb.Level0OrMore {
		t.Errorf("expected adsbLevel to be returned as Level0OrMore, but got %v instead", adsbResult.ToString())
	}
}

func TestDetectAndReadReadFormat19GroundNormalValid(t *testing.T) {

	msg, adsbResult, err := ReadBDS09(adsb.Level2, buildValidFormat19GroundNormalMessage())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format19GroundNormal); !ok {
		t.Errorf("expected a message of type Format19GroundNormal, but got %v instead", msg.GetMessageFormat().ToString())
	}
	if adsbResult != adsb.Level2 {
		t.Errorf("expected adsbLevel to be returned as Level2, but got %v instead", adsbResult.ToString())
	}
}

func TestDetectAndReadReadFormat19GroundSupersonicValid(t *testing.T) {

	msg, adsbResult, err := ReadBDS09(adsb.Level0OrMore, buildValidFormat19GroundSupersonicMessage())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format19GroundSupersonic); !ok {
		t.Errorf("expected a message of type Format19GroundSupersonic, but got %v instead", msg.GetMessageFormat().ToString())
	}
	if adsbResult != adsb.Level0OrMore {
		t.Errorf("expected adsbLevel to be returned as Level0OrMore, but got %v instead", adsbResult.ToString())
	}
}

func TestDetectBadFormat(t *testing.T) {

	// make a Format 05 message
	data := buildValidFormat19GroundSupersonicMessage()
	data[0] = (data[0] & 0x07) | 0x28

	_, _, err := ReadBDS09(adsb.Level0OrMore, data)
	if err == nil {
		t.Fatal("Expected an error while reading a message with format 5, but message was read")
	}
}
