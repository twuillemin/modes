package fields

import "fmt"

// ApproachModeEngaged is the Altitude Hold Mode Engaged definition
//
// Specified in Doc 9871 / C.2.3.9.15
type ApproachModeEngaged byte

const (
	// AMENotEngaged indicates that approach mode is not active
	AMENotEngaged ApproachModeEngaged = 0
	// AMEngaged indicates that approach mode is active
	AMEngaged ApproachModeEngaged = 1
)

// ToString returns a basic, but readable, representation of the field
func (status ApproachModeEngaged) ToString() string {

	switch status {
	case AMENotEngaged:
		return "0 - approach mode is not active"
	case AMEngaged:
		return "1 - approach mode is active"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadApproachModeEngaged reads the ApproachModeEngaged from a 56 bits data field
func ReadApproachModeEngaged(data []byte) ApproachModeEngaged {
	bits := (data[6] & 0x01) >> 4
	return ApproachModeEngaged(bits)
}
