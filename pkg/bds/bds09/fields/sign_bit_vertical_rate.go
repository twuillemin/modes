package fields

import "fmt"

// SignBitVerticalRate is the Sign Bit for Vertical Rate definition
//
// Specified in Doc 9871 / Table A-2-9
type SignBitVerticalRate byte

const (
	// SBVRUp indicates Up
	SBVRUp SignBitVerticalRate = 0
	// SBVRDown indicates Down
	SBVRDown SignBitVerticalRate = 1
)

// ToString returns a basic, but readable, representation of the field
func (bit SignBitVerticalRate) ToString() string {

	switch bit {
	case SBVRUp:
		return "0 - up"
	case SBVRDown:
		return "1 - down"
	default:
		return fmt.Sprintf("%v - Unknown code", bit)
	}
}

// ReadSignBitVerticalRate reads the SignBitVerticalRate from a 56 bits data field
func ReadSignBitVerticalRate(data []byte) SignBitVerticalRate {
	bits := (data[4] & 0x08) >> 3
	return SignBitVerticalRate(bits)
}
