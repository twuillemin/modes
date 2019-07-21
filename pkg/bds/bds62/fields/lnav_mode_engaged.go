package fields

import "fmt"

// LNAVModeEngaged is the LNAV Mode Engaged definition
//
// Specified in Doc 9871 / C.2.3.9.19
type LNAVModeEngaged byte

const (
	// LMENotEngaged indicates that LNAV Mode is NOT Active
	LMENotEngaged LNAVModeEngaged = 0
	// LMEngaged indicates that LNAV Mode is Active
	LMEngaged LNAVModeEngaged = 1
)

// ToString returns a basic, but readable, representation of the field
func (status LNAVModeEngaged) ToString() string {

	switch status {
	case LMENotEngaged:
		return "0 - LNAV mode is not active"
	case LMEngaged:
		return "1 - LNAV mode is active"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadLNAVModeEngaged reads the LNAVModeEngaged from a 56 bits data field
func ReadLNAVModeEngaged(data []byte) LNAVModeEngaged {
	bits := (data[6] & 0x40) >> 2
	return LNAVModeEngaged(bits)
}
