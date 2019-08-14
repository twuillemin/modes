package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds09/fields"
	"testing"
)

func TestReadFormat19GroundNormalValid(t *testing.T) {

	msg, err := readFormat19GroundNormal(buildValidFormat19GroundNormalMessage())
	if err != nil {
		t.Fatal(err)
	}

	if msg.GetMessageFormat() != adsb.Format19V0OrMore {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format19V0OrMore.ToString(),
			msg.GetMessageFormat().ToString())
	}

	if msg.GetRegister().GetId() != bds.BDS09.GetId() {
		t.Errorf("Expected Register \"%v\", got \"%v\"",
			bds.BDS09.GetId(),
			msg.GetRegister().GetId())
	}

	if msg.GetSubtype() != fields.SubtypeGroundSpeedNormal {
		t.Errorf("Expected Subtype \"%v\", got \"%v\"",
			fields.SubtypeGroundSpeedNormal.ToString(),
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

	if msg.DirectionEastWest != fields.DEWWest {
		t.Errorf("Expected Direction East West \"%v\", got \"%v\"",
			fields.DEWWest.ToString(),
			msg.DirectionEastWest.ToString())
	}

	if msg.VelocityEWNormal.GetVelocity() != 340 {
		t.Errorf("Expected Velocity EW to be 340, got \"%v\"",
			msg.VelocityEWNormal.GetVelocity())
	}

	if msg.DirectionNorthSouth != fields.DNSSouth {
		t.Errorf("Expected Direction North South \"%v\", got \"%v\"",
			fields.DNSSouth.ToString(),
			msg.DirectionNorthSouth.ToString())
	}

	if msg.VelocityNSNormal.GetVelocity() != 340 {
		t.Errorf("Expected Velocity NS to be 340, got \"%v\"",
			msg.VelocityNSNormal.GetVelocity())
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

func TestReadFormat19GroundNormalTooShort(t *testing.T) {

	// Get too short data
	data := buildValidFormat19GroundNormalMessage()[:6]

	_, err := readFormat19GroundNormal(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat19GroundNormalBadCode(t *testing.T) {

	// Change code to 2
	data := buildValidFormat19GroundNormalMessage()
	data[0] = (data[0] & 0x07) | 0x10

	_, err := readFormat19GroundNormal(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat19GroundNormalBadSubType(t *testing.T) {

	// Change subtype to reserved (0)
	data := buildValidFormat19GroundNormalMessage()
	data[0] = data[0] & 0xF8

	_, err := readFormat19GroundNormal(data)
	if err == nil {
		t.Error(err)
	}
}

func buildValidFormat19GroundNormalMessage() []byte {
	data := make([]byte, 7)

	// 1001 1001: code 19 (10011) + subtype 1 (001)
	data[0] = 0x99

	// 1101 0101: Intent Change Flag (1) + IFR capability Flag (1) + NUC < 3 m/s (010) + Direction EW West (1) +
	// Velocity EW 340kt -> 341 (01 [0101 0101])
	data[1] = 0xD5

	// 0101 0101: Velocity 340kt -> 341 ([01] 0101 0101)
	data[2] = 0x55

	// 1010 1010: + Direction NS South (1) + Velocity NS 340kt -> 341 (010 1010 [101])
	data[3] = 0xAA

	// 1011 1010:  Velocity NS 340kt -> 341 ([010 1010]101) + Source Vertical Baro (1) + Sign Vertical Down (1) +
	// Vertical Rate 10816 -> 170 (010[101010])
	data[4] = 0xBA

	// 1010 1000: Vertical Rate 10816 -> 170 ([010]101010) + reserved (00)
	data[5] = 0xA8

	// 1101 0101: Difference Sign Bit below (1) + 2150 -> 85(101 0101)
	data[6] = 0xD5

	return data
}
