package fields

import (
	"github.com/twuillemin/modes/pkg/bitutils"
)

// ReadVerticalRate reads the Vertical Rate from a 56 bits data field
func ReadVerticalRate(data []byte) (int16, NumericValueStatus) {

	negative := data[4]&0x08 != 0
	byte1 := data[4] & 0x07
	byte2 := data[5] & 0xFC
	rate := int16(bitutils.Pack2Bytes(byte1, byte2) >> 2)

	if rate == 0 {
		return 0, NVSNoInformation
	} else if rate >= 511 {
		if negative {
			return -32640, NVSMaximum
		} else {
			return 32640, NVSMaximum
		}
	}

	rate = (rate - 1) * 64

	if negative {
		rate = -rate
	}

	return rate, NVSRegular
}
