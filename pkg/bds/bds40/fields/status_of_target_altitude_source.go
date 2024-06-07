package fields

import "fmt"

// StatusOfTargetAltitudeSource is the status of mode bits
//
// Specified in Doc 9871 / Table A-2-64
type StatusOfTargetAltitudeSource byte

const (
	// STASNoSourceInformation indicates No source information provided.
	STASNoSourceInformation StatusOfTargetAltitudeSource = 0
	// STASInformationProvided indicates Source information deliberately provided.
	STASInformationProvided StatusOfTargetAltitudeSource = 1
)

// ToString returns a basic, but readable, representation of the field
func (smb StatusOfTargetAltitudeSource) ToString() string {

	switch smb {
	case STASNoSourceInformation:
		return "0 - No source information provided"
	case STASInformationProvided:
		return "1 - Source information deliberately provided"
	default:
		return fmt.Sprintf("%v - Unknown code", smb)
	}
}

// ReadStatusOfTargetAltitudeSource reads the StatusOfTargetAltitudeSource from a 56 bits data field
func ReadStatusOfTargetAltitudeSource(data []byte) StatusOfTargetAltitudeSource {
	bits := (data[6] & 0x04) >> 2
	return StatusOfTargetAltitudeSource(bits)
}
