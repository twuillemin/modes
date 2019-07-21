package fields

import "fmt"

// NICSupplementC is the NIC Supplement definition
//
// Specified in Doc 9871 / C.2.3.10.3
type NICSupplementC byte

const (
	// NICCSZero indicates NICSupplementC = 0
	NICCSZero NICSupplementC = 0
	// NICCSOne indicates NICSupplementC = 1
	NICCSOne NICSupplementC = 1
)

// ToString returns a basic, but readable, representation of the field
func (supplement NICSupplementC) ToString() string {

	switch supplement {
	case NICCSZero:
		return "0"
	case NICCSOne:
		return "1"
	default:
		return fmt.Sprintf("%v - Unknown code", supplement)
	}
}

// ReadNICSupplementC reads the NICSupplementC from a 56 bits data field
func ReadNICSupplementC(data []byte) NICSupplementC {
	bits := (data[2] & 0x10) >> 4
	return NICSupplementC(bits)
}
