package fields

import "fmt"

// SubtypeV1 is the Airborne Velocity Subtype Code definition
//
// Specified in Doc 9871 / Table A-2-9
type SubtypeV1 byte

const (
	// SubtypeV1NoInformation indicates No Information
	SubtypeV1NoInformation SubtypeV1 = 0
	// SubtypeV1EmergencyPriorityStatus indicates Emergency/priority status and Mode A Code
	SubtypeV1EmergencyPriorityStatus SubtypeV1 = 1
	// SubtypeV1RABroadcast indicates TCAS/ACAS RA Broadcast
	SubtypeV1RABroadcast SubtypeV1 = 2
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
	case SubtypeV1NoInformation:
		return "0 - no information"
	case SubtypeV1EmergencyPriorityStatus:
		return "1 - emergency/priority status"
	case SubtypeV1RABroadcast:
		return "2 - TCAS/ACAS RA Broadcast"
	case SubtypeV1Reserved3, SubtypeV1Reserved4, SubtypeV1Reserved5, SubtypeV1Reserved6, SubtypeV1Reserved7:
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
