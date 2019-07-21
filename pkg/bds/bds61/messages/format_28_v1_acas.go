package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
)

// Format28V1ACAS is a message at the format BDS 6,1
type Format28V1ACAS struct {
	Subtype  fields.SubtypeV1
	ACASData []byte
}

// GetName returns the name of the message
func (message *Format28V1ACAS) GetName() string {
	return bds61Name
}

// GetBDS returns the binary data format
func (message *Format28V1ACAS) GetBDS() string {
	return bds61Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format28V1ACAS) GetFormatTypeCode() byte {
	return 28
}

// GetSubtype returns the Subtype
func (message *Format28V1ACAS) GetSubtype() fields.Subtype {
	return message.Subtype
}

// GetACASData returns the ACAS Data
func (message *Format28V1ACAS) GetACASData() []byte {
	return message.ACASData
}

// ToString returns a basic, but readable, representation of the message
func (message *Format28V1ACAS) ToString() string {
	return fmt.Sprintf("Message:                   %v - %v (BDS: %v)\n"+
		"Subtype:                   %v\n"+
		"ACAS Data                  : %02X %02X %02X %02X %02X %02X",
		message.GetFormatTypeCode(), message.GetName(), message.GetBDS(),
		message.GetSubtype().ToString(),
		message.ACASData[0], message.ACASData[1], message.ACASData[2], message.ACASData[3], message.ACASData[4], message.ACASData[5])
}

// readFormat28V1ACAS reads a message at the format BDS 6,1
func readFormat28V1ACAS(data []byte) (*Format28V1ACAS, error) {

	// Copy data to not keep a reference to the given data
	acasData := make([]byte, 6)
	for i := 0; i < 6; i++ {
		acasData[i] = data[i+1]
	}

	return &Format28V1ACAS{
		Subtype:  fields.ReadSubtypeV1(data),
		ACASData: acasData,
	}, nil
}
