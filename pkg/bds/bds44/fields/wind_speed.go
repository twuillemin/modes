package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadWindSpeed(data []byte) (bool, uint32) {
	status := (data[0] & 0x08) != 0

	byte1 := data[0] & 0x03
	byte2 := data[1] & 0xFC
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 2
	speed := uint32(allBits)

	return status, speed
}
