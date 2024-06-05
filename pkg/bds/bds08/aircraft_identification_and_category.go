package bds08

import (
	"errors"
	"fmt"
	"strings"

	"github.com/twuillemin/modes/pkg/bds/bds08/fields"
	"github.com/twuillemin/modes/pkg/bds/encoding"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// AircraftIdentificationAndCategory is a message at the format BDS 0,6
type AircraftIdentificationAndCategory struct {
	FormatTypeCode byte
	Category       fields.AircraftCategory
	Identification string
}

// GetRegister returns the Register the message
func (message AircraftIdentificationAndCategory) GetRegister() register.Register {
	return register.BDS08
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message AircraftIdentificationAndCategory) CheckCoherency() error {

	if strings.ContainsAny(message.Identification, "#") {
		return errors.New("field Identification contains a non valid character")
	}

	err := message.Category.CheckCoherency()
	if err != nil {
		return err
	}

	return nil
}

func (message AircraftIdentificationAndCategory) ToString() string {
	return fmt.Sprintf(""+
		"Message:                           %v\n"+
		"Format Type Code:                  %v\n"+
		"Category:                          %v\n"+
		"Identification:                    %v",
		message.GetRegister().ToString(),
		message.FormatTypeCode,
		message.Category.ToString(),
		message.Identification)
}

// ReadAircraftIdentificationAndCategory reads a message at the format Format09V1
func ReadAircraftIdentificationAndCategory(data []byte) (*AircraftIdentificationAndCategory, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode < 1 || formatTypeCode > 4 {
		return nil, fmt.Errorf("the field FormatTypeCode must be comprised between 1 and 4 included, got %v", formatTypeCode)
	}

	var category fields.AircraftCategory
	switch formatTypeCode {
	case 1:
		category = fields.AircraftCategorySetD(data[0] & 0x07)
		break
	case 2:
		category = fields.AircraftCategorySetC(data[0] & 0x07)
		break
	case 3:
		category = fields.AircraftCategorySetB(data[0] & 0x07)
		break
	default:
		category = fields.AircraftCategorySetA(data[0] & 0x07)
		break
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
		chars[i] = encoding.IdentificationCharacterCoding[code]
	}

	return &AircraftIdentificationAndCategory{
		FormatTypeCode: formatTypeCode,
		Category:       category,
		Identification: string(chars),
	}, nil
}
