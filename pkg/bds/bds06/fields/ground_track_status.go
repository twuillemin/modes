package fields

import "fmt"

// GroundTrackStatus is the Ground Track Status definition
//
// Specified in Doc 9871 / Table A-2-6
type GroundTrackStatus byte

const (
	// GTSInvalid indicates invalid
	GTSInvalid GroundTrackStatus = 0
	// GTSValid indicates valid
	GTSValid GroundTrackStatus = 1
)

// ToString returns a basic, but readable, representation of the field
func (status GroundTrackStatus) ToString() string {

	switch status {
	case GTSInvalid:
		return "0 - invalid"
	case GTSValid:
		return "1 - valid"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadGroundTrackStatus reads the GroundTrackStatus from a 56 bits data field
func ReadGroundTrackStatus(data []byte) GroundTrackStatus {
	bits := (data[1] & 0x04) >> 2
	return GroundTrackStatus(bits)
}
