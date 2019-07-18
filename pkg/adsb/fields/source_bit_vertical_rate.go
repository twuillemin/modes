package fields

import "fmt"

// SourceBitVerticalRate is the Source Bit for Vertical Rate definition
//
// Specified in Doc 9871 / Table A-2-9
type SourceBitVerticalRate byte

const (
	// SBVRGNSS indicates GNSS
	SBVRGNSS SourceBitVerticalRate = 0
	// SBVRBaro indicates Baro
	SBVRBaro SourceBitVerticalRate = 1
)

// ToString returns a basic, but readable, representation of the field
func (bit SourceBitVerticalRate) ToString() string {

	switch bit {
	case SBVRGNSS:
		return "0 - GNSS"
	case SBVRBaro:
		return "1 - Baro"
	default:
		return fmt.Sprintf("%v - Unknown code", bit)
	}
}

// ReadSourceBitVerticalRate reads the SourceBitVerticalRate from a 56 bits data field
func ReadSourceBitVerticalRate(data []byte) SourceBitVerticalRate {
	bits := (data[4] & 0x10) >> 4
	return SourceBitVerticalRate(bits)
}
