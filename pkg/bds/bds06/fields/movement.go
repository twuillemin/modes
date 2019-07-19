package fields

import "fmt"

// Movement is the Movement definition
//
// Specified in Doc 9871 / A.2.3.3.1
type Movement byte

// GetStatus returns the status of the altitude
func (movement Movement) GetStatus() MovementStatus {
	if movement == 0 {
		return MSNoInformation
	} else if movement == 124 {
		return MSAboveMaximum
	} else if movement > 124 {
		return MSReserved
	}

	return MSValid
}

// ToString returns a basic, but readable, representation of the field
func (movement Movement) ToString() string {

	if movement == 0 {
		return "0 - no information available"
	} else if movement == 1 {
		return "1 - aircraft stopped (ground speed < 0.2315 km/h (0.125 kt))"
	} else if movement == 124 {
		return "124 - Ground speed ≥ 324.1 km/h (175 kt)"
	} else if movement > 124 {
		return fmt.Sprintf("%v - reserved", movement)
	}

	return fmt.Sprintf("%v km/h", movement.GetMovement())

}

// GetMovement returns the Movement in km/h. Note that the returned value will be 0 if movement status is
// MSNoInformation or MSReserved and returned 324.1 if movement status is MsAboveMaximum.
func (movement Movement) GetMovement() float32 {

	if movement == 0 || movement == 1 || movement > 124 {
		return 0
	} else if 2 <= movement && movement <= 8 {
		return 0.2315 + float32(2-movement)*0.2315
	} else if 9 <= movement && movement <= 12 {
		return 1.852 + float32(9-movement)*0.463
	} else if 13 <= movement && movement <= 38 {
		return 3.704 + float32(13-movement)*0.926
	} else if 39 <= movement && movement <= 93 {
		return 27.78 + float32(39-movement)*1.852
	} else if 94 <= movement && movement <= 108 {
		return 129.64 + float32(94-movement)*3.704
	} else if 109 <= movement && movement <= 123 {
		return 185.2 + float32(109-movement)*9.26
	}

	// Movement max
	return 324.1
}

// ReadMovement reads the Movement from a 56 bits data field
func ReadMovement(data []byte) Movement {
	bits := (data[0]&0x07)<<3 + (data[1]&0xF0)>>4
	return Movement(bits)
}
