package fields

// AircraftIdentification is the identification (the name) of an aircraft
//
// Specified in Annex 10, Volume IV, Table 3-8
type AircraftIdentification string

var identificationCharacterCoding = []byte{
	'#', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O',
	'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '#', '#', '#', '#', '#',
	' ', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '#', '#', '#', '#', '#', '#',
}

// ReadAircraftIdentification reads the aircraft identification from a 56 bits data field
func ReadAircraftIdentification(data []byte) AircraftIdentification {

	// Get the codes
	codes := make([]byte, 8)
	codes[0] = (data[1] & 0xFC) >> 2
	codes[1] = (data[1]&0x03)<<4 + (data[2]&0xF0)>>4
	codes[2] = (data[2]&0x0F)<<2 + (data[3]&0xC0)>>6
	codes[3] = data[3] & 0x3F
	codes[4] = (data[4] & 0xFC) >> 2
	codes[5] = (data[4]&0x03)<<4 + (data[5]&0xF0)>>4
	codes[6] = (data[5]&0x0F)<<2 + (data[6]&0xC0)>>6
	codes[7] = data[6] & 0x3F

	// Convert the codes to actual char
	chars := make([]byte, 8)
	for i, code := range codes {
		chars[i] = identificationCharacterCoding[code]
	}

	// Return the value
	return AircraftIdentification(chars)
}
