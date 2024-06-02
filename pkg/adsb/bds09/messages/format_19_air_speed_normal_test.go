package messages

import (
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/bds09/fields"
	"testing"
)

func TestReadFormat19AirspeedNormalValid(t *testing.T) {

	msg, err := ReadFormat19AirSpeedNormal(buildValidFormat19AirspeedNormalMessage())
	if err != nil {
		t.Fatal(err)
	}

	if msg.GetMessageFormat() != adsb.Format19 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format19.ToString(),
			msg.GetMessageFormat().ToString())
	}

	if msg.GetSubtype() != fields.SubtypeAirspeedNormal {
		t.Errorf("Expected Subtype \"%v\", got \"%v\"",
			fields.SubtypeAirspeedNormal.ToString(),
			msg.GetSubtype().ToString())
	}

	if msg.IntentChange != fields.ICChangeInIntent {
		t.Errorf("Expected Intent Change \"%v\", got \"%v\"",
			fields.ICChangeInIntent.ToString(),
			msg.IntentChange.ToString())
	}

	if msg.IFRCapability != fields.IFRCapable {
		t.Errorf("Expected IFR Capability \"%v\", got \"%v\"",
			fields.IFRCapable.ToString(),
			msg.IFRCapability.ToString())
	}

	if msg.NavigationUncertaintyCategory != fields.NUCPHorizontalLowerThan3VerticalLowerThan4Point6 {
		t.Errorf("Expected Navigation Uncertainty Category \"%v\", got \"%v\"",
			fields.NUCPHorizontalLowerThan3VerticalLowerThan4Point6.ToString(),
			msg.IFRCapability.ToString())
	}

	if msg.MagneticHeadingStatus != fields.AVSBAvailable {
		t.Errorf("Expected Magnetic Heading Status \"%v\", got \"%v\"",
			fields.AVSBAvailable.ToString(),
			msg.MagneticHeadingStatus.ToString())
	}

	if msg.MagneticHeading < 119.0 || msg.MagneticHeading > 120.0 {
		t.Errorf("Expected Magnetic Heading between 119 and 120, got \"%v\"",
			msg.MagneticHeading)
	}

	if msg.AirspeedType != fields.ATTrueAirSpeed {
		t.Errorf("Expected AirspeedType \"%v\", got \"%v\"",
			fields.ATTrueAirSpeed.ToString(),
			msg.AirspeedType.ToString())
	}

	if msg.AirspeedNormal.GetAirspeed() != 340 {
		t.Errorf("Expected Airspeed to be 340, got \"%v\"",
			msg.AirspeedNormal.GetAirspeed())
	}

	if msg.VerticalRateSource != fields.VRSBaro {
		t.Errorf("Expected Vertical Rate Source \"%v\", got \"%v\"",
			fields.VRSBaro.ToString(),
			msg.VerticalRateSource.ToString())
	}

	if msg.VerticalRateSign != fields.VRSDown {
		t.Errorf("Expected Vertical Rate Sign \"%v\", got \"%v\"",
			fields.VRSDown.ToString(),
			msg.VerticalRateSign.ToString())
	}

	if msg.DifferenceGNSSBaroSign != fields.GASBAboveBaro {
		t.Errorf("Expected Difference GNSS Baro Sign \"%v\", got \"%v\"",
			fields.GASBAboveBaro.ToString(),
			msg.DifferenceGNSSBaroSign.ToString())
	}

	if msg.DifferenceGNSSBaro.GetDifferenceGNSSBaro() != 2100 {
		t.Errorf("Expected Difference GNSS Baro to be 2100, got \"%v\"",
			msg.DifferenceGNSSBaro.GetDifferenceGNSSBaro())
	}

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func TestReadFormat19AirspeedNormalTooShort(t *testing.T) {

	// Get too short data
	data := buildValidFormat19AirspeedNormalMessage()[:6]

	_, err := ReadFormat19AirSpeedNormal(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat19AirspeedNormalBadCode(t *testing.T) {

	// Change code to 2
	data := buildValidFormat19AirspeedNormalMessage()
	data[0] = (data[0] & 0x07) | 0x10

	_, err := ReadFormat19AirSpeedNormal(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat19AirspeedNormalBadSubType(t *testing.T) {

	// Change subtype to reserved (0)
	data := buildValidFormat19AirspeedNormalMessage()
	data[0] = data[0] & 0xF8

	_, err := ReadFormat19AirSpeedNormal(data)
	if err == nil {
		t.Error(err)
	}
}

func buildValidFormat19AirspeedNormalMessage() []byte {
	data := make([]byte, 7)

	// 1001 1011: code 19 (10011) + subtype 3 (011)
	data[0] = 0x9B

	// 1101 0101: Intent Change Flag (1) + IFR capability Flag (1) + NUC < 3 m/s (010) + Magnetic available (1) +
	// Magnetic heading: 119.88 deg => 341/1024 (01 [0101 0101])
	data[1] = 0xD5

	// 0101 0101: Magnetic heading: 119.88 deg => 341/1024 ([01] 0101 0101)
	data[2] = 0x55

	// 1010 1010: AirSpeed Type TAs (1) + Velocity 340 kt => 341 (010 1010 [101])
	data[3] = 0xAA

	// 1011 1010:  Velocity 340 kt => 341 ([010 1010] 101) + Source Vertical Baro (1) + Sign Vertical Down (1) +
	// Vertical Rate 10816 -> 170 (010[101010])
	data[4] = 0xBA

	// 1010 1000: Vertical Rate 10816 -> 170 ([010]101010) + reserved (00)
	data[5] = 0xA8

	// 1101 0101: Difference Sign Bit below (1) + 2150 -> 85(101 0101)
	data[6] = 0xD5

	return data
}