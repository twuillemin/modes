package fields

import "fmt"

// SingleAntennaFlag is the single antenna flag definition
//
// Specified in Doc 9871 / A.2.3.2.5
type SingleAntennaFlag byte

const (
	// SAFDual indicates dual transmit antenna system.
	SAFDual SingleAntennaFlag = 0
	// SAFSingle indicates single transmit antenna system.
	SAFSingle SingleAntennaFlag = 1
)

// ToString returns a basic, but readable, representation of the field
func (saf SingleAntennaFlag) ToString() string {

	switch saf {
	case SAFDual:
		return "0 - dual transmit antenna system"
	case SAFSingle:
		return "1 - single transmit antenna system"
	default:
		return fmt.Sprintf("%v - Unknown code", saf)
	}
}

// ReadSingleAntennaFlag reads the SingleAntennaFlag from a 56 bits data field
func ReadSingleAntennaFlag(data []byte) SingleAntennaFlag {
	bits := data[0] & 0x01
	return SingleAntennaFlag(bits)
}
