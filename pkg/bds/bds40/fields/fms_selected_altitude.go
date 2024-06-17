package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadFMSSelectedAltitude(data []byte) (bool, uint32) {
	status := (data[1] & 0x04) != 0

	byte1 := data[0] & 0x7F
	byte2 := data[1] & 0xF8
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 3
	fmsAltitude := uint32(allBits) * 16

	return status, fmsAltitude
}
