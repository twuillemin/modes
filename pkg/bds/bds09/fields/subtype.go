package fields

import (
	"fmt"
)

// Subtype is the Airborne Velocity Subtype Code definition
//
// Specified in Doc 9871 / Table A-2-9
type Subtype byte

const (
	// SubtypeGroundSpeedNormal indicates Surface Status Message
	SubtypeGroundSpeedNormal Subtype = 1
	// SubtypeGroundSpeedSupersonic is reserved
	SubtypeGroundSpeedSupersonic Subtype = 2
	// SubtypeAirspeedNormal is reserved
	SubtypeAirspeedNormal Subtype = 3
	// SubtypeAirspeedSupersonic is reserved
	SubtypeAirspeedSupersonic Subtype = 4
)

// ToString returns a basic, but readable, representation of the field
func (code Subtype) ToString() string {

	switch code {
	case SubtypeGroundSpeedNormal:
		return "1 - GroundSpeed / Normal"
	case SubtypeGroundSpeedSupersonic:
		return "2 - GroundSpeed / Supersonic"
	case SubtypeAirspeedNormal:
		return "3 - Airspeed,Heading / Normal"
	case SubtypeAirspeedSupersonic:
		return "4 - Airspeed,Heading / Supersonic"
	default:
		return fmt.Sprintf("%v - Unknown code", code)
	}
}

// ReadSubtype reads the Subtype from a 56 bits data field
func ReadSubtype(data []byte) Subtype {
	bits := data[0] & 0x07
	return Subtype(bits)
}
