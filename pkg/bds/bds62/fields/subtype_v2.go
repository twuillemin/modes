package fields

import "fmt"

// SubtypeV2 is the Airborne Velocity Subtype Code definition
//
// Specified in Doc 9871 / Table A-2-9
type SubtypeV2 byte

const (
	// SubtypeV2Subtype0 indicates Subtype 0
	SubtypeV2Subtype0 SubtypeV2 = 0
	// SubtypeV2Subtype1 indicates Subtype 1
	SubtypeV2Subtype1 SubtypeV2 = 1
	// SubtypeV2Reserved2 is reserved
	SubtypeV2Reserved2 SubtypeV2 = 2
	// SubtypeV2Reserved3 is reserved
	SubtypeV2Reserved3 SubtypeV2 = 3
	// SubtypeV2Reserved4 is reserved
	SubtypeV2Reserved4 SubtypeV2 = 4
	// SubtypeV2Reserved5 is reserved
	SubtypeV2Reserved5 SubtypeV2 = 5
	// SubtypeV2Reserved6 is reserved
	SubtypeV2Reserved6 SubtypeV2 = 6
	// SubtypeV2Reserved7 is reserved
	SubtypeV2Reserved7 SubtypeV2 = 7
)

// ToSubtype returns the subtype itself
func (subtype SubtypeV2) ToSubtype() Subtype {
	return subtype
}

// ToString returns a basic, but readable, representation of the field
func (subtype SubtypeV2) ToString() string {

	switch subtype {
	case SubtypeV2Subtype0:
		return "0 - Subtype 0"
	case SubtypeV2Subtype1:
		return "0 - Subtype 1"
	case SubtypeV2Reserved2, SubtypeV2Reserved3,
		SubtypeV2Reserved4, SubtypeV2Reserved5, SubtypeV2Reserved6,
		SubtypeV2Reserved7:
		return fmt.Sprintf("%v - reserved", subtype)
	default:
		return fmt.Sprintf("%v - Unknown code", subtype)
	}
}

// ReadSubtypeV2 reads the Subtype from a 56 bits data field
func ReadSubtypeV2(data []byte) SubtypeV2 {
	bits := (data[0] & 0x06) >> 1
	return SubtypeV2(bits)
}
