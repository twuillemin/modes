package messages

import (
	fields2 "github.com/twuillemin/modes/pkg/bds/bds08/fields"
)

// Format01 is a message at the format BDS 0,8
type Format01 struct {
	AircraftCategory       fields2.AircraftCategorySetD
	AircraftIdentification fields2.AircraftIdentification
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
func (message *Format01) GetAircraftCategory() fields2.AircraftCategory {
	return message.AircraftCategory
}

// GetAircraftIdentification returns the identification of the aircraft
func (message *Format01) GetAircraftIdentification() fields2.AircraftIdentification {
	return message.AircraftIdentification
}

// readFormat01 reads a message at the format BDS 0,8
func readFormat01(data []byte) (*Format01, error) {

	return &Format01{
		AircraftCategory:       fields2.ReadAircraftCategorySetD(data),
		AircraftIdentification: fields2.ReadAircraftIdentification(data),
	}, nil
}
