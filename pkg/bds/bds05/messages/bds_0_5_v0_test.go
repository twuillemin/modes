package messages

import (
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
	"testing"
)

// --------------------------------------------------------------------------------
//
// Generic methods valid for all messages
//
// --------------------------------------------------------------------------------

func isBDS05V0Valid(t *testing.T, msg MessageBDS05V0) {

	if msg.GetRegister().GetId() != bds.BDS05.GetId() {
		t.Errorf("Expected Register \"%v\", got \"%v\"",
			bds.BDS05.GetId(),
			msg.GetRegister().GetId())
	}

	if msg.GetSurveillanceStatus() != fields.SSTemporaryAlert {
		t.Errorf("Expected Time \"%v\", got \"%v\"",
			fields.SSTemporaryAlert.ToString(),
			msg.GetSurveillanceStatus().ToString())
	}

	if msg.GetSingleAntennaFlag() != fields.SAFSingle {
		t.Errorf("Expected Time \"%v\", got \"%v\"",
			fields.SAFSingle.ToString(),
			msg.GetSingleAntennaFlag().ToString())
	}

	if msg.GetAltitude().AltitudeInFeet != 33250 {
		t.Errorf("Expected Altitude In Feet to be 33250, got \"%v\"",
			msg.GetAltitude().AltitudeInFeet)
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

func buildValidBDS05V0Message() []byte {
	data := make([]byte, 7)

	// Altitude in 25 feet Encoded as C1 A1 C2 A2 C4 A4 <<M>> B1 <Q> B2 D2 B4 D4 with
	// <<M>> : Skipped
	// <Q>: 1 : 25 feet increment
	// 1 0 1 0 1 0 - _ - 1 - 1 - 1 0 1 0 -> 10101011010 = 1370 -> Altitude = 1370 * 25 - 1000 = 33250

	// 0000 0101: code empty (00000) + Surveillance Status: Temporary alert (10) + Single Antenna (1)
	data[0] = 0x05

	// 1010 1011: Altitude: 33250 ft (10101011 [1010])
	data[1] = 0xAB

	// 1010 1101: Altitude: 33250 ft ([10101011] 1010) + Time: Synchro UTC (1) + CPR format Odd (1) + Latitude (01 [...])
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
