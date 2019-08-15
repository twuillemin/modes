package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
)

// Subtype is Subtype of the message
//
// Specified in Doc 9871 / Table A-2-9
type Subtype byte

const (
	// SubtypeNoInformation indicates No Information
	SubtypeNoInformation Subtype = 0
	// SubtypeEmergencyPriorityStatus indicates Emergency/priority status
	SubtypeEmergencyPriorityStatus Subtype = 1
	// SubtypeRABroadcast indicates TCAS/ACAS RA Broadcast
	SubtypeRABroadcast Subtype = 2
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
	case SubtypeNoInformation:
		return "0 - no information"
	case SubtypeEmergencyPriorityStatus:
		return "1 - emergency/priority status"
	case SubtypeRABroadcast:
		return "2 - TCAS/ACAS RA Broadcast"
	case SubtypeReserved3, SubtypeReserved4, SubtypeReserved5, SubtypeReserved6, SubtypeReserved7:
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

// ToSubtype returns the ADSB generic subtype
func (code Subtype) ToSubtype() adsb.Subtype {
	return code
}
