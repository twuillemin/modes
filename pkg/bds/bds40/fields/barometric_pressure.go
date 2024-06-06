package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadBarometricPressure(data []byte) (bool, float32) {
	if (data[0] & 0x80) == 0 {
		return false, 0
	}

	byte1 := data[3] & 0x1F
	byte2 := data[4] & 0xFE
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 1
	return true, (float32(allBits) * 0.1) + 800
}
