package fields

import "fmt"

// StatusOfTargetSource is the status of mode bits
//
// Specified in Doc 9871 / Table A-2-64
type StatusOfTargetSource byte

const (
	// STSNoSourceInformation indicates No source information provided.
	STSNoSourceInformation StatusOfTargetSource = 0
	// STSInformationProvided indicates Source information deliberately provided.
	STSInformationProvided StatusOfTargetSource = 1
)

// ToString returns a basic, but readable, representation of the field
func (smb StatusOfTargetSource) ToString() string {

	switch smb {
	case STSNoSourceInformation:
		return "0 - No source information provided"
	case STSInformationProvided:
		return "1 - Source information deliberately provided"
	default:
		return fmt.Sprintf("%v - Unknown code", smb)
	}
}

// ReadStatusOfTargetSource reads the StatusOfTargetSource from a 56 bits data field
func ReadStatusOfTargetSource(data []byte) StatusOfTargetSource {
	bits := (data[6] & 0x04) >> 2
	return StatusOfTargetSource(bits)
}
