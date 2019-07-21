package fields

import "fmt"

// VNAVModeEngaged is the VNAV Mode Engaged definition
//
// Specified in Doc 9871 / C.2.3.9.14
type VNAVModeEngaged byte

const (
	// VMENotEngaged indicates that VNAV Mode is NOT Active
	VMENotEngaged VNAVModeEngaged = 0
	// VMEngaged indicates that VNAV Mode is Active
	VMEngaged VNAVModeEngaged = 1
)

// ToString returns a basic, but readable, representation of the field
func (status VNAVModeEngaged) ToString() string {

	switch status {
	case VMENotEngaged:
		return "0 - VNAV mode is not active"
	case VMEngaged:
		return "1 - VNAV mode is active"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadVNAVModeEngaged reads the VNAVModeEngaged from a 56 bits data field
func ReadVNAVModeEngaged(data []byte) VNAVModeEngaged {
	bits := (data[6] & 0x80) >> 7
	return VNAVModeEngaged(bits)
}
