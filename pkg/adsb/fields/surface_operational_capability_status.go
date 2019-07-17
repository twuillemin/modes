package fields

import "fmt"

// SurfaceOperationalCapabilityStatus is the Surface Operational Capability Status format definition
//
// Specified in Doc 9871 / A.2.3.11.10
type SurfaceOperationalCapabilityStatus byte

const (
	// SOCSReserved0 is reserved
	SOCSReserved0 SurfaceOperationalCapabilityStatus = 0
	// SOCSReserved1 is reserved
	SOCSReserved1 SurfaceOperationalCapabilityStatus = 1
	// SOCSReserved2 is reserved
	SOCSReserved2 SurfaceOperationalCapabilityStatus = 2
	// SOCSReserved3 is reserved
	SOCSReserved3 SurfaceOperationalCapabilityStatus = 3
	// SOCSReserved4 is reserved
	SOCSReserved4 SurfaceOperationalCapabilityStatus = 4
	// SOCSReserved5 is reserved
	SOCSReserved5 SurfaceOperationalCapabilityStatus = 5
	// SOCSReserved6 is reserved
	SOCSReserved6 SurfaceOperationalCapabilityStatus = 6
	// SOCSReserved7 is reserved
	SOCSReserved7 SurfaceOperationalCapabilityStatus = 7
	// SOCSReserved8 is reserved
	SOCSReserved8 SurfaceOperationalCapabilityStatus = 8
	// SOCSReserved9 is reserved
	SOCSReserved9 SurfaceOperationalCapabilityStatus = 9
	// SOCSReserved10 is reserved
	SOCSReserved10 SurfaceOperationalCapabilityStatus = 10
	// SOCSReserved11 is reserved
	SOCSReserved11 SurfaceOperationalCapabilityStatus = 11
	// SOCSReserved12 is reserved
	SOCSReserved12 SurfaceOperationalCapabilityStatus = 12
	// SOCSReserved13 is reserved
	SOCSReserved13 SurfaceOperationalCapabilityStatus = 13
	// SOCSReserved14 is reserved
	SOCSReserved14 SurfaceOperationalCapabilityStatus = 14
	// SOCSReserved15 is reserved
	SOCSReserved15 SurfaceOperationalCapabilityStatus = 15
)

// ToString returns a basic, but readable, representation of the field
func (capabilities SurfaceOperationalCapabilityStatus) ToString() string {

	if capabilities <= 15 {
		return fmt.Sprintf("%v - reserved", capabilities)
	} else {
		return fmt.Sprintf("%v - Unknown code", capabilities)
	}
}

// ReadSurfaceOperationalCapabilityStatus reads the SurfaceOperationalCapabilityStatus from a 56 bits data field
func ReadSurfaceOperationalCapabilityStatus(data []byte) SurfaceOperationalCapabilityStatus {
	bits := data[4] & 0x0F
	return SurfaceOperationalCapabilityStatus(bits)
}
