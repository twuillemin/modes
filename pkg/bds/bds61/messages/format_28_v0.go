package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
)

// Format28 is a message at the format BDS 6,1
type Format28 struct {
	Subtype                 fields.SubtypeV0
	EmergencyPriorityStatus fields.EmergencyPriorityStatus
}

// GetName returns the name of the message
func (message *Format28) GetName() string {
	return bds61Name
}

// GetBDS returns the binary data format
func (message *Format28) GetBDS() string {
	return bds61Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format28) GetFormatTypeCode() byte {
	return 28
}

// GetSubtype returns the Subtype
func (message *Format28) GetSubtype() fields.Subtype {
	return message.Subtype
}

// GetEmergencyPriorityStatus returns the EmergencyPriorityStatus
func (message *Format28) GetEmergencyPriorityStatus() fields.EmergencyPriorityStatus {
	return message.EmergencyPriorityStatus
}

// ToString returns a basic, but readable, representation of the message
func (message *Format28) ToString() string {
	return fmt.Sprintf("Message:                   %v (%v)\n"+
		"Format Type Code:          %v\n"+
		"Subtype:                   %v\n"+
		"Emergency Priority Status: %v",
		message.GetBDS(),
		message.GetName(),
		message.GetFormatTypeCode(),
		message.GetSubtype(),
		message.GetEmergencyPriorityStatus().ToString())
}

// ReadFormat28V0 reads a message at the format BDS 6,1
func ReadFormat28V0(data []byte) (*Format28, error) {

	return &Format28{
		Subtype:                 fields.ReadSubtypeV0(data),
		EmergencyPriorityStatus: fields.ReadEmergencyPriorityStatus(data),
	}, nil
}
