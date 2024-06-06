package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadFMSSelectedAltitude(data []byte) (bool, uint32) {
	if (data[1] & 0x04) == 0 {
		return false, 0
	}

	byte1 := data[0] & 0x7F
	byte2 := data[1] & 0xF8
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 3
	return true, uint32(allBits) * 16
}
