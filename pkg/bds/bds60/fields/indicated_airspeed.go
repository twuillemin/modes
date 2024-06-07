package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadIndicatedAirSpeed(data []byte) (bool, uint32) {
	status := (data[1] & 0x08) != 0

	byte1 := data[1] & 0x07
	byte2 := data[2] & 0xFE
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 1
	speed := uint32(allBits)

	return status, speed
}
