package messages

import (
	"github.com/twuillemin/modes/pkg/adsb/fields"
)

// Format02 is a message at the format BDS 0,8
type Format02 struct {
	AircraftCategory       fields.AircraftCategorySetC
	AircraftIdentification fields.AircraftIdentification
}

// GetName returns the name of the message
func (message *Format02) GetName() string {
	return bds08Name
}

// GetBDS returns the binary data format
func (message *Format02) GetBDS() string {
	return bds08Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format02) GetFormatTypeCode() byte {
	return 2
}

// ToString returns a basic, but readable, representation of the field
func (message *Format02) ToString() string {
	return bds08ToString(message)
}

// GetAircraftCategory returns the category of the aircraft
func (message *Format02) GetAircraftCategory() fields.AircraftCategory {
	return message.AircraftCategory
}

// GetAircraftIdentification returns the identification of the aircraft
func (message *Format02) GetAircraftIdentification() fields.AircraftIdentification {
	return message.AircraftIdentification
}

// ReadFormat02 reads a message at the format BDS 0,8
func ReadFormat02(data []byte) (*Format02, error) {

	return &Format02{
		AircraftCategory:       fields.ReadAircraftCategorySetC(data),
		AircraftIdentification: fields.ReadAircraftIdentification(data),
	}, nil
}
