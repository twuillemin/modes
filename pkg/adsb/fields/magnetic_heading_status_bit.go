package fields

import "fmt"

// MagneticHeadingStatus is the Magnetic Heading Status definition
//
// Specified in Doc 9871 / Table A-2-9
type MagneticHeadingStatus byte

const (
	// AVSBNotAvailable indicates magnetic heading not available
	AVSBNotAvailable MagneticHeadingStatus = 0
	// AVSBAvailable indicates magnetic heading available
	AVSBAvailable MagneticHeadingStatus = 1
)

// ToString returns a basic, but readable, representation of the field
func (bit MagneticHeadingStatus) ToString() string {

	switch bit {
	case AVSBNotAvailable:
		return "0 - magnetic heading not available"
	case AVSBAvailable:
		return "1 - magnetic heading available"
	default:
		return fmt.Sprintf("%v - Unknown code", bit)
	}
}

// ReadMagneticHeadingStatus reads the MagneticHeadingStatus from a 56 bits data field
func ReadMagneticHeadingStatus(data []byte) MagneticHeadingStatus {
	bits := (data[1] & 0x04) >> 2
	return MagneticHeadingStatus(bits)
}
