package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
)

// Subtype is the Airborne Velocity Subtype Code definition
//
// Specified in Doc 9871 / Table A-2-9 and Table B-2-98
type Subtype byte

const (
	// Subtype0 indicates Subtype 0
	Subtype0 Subtype = 0
	// Subtype1 indicates Subtype 1
	Subtype1 Subtype = 1
	// SubtypeReserved2 is reserved
	SubtypeReserved2 Subtype = 2
	// SubtypeReserved3 is reserved
	SubtypeReserved3 Subtype = 3
)

// ToString returns a basic, but readable, representation of the field
func (subtype Subtype) ToString() string {

	switch subtype {
	case Subtype0:
		return "0 - Subtype 0"
	case Subtype1:
		return "1 - Subtype 1"
	case SubtypeReserved2, SubtypeReserved3:
		return fmt.Sprintf("%v - reserved", subtype)
	default:
		return fmt.Sprintf("%v - Unknown code", subtype)
	}
}

// ReadSubtype reads the Subtype from a 56 bits data field
func ReadSubtype(data []byte) Subtype {
	bits := (data[0] & 0x06) >> 1
	return Subtype(bits)
}

// ToSubtype returns the subtype itself
func (subtype Subtype) ToSubtype() adsb.Subtype {
	return subtype
}
