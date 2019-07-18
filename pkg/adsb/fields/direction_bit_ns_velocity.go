package fields

import "fmt"

// DirectionBitNSVelocity is the Direction Bit NS Velocity definition
//
// Specified in Doc 9871 / Table A-2-9
type DirectionBitNSVelocity byte

const (
	// DBVNorth indicates North
	DBVNorth DirectionBitNSVelocity = 0
	// DBVSouth indicates South
	DBVSouth DirectionBitNSVelocity = 1
)

// ToString returns a basic, but readable, representation of the field
func (bit DirectionBitNSVelocity) ToString() string {

	switch bit {
	case DBVNorth:
		return "0 - north"
	case DBVSouth:
		return "1 - south"
	default:
		return fmt.Sprintf("%v - Unknown code", bit)
	}
}

// ReadDirectionBitNSVelocity reads the DirectionBitNSVelocity from a 56 bits data field
func ReadDirectionBitNSVelocity(data []byte) DirectionBitNSVelocity {
	bits := (data[3] & 0x80) >> 7
	return DirectionBitNSVelocity(bits)
}
