package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadWindDirection(data []byte) (bool, float32) {
	status := (data[0] & 0x08) != 0

	byte1 := data[1] & 0x03
	byte2 := data[2] & 0xFE
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 1
	direction := float32(allBits) * 180 / 256

	return status, direction
}
