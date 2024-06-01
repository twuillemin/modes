package fields

import "fmt"

// StatusMCPFCUBits is the Status of MCP/FCU Mode Bits definition
//
// Specified in Doc 9871 / C.2.3.9.12
type StatusMCPFCUBits byte

const (
	// SMFBNoInformationProvided indicates that No Mode Information is being provided in “ME” bits 48, 49, 50,
	// 52 or 54 (Message bits 80, 81, 82, 84 or 86)
	SMFBNoInformationProvided StatusMCPFCUBits = 0
	// SMFBInformationProvided indicates that the Mode Information is deliberately being provided in “ME” bits 48,
	// 49, 50, 52, or 54 (Message bits 80, 81, 82, 84, or 86)
	SMFBInformationProvided StatusMCPFCUBits = 1
)

// ToString returns a basic, but readable, representation of the field
func (status StatusMCPFCUBits) ToString() string {

	switch status {
	case SMFBNoInformationProvided:
		return "0 - no information provided for MCP/FCU mode bits"
	case SMFBInformationProvided:
		return "1 - information provided for MCP/FCU mode bits"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadStatusMCPFCUBits reads the StatusMCPFCUBits from a 56 bits data field
func ReadStatusMCPFCUBits(data []byte) StatusMCPFCUBits {
	bits := (data[5] & 0x02) >> 1
	return StatusMCPFCUBits(bits)
}
