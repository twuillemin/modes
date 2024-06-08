package fields

// ReadHeightDifference reads the Height difference from a 56 bits data field
func ReadHeightDifference(data []byte) (int16, NumericValueStatus) {

	negative := data[6]&0x80 != 0
	difference := int16(data[6] & 0x7F)

	if difference == 0 {
		return 0, NVSNoInformation
	} else if difference >= 127 {
		if negative {
			return -3150, NVSMaximum
		} else {
			return 3150, NVSMaximum
		}
	}

	difference = (difference - 1) * 25

	if negative {
		difference = -difference
	}

	return difference, NVSRegular
}
