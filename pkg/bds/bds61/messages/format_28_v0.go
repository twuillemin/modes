package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
)

// Format28V0 is a message at the format BDS 6,1
type Format28V0 struct {
	Subtype                 fields.SubtypeV0
	EmergencyPriorityStatus fields.EmergencyPriorityStatus
}

// GetName returns the name of the message
func (message *Format28V0) GetName() string {
	return bds61Name
}

// GetBDS returns the binary data format
func (message *Format28V0) GetBDS() string {
	return bds61Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format28V0) GetFormatTypeCode() byte {
	return 28
}

// GetSubtype returns the Subtype
func (message *Format28V0) GetSubtype() fields.Subtype {
	return message.Subtype
}

// GetEmergencyPriorityStatus returns the EmergencyPriorityStatus
func (message *Format28V0) GetEmergencyPriorityStatus() fields.EmergencyPriorityStatus {
	return message.EmergencyPriorityStatus
}

// ToString returns a basic, but readable, representation of the message
func (message *Format28V0) ToString() string {
	return fmt.Sprintf("Message:                   %v - %v (BDS: %v)\n"+
		"Subtype:                   %v\n"+
		"Emergency Priority Status: %v",
		message.GetFormatTypeCode(),
		message.GetName(),
		message.GetBDS(),
		message.GetSubtype().ToString(),
		message.GetEmergencyPriorityStatus().ToString())
}

// ReadFormat28V0 reads a message at the format BDS 6,1
func ReadFormat28V0(data []byte) (*Format28V0, error) {

	return &Format28V0{
		Subtype:                 fields.ReadSubtypeV0(data),
		EmergencyPriorityStatus: fields.ReadEmergencyPriorityStatus(data),
	}, nil
}
