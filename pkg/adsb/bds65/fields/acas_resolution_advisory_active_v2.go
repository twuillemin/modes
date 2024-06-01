package fields

import "fmt"

// ACASResolutionAdvisoryActiveV2 is the ACAS Resolution Advisory Active definition
//
// Specified in Doc 9871 / C.2.3.10.4
type ACASResolutionAdvisoryActiveV2 byte

const (
	// ARAV2NotActive indicates TCAS II or ACAS RA not active
	ARAV2NotActive ACASResolutionAdvisoryActiveV2 = 0
	// ARAV2Active indicates TCAS/ACAS RA is active
	ARAV2Active ACASResolutionAdvisoryActiveV2 = 1
)

// ToString returns a basic, but readable, representation of the field
func (status ACASResolutionAdvisoryActiveV2) ToString() string {

	switch status {
	case ARAV2NotActive:
		return "0 - TCAS II or ACAS RA not active"
	case ARAV2Active:
		return "1 - TCAS/ACAS RA is active"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadACASResolutionAdvisoryActiveV2 reads the ACASResolutionAdvisoryActiveV2 from a 56 bits data field
func ReadACASResolutionAdvisoryActiveV2(data []byte) ACASResolutionAdvisoryActiveV2 {
	bits := (data[3] & 0x20) >> 5
	return ACASResolutionAdvisoryActiveV2(bits)
}
