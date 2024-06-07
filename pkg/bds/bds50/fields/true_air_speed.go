package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadTrueAirSpeed(data []byte) (bool, float32) {
	status := (data[5] & 0x04) != 0

	byte1 := data[5] & 0x03
	byte2 := data[6] & 0xFF
	allBits := bitutils.Pack2Bytes(byte1, byte2)
	speed := float32(allBits) * 2

	return status, speed
}
