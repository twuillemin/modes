package fields

import "fmt"

// SelectedAltitudeType is the Selected Altitude Type definition
//
// Specified in Doc 9871 / C.2.3.9.3
type SelectedAltitudeType byte

const (
	// SATByMCPFCU indicates MCP/FCU (Mode Control Panel / Flight Control Unit)
	SATByMCPFCU SelectedAltitudeType = 0
	// SATByFMS indicates FMS (Flight Management System)
	SATByFMS SelectedAltitudeType = 1
)

// ToString returns a basic, but readable, representation of the field
func (status SelectedAltitudeType) ToString() string {

	switch status {
	case SATByMCPFCU:
		return "0 - MCP/FCU (Mode Control Panel / Flight Control Unit)"
	case SATByFMS:
		return "1 - FMS (Flight Management System)"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadSelectedAltitudeType reads the SelectedAltitudeType from a 56 bits data field
func ReadSelectedAltitudeType(data []byte) SelectedAltitudeType {
	bits := data[0] & 0x01
	return SelectedAltitudeType(bits)
}
