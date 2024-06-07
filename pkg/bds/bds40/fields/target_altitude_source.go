package fields

import "fmt"

// TargetAltitudeSource is the status of mode bits
//
// Specified in Doc 9871 / Table A-2-64
type TargetAltitudeSource byte

const (
	// TASUnknown indicates Unknown.
	TASUnknown TargetAltitudeSource = 0
	// TASAircraftAltitude indicates Aircraft altitude.
	TASAircraftAltitude TargetAltitudeSource = 1
	// TASFCUMCPSelectedAltitude indicates FCU/MCP selected altitude.
	TASFCUMCPSelectedAltitude TargetAltitudeSource = 2
	// TASFMSSelectedAltitude indicates FMS selected altitude.
	TASFMSSelectedAltitude TargetAltitudeSource = 3
)

// ToString returns a basic, but readable, representation of the field
func (smb TargetAltitudeSource) ToString() string {

	switch smb {
	case TASUnknown:
		return "0 - Unknown"
	case TASAircraftAltitude:
		return "1 - Aircraft altitude"
	case TASFCUMCPSelectedAltitude:
		return "2 - FCU/MCP selected altitude"
	case TASFMSSelectedAltitude:
		return "3 - FMS selected altitude"
	default:
		return fmt.Sprintf("%v - Unknown code", smb)
	}
}

// ReadTargetAltitudeSource reads the TargetAltitudeSource from a 56 bits data field
func ReadTargetAltitudeSource(data []byte) TargetAltitudeSource {
	bits := data[6] & 0x03
	return TargetAltitudeSource(bits)
}
