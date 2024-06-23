package fields

import (
	"testing"
)

func TestAltitudeInvalid(t *testing.T) {

	data := make([]byte, 6)
	for i := 0; i < len(data); i++ {
		data[i] = 0
	}

	altitude, _ := ReadThreatIdentityAltitude(data)

	if altitude.AltitudeStatus != AltitudeInvalid {
		t.Errorf("Expected Altitude Status \"%v\", got \"%v\"",
			AltitudeInvalid.ToString(),
			altitude.AltitudeStatus.ToString())
	}
}

func TestAltitudeLow(t *testing.T) {

	// Altitude code is a 13 bits fields, so read a uint16
	// byte         data[2]    |        data[3]        |   data[4]
	// bit       19 20 21 22 23|24 25 26 27 28 29 30 31|32 33 34 35 36
	// value     _  _  _  C1 A1|C2 A2 C4 A4  0 B1 D1 B2|D2 B4 D4  _  _

	// ALTITUDE  A1   A2   A4   B1   B2   B4   C1   C2   C4   D1   D2   D4   SQUAWK
	//	-1200     0    0    0    0    0    0    0    0    1    0    0    0    0040

	data := make([]byte, 6)
	for i := 0; i < len(data); i++ {
		data[i] = 0
	}

	// Set C4 to 1
	data[3] = data[3] | 0x20

	altitude, _ := ReadThreatIdentityAltitude(data)

	if altitude.AltitudeStatus != AltitudeValid {
		t.Errorf("Expected Altitude Status \"%v\", got \"%v\"",
			AltitudeValid.ToString(),
			altitude.AltitudeStatus.ToString())
	}

	if altitude.AltitudeInFeet != -1200 {
		t.Errorf("Expected Altitude In Feet \"-1200\", got \"%v\"", altitude.AltitudeInFeet)
	}
}

func TestAltitudeZero(t *testing.T) {

	// Altitude code is a 13 bits fields, so read a uint16
	// byte         data[2]    |        data[3]        |   data[4]
	// bit       19 20 21 22 23|24 25 26 27 28 29 30 31|32 33 34 35 36
	// value     _  _  _  C1 A1|C2 A2 C4 A4  0 B1 D1 B2|D2 B4 D4  _  _

	// ALTITUDE  A1   A2   A4   B1   B2   B4   C1   C2   C4   D1   D2   D4   SQUAWK
	//  0         0    0    0    0    1    1    0    1    0    0    0    0    0620

	// => 00 1000 0001 010

	data := make([]byte, 6)
	for i := 0; i < len(data); i++ {
		data[i] = 0
	}

	// Set B2
	data[3] = data[3] | 0x01
	// Set B4
	data[4] = data[4] | 0x40
	// Set C2
	data[3] = data[3] | 0x80

	altitude, _ := ReadThreatIdentityAltitude(data)

	if altitude.AltitudeStatus != AltitudeValid {
		t.Errorf("Expected Altitude Status \"%v\", got \"%v\"",
			AltitudeValid.ToString(),
			altitude.AltitudeStatus.ToString())
	}

	if altitude.AltitudeInFeet != 0 {
		t.Errorf("Expected Altitude In Feet \"0\", got \"%v\"", altitude.AltitudeInFeet)
	}
}

func TestAltitudeHigh(t *testing.T) {

	// Altitude code is a 13 bits fields, so read a uint16
	// byte         data[2]    |        data[3]        |   data[4]
	// bit       19 20 21 22 23|24 25 26 27 28 29 30 31|32 33 34 35 36
	// value     _  _  _  C1 A1|C2 A2 C4 A4  0 B1 D1 B2|D2 B4 D4  _  _

	// ALTITUDE  A1   A2   A4   B1   B2   B4   C1   C2   C4   D1   D2   D4   SQUAWK
	//  126700    0    0    0    0    0    0    0    0    1    0    1    0    0042

	data := make([]byte, 6)
	for i := 0; i < len(data); i++ {
		data[i] = 0
	}

	// Set C4
	data[3] = data[3] | 0x20
	// Set D2
	data[4] = data[4] | 0x80

	altitude, _ := ReadThreatIdentityAltitude(data)

	if altitude.AltitudeStatus != AltitudeValid {
		t.Errorf("Expected Altitude Status \"%v\", got \"%v\"",
			AltitudeValid.ToString(),
			altitude.AltitudeStatus.ToString())
	}

	if altitude.AltitudeInFeet != 126700 {
		t.Errorf("Expected Altitude In Feet \"126700\", got \"%v\"", altitude.AltitudeInFeet)
	}
}
