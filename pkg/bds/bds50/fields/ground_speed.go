package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadGroundSpeed(data []byte) (bool, float32) {
	if (data[2] & 0x01) == 0 {
		return false, 0
	}

	byte1 := data[3] & 0xFF
	byte2 := data[4] & 0xC0
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 6
	speed := float32(allBits) * 1024.0 / 512.0

	return true, speed
}
