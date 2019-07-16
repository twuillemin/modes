package messages

import (
	"github.com/twuillemin/modes/pkg/adsb/fields"
)

// Format04 is a message at the format BDS 0,8
type Format04 struct {
	AircraftCategory       fields.AircraftCategorySetA
	AircraftIdentification fields.AircraftIdentification
}

// GetName returns the name of the message
func (message *Format04) GetName() string {
	return bds08Name
}

// GetBDS returns the binary data format
func (message *Format04) GetBDS() string {
	return bds08Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format04) GetFormatTypeCode() byte {
	return 4
}

// ToString returns a basic, but readable, representation of the field
func (message *Format04) ToString() string {
	return bds08ToString(message)
}

// GetAircraftCategory returns the category of the aircraft
func (message *Format04) GetAircraftCategory() fields.AircraftCategory {
	return message.AircraftCategory
}

// GetAircraftIdentification returns the identification of the aircraft
func (message *Format04) GetAircraftIdentification() fields.AircraftIdentification {
	return message.AircraftIdentification
}

// ReadFormat04 reads a message at the format BDS 0,8
func ReadFormat04(data []byte) (*Format04, error) {

	return &Format04{
		AircraftCategory:       fields.ReadAircraftCategorySetA(data),
		AircraftIdentification: fields.ReadAircraftIdentification(data),
	}, nil
}
