package fields

import "fmt"

// VerticalRateSign is the Sign Bit for Vertical Rate definition
//
// Specified in Doc 9871 / Table A-2-9
type VerticalRateSign byte

const (
	// VRSUp indicates Up
	VRSUp VerticalRateSign = 0
	// VRSDown indicates Down
	VRSDown VerticalRateSign = 1
)

// ToString returns a basic, but readable, representation of the field
func (bit VerticalRateSign) ToString() string {

	switch bit {
	case VRSUp:
		return "0 - up"
	case VRSDown:
		return "1 - down"
	default:
		return fmt.Sprintf("%v - Unknown code", bit)
	}
}

// ReadVerticalRateSign reads the VerticalRateSign from a 56 bits data field
func ReadVerticalRateSign(data []byte) VerticalRateSign {
	bits := (data[4] & 0x08) >> 3
	return VerticalRateSign(bits)
}
