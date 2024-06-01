package fields

import "fmt"

// SingleAntennaFlag is the Single Antenna Flag definition
//
// Specified in Doc 9871 / C.2.3.10.4
type SingleAntennaFlag byte

const (
	// SAFTwoAntennas indicates Ident switch not active
	SAFTwoAntennas SingleAntennaFlag = 0
	// SAFOneAntenna indicates Ident switch active - retained for 18 ±1 seconds
	SAFOneAntenna SingleAntennaFlag = 1
)

// ToString returns a basic, but readable, representation of the field
func (status SingleAntennaFlag) ToString() string {

	switch status {
	case SAFTwoAntennas:
		return "0 - Systems with two functioning antennas"
	case SAFOneAntenna:
		return "1 - Systems that use only one antenna"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadSingleAntennaFlag reads the SingleAntennaFlag from a 56 bits data field
func ReadSingleAntennaFlag(data []byte) SingleAntennaFlag {
	bits := (data[3] & 0x04) >> 2
	return SingleAntennaFlag(bits)
}
