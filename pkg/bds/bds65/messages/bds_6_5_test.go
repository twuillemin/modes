package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"testing"
)

func TestDetectAndReadReadFormat31V0Valid(t *testing.T) {

	msg, adsbResult, err := ReadBDS65(adsb.Level0OrMore, buildValidFormat31V0Message())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format31V0); !ok {
		t.Errorf("expected a message of type Format31V0, but got %v instead", msg.GetMessageFormat().ToString())
	}
	if adsbResult != adsb.Level0Exactly {
		t.Errorf("expected adsbLevel to be returned as Level0Exactly, but got %v instead", adsbResult.ToString())
	}
}

func TestDetectAndReadReadFormat31V1AirborneValid(t *testing.T) {

	msg, adsbResult, err := ReadBDS65(adsb.Level0OrMore, buildValidFormat31V1AirborneMessage())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format31V1Airborne); !ok {
		t.Errorf("expected a message of type Format31V1Airborne, but got %v instead", msg.GetMessageFormat().ToString())
	}
	if adsbResult != adsb.Level1Exactly {
		t.Errorf("expected adsbLevel to be returned as Level1Exactly, but got %v instead", adsbResult.ToString())
	}
}

func TestDetectAndReadReadFormat31V1SurfaceValid(t *testing.T) {

	msg, adsbResult, err := ReadBDS65(adsb.Level0OrMore, buildValidFormat31V1SurfaceMessage())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format31V1Surface); !ok {
		t.Errorf("expected a message of type Format31V1Surface, but got %v instead", msg.GetMessageFormat().ToString())
	}
	if adsbResult != adsb.Level1Exactly {
		t.Errorf("expected adsbLevel to be returned as Level1Exactly, but got %v instead", adsbResult.ToString())
	}
}

func TestDetectAndReadReadFormat31V2AirborneValid(t *testing.T) {

	msg, adsbResult, err := ReadBDS65(adsb.Level0OrMore, buildValidFormat31V2AirborneMessage())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format31V2Airborne); !ok {
		t.Errorf("expected a message of type Format31V2Airborne, but got %v instead", msg.GetMessageFormat().ToString())
	}
	if adsbResult != adsb.Level2 {
		t.Errorf("expected adsbLevel to be returned as Level2, but got %v instead", adsbResult.ToString())
	}
}

func TestDetectAndReadReadFormat31V2SurfaceValid(t *testing.T) {

	msg, adsbResult, err := ReadBDS65(adsb.Level0OrMore, buildValidFormat31V2SurfaceMessage())
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format31V2Surface); !ok {
		t.Errorf("expected a message of type Format31V2Surface, but got %v instead", msg.GetMessageFormat().ToString())
	}
	if adsbResult != adsb.Level2 {
		t.Errorf("expected adsbLevel to be returned as Level2, but got %v instead", adsbResult.ToString())
	}
}

func TestDetectBadFormat(t *testing.T) {

	// make a Format 05 message
	data := buildValidFormat31V2SurfaceMessage()
	data[0] = (data[0] & 0x07) | 0x28

	_, _, err := ReadBDS65(adsb.Level0OrMore, data)
	if err == nil {
		t.Fatal("Expected an error while reading a message with format 5, but message was read")
	}
}
