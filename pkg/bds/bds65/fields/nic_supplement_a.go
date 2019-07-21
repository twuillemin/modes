package fields

import "fmt"

// NICSupplementA is the NIC Supplement A definition
//
// Specified in Doc 9871 / B.2.3.10.6
type NICSupplementA byte

const (
	// NICAZero indicates NICSupplementA = 0
	NICAZero NICSupplementA = 0
	// NICAOne indicates NICSupplementA = 1
	NICAOne NICSupplementA = 1
)

// ToString returns a basic, but readable, representation of the field
func (supplement NICSupplementA) ToString() string {

	switch supplement {
	case NICAZero:
		return "0"
	case NICAOne:
		return "1"
	default:
		return fmt.Sprintf("%v - Unknown code", supplement)
	}
}

// ReadNICSupplementA reads the NICSupplementA from a 56 bits data field
func ReadNICSupplementA(data []byte) NICSupplementA {
	bits := (data[5] & 0x20) >> 5
	return NICSupplementA(bits)
}
