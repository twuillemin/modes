package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadMCPSelectedAltitude(data []byte) (bool, uint32) {
	status := (data[0] & 0x80) != 0

	byte1 := data[1] & 0x03
	byte2 := data[2]
	byte3 := data[3] & 0xC0
	allBits := bitutils.Pack3Bytes(byte1, byte2, byte3) >> 6
	altitude := allBits * 16

	return status, altitude
}
