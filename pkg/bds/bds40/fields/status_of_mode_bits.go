package fields

import "fmt"

// StatusOfModeBits is the status of mode bits
//
// Specified in Doc 9871 / Table A-2-64
type StatusOfModeBits byte

const (
	// SMBNoModeInformation indicates a transponder Level 2 to 4.
	SMBNoModeInformation StatusOfModeBits = 0
	// SMBInformationProvided indicates a transponder Level 5.
	SMBInformationProvided StatusOfModeBits = 1
)

// ToString returns a basic, but readable, representation of the field
func (smb StatusOfModeBits) ToString() string {

	switch smb {
	case SMBNoModeInformation:
		return "0 - No mode information provided"
	case SMBInformationProvided:
		return "1 - Mode information deliberately provided"
	default:
		return fmt.Sprintf("%v - Unknown code", smb)
	}
}

// ReadStatusOfModeBits reads the StatusOfModeBits from a 56 bits data field
func ReadStatusOfModeBits(data []byte) StatusOfModeBits {
	bits := data[5] & 0x01
	return StatusOfModeBits(bits)
}
