package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/fields"
)

// -------------------------------------------------------------------------------------
//                                         DF24
// -------------------------------------------------------------------------------------

// MessageDF24 is a message at the format DF24
type MessageDF24 struct {
	common.MessageData
	ControlELM       fields.ControlELM
	NumberOfDSegment fields.NumberOfDSegment
	MessageCommD     fields.MessageCommD
}

// GetName returns the name of the message
func (message *MessageDF24) GetName() string {
	return "Comm-D (ELM)"
}

// GetDownLinkFormat returns the downlink format of the message
func (message *MessageDF24) GetDownLinkFormat() int {
	return 24
}

// ParseDF24 parses a message at the DF24 format
func ParseDF24(message common.MessageData) (*MessageDF24, error) {

	// Format of the message is as follow:
	//
	//  DF _ KE   ND   |  Comm-B |   AP
	// 1 1 _ k n n n n | 80 bits | 24bits

	if message.DownLinkFormat&0x18 == 0x18 {
		return nil, errors.New("DF24 message must have a DownLinkFormat of 24")
	}
	if len(message.Payload) != 10 {
		return nil, errors.New("DF24 message must be 14 bytes long")
	}

	return &MessageDF24{
		MessageData:      message,
		ControlELM:       fields.ReadControlELM(message),
		NumberOfDSegment: fields.ReadNumberOfDSegment(message),
		MessageCommD:     fields.ReadMessageCommD(message),
	}, nil
}

// ToString returns a basic, but readable, representation of the field
func (message *MessageDF24) ToString() string {
	return fmt.Sprintf("Downlink format:  %v - %v\n"+
		"ControlELM:       %v\n"+
		"NumberOfDSegment: %v\n"+
		"MessageCommD:     %v",
		message.GetDownLinkFormat(),
		message.GetName(),
		message.ControlELM.ToString(),
		message.NumberOfDSegment.ToString(),
		message.MessageCommD.ToString())
}
