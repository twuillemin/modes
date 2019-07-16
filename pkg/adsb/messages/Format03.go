package messages

import (
	"github.com/twuillemin/modes/pkg/adsb/fields"
)

// Format03 is a message at the format BDS 0,8
type Format03 struct {
	AircraftCategory       fields.AircraftCategorySetB
	AircraftIdentification fields.AircraftIdentification
}

// GetName returns the name of the message
func (message *Format03) GetName() string {
	return bds08Name
}

// GetBDS returns the binary data format
func (message *Format03) GetBDS() string {
	return bds08Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format03) GetFormatTypeCode() byte {
	return 3
}

// ToString returns a basic, but readable, representation of the field
func (message *Format03) ToString() string {
	return bds08ToString(message)
}

// GetAircraftCategory returns the category of the aircraft
func (message *Format03) GetAircraftCategory() fields.AircraftCategory {
	return message.AircraftCategory
}

// GetAircraftIdentification returns the identification of the aircraft
func (message *Format03) GetAircraftIdentification() fields.AircraftIdentification {
	return message.AircraftIdentification
}

// ReadFormat03 reads a message at the format BDS 0,8
func ReadFormat03(data []byte) (*Format03, error) {

	return &Format03{
		AircraftCategory:       fields.ReadAircraftCategorySetB(data),
		AircraftIdentification: fields.ReadAircraftIdentification(data),
	}, nil
}
