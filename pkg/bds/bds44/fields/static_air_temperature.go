package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadStaticAirTemperature(data []byte) (bool, float32) {
	status := (data[2] & 0x01) != 0

	negative := (data[3] & 0x80) != 0

	byte1 := data[3] & 0x7F
	byte2 := data[4] & 0xE0
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 5
	temperature := float32(allBits) * 0.125

	if negative {
		temperature -= temperature
	}

	return status, temperature
}
