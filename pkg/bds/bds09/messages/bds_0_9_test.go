package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"testing"
)

func TestDetectAndReadFormat19AirspeedNormalValid(t *testing.T) {

	msg, err := ReadBDS09(adsb.ReaderLevel0, buildValidFormat19AirspeedNormalMessage())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format19AirSpeedNormal); !ok {
		t.Errorf("expected a message of type Format19AirSpeedNormal, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat19AirspeedSupersonicValid(t *testing.T) {

	msg, err := ReadBDS09(adsb.ReaderLevel0, buildValidFormat19AirspeedSupersonicMessage())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format19AirSpeedSupersonic); !ok {
		t.Errorf("expected a message of type Format19AirSpeedSupersonic, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadReadFormat19GroundNormalValid(t *testing.T) {

	msg, err := ReadBDS09(adsb.ReaderLevel2, buildValidFormat19GroundNormalMessage())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format19GroundSpeedNormal); !ok {
		t.Errorf("expected a message of type Format19GroundSpeedNormal, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadReadFormat19GroundSupersonicValid(t *testing.T) {

	msg, err := ReadBDS09(adsb.ReaderLevel0, buildValidFormat19GroundSupersonicMessage())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format19GroundSpeedSupersonic); !ok {
		t.Errorf("expected a message of type Format19GroundSpeedSupersonic, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectBadFormat(t *testing.T) {

	// make a Format 05 message
	data := buildValidFormat19GroundSupersonicMessage()
	data[0] = (data[0] & 0x07) | 0x28

	_, err := ReadBDS09(adsb.ReaderLevel0, data)
	if err == nil {
		t.Fatal("Expected an error while reading a message with format 5, but message was read")
	}
}
