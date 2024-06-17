package fields

import "github.com/twuillemin/modes/pkg/bitutils"

// ReadWindDirectionV0 as given by Table A-2-68
func ReadWindDirectionV0(data []byte) (bool, float32) {
	status := (data[0] & 0x08) != 0

	byte1 := data[1] & 0x03
	byte2 := data[2] & 0xFE
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 1
	direction := float32(allBits) * 180 / 256

	return status, direction
}

// ReadWindDirectionV1AndV2 as given by Table E-2-68
func ReadWindDirectionV1AndV2(data []byte) (bool, float32) {
	status := (data[1] & 0x02) != 0

	byte1 := data[1] & 0x01
	byte2 := data[2] & 0xFE
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 1
	direction := float32(allBits) * 180 / 128

	return status, direction
}
