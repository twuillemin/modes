package fields

import "fmt"

// SubtypeV0 is the Airborne Velocity Subtype Code definition
//
// Specified in Doc 9871 / Table A-2-9
type SubtypeV0 byte

const (
	// SubtypeV0NoInformation indicates No Information
	SubtypeV0NoInformation SubtypeV0 = 0
	// SubtypeV0EmergencyPriorityStatus indicates Emergency/priority status
	SubtypeV0EmergencyPriorityStatus SubtypeV0 = 1
	// SubtypeV0Reserved2 is reserved
	SubtypeV0Reserved2 SubtypeV0 = 2
	// SubtypeV0Reserved3 is reserved
	SubtypeV0Reserved3 SubtypeV0 = 3
	// SubtypeV0Reserved4 is reserved
	SubtypeV0Reserved4 SubtypeV0 = 4
	// SubtypeV0Reserved5 is reserved
	SubtypeV0Reserved5 SubtypeV0 = 5
	// SubtypeV0Reserved6 is reserved
	SubtypeV0Reserved6 SubtypeV0 = 6
	// SubtypeV0Reserved7 is reserved
	SubtypeV0Reserved7 SubtypeV0 = 7
)

// GetSubtype returns the subtype itself
func (subtype SubtypeV0) GetSubtype() Subtype {
	return subtype
}

// ToString returns a basic, but readable, representation of the field
func (subtype SubtypeV0) ToString() string {

	switch subtype {
	case SubtypeV0NoInformation:
		return "0 - no information"
	case SubtypeV0EmergencyPriorityStatus:
		return "1 - emergency/priority status"
	case SubtypeV0Reserved2, SubtypeV0Reserved3, SubtypeV0Reserved4,
		SubtypeV0Reserved5, SubtypeV0Reserved6, SubtypeV0Reserved7:
		return fmt.Sprintf("%v - reserved", subtype)
	default:
		return fmt.Sprintf("%v - Unknown code", subtype)
	}
}

// ReadSubtypeV0 reads the Subtype from a 56 bits data field
func ReadSubtypeV0(data []byte) SubtypeV0 {
	bits := data[0] & 0x07
	return SubtypeV0(bits)
}
