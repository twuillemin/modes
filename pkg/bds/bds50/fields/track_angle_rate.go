package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadTrackAngleRate(data []byte) (bool, float32) {
	status := (data[4] & 0x20) != 0

	byte1 := data[4] & 0x0F
	byte2 := data[5] & 0xF8
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 4
	angleRate := float32(allBits) * 8.0 / 256.0

	if (data[4] & 0x10) != 0 {
		angleRate = -angleRate
	}

	return status, angleRate
}
