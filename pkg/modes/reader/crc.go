package reader

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/messages"
)

// CheckCRC checks that the CRC of a message is valid and/or return the ICAO address / Interrogator Identifier of the
// message. As the message parity is a XOR of the CRC and the Address or Interrogator Identifier (except for DF17 and
// DF18), it is not possible to ensure that a message is correct without previously known valid Address or Interrogator
// Identifier. Only the messages DF18 and DF17 always give a valid Address.
//
// Params:
//   - message: The message to check
//   - data: The raw data of the message
//   - allowedAddresses: For the messages that have uncertainty when computing the Address, allows to reject the
//     messages having an unknown Address. Leave to nil to ignore.
//   - allowedInterrogatorIdentifiers: For the messages that have uncertainty when computing Interrogator Identifiers,
//     allows to reject the messages having an unknown Interrogator Identifier.
//
// Notes:
//   - the allowedAddresses is not used for messages DF17 and DF18 are always giving a valid address.
//
// Returns the ICAO Interrogator Identifiers for messages DF11 and Address for all others.
func CheckCRC(
	message messages.ModeSMessage,
	data []byte,
	allowedAddresses map[common.ICAOAddress]bool,
	allowedInterrogatorIdentifiers map[common.ICAOAddress]bool) (common.ICAOAddress, error) {

	switch message.GetDownLinkFormat() {
	case 11:
		return checkCRCDF11(data, allowedInterrogatorIdentifiers)
	case 17, 18:
		return checkCRCDF17And18(data)
	default:
		return checkCRCOther(data, allowedAddresses)
	}
}

func checkCRCDF11(
	data []byte,
	allowedInterrogatorIdentifiers map[common.ICAOAddress]bool) (common.ICAOAddress, error) {

	// For DF11, the ICAO code is not returned (as it is the base of the message). Instead, the interrogator id (II)
	// is used to XOR the parity. So, an interrogator can detect if a message is a reply to its interrogation.
	contentParity := computeParity(data[:4])
	messageParity := bitutils.Pack3Bytes(data[4], data[5], data[6])

	interrogatorIdentifier := common.ICAOAddress(contentParity ^ messageParity)

	// If the interrogator is not valid
	if len(allowedInterrogatorIdentifiers) > 0 {
		if _, ok := allowedInterrogatorIdentifiers[interrogatorIdentifier]; !ok {
			return 0, fmt.Errorf("the message parity resolves to an unknown Interrogator Identifier")
		}
	}

	return interrogatorIdentifier, nil
}

func checkCRCDF17And18(data []byte) (common.ICAOAddress, error) {

	// For DF17 and DF18 (extended squitter), the ICAO address is returned as the first 3 bytes of the payload.
	messageICAO := common.ICAOAddress(bitutils.Pack3Bytes(data[1], data[2], data[3]))

	// The parity is XORed against an Interrogator ID equals to 0
	contentParity := computeParity(data[:11])
	messageParity := bitutils.Pack3Bytes(data[11], data[12], data[13])

	if contentParity != messageParity {
		return 0, fmt.Errorf("the message does not have a valid CRC")
	}

	return messageICAO, nil
}

func checkCRCOther(
	data []byte,
	allowedAddresses map[common.ICAOAddress]bool) (common.ICAOAddress, error) {

	messageLength := len(data)

	// Compute parity on the whole message, except the 3 last bytes
	contentParity := computeParity(data[:messageLength-3])
	messageParity := bitutils.Pack3Bytes(data[messageLength-3], data[messageLength-2], data[messageLength-1])

	address := common.ICAOAddress(contentParity ^ messageParity)

	// If the address is not valid
	if len(allowedAddresses) > 0 {
		if _, ok := allowedAddresses[address]; !ok {
			return 0, fmt.Errorf("the message parity resolves to an unknown Address")
		}
	}

	return address, nil
}

//		crcPolynomial is the polynomial for the CRC redundancy check
//	 Note: we assume that the degree of the polynomial is divisible by 8 (holds for Mode S) and the msb is left out
//
// Values defined according to Annex 10 V4
var crcPolynomial = []uint8{0xFF, 0xF4, 0x09}

// computeParity computes the parity of a slice of byte as 3-byte array. We used the implementation from:
// http://www.eurocontrol.int/eec/gallery/content/public/document/eec/report/1994/022_CRC_calculations_for_Mode_S.pdf
//
// params:
//   - data: The data for which to compute parity
//
// Returns the CRC (3 bytes)
func computeParity(data []byte) uint32 {

	crcLength := len(crcPolynomial)

	// Initialize with the beginning of the message
	crc := make([]byte, crcLength)
	for i := 0; i < crcLength; i++ {
		crc[i] = data[i]
	}

	// For all bit
	for i := 0; i < len(data)*8; i++ {

		// Keep msb
		invert := (crc[0] & 0x80) != 0

		// Shift left
		crc[0] <<= 1
		for b := 1; b < crcLength; b++ {
			crc[b-1] |= (crc[b] >> 7) & 0x1
			crc[b] <<= 1
		}

		// Get next bit from message
		byteIdx := (crcLength*8 + i) / 8
		bitShift := uint(7 - (i % 8))
		if byteIdx < len(data) {
			crc[len(crc)-1] |= (data[byteIdx] >> bitShift) & 0x1
		}

		// xor
		if invert {
			for b := 0; b < crcLength; b++ {
				crc[b] ^= crcPolynomial[b]
			}
		}
	}

	return bitutils.Pack3Bytes(crc[0], crc[1], crc[2])
}
