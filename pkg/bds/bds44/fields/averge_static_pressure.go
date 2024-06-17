package fields

import "github.com/twuillemin/modes/pkg/bitutils"

// ReadAverageStaticPressureV0 as given by Table A-2-68
func ReadAverageStaticPressureV0(data []byte) (bool, uint32) {
	status := (data[4] & 0x20) != 0

	byte1 := data[4] & 0x1F
	byte2 := data[5] & 0xFC
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 2
	pressure := uint32(allBits)

	return status, pressure
}

// ReadAverageStaticPressureV1AndV2 as given by Table E-2-68
func ReadAverageStaticPressureV1AndV2(data []byte) (bool, uint32) {
	status := (data[4] & 0x10) != 0

	byte1 := data[4] & 0x0F
	byte2 := data[5] & 0xFE
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 1
	pressure := uint32(allBits)

	return status, pressure
}
