package fields

// Field is the basic interface that Mode S messages fields are expected to implement
type Field interface {
	// ToString returns a basic, but readable, representation of the field
	ToString() string
	// ToExtendedString returns a complete representation of the field
	ToExtendedString() string
}

// -----------------------------------------------------------------------------------------
//
//                                 Miscellaneous functions
//
// -----------------------------------------------------------------------------------------

// grayToBinary convert the given bits (b0 being the MSB and b7 the LSB) from gray code to "classical" binary
func grayToBinary(b0, b1, b2, b3, b4, b5, b6, b7 bool) uint8 {

	num := uint8(0)
	if b0 {
		num |= 0x80
	}
	if b1 {
		num |= 0x40
	}
	if b2 {
		num |= 0x20
	}
	if b3 {
		num |= 0x10
	}
	if b4 {
		num |= 0x08
	}
	if b5 {
		num |= 0x04
	}
	if b6 {
		num |= 0x02
	}
	if b7 {
		num |= 0x01
	}

	mask := num >> 1

	for mask != 0 {
		num = num ^ mask
		mask = num >> 1
	}

	return num
}
