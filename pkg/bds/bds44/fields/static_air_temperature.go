package fields

import "github.com/twuillemin/modes/pkg/bitutils"

// ReadStaticAirTemperatureV0 as given by Table A-2-68
func ReadStaticAirTemperatureV0(data []byte) float32 {

	byte1 := data[3] & 0xFF
	byte2 := data[4] & 0xC0
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 6
	temperature := float32(allBits) * 0.25

	if (data[3] & 0x01) != 0 {
		temperature = temperature - 128
	}

	return temperature
}

// ReadStaticAirTemperatureV1AndV2 as given by Table E-2-68
func ReadStaticAirTemperatureV1AndV2(data []byte) (bool, float32) {
	status := (data[2] & 0x01) != 0

	byte1 := data[3] & 0x7F
	byte2 := data[4] & 0xE0
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 5
	temperature := float32(allBits) * 0.125

	if (data[3] & 0x80) != 0 {
		temperature = temperature - 128
	}

	return status, temperature
}
