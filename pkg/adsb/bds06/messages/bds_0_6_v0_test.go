package messages

import (
	"github.com/twuillemin/modes/pkg/adsb/bds06/fields"
	"github.com/twuillemin/modes/pkg/bds"
	"testing"
)

func isBDS06V0Valid(t *testing.T, msg MessageBDS06) {

	if msg.GetRegister().GetId() != bds.BDS06.GetId() {
		t.Errorf("Expected Register \"%v\", got \"%v\"",
			bds.BDS06.GetId(),
			msg.GetRegister().GetId())
	}

	if movement := msg.GetMovement().GetMovement(); movement < 112.5 || movement > 113.0 {
		t.Errorf("Expected Movement between 112.5 and 113 km/h, got \"%v\"",
			movement)
	}

	if msg.GetGroundTrackStatus() != fields.GTSValid {
		t.Errorf("Expected Ground Track Status \"%v\", got \"%v\"",
			fields.GTSValid.ToString(),
			msg.GetGroundTrackStatus().ToString())
	}

	if msg.GetGroundTrack().GetGroundTrack() != 118.125 {
		t.Errorf("Expected Ground Track to be 118.125, got \"%v\"",
			msg.GetGroundTrack().GetGroundTrack())
	}

	if msg.GetTime() != fields.TSynchronizedUTC {
		t.Errorf("Expected Time \"%v\", got \"%v\"",
			fields.TSynchronizedUTC.ToString(),
			msg.GetTime().ToString())
	}

	if msg.GetCPRFormat() != fields.CPRFormatOdd {
		t.Errorf("Expected CPR Format \"%v\", got \"%v\"",
			fields.CPRFormatOdd.ToString(),
			msg.GetCPRFormat().ToString())
	}

	if msg.GetEncodedLatitude() != 43722 {
		t.Errorf("Expected Latitude to be 43722, got \"%v\"",
			msg.GetEncodedLatitude())
	}

	if msg.GetEncodedLongitude() != 87466 {
		t.Errorf("Expected Longitude to be 87466, got \"%v\"",
			msg.GetEncodedLatitude())
	}

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func buildValidBDS06V0Message() []byte {
	data := make([]byte, 7)

	// Movement 1010101-> 85 decimal. In range 39-93 So value is (85-39) * 1.852 + 27.78 = 112.972
	// Ground track 0101010 -> 42 decimal. Factor 360 / 128: 118.125
	// Latitude: 43722: 01 01010101 1001010 (just a number)
	// Longitude: 87466: 1 01010101 10101010 (just a number)

	// 0000 0101: code empty (00000) + Movement 112.972 km/h -> (101[0101])
	data[0] = 0x05

	// 0101 1010: Movement 112.972 km/h -> ([101]0101) + Ground Track: Valid (1) + Ground track :118.125 (010 [1010])
	data[1] = 0x5A

	// 1010 1101: Ground track :118.125 ([010] 1010) + Time: Synchro UTC (1) + CPR format Odd (1) + Latitude (01 [...])
	data[2] = 0xAD

	// 0101 0101: Latitude ([01] 01010101 [1001010])
	data[3] = 0x55

	// 1001 0101: Latitude ([01 01010101] 1001010) + Longitude (1 [...])
	data[4] = 0x95

	// 0101 0101: Longitude ([1] 01010101 [10101010])
	data[5] = 0x55

	// 1010 1010: Longitude ([1 01010101] 10101010)
	data[6] = 0xAA

	return data
}
