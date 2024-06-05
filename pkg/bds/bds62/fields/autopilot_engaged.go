package fields

import "fmt"

// AutopilotEngaged is the Autopilot Engaged definition
//
// Specified in Doc 9871 / C.2.3.9.13
type AutopilotEngaged byte

const (
	// AENotEngaged indicates that autopilot is not engaged (e.g., not actively coupled and flying the aircraft)
	AENotEngaged AutopilotEngaged = 0
	// AEEngaged indicates that autopilot is engaged (e.g., actively coupled and flying the aircraft)
	AEEngaged AutopilotEngaged = 1
)

// ToString returns a basic, but readable, representation of the field
func (status AutopilotEngaged) ToString() string {

	switch status {
	case AENotEngaged:
		return "0 - autopilot is not engaged (e.g., not actively coupled and flying the aircraft)"
	case AEEngaged:
		return "1 - autopilot is engaged (e.g., actively coupled and flying the aircraft)"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadAutopilotEngaged reads the AutopilotEngaged from a 56 bits data field
func ReadAutopilotEngaged(data []byte) AutopilotEngaged {
	bits := data[5] & 0x01
	return AutopilotEngaged(bits)
}
