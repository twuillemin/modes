package fields

import "fmt"

// DifferenceGNSSBaroSign is the GNSS Altitude Sign Bit definition
//
// Specified in Doc 9871 / Table A-2-9
type DifferenceGNSSBaroSign byte

const (
	// GASBAboveBaro indicates above baro alt.
	GASBAboveBaro DifferenceGNSSBaroSign = 0
	// GASBBelowBaro indicates below baro alt.
	GASBBelowBaro DifferenceGNSSBaroSign = 1
)

// ToString returns a basic, but readable, representation of the field
func (bit DifferenceGNSSBaroSign) ToString() string {

	switch bit {
	case GASBAboveBaro:
		return "0 - above baro alt."
	case GASBBelowBaro:
		return "1 - below baro alt."
	default:
		return fmt.Sprintf("%v - Unknown code", bit)
	}
}

// ReadDifferenceGNSSBaroSign reads the DifferenceGNSSBaroSign from a 56 bits data field
func ReadDifferenceGNSSBaroSign(data []byte) DifferenceGNSSBaroSign {
	bits := (data[6] & 0x08) >> 3
	return DifferenceGNSSBaroSign(bits)
}
