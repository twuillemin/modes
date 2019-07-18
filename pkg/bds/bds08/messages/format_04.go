package messages

import (
	fields2 "github.com/twuillemin/modes/pkg/bds/bds08/fields"
)

// Format04 is a message at the format BDS 0,8
type Format04 struct {
	AircraftCategory       fields2.AircraftCategorySetA
	AircraftIdentification fields2.AircraftIdentification
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

// ToString returns a basic, but readable, representation of the message
func (message *Format04) ToString() string {
	return bds08ToString(message)
}

// GetAircraftCategory returns the category of the aircraft
func (message *Format04) GetAircraftCategory() fields2.AircraftCategory {
	return message.AircraftCategory
}

// GetAircraftIdentification returns the identification of the aircraft
func (message *Format04) GetAircraftIdentification() fields2.AircraftIdentification {
	return message.AircraftIdentification
}

// ReadFormat04 reads a message at the format BDS 0,8
func ReadFormat04(data []byte) (*Format04, error) {

	return &Format04{
		AircraftCategory:       fields2.ReadAircraftCategorySetA(data),
		AircraftIdentification: fields2.ReadAircraftIdentification(data),
	}, nil
}
