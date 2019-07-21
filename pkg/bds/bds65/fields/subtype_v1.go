package fields

import "fmt"

// SubtypeV1 is the Aircraft Operational Status subtypes
//
// Specified in Doc 9871
type SubtypeV1 byte

const (
	// SubtypeV1Airborne indicates airborne data
	SubtypeV1Airborne SubtypeV1 = 0
	// SubtypeV1Surface indicates surface data
	SubtypeV1Surface SubtypeV1 = 1
	// SubtypeV1Reserved2 is reserved
	SubtypeV1Reserved2 SubtypeV1 = 2
	// SubtypeV1Reserved3 is reserved
	SubtypeV1Reserved3 SubtypeV1 = 3
	// SubtypeV1Reserved4 is reserved
	SubtypeV1Reserved4 SubtypeV1 = 4
	// SubtypeV1Reserved5 is reserved
	SubtypeV1Reserved5 SubtypeV1 = 5
	// SubtypeV1Reserved6 is reserved
	SubtypeV1Reserved6 SubtypeV1 = 6
	// SubtypeV1Reserved7 is reserved
	SubtypeV1Reserved7 SubtypeV1 = 7
)

// ToSubtype returns the subtype itself
func (subtype SubtypeV1) ToSubtype() Subtype {
	return subtype
}

// ToString returns a basic, but readable, representation of the field
func (subtype SubtypeV1) ToString() string {

	switch subtype {
	case SubtypeV1Airborne:
		return "0 - Airborne Status Message"
	case SubtypeV1Surface:
		return "1 - Surface Status Message"
	case SubtypeV1Reserved2, SubtypeV1Reserved3,
		SubtypeV1Reserved4, SubtypeV1Reserved5, SubtypeV1Reserved6,
		SubtypeV1Reserved7:
		return fmt.Sprintf("%v - reserved", subtype)
	default:
		return fmt.Sprintf("%v - Unknown code", subtype)
	}
}

// ReadSubtypeV1 reads the Subtype from a 56 bits data field
func ReadSubtypeV1(data []byte) SubtypeV1 {
	bits := data[0] & 0x07
	return SubtypeV1(bits)
}
