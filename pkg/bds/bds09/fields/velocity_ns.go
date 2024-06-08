package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// DirectionNorthSouth is the Direction Bit NS Velocity definition
//
// Specified in Doc 9871 / Table A-2-9
type DirectionNorthSouth byte

const (
	// DNSNorth indicates North
	DNSNorth DirectionNorthSouth = 0
	// DNSSouth indicates South
	DNSSouth DirectionNorthSouth = 1
)

// ToString returns a basic, but readable, representation of the field
func (bit DirectionNorthSouth) ToString() string {

	switch bit {
	case DNSNorth:
		return "0 - north"
	case DNSSouth:
		return "1 - south"
	default:
		return fmt.Sprintf("%v - Unknown code", bit)
	}
}

// ReadDirectionNorthSouth reads the DirectionNorthSouth from a 56 bits data field
func ReadDirectionNorthSouth(data []byte) DirectionNorthSouth {
	bits := (data[3] & 0x80) >> 7
	return DirectionNorthSouth(bits)
}

// ReadVelocityNSNormal reads the Velocity NS Normal from a 56 bits data field
func ReadVelocityNSNormal(data []byte) (uint16, NumericValueStatus) {
	byte1 := data[3] & 0x7F
	byte2 := data[4] & 0xE0
	velocity := bitutils.Pack2Bytes(byte1, byte2) >> 5

	if velocity == 0 {
		return 0, NVSNoInformation
	} else if velocity >= 1023 {
		return 1023, NVSMaximum
	}

	return velocity - 1, NVSRegular
}

// ReadVelocityNSSupersonic reads the Velocity NS Supersonic from a 56 bits data field
func ReadVelocityNSSupersonic(data []byte) (uint16, NumericValueStatus) {
	byte1 := data[3] & 0x7F
	byte2 := data[4] & 0xE0
	velocity := bitutils.Pack2Bytes(byte1, byte2) >> 5

	if velocity == 0 {
		return 0, NVSNoInformation
	} else if velocity >= 1023 {
		return 4088, NVSMaximum
	}

	return (velocity - 1) * 4, NVSRegular
}
