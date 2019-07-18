package fields

import "fmt"

// DirectionBitEWVelocity is the Direction Bit EW Velocity definition
//
// Specified in Doc 9871 / Table A-2-9
type DirectionBitEWVelocity byte

const (
	// DBVEast indicates East
	DBVEast DirectionBitEWVelocity = 0
	// DBVWest indicates West
	DBVWest DirectionBitEWVelocity = 1
)

// ToString returns a basic, but readable, representation of the field
func (bit DirectionBitEWVelocity) ToString() string {

	switch bit {
	case DBVEast:
		return "0 - east"
	case DBVWest:
		return "1 - west"
	default:
		return fmt.Sprintf("%v - Unknown code", bit)
	}
}

// ReadDirectionBitEWVelocity reads the DirectionBitEWVelocity from a 56 bits data field
func ReadDirectionBitEWVelocity(data []byte) DirectionBitEWVelocity {
	bits := (data[1] & 0x04) >> 2
	return DirectionBitEWVelocity(bits)
}
