package fields

import "fmt"

// AirspeedType is the Airspeed Type definition
//
// Specified in Doc 9871 / Table A-2-9
type AirspeedType byte

const (
	// ATIndicatedAirSpeed indicates IAS
	ATIndicatedAirSpeed AirspeedType = 0
	// ATTrueAirSpeed indicates TAS
	ATTrueAirSpeed AirspeedType = 1
)

// ToString returns a basic, but readable, representation of the field
func (bit AirspeedType) ToString() string {

	switch bit {
	case ATIndicatedAirSpeed:
		return "0 - IAS"
	case ATTrueAirSpeed:
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
