package fields

import "fmt"

// CapabilityModeCode is the Capability Mode Code definition
//
// Specified in Doc 9871 / B.2.3.9.15
type CapabilityModeCode byte

const (
	// CMCOperationalNoRAActive indicates TCAS/ACAS operational or unknown + No TCAS/ACAS Resolution Advisory active
	CMCOperationalNoRAActive CapabilityModeCode = 0
	// CMCOperationalRAActive indicates TCAS/ACAS operational or unknown + TCAS/ACAS Resolution Advisory active
	CMCOperationalRAActive CapabilityModeCode = 1
	// CMCNotOperationalNoRAActive indicates TCAS/ACAS not operational + No TCAS/ACAS Resolution Advisory active
	CMCNotOperationalNoRAActive CapabilityModeCode = 2
	// CMCNotOperationalRAActive indicates TCAS/ACAS not operational + TCAS/ACAS Resolution Advisory active
	CMCNotOperationalRAActive CapabilityModeCode = 3
)

// ToString returns a basic, but readable, representation of the field
func (capability CapabilityModeCode) ToString() string {

	switch capability {
	case CMCOperationalNoRAActive:
		return "0 - TCAS/ACAS operational or unknown + No TCAS/ACAS Resolution Advisory active"
	case CMCOperationalRAActive:
		return "1 - TCAS/ACAS operational or unknown + TCAS/ACAS Resolution Advisory active"
	case CMCNotOperationalNoRAActive:
		return "2 - TCAS/ACAS not operational + No TCAS/ACAS Resolution Advisory active"
	case CMCNotOperationalRAActive:
		return "3 - TCAS/ACAS not operational + TCAS/ACAS Resolution Advisory active"
	default:
		return fmt.Sprintf("%v - Unknown code", capability)
	}
}

// ReadCapabilityModeCode reads the CapabilityModeCode from a 56 bits data field
func ReadCapabilityModeCode(data []byte) CapabilityModeCode {
	bits := (data[6] & 0x18) >> 3
	return CapabilityModeCode(bits)
}
