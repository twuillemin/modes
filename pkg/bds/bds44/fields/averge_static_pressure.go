package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadAverageStaticPressure(data []byte) (bool, uint32) {
	status := (data[4] & 0x01) != 0

	byte1 := data[4] & 0x0F
	byte2 := data[5] & 0xFE
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 1
	pressure := uint32(allBits)

	return status, pressure
}
