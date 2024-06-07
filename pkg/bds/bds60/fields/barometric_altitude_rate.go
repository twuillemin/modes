package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadBarometricAltitudeRate(data []byte) (bool, int32) {
	status := (data[4] & 0x20) != 0

	byte1 := data[4] & 0x0F
	byte2 := data[5] & 0xF8
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 3
	rate := int32(allBits) * 32

	if (data[4] & 0x10) != 0 {
		rate = rate - 16384
	}

	return status, rate
}
