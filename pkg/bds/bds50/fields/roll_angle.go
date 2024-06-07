package fields

import (
	"github.com/twuillemin/modes/pkg/bitutils"
)

func ReadRollAngle(data []byte) (bool, float32) {
	status := (data[0] & 0x80) != 0

	byte1 := data[0] & 0x3F
	byte2 := data[1] & 0xE0
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 5
	rollAngle := float32(allBits) * 45.0 / 256.0

	if (data[0]&0x40)>>6 != 0 {
		rollAngle = rollAngle - 90
	}

	return status, rollAngle
}
