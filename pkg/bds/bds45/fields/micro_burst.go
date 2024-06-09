package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadMicroBurstV0(data []byte) (bool, HazardLevel) {
	status := (data[0] & 0x02) != 0

	byte1 := data[0] & 0x01
	byte2 := data[1] & 0x80
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 7
	level := HazardLevel(allBits)

	return status, level
}

func ReadMicroBurstV1(data []byte) (bool, HazardLevel) {
	status := (data[0] & 0x10) != 0
	level := HazardLevel((data[0] & 0x0C) >> 2)
	return status, level
}
