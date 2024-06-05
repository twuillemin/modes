package bds20

import (
	"errors"
	"fmt"
	"strings"

	"github.com/twuillemin/modes/pkg/bds/register"
)

// AircraftIdentification is a message at the format BDS 2,0
//
// Specified in Doc 9871 / D.2.4.3
type AircraftIdentification struct {
	Identification string
}

// GetRegister returns the Register the message
func (message AircraftIdentification) GetRegister() register.Register {
	return register.BDS20
}

// ToString returns a basic, but readable, representation of the message
func (message AircraftIdentification) ToString() string {
	return fmt.Sprintf(""+
		"Message:                                 %v\n"+
		"Identification:                          %v\n",
		message.GetRegister().ToString(),
		message.Identification)
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message AircraftIdentification) CheckCoherency() error {
	if strings.ContainsAny(message.Identification, "#") {
		return errors.New("field Identification contains a non valid character")
	}

	return nil
}

var identificationCharacterCoding = []byte{
	'#', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O',
	'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '#', '#', '#', '#', '#',
	' ', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '#', '#', '#', '#', '#', '#',
}

// ReadAircraftIdentification reads a message as a AircraftIdentification
func ReadAircraftIdentification(data []byte) (*AircraftIdentification, error) {
	err := CheckIfDataReadable(data)
	if err != nil {
		return nil, err
	}

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

	return &AircraftIdentification{
		Identification: string(chars),
	}, nil
}

// CheckIfDataReadable checks if the data can be read as a DataLinkCapabilityReport
func CheckIfDataReadable(data []byte) error {
	if len(data) != 7 {
		return errors.New("the data for Comm-B AircraftIdentification message must be 7 bytes long")
	}

	// First byte is simply the BDS format 0010 0000
	if data[0] != 0x20 {
		return errors.New("the first byte of data is not 0x20")
	}

	return nil
}
