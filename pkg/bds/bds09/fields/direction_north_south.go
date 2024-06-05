package fields

import "fmt"

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
