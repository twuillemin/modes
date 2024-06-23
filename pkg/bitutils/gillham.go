package bitutils

import "errors"

// GillhamToAltitude convert an altitude given in Gillham bits to an altitude in feet
func GillhamToAltitude(d1, d2, d4, a1, a2, a4, b1, b2, b4, c1, c2, c4 bool) (int32, error) {

	fiveHundredBits := uint16(0)
	if d1 {
		fiveHundredBits |= 0x0100
	}
	if d2 {
		fiveHundredBits |= 0x0080
	}
	if d4 {
		fiveHundredBits |= 0x0040
	}
	if a1 {
		fiveHundredBits |= 0x0020
	}
	if a2 {
		fiveHundredBits |= 0x0010
	}
	if a4 {
		fiveHundredBits |= 0x0008
	}
	if b1 {
		fiveHundredBits |= 0x0004
	}
	if b2 {
		fiveHundredBits |= 0x0002
	}
	if b4 {
		fiveHundredBits |= 0x0001
	}

	oneHundredBits := uint16(0)
	if c1 {
		oneHundredBits |= 0x0004
	}
	if c2 {
		oneHundredBits |= 0x0002
	}
	if c4 {
		oneHundredBits |= 0x0001
	}

	oneHundred := int32(grayToBinary(oneHundredBits))
	fiveHundred := int32(grayToBinary(fiveHundredBits))

	// Check for invalid codes.
	if oneHundred == 5 || oneHundred == 6 || oneHundred == 0 {
		return 0, errors.New("the bits C1 to to C4 are incorrect")
	}

	// Remove 7s from OneHundreds.
	if oneHundred == 7 {
		oneHundred = 5
	}

	// Correct order of OneHundreds.
	if fiveHundred%2 != 0 {
		oneHundred = 6 - oneHundred
	}

	// Convert to feet and apply altitude datum offset.
	return (int32(fiveHundred)*500 + int32(oneHundred)*100) - 1300, nil
}

func grayToBinary(num uint16) uint16 {
	temp := uint16(0)

	temp = num ^ (num >> 8)
	temp ^= temp >> 4
	temp ^= temp >> 2
	temp ^= temp >> 1

	return temp
}
