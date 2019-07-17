package fields

import "fmt"

// NICSupplement is the NIC Supplement definition
//
// Specified in Doc 9871 / B.2.3.10.6
type NICSupplement byte

const (
	// NICSZero indicates NICSupplement = 0
	NICSZero NICSupplement = 0
	// NICSOne indicates NICSupplement = 1
	NICSOne NICSupplement = 1
)

// ToString returns a basic, but readable, representation of the field
func (supplement NICSupplement) ToString() string {

	switch supplement {
	case NICSZero:
		return "0"
	case NICSOne:
		return "1"
	default:
		return fmt.Sprintf("%v - Unknown code", supplement)
	}
}

// ReadNICSupplement reads the NICSupplement from a 56 bits data field
func ReadNICSupplement(data []byte) NICSupplement {
	bits := (data[5] & 0x20) >> 5
	return NICSupplement(bits)
}
