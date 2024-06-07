package fields

import (
	"github.com/twuillemin/modes/pkg/bitutils"
)

func ReadMagneticHeading(data []byte) (bool, float32) {
	status := (data[0] & 0x80) != 0

	byte1 := data[0] & 0x3F
	byte2 := data[1] & 0xF0
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 4
	magneticHeading := float32(allBits) * 90.0 / 512.0

	if (data[0] & 0x40) != 0 {
		magneticHeading = magneticHeading - 180
	}

	return status, magneticHeading
}
