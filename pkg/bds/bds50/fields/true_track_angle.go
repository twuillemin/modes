package fields

import (
	"github.com/twuillemin/modes/pkg/bitutils"
)

func ReadTrueTrackAngle(data []byte) (bool, float32) {
	status := (data[1] & 0x10) != 0

	byte1 := data[1] & 0x07
	byte2 := data[2] & 0xFE
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 1
	angle := float32(allBits) * 90.0 / 512.0

	if (data[1] & 0x08) != 0 {
		angle = angle - 180
	}

	return status, angle
}
