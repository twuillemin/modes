package fields

import "fmt"

// ApproachLandingOperationalCapabilities is the Approach Landing Operational Capabilities
//
// Specified in Doc 9871 / A.2.3.11.5
type ApproachLandingOperationalCapabilities byte

const (
	// ALOCReserved0 is reserved
	ALOCReserved0 ApproachLandingOperationalCapabilities = 0
	// ALOCReserved1 is reserved
	ALOCReserved1 ApproachLandingOperationalCapabilities = 1
	// ALOCReserved2 is reserved
	ALOCReserved2 ApproachLandingOperationalCapabilities = 2
	// ALOCReserved3 is reserved
	ALOCReserved3 ApproachLandingOperationalCapabilities = 3
	// ALOCReserved4 is reserved
	ALOCReserved4 ApproachLandingOperationalCapabilities = 4
	// ALOCReserved5 is reserved
	ALOCReserved5 ApproachLandingOperationalCapabilities = 5
	// ALOCReserved6 is reserved
	ALOCReserved6 ApproachLandingOperationalCapabilities = 6
	// ALOCReserved7 is reserved
	ALOCReserved7 ApproachLandingOperationalCapabilities = 7
	// ALOCReserved8 is reserved
	ALOCReserved8 ApproachLandingOperationalCapabilities = 8
	// ALOCReserved9 is reserved
	ALOCReserved9 ApproachLandingOperationalCapabilities = 9
	// ALOCReserved10 is reserved
	ALOCReserved10 ApproachLandingOperationalCapabilities = 10
	// ALOCReserved11 is reserved
	ALOCReserved11 ApproachLandingOperationalCapabilities = 11
	// ALOCReserved12 is reserved
	ALOCReserved12 ApproachLandingOperationalCapabilities = 12
	// ALOCReserved13 is reserved
	ALOCReserved13 ApproachLandingOperationalCapabilities = 13
	// ALOCReserved14 is reserved
	ALOCReserved14 ApproachLandingOperationalCapabilities = 14
	// ALOCReserved15 is reserved
	ALOCReserved15 ApproachLandingOperationalCapabilities = 15
)

// ToString returns a basic, but readable, representation of the field
func (capabilities ApproachLandingOperationalCapabilities) ToString() string {

	switch capabilities {

	case ALOCReserved0, ALOCReserved1, ALOCReserved2, ALOCReserved3,
		ALOCReserved4, ALOCReserved5, ALOCReserved6, ALOCReserved7,
		ALOCReserved8, ALOCReserved9, ALOCReserved10, ALOCReserved11,
		ALOCReserved12, ALOCReserved13, ALOCReserved14, ALOCReserved15:

		return fmt.Sprintf("%v - reserved", capabilities)

	default:
		return fmt.Sprintf("%v - Unknown code", capabilities)
	}
}

// ReadApproachLandingOperationalCapabilities reads the ApproachLandingOperationalCapabilities from a 56 bits data field
func ReadApproachLandingOperationalCapabilities(data []byte) ApproachLandingOperationalCapabilities {
	bits := (data[2] & 0xF0) >> 4
	return ApproachLandingOperationalCapabilities(bits)
}
