package bitutils

// Pack2Bytes packs 2 bytes in a single uint16.
//
// Params:
//    - b1: the first bit
//    - b0: the second bit
//
// The two bytes are packed as b1 b1 ... b1 b1 b0 b0 ... b0 b0
//
// Return the packed byte
func Pack2Bytes(b1 byte, b0 byte) uint16 {
	return uint16(b1)<<8 + uint16(b0)
}

// Pack3Bytes packs 3 bytes in a single uint32.
//
// Params:
//    - b2: the first bit
//    - b1: the second bit
//    - b0: the third bit
//
// The two bytes are packed as b2 b2 ... b2 b2 b1 b1 ... b1 b1 b0 b0 ... b0 b0
//
// Return the packed byte
func Pack3Bytes(b2 byte, b1 byte, b0 byte) uint32 {
	return uint32(b2)<<16 + uint32(b1)<<8 + uint32(b0)
}

// Pack4Bytes packs 4 bytes in a single uint32.
//
// Params:
//    - b3: the first bit
//    - b2: the second bit
//    - b1: the third bit
//    - b0: the fourth bit
//
// The two bytes are packed as b3 ... b3 b2 ... b2 b1 ... b1 b0 ... b0
//
// Return the packed byte
func Pack4Bytes(b3 byte, b2 byte, b1 byte, b0 byte) uint32 {
	return uint32(b3)<<24 + uint32(b2)<<16 + uint32(b1)<<8 + uint32(b0)
}

// Pack8Bits packs 8 bytes in a single byte.
//
// Params:
//    - b7: the first bit (MSB)
//    - bn: ...
//    - b0: the last bit (LSB)
//
// The bits are packed as b7 b6 ... b1 b0
//
// Return the packed byte
func Pack8Bits(b7, b6, b5, b4, b3, b2, b1, b0 bool) byte {

	result := byte(0)
	if b7 {
		result |= 0x80
	}
	if b6 {
		result |= 0x40
	}
	if b5 {
		result |= 0x20
	}
	if b4 {
		result |= 0x10
	}
	if b3 {
		result |= 0x08
	}
	if b2 {
		result |= 0x04
	}
	if b1 {
		result |= 0x02
	}
	if b0 {
		result |= 0x01
	}
	return result
}

// Pack16Bits packs 16 bytes in a single uint16.
//
// Params:
//    - b15: the first bit (MSB)
//    - bn: ...
//    - b0: the last bit (LSB)
//
// The bits are packed as b15 b14 ... b1 b0
//
// Return the packed byte
func Pack16Bits(b15, b14, b13, b12, b11, b10, b9, b8, b7, b6, b5, b4, b3, b2, b1, b0 bool) uint16 {

	result := uint16(0)
	if b15 {
		result |= 0x8000
	}
	if b14 {
		result |= 0x4000
	}
	if b13 {
		result |= 0x2000
	}
	if b12 {
		result |= 0x1000
	}
	if b11 {
		result |= 0x0800
	}
	if b10 {
		result |= 0x0400
	}
	if b9 {
		result |= 0x0200
	}
	if b8 {
		result |= 0x0100
	}
	if b7 {
		result |= 0x0080
	}
	if b6 {
		result |= 0x0040
	}
	if b5 {
		result |= 0x0020
	}
	if b4 {
		result |= 0x0010
	}
	if b3 {
		result |= 0x0008
	}
	if b2 {
		result |= 0x0004
	}
	if b1 {
		result |= 0x0002
	}
	if b0 {
		result |= 0x0001
	}
	return result
}
