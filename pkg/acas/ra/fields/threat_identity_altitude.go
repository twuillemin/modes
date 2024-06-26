package fields

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// AltitudeStatus is the status of the Altitude
type AltitudeStatus byte

const (
	// AltitudeInvalid signifies that altitude information is not available or that the altitude has been determined invalid.
	AltitudeInvalid AltitudeStatus = 0
	// AltitudeValid signifies that altitude is valid.
	AltitudeValid AltitudeStatus = 1
)

// ThreatIdentityAltitude is the altitude of the threat. It is given in 100 feet increment
type ThreatIdentityAltitude struct {
	AltitudeStatus AltitudeStatus
	AltitudeInFeet int32
}

// ReadThreatIdentityAltitude reads the altitude code from a message
func ReadThreatIdentityAltitude(data []byte) (ThreatIdentityAltitude, error) {

	// Altitude code is a 13 bits fields, so read a uint16
	// byte         data[2]    |        data[3]        |   data[4]
	// bit       19 20 21 22 23|24 25 26 27 28 29 30 31|32 33 34 35 36
	// value     _  _  _  C1 A1|C2 A2 C4 A4  0 B1 D1 B2|D2 B4 D4  _  _

	// Start by D2 B4 D4
	altitudeCode := uint16(data[4]&0xE0) >> 5
	// Then pack B1 D1 B2
	altitudeCode += uint16(data[3]&0x07) << 3
	// Then C2 A2 C4 A4
	altitudeCode += uint16(data[3]&0xF0) << 2
	// Then C1 A1
	altitudeCode += uint16(data[2]&0x03) << 2

	// Detect invalid altitude
	if altitudeCode == 0 {
		return ThreatIdentityAltitude{
			AltitudeInvalid,
			0,
		}, nil
	}

	c1 := (altitudeCode & 0x0800) != 0
	a1 := (altitudeCode & 0x0400) != 0
	c2 := (altitudeCode & 0x0200) != 0
	a2 := (altitudeCode & 0x0100) != 0
	c4 := (altitudeCode & 0x0080) != 0
	a4 := (altitudeCode & 0x0040) != 0
	b1 := (altitudeCode & 0x0020) != 0
	d1 := (altitudeCode & 0x0010) != 0
	b2 := (altitudeCode & 0x0008) != 0
	d2 := (altitudeCode & 0x0004) != 0
	b4 := (altitudeCode & 0x0002) != 0
	d4 := (altitudeCode & 0x0001) != 0

	altitudeFeet, err := bitutils.GillhamToAltitude(d1, d2, d4, a1, a2, a4, b1, b2, b4, c1, c2, c4)
	if err != nil {
		return ThreatIdentityAltitude{
			AltitudeInvalid,
			0,
		}, errors.New("the Altitude field is malformed")
	}

	return ThreatIdentityAltitude{AltitudeValid, altitudeFeet}, nil
}

// ToString returns a basic, but readable, representation of the field
func (altitudeStatus AltitudeStatus) ToString() string {
	switch altitudeStatus {
	case AltitudeInvalid:
		return "Not Available or Invalid"
	case AltitudeValid:
		return "Valid"
	default:
		return fmt.Sprintf("%v - Unknown code", altitudeStatus)
	}
}

// ToString returns a basic, but readable, representation of the field
func (altitude ThreatIdentityAltitude) ToString() string {
	return fmt.Sprintf("%v ft / Status: %v", altitude.AltitudeInFeet, altitude.AltitudeStatus.ToString())
}
