package messages

import (
	"github.com/twuillemin/modes/pkg/adsb/fields"
)

// Format01 is a message at the format BDS 0,8
type Format01 struct {
	AircraftCategory       fields.AircraftCategorySetD
	AircraftIdentification fields.AircraftIdentification
}

// GetName returns the name of the message
func (message *Format01) GetName() string {
	return bds08Name
}

// GetBDS returns the binary data format
func (message *Format01) GetBDS() string {
	return bds08Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format01) GetFormatTypeCode() byte {
	return 1
}

// ToString returns a basic, but readable, representation of the message
func (message *Format01) ToString() string {
	return bds08ToString(message)
}

// GetAircraftCategory returns the category of the aircraft
func (message *Format01) GetAircraftCategory() fields.AircraftCategory {
	return message.AircraftCategory
}

// GetAircraftIdentification returns the identification of the aircraft
func (message *Format01) GetAircraftIdentification() fields.AircraftIdentification {
	return message.AircraftIdentification
}

// ReadFormat01 reads a message at the format BDS 0,8
func ReadFormat01(data []byte) (*Format01, error) {

	return &Format01{
		AircraftCategory:       fields.ReadAircraftCategorySetD(data),
		AircraftIdentification: fields.ReadAircraftIdentification(data),
	}, nil
}
