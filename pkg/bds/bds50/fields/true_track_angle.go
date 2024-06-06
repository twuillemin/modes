package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// TrueTrackOrientation is the orientation (east/west) of the true track angle
//
// Specified in Doc 9871 / Table A-2-80
type TrueTrackOrientation byte

const (
	// TTEast indicates east.
	TTEast TrueTrackOrientation = 0
	// TTWest indicates west.
	TTWest TrueTrackOrientation = 1
)

// ToString returns a basic, but readable, representation of the field
func (tto TrueTrackOrientation) ToString() string {

	switch tto {
	case TTEast:
		return "0 - East"
	case TTWest:
		return "1 - West"
	default:
		return fmt.Sprintf("%v - Unknown code", tto)
	}
}

func ReadTrueTrackAngle(data []byte) (bool, TrueTrackOrientation, float32) {
	status := (data[1] & 0x10) == 0

	orientation := TrueTrackOrientation((data[1] & 0x80) >> 3)

	byte1 := data[0] & 0x03
	byte2 := data[1] & 0xFE
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 1
	angle := float32(allBits) * 90.0 / 512.0

	return status, orientation, angle
}
