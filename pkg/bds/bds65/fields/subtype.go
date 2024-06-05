package fields

import (
	"fmt"
)

// Subtype is Subtype of the message
//
// Specified in Doc 9871 / Figure C-10
type Subtype byte

const (
	// SubtypeAirborne indicates airborne data
	SubtypeAirborne Subtype = 0
	// SubtypeSurface indicates surface data
	SubtypeSurface Subtype = 1
)

// ToString returns a basic, but readable, representation of the field
func (subtype Subtype) ToString() string {

	switch subtype {
	case SubtypeAirborne:
		return "0 - SubtypeAirborne Status Message"
	case SubtypeSurface:
		return "1 - SubtypeSurface Status Message"
	default:
		return fmt.Sprintf("%v - Unknown code", subtype)
	}
}

// ReadSubtype reads the Subtype from a 56 bits data field
func ReadSubtype(data []byte) Subtype {
	bits := data[0] & 0x07
	return Subtype(bits)
}
