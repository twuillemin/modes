package fields

import "fmt"

// AltitudeTypeSubfield is the Altitude type subfield definition
//
// Specified in Doc 9871 / A.2.7
type AltitudeTypeSubfield byte

const (
	// BarometricAltitude indicates Barometric altitude.
	BarometricAltitude AltitudeTypeSubfield = 0
	// GNSSHeight indicates GNSS height.
	GNSSHeight AltitudeTypeSubfield = 1
)

// ToString returns a basic, but readable, representation of the field
func (ats AltitudeTypeSubfield) ToString() string {

	switch ats {
	case BarometricAltitude:
		return "0 - Barometric altitude"
	case GNSSHeight:
		return "1 - GNSS height (HAE)"
	default:
		return fmt.Sprintf("%v - Unknown code", ats)
	}
}

// ReadAltitudeTypeSubfield reads the AltitudeTypeSubfield from a 56 bits data field
func ReadAltitudeTypeSubfield(data []byte) AltitudeTypeSubfield {
	bits := (data[0] & 0x20) >> 5
	return AltitudeTypeSubfield(bits)
}
