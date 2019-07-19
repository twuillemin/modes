package fields

import "fmt"

// VerticalRateSource is the Source Bit for Vertical Rate definition
//
// Specified in Doc 9871 / Table A-2-9
type VerticalRateSource byte

const (
	// VRSGNSS indicates GNSS
	VRSGNSS VerticalRateSource = 0
	// VRSBaro indicates Baro
	VRSBaro VerticalRateSource = 1
)

// ToString returns a basic, but readable, representation of the field
func (bit VerticalRateSource) ToString() string {

	switch bit {
	case VRSGNSS:
		return "0 - GNSS"
	case VRSBaro:
		return "1 - Baro"
	default:
		return fmt.Sprintf("%v - Unknown code", bit)
	}
}

// ReadVerticalRateSource reads the VerticalRateSource from a 56 bits data field
func ReadVerticalRateSource(data []byte) VerticalRateSource {
	bits := (data[4] & 0x10) >> 4
	return VerticalRateSource(bits)
}
