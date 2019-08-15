package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
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
	// SubtypeReserved2 is reserved
	SubtypeReserved2 Subtype = 2
	// SubtypeReserved3 is reserved
	SubtypeReserved3 Subtype = 3
	// SubtypeReserved4 is reserved
	SubtypeReserved4 Subtype = 4
	// SubtypeReserved5 is reserved
	SubtypeReserved5 Subtype = 5
	// SubtypeReserved6 is reserved
	SubtypeReserved6 Subtype = 6
	// SubtypeReserved7 is reserved
	SubtypeReserved7 Subtype = 7
)

// ToString returns a basic, but readable, representation of the field
func (subtype Subtype) ToString() string {

	switch subtype {
	case SubtypeAirborne:
		return "0 - SubtypeAirborne Status Message"
	case SubtypeSurface:
		return "1 - SubtypeSurface Status Message"
	case SubtypeReserved2, SubtypeReserved3, SubtypeReserved4,
		SubtypeReserved5, SubtypeReserved6, SubtypeReserved7:
		return fmt.Sprintf("%v - reserved", subtype)
	default:
		return fmt.Sprintf("%v - Unknown code", subtype)
	}
}

// ReadSubtype reads the Subtype from a 56 bits data field
func ReadSubtype(data []byte) Subtype {
	bits := data[0] & 0x07
	return Subtype(bits)
}

// ToSubtype returns the subtype itself
func (subtype Subtype) ToSubtype() adsb.Subtype {
	return subtype
}
