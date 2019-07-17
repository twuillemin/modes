package fields

import "fmt"

// ACASResolutionAdvisoryActive is the ACAS Resolution Advisory Active definition
//
// Specified in Doc 9871 / B.2.3.10.4
type ACASResolutionAdvisoryActive byte

const (
	// ARAANotActive indicates ACAS II or ACAS RA not active
	ARAANotActive ACASResolutionAdvisoryActive = 0
	// ARAAActive indicates ACAS RA is active
	ARAAActive ACASResolutionAdvisoryActive = 1
)

// ToString returns a basic, but readable, representation of the field
func (status ACASResolutionAdvisoryActive) ToString() string {

	switch status {
	case ARAANotActive:
		return "0 - ACAS II or ACAS RA not active"
	case ARAAActive:
		return "1 - ACAS RA is active"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadACASResolutionAdvisoryActive reads the ACASResolutionAdvisoryActive from a 56 bits data field
func ReadACASResolutionAdvisoryActive(data []byte) ACASResolutionAdvisoryActive {
	bits := (data[3] & 0x20) >> 5
	return ACASResolutionAdvisoryActive(bits)
}
