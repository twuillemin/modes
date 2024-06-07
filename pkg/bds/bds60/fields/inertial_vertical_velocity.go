package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadInertialVerticalVelocity(data []byte) (bool, int32) {
	status := (data[5] & 0x04) != 0

	byte1 := data[5] & 0x0F
	byte2 := data[6] & 0xFF
	allBits := bitutils.Pack2Bytes(byte1, byte2)
	rate := int32(allBits) * 8192 / 256

	if (data[5] & 0x02) != 0 {
		rate = rate - 16384
	}

	return status, rate
}
