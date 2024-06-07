package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadRadioHeight(data []byte) (bool, uint32) {
	status := (data[4] & 0x02) != 0

	byte1 := data[3] & 0x01
	byte2 := data[5] & 0xFF
	byte3 := data[6] & 0xE0
	allBits := bitutils.Pack3Bytes(byte1, byte2, byte3) >> 5
	height := uint32(allBits) * 16

	return status, height
}
