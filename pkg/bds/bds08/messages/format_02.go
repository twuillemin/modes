package messages

import (
	fields2 "github.com/twuillemin/modes/pkg/bds/bds08/fields"
)

// Format02 is a message at the format BDS 0,8
type Format02 struct {
	AircraftCategory       fields2.AircraftCategorySetC
	AircraftIdentification fields2.AircraftIdentification
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

// ToString returns a basic, but readable, representation of the message
func (message *Format02) ToString() string {
	return bds08ToString(message)
}

// GetAircraftCategory returns the category of the aircraft
func (message *Format02) GetAircraftCategory() fields2.AircraftCategory {
	return message.AircraftCategory
}

// GetAircraftIdentification returns the identification of the aircraft
func (message *Format02) GetAircraftIdentification() fields2.AircraftIdentification {
	return message.AircraftIdentification
}

// ReadFormat02 reads a message at the format BDS 0,8
func ReadFormat02(data []byte) (*Format02, error) {

	return &Format02{
		AircraftCategory:       fields2.ReadAircraftCategorySetC(data),
		AircraftIdentification: fields2.ReadAircraftIdentification(data),
	}, nil
}
