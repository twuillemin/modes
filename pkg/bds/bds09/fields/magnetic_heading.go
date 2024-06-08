package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// MagneticHeadingStatus is the Magnetic Heading Status definition
//
// Specified in Doc 9871 / Table A-2-9
type MagneticHeadingStatus byte

const (
	// MHSNotAvailable indicates magnetic heading not available
	MHSNotAvailable MagneticHeadingStatus = 0
	// MHSAvailable indicates magnetic heading available
	MHSAvailable MagneticHeadingStatus = 1
)

// ToString returns a basic, but readable, representation of the field
func (bit MagneticHeadingStatus) ToString() string {

	switch bit {
	case MHSNotAvailable:
		return "0 - Magnetic heading not available"
	case MHSAvailable:
		return "1 - Magnetic heading available"
	default:
		return fmt.Sprintf("%v - Unknown code", bit)
	}
}

// ReadMagneticHeading reads the MagneticHeading from a 56 bits data field
func ReadMagneticHeading(data []byte) (float32, MagneticHeadingStatus) {

	status := MagneticHeadingStatus((data[1] & 0x04) >> 2)

	byte1 := data[1] & 0x03
	byte2 := data[2]

	value := float64(bitutils.Pack2Bytes(byte1, byte2)) * 360 / 1024.0

	return float32(value), status
}
