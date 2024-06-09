package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadIcingV0(data []byte) (bool, HazardLevel) {
	status := (data[1] & 0x40) != 0
	level := HazardLevel((data[1] & 0x30) >> 4)
	return status, level
}

func ReadIcingV1(data []byte) (bool, HazardLevel) {
	status := (data[0] & 0x02) != 0

	byte1 := data[0] & 0x01
	byte2 := data[1] & 0x80
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 7
	level := HazardLevel(allBits)

	return status, level
}
