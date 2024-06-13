package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadRadioHeightV0(data []byte) (bool, uint32) {
	status := (data[4] & 0x02) != 0

	byte1 := data[4] & 0x01
	byte2 := data[5] & 0xFF
	byte3 := data[6] & 0xE0
	allBits := bitutils.Pack3Bytes(byte1, byte2, byte3) >> 5
	height := uint32(allBits) * 16

	return status, height
}

func ReadRadioHeightV1(data []byte) (bool, uint32) {
	status := (data[3] & 0x80) != 0

	byte1 := data[3] & 0x7F
	byte2 := data[4] & 0xF8
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 3
	height := uint32(allBits) * 2

	return status, height
}
