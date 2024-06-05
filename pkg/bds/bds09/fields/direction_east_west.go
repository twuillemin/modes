package fields

import "fmt"

// DirectionEastWest is the Direction Bit EW Velocity definition
//
// Specified in Doc 9871 / Table A-2-9
type DirectionEastWest byte

const (
	// DEWEast indicates East
	DEWEast DirectionEastWest = 0
	// DEWWest indicates West
	DEWWest DirectionEastWest = 1
)

// ToString returns a basic, but readable, representation of the field
func (bit DirectionEastWest) ToString() string {

	switch bit {
	case DEWEast:
		return "0 - east"
	case DEWWest:
		return "1 - west"
	default:
		return fmt.Sprintf("%v - Unknown code", bit)
	}
}

// ReadDirectionEastWest reads the DirectionEastWest from a 56 bits data field
func ReadDirectionEastWest(data []byte) DirectionEastWest {
	bits := (data[1] & 0x04) >> 2
	return DirectionEastWest(bits)
}
