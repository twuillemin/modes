package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadStaticAirTemperature(data []byte) (bool, float32) {
	status := (data[1] & 0x01) != 0

	negative := (data[2] & 0x80) != 0

	byte1 := data[2] & 0x7F
	byte2 := data[3] & 0xC0
	allBits := bitutils.Pack2Bytes(byte1, byte2) >> 6
	temperature := float32(allBits) * 0.25

	if negative {
		temperature -= temperature
	}

	return status, temperature
}
