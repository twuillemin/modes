package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
)

// Format28NoInformation is a message at the format BDS 6,1
type Format28NoInformation struct {
	Subtype fields.SubtypeV0
}

// GetName returns the name of the message
func (message *Format28NoInformation) GetName() string {
	return bds61Name
}

// GetBDS returns the binary data format
func (message *Format28NoInformation) GetBDS() string {
	return bds61Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format28NoInformation) GetFormatTypeCode() byte {
	return 28
}

// GetSubtype returns the Subtype
func (message *Format28NoInformation) GetSubtype() fields.Subtype {
	return message.Subtype
}

// ToString returns a basic, but readable, representation of the message
func (message *Format28NoInformation) ToString() string {
	return fmt.Sprintf("Message:                   %v - %v (BDS: %v)\n"+
		"Subtype:                   %v - No Information",
		message.GetFormatTypeCode(),
		message.GetName(),
		message.GetBDS(),
		message.GetSubtype().ToString())
}

// readFormat28NoInformation reads a message at the format BDS 6,1 / subtype 0
func readFormat28NoInformation(data []byte) (*Format28NoInformation, error) {

	return &Format28NoInformation{
		Subtype: fields.ReadSubtypeV0(data),
	}, nil
}
