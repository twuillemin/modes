package fields

import "fmt"

// Movement is the Movement definition
//
// Specified in Doc 9871 / A.2.3.3.1

// MovementStatus is the status of the Movement information
type MovementStatus int

const (
	// MSNoInformation indicates no information
	MSNoInformation MovementStatus = 0
	// MSAircraftStopped indicates that the aircraft is stopped
	MSAircraftStopped MovementStatus = 1
	// MSAboveMaximum indicates that the Movement is above the maximum
	MSAboveMaximum MovementStatus = 124
	// MSReservedDecelerating indicates that the value is reserved
	MSReservedDecelerating MovementStatus = 125
	// MSReservedAccelerating indicates that the value is reserved
	MSReservedAccelerating MovementStatus = 126
	// MSReservedBackingUp indicates that the value is reserved
	MSReservedBackingUp MovementStatus = 127
)

// ToString returns a basic, but readable, representation of the field
func (movementStatus MovementStatus) ToString() string {

	switch movementStatus {
	case MSNoInformation:
		return "0 - no information available"
	case MSAircraftStopped:
		return "1 - aircraft stopped (ground speed < 0.2315 km/h (0.125 kt))"
	case MSAboveMaximum:
		return "124 - ground speed ≥ 324.1 km/h (175 kt)"
	case MSReservedDecelerating:
		return "125 - reserved for A/C Decelerating"
	case MSReservedAccelerating:
		return "126 - reserved for A/C Accelerating"
	case MSReservedBackingUp:
		return "127 - reserved for Backing Up"
	default:
		return fmt.Sprintf("%v - valid", movementStatus)
	}
}

// GetSpeed returns the Movement in knots. Note that the returned value will be 0 if movement status is
// MSNoInformation or MSReserved* and returned 324.1 if movement status is MsAboveMaximum.
func (movementStatus MovementStatus) getSpeed() float32 {

	if movementStatus == 0 || movementStatus == 1 || movementStatus > 124 {
		return 0
	} else if 2 <= movementStatus && movementStatus <= 8 {
		return 0.125 + float32(movementStatus-2)*0.125
	} else if 9 <= movementStatus && movementStatus <= 12 {
		return 1 + float32(movementStatus-9)*0.25
	} else if 13 <= movementStatus && movementStatus <= 38 {
		return 2 + float32(movementStatus-13)*0.5
	} else if 39 <= movementStatus && movementStatus <= 93 {
		return 15 + float32(movementStatus-39)*1.0
	} else if 94 <= movementStatus && movementStatus <= 108 {
		return 70 + float32(movementStatus-94)*2.0
	} else if 109 <= movementStatus && movementStatus <= 123 {
		return 100 + float32(movementStatus-109)*5.0
	}

	// Movement max
	return 175
}

// ReadMovement reads the Movement from a 56 bits data field
func ReadMovement(data []byte) (float32, MovementStatus) {
	bits := (data[0]&0x07)<<4 + (data[1]&0xF0)>>4
	status := MovementStatus(bits)
	return status.getSpeed(), status
}
