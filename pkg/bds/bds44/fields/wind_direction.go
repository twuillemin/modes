package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadWindDirection(data []byte) (bool, float32) {
	status := (data[2] & 0x02) != 0

	byte1 := data[1] & 0x01
	byte2 := data[2] & 0xFE
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 1
	direction := float32(allBits) * 180 / 128

	return status, direction
}
