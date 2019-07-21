package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
)

// Format28V1Status is a message at the format BDS 6,1
type Format28V1Status struct {
	Subtype                 fields.SubtypeV1
	EmergencyPriorityStatus fields.EmergencyPriorityStatus
}

// GetName returns the name of the message
func (message *Format28V1Status) GetName() string {
	return bds61Name
}

// GetBDS returns the binary data format
func (message *Format28V1Status) GetBDS() string {
	return bds61Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format28V1Status) GetFormatTypeCode() byte {
	return 28
}

// GetSubtype returns the Subtype
func (message *Format28V1Status) GetSubtype() fields.Subtype {
	return message.Subtype
}

// GetEmergencyPriorityStatus returns the EmergencyPriorityStatus
func (message *Format28V1Status) GetEmergencyPriorityStatus() fields.EmergencyPriorityStatus {
	return message.EmergencyPriorityStatus
}

// ToString returns a basic, but readable, representation of the message
func (message *Format28V1Status) ToString() string {
	return fmt.Sprintf("Message:                   %v - %v (BDS: %v)\n"+
		"Subtype:                   %v\n"+
		"Emergency Priority Status: %v",
		message.GetFormatTypeCode(),
		message.GetName(),
		message.GetBDS(),
		message.GetSubtype().ToString(),
		message.GetEmergencyPriorityStatus().ToString())
}

// readFormat28V1Status reads a message at the format BDS 6,1
func readFormat28V1Status(data []byte) (*Format28V1Status, error) {

	return &Format28V1Status{
		Subtype:                 fields.ReadSubtypeV1(data),
		EmergencyPriorityStatus: fields.ReadEmergencyPriorityStatus(data),
	}, nil
}
