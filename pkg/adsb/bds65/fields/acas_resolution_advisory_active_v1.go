package fields

import "fmt"

// ACASResolutionAdvisoryActiveV1 is the ACAS Resolution Advisory Active definition
//
// Specified in Doc 9871 / B.2.3.10.4
type ACASResolutionAdvisoryActiveV1 byte

const (
	// ARAV1NotActive indicates ACAS II or ACAS RA not active
	ARAV1NotActive ACASResolutionAdvisoryActiveV1 = 0
	// ARAV1Active indicates ACAS RA is active
	ARAV1Active ACASResolutionAdvisoryActiveV1 = 1
)

// ToString returns a basic, but readable, representation of the field
func (status ACASResolutionAdvisoryActiveV1) ToString() string {

	switch status {
	case ARAV1NotActive:
		return "0 - ACAS II or ACAS RA not active"
	case ARAV1Active:
		return "1 - ACAS RA is active"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadACASResolutionAdvisoryActiveV1 reads the ACASResolutionAdvisoryActiveV1 from a 56 bits data field
func ReadACASResolutionAdvisoryActiveV1(data []byte) ACASResolutionAdvisoryActiveV1 {
	bits := (data[3] & 0x20) >> 5
	return ACASResolutionAdvisoryActiveV1(bits)
}
