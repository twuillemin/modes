package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb/fields"
)

// MessageBDS08 is a message at the format BDS0,8
type MessageBDS08 struct {
	FormatTypeCode         byte
	AircraftCategory       fields.AircraftCategory
	AircraftIdentification fields.AircraftIdentification
}

// GetName returns the name of the message
func (message *MessageBDS08) GetName() string {
	return "Extended squitter aircraft identification and category"
}

// GetBDS returns the binary data format
func (message *MessageBDS08) GetBDS() string {
	return "BDS 0,8"
}

// GetFormatTypeCode returns the Format Type Code
func (message *MessageBDS08) GetFormatTypeCode() byte {
	return message.FormatTypeCode
}

// ToString returns a basic, but readable, representation of the field
func (message *MessageBDS08) ToString() string {
	return fmt.Sprintf("Message: %v (%v)\n"+
		"FormatTypeCode: %v\n"+
		"AircraftCategory: %v (%v)\n"+
		"AircraftIdentification: %v",
		message.GetBDS(),
		message.GetName(),
		message.FormatTypeCode,
		message.AircraftCategory.ToString(),
		message.AircraftCategory.GetCategorySetName(),
		message.AircraftIdentification)
}

// ReadBDS08 reads a message at the format BDS 0,8
func ReadBDS08(data []byte) (*MessageBDS08, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3

	var category fields.AircraftCategory

	switch formatTypeCode {
	case 1:
		category = fields.ReadAircraftCategorySetD(data)
	case 2:
		category = fields.ReadAircraftCategorySetC(data)
	case 3:
		category = fields.ReadAircraftCategorySetB(data)
	case 4:
		category = fields.ReadAircraftCategorySetA(data)
	}

	return &MessageBDS08{
		FormatTypeCode:         formatTypeCode,
		AircraftCategory:       category,
		AircraftIdentification: fields.ReadAircraftIdentification(data),
	}, nil
}
