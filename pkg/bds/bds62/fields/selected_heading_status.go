package fields

import "fmt"

// SelectedHeadingStatus is the Selected Heading Status definition
//
// Specified in Doc 9871 / C.2.3.9.6
type SelectedHeadingStatus byte

const (
	// SHSInvalid indicates that the Selected Heading is invalid or not available
	SHSInvalid = 0
	// SHSValid indicates that the Selected Heading is invalid or not available
	SHSValid SelectedHeadingStatus = 1
)

// ToString returns a basic, but readable, representation of the field
func (status SelectedHeadingStatus) ToString() string {

	switch status {
	case SHSInvalid:
		return "0 - selected heading data is either not available or is invalid"
	case SHSValid:
		return "1 - selected heading data is valid"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadSelectedHeadingStatus reads the SelectedHeadingStatus from a 56 bits data field
func ReadSelectedHeadingStatus(data []byte) SelectedHeadingStatus {
	bits := (data[3] & 0x04) >> 2
	return SelectedHeadingStatus(bits)
}
