package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb/fields"
	"github.com/twuillemin/modes/pkg/bds/bds65/messages"
)

// Format31V1Airborne is a message at the format BDS 6,5 the ADSB V1 / Airborne
type Format31V1Airborne struct {
	AirborneCapabilityClass fields.CapabilityClassAirborne
	OperationalMode         fields.OperationalMode
	VersionNumber           fields.VersionNumber
}

// GetName returns the name of the message
func (message *Format31V1Airborne) GetName() string {
	return messages.bds65Name
}

// GetBDS returns the binary data format
func (message *Format31V1Airborne) GetBDS() string {
	return messages.bds65Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format31V1Airborne) GetFormatTypeCode() byte {
	return 31
}

// GetOperationalStatusSubTypeCode returns the code of the Operational Status Sub Type
func (message *Format31V1Airborne) GetOperationalStatusSubTypeCode() byte {
	return 0
}

// ToString returns a basic, but readable, representation of the message
func (message Format31V1Airborne) ToString() string {
	return fmt.Sprintf("Message:          %v (%v)\n"+
		"SubType:          0 - Airborne\n"+
		"VersionNumber:    %v\n"+
		"AirborneCapabilityClass:\n%v\n"+
		"OperationalMode:\n%v",
		message.GetBDS(),
		message.GetName(),
		message.VersionNumber.ToString(),
		message.AirborneCapabilityClass.ToString(),
		message.OperationalMode.ToString())
}

// ReadFormat31V1Airborne reads a message at the format Format31V1Airborne
func ReadFormat31V1Airborne(data []byte) (*Format31V1Airborne, error) {

	return &Format31V1Airborne{
		AirborneCapabilityClass: fields.ReadCapabilityClassAirborne(data),
		OperationalMode:         fields.ReadOperationalMode(data),
		VersionNumber:           fields.ReadVersionNumber(data),
	}, nil
}
