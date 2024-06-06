package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// MagneticHeading is the magnetic heading
//
// Specified in Doc 9871 / Table A-2-96
type MagneticHeading byte

const (
	// MHEast indicates east.
	MHEast MagneticHeading = 0
	// MHWest indicates west.
	MHWest MagneticHeading = 1
)

// ToString returns a basic, but readable, representation of the field
func (mh MagneticHeading) ToString() string {

	switch mh {
	case MHEast:
		return "0 - East"
	case MHWest:
		return "1 - West"
	default:
		return fmt.Sprintf("%v - Unknown code", mh)
	}
}

func ReadMagneticHeading(data []byte) (bool, MagneticHeading, float32) {
	status := (data[0] & 0x80) == 0

	magneticHeading := MagneticHeading((data[0] & 0x40) >> 6)

	byte1 := data[0] & 0x3F
	byte2 := data[1] & 0xF0
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 4
	rollAngle := float32(allBits) * 90.0 / 512.0

	return status, magneticHeading, rollAngle
}
