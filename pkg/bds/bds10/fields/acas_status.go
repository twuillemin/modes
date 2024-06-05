package fields

import "fmt"

// ACASStatus is the ACAS status definition
//
// Specified in Doc 9871 / D.2.4.1
type ACASStatus byte

const (
	// ACASFailedOrStandBy indicates that ACAS has failed or is on standby.
	ACASFailedOrStandBy ACASStatus = 0
	// ACASOperational indicates that ACAS is operational .
	ACASOperational ACASStatus = 1
)

// ToString returns a basic, but readable, representation of the field
func (ac ACASStatus) ToString() string {

	switch ac {
	case ACASFailedOrStandBy:
		return "0 - ACAS Failed or On Standby"
	case ACASOperational:
		return "1 - ACAS Operational"
	default:
		return fmt.Sprintf("%v - Unknown code", ac)
	}
}

// ReadACASStatus reads the ACASStatus from a 56 bits data field
func ReadACASStatus(data []byte) ACASStatus {
	bits := data[1] & 0x01
	return ACASStatus(bits)
}
