package fields

import "fmt"

// GNSSAltitudeSignBit is the GNSS Altitude Sign Bit definition
//
// Specified in Doc 9871 / Table A-2-9
type GNSSAltitudeSignBit byte

const (
	// GASBAboveBaro indicates above baro alt.
	GASBAboveBaro GNSSAltitudeSignBit = 0
	// GASBBelowBaro indicates below baro alt.
	GASBBelowBaro GNSSAltitudeSignBit = 1
)

// ToString returns a basic, but readable, representation of the field
func (bit GNSSAltitudeSignBit) ToString() string {

	switch bit {
	case GASBAboveBaro:
		return "0 - above baro alt."
	case GASBBelowBaro:
		return "1 - below baro alt."
	default:
		return fmt.Sprintf("%v - Unknown code", bit)
	}
}

// ReadGNSSAltitudeSignBit reads the GNSSAltitudeSignBit from a 56 bits data field
func ReadGNSSAltitudeSignBit(data []byte) GNSSAltitudeSignBit {
	bits := (data[6] & 0x08) >> 3
	return GNSSAltitudeSignBit(bits)
}
