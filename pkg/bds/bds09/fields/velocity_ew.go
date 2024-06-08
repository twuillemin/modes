package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

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
		return "0 - East"
	case DEWWest:
		return "1 - West"
	default:
		return fmt.Sprintf("%v - Unknown code", bit)
	}
}

// ReadDirectionEastWest reads the DirectionEastWest from a 56 bits data field
func ReadDirectionEastWest(data []byte) DirectionEastWest {
	bits := (data[1] & 0x04) >> 2
	return DirectionEastWest(bits)
}

// ReadVelocityEWNormal reads the Velocity EW Normal from a 56 bits data field
func ReadVelocityEWNormal(data []byte) (uint16, NumericValueStatus) {
	byte1 := data[1] & 0x03
	byte2 := data[2]
	velocity := bitutils.Pack2Bytes(byte1, byte2)

	if velocity == 0 {
		return 0, NVSNoInformation
	} else if velocity >= 1023 {
		return 1023, NVSMaximum
	}

	return velocity - 1, NVSRegular
}

// ReadVelocityEWSupersonic reads the Velocity EW Supersonic from a 56 bits data field
func ReadVelocityEWSupersonic(data []byte) (uint16, NumericValueStatus) {
	byte1 := data[1] & 0x03
	byte2 := data[2]
	velocity := bitutils.Pack2Bytes(byte1, byte2)

	if velocity == 0 {
		return 0, NVSNoInformation
	} else if velocity >= 1023 {
		return 4088, NVSMaximum
	}

	return (velocity - 1) * 4, NVSRegular
}
