package fields

import "fmt"

// ACASOperational is the Not ACAS flag definition
//
// Specified in Doc 9871 / C.2.3.9.18
type ACASOperational byte

const (
	// AONotOperational indicates TCAS/ACAS is NOT Operational
	AONotOperational ACASOperational = 0
	// AOOperational indicates TCAS/ACAS IS Operational
	AOOperational ACASOperational = 1
)

// ToString returns a basic, but readable, representation of the field
func (status ACASOperational) ToString() string {

	switch status {
	case AONotOperational:
		return "0 - TCAS/ACAS is not operational"
	case AOOperational:
		return "1 - TCAS/ACAS is operational"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadACASOperational reads the ACASOperational from a 56 bits data field
func ReadACASOperational(data []byte) ACASOperational {
	bits := (data[6] & 0x08) >> 3
	return ACASOperational(bits)
}
