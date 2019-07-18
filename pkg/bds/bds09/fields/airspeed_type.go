package fields

import "fmt"

// AirspeedType is the Direction Airspeed Type definition
//
// Specified in Doc 9871 / Table A-2-9
type AirspeedType byte

const (
	// ATIAS indicates IAS
	ATIAS AirspeedType = 0
	// ATTAS indicates TAS
	ATTAS AirspeedType = 1
)

// ToString returns a basic, but readable, representation of the field
func (bit AirspeedType) ToString() string {

	switch bit {
	case ATIAS:
		return "0 - IAS"
	case ATTAS:
		return "1 - TAS"
	default:
		return fmt.Sprintf("%v - Unknown code", bit)
	}
}

// ReadAirspeedType reads the AirspeedType from a 56 bits data field
func ReadAirspeedType(data []byte) AirspeedType {
	bits := (data[3] & 0x80) >> 7
	return AirspeedType(bits)
}
