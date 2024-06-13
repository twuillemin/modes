package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadStaticAirTemperatureV0(data []byte) (bool, float32) {
	status := (data[1] & 0x01) != 0

	byte1 := data[2] & 0x7F
	byte2 := data[3] & 0xC0
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 6
	temperature := float32(allBits) * 0.25

	if (data[2] & 0x80) != 0 {
		temperature = temperature - 128
	}

	return status, temperature
}

func ReadStaticAirTemperatureV1(data []byte) (bool, float32) {
	status := (data[1] & 0x08) != 0

	byte1 := data[1] & 0x03
	byte2 := data[2] & 0xFF
	allBits := bitutils.Pack2Bytes(byte1, byte2)
	temperature := float32(allBits) * 0.125

	if (data[1] & 0x04) != 0 {
		temperature = temperature - 128
	}

	return status, temperature
}
