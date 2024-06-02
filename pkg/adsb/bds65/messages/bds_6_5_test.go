package messages

import (
	"github.com/twuillemin/modes/pkg/adsb"
	"testing"
)

func TestDetectAndReadReadFormat31V0Valid(t *testing.T) {

	msg, err := ReadBDS65(adsb.ReaderLevel0, buildValidFormat31ReservedMessage())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format31Reserved); !ok {
		t.Errorf("expected a message of type Format31Reserved, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadReadFormat31AirborneV1Valid(t *testing.T) {

	msg, err := ReadBDS65(adsb.ReaderLevel0, buildValidFormat31AirborneV1Message())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format31AirborneV1); !ok {
		t.Errorf("expected a message of type Format31AirborneV1, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadReadFormat31SurfaceV1Valid(t *testing.T) {

	msg, err := ReadBDS65(adsb.ReaderLevel0, buildValidFormat31SurfaceV1Message())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format31SurfaceV1); !ok {
		t.Errorf("expected a message of type Format31SurfaceV1, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadReadFormat31AirborneV2Valid(t *testing.T) {

	msg, err := ReadBDS65(adsb.ReaderLevel0, buildValidFormat31AirborneV2Message())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format31AirborneV2); !ok {
		t.Errorf("expected a message of type Format31AirborneV2, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadReadFormat31SurfaceV2Valid(t *testing.T) {

	msg, err := ReadBDS65(adsb.ReaderLevel0, buildValidFormat31SurfaceV2Message())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format31SurfaceV2); !ok {
		t.Errorf("expected a message of type Format31SurfaceV2, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectBadFormat(t *testing.T) {

	// make a Format 05 message
	data := buildValidFormat31SurfaceV2Message()
	data[0] = (data[0] & 0x07) | 0x28

	_, err := ReadBDS65(adsb.ReaderLevel0, data)
	if err == nil {
		t.Fatal("Expected an error while reading a message with format 5, but message was read")
	}
}