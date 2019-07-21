package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// Format28V2Status is a message at the format BDS 6,1
type Format28V2Status struct {
	Subtype                 fields.SubtypeV2
	EmergencyPriorityStatus fields.EmergencyPriorityStatus
	ModeACode               uint16
}

// GetName returns the name of the message
func (message *Format28V2Status) GetName() string {
	return bds61Name
}

// GetBDS returns the binary data format
func (message *Format28V2Status) GetBDS() string {
	return bds61Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format28V2Status) GetFormatTypeCode() byte {
	return 28
}

// GetSubtype returns the Subtype
func (message *Format28V2Status) GetSubtype() fields.Subtype {
	return message.Subtype
}

// GetEmergencyPriorityStatus returns the EmergencyPriorityStatus
func (message *Format28V2Status) GetEmergencyPriorityStatus() fields.EmergencyPriorityStatus {
	return message.EmergencyPriorityStatus
}

// ToString returns a basic, but readable, representation of the message
func (message *Format28V2Status) ToString() string {
	return fmt.Sprintf("Message:                   %v - %v (BDS: %v)\n"+
		"Subtype:                   %v\n"+
		"Emergency Priority Status: %v\n"+
		"Mode A Code:               %v\n",
		message.GetFormatTypeCode(),
		message.GetName(),
		message.GetBDS(),
		message.GetSubtype().ToString(),
		message.GetEmergencyPriorityStatus().ToString(),
		message.ModeACode)
}

// readFormat28V2Status reads a message at the format BDS 6,1
func readFormat28V2Status(data []byte) (*Format28V2Status, error) {

	byte1 := data[1] & 0x1F
	byte2 := data[2]

	return &Format28V2Status{
		Subtype:                 fields.ReadSubtypeV2(data),
		EmergencyPriorityStatus: fields.ReadEmergencyPriorityStatus(data),
		ModeACode:               bitutils.Pack2Bytes(byte1, byte2),
	}, nil
}
