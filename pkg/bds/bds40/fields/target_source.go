package fields

import "fmt"

// TargetSource is the status of mode bits
//
// Specified in Doc 9871 / Table A-2-64
type TargetSource byte

const (
	// TSUnknown indicates Unknown.
	TSUnknown TargetSource = 0
	// TSAircraftAltitude indicates Aircraft altitude.
	TSAircraftAltitude TargetSource = 1
	// TSFCUMCPSelectedAltitude indicates FCU/MCP selected altitude.
	TSFCUMCPSelectedAltitude TargetSource = 2
	// TSFMSSelectedAltitude indicates FMS selected altitude.
	TSFMSSelectedAltitude TargetSource = 3
)

// ToString returns a basic, but readable, representation of the field
func (smb TargetSource) ToString() string {

	switch smb {
	case TSUnknown:
		return "0 - Unknown"
	case TSAircraftAltitude:
		return "1 - Aircraft altitude"
	case TSFCUMCPSelectedAltitude:
		return "2 - FCU/MCP selected altitude"
	case TSFMSSelectedAltitude:
		return "3 - FMS selected altitude"
	default:
		return fmt.Sprintf("%v - Unknown code", smb)
	}
}

// ReadTargetSource reads the TargetSource from a 56 bits data field
func ReadTargetSource(data []byte) TargetSource {
	bits := data[6] & 0x03
	return TargetSource(bits)
}
