package fields

import "fmt"

// IdentSwitchActive is the Ident Switch Active definition
//
// Specified in Doc 9871 / B.2.3.10.4
type IdentSwitchActive byte

const (
	// ISANotActive indicates Ident switch not active
	ISANotActive IdentSwitchActive = 0
	// ISAActive indicates Ident switch active - retained for 18 ±1 seconds
	ISAActive IdentSwitchActive = 1
)

// ToString returns a basic, but readable, representation of the field
func (status IdentSwitchActive) ToString() string {

	switch status {
	case ISANotActive:
		return "0 - Ident switch not active"
	case ISAActive:
		return "1 - Ident switch active - retained for 18 +/-1 seconds"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadIdentSwitchActive reads the IdentSwitchActive from a 56 bits data field
func ReadIdentSwitchActive(data []byte) IdentSwitchActive {
	bits := (data[3] & 0x10) >> 4
	return IdentSwitchActive(bits)
}
