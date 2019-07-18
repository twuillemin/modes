package fields

import "fmt"

// ApproachLandingOperationalCapabilityStatus is the Approach Landing Operational Capability Status
//
// Specified in Doc 9871 / A.2.3.11.9
type ApproachLandingOperationalCapabilityStatus byte

const (
	// ALOCSReserved0 is reserved
	ALOCSReserved0 ApproachLandingOperationalCapabilityStatus = 0
	// ALOCSReserved1 is reserved
	ALOCSReserved1 ApproachLandingOperationalCapabilityStatus = 1
	// ALOCSReserved2 is reserved
	ALOCSReserved2 ApproachLandingOperationalCapabilityStatus = 2
	// ALOCSReserved3 is reserved
	ALOCSReserved3 ApproachLandingOperationalCapabilityStatus = 3
	// ALOCSReserved4 is reserved
	ALOCSReserved4 ApproachLandingOperationalCapabilityStatus = 4
	// ALOCSReserved5 is reserved
	ALOCSReserved5 ApproachLandingOperationalCapabilityStatus = 5
	// ALOCSReserved6 is reserved
	ALOCSReserved6 ApproachLandingOperationalCapabilityStatus = 6
	// ALOCSReserved7 is reserved
	ALOCSReserved7 ApproachLandingOperationalCapabilityStatus = 7
	// ALOCSReserved8 is reserved
	ALOCSReserved8 ApproachLandingOperationalCapabilityStatus = 8
	// ALOCSReserved9 is reserved
	ALOCSReserved9 ApproachLandingOperationalCapabilityStatus = 9
	// ALOCSReserved10 is reserved
	ALOCSReserved10 ApproachLandingOperationalCapabilityStatus = 10
	// ALOCSReserved11 is reserved
	ALOCSReserved11 ApproachLandingOperationalCapabilityStatus = 11
	// ALOCSReserved12 is reserved
	ALOCSReserved12 ApproachLandingOperationalCapabilityStatus = 12
	// ALOCSReserved13 is reserved
	ALOCSReserved13 ApproachLandingOperationalCapabilityStatus = 13
	// ALOCSReserved14 is reserved
	ALOCSReserved14 ApproachLandingOperationalCapabilityStatus = 14
	// ALOCSReserved15 is reserved
	ALOCSReserved15 ApproachLandingOperationalCapabilityStatus = 15
)

// ToString returns a basic, but readable, representation of the field
func (status ApproachLandingOperationalCapabilityStatus) ToString() string {

	switch status {

	case ALOCSReserved0, ALOCSReserved1, ALOCSReserved2, ALOCSReserved3,
		ALOCSReserved4, ALOCSReserved5, ALOCSReserved6, ALOCSReserved7,
		ALOCSReserved8, ALOCSReserved9, ALOCSReserved10, ALOCSReserved11,
		ALOCSReserved12, ALOCSReserved13, ALOCSReserved14, ALOCSReserved15:

		return fmt.Sprintf("%v - reserved", status)

	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadApproachLandingOperationalCapabilityStatus reads the ApproachLandingOperationalCapabilityStatus from a 56 bits data field
func ReadApproachLandingOperationalCapabilityStatus(data []byte) ApproachLandingOperationalCapabilityStatus {
	bits := (data[4] & 0xF0) >> 4
	return ApproachLandingOperationalCapabilityStatus(bits)
}
