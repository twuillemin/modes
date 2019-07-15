package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/fields"
)

// -------------------------------------------------------------------------------------
//                                         DF19
// -------------------------------------------------------------------------------------

// MessageDF19 is a message at the format DF19
type MessageDF19 struct {
	common.MessageData
	ApplicationField fields.ApplicationField
}

// GetName returns the name of the message
func (message *MessageDF19) GetName() string {
	return "Military extended squitter"
}

// GetDownLinkFormat returns the downlink format of the message
func (message *MessageDF19) GetDownLinkFormat() int {
	return 19
}

// ParseDF19 parses a message at the DF19 format
func ParseDF19(message common.MessageData) (*MessageDF19, error) {

	// Format of the message is as follow:
	//
	//     DF     AF   | Military use
	// 1 0 0 1 1 a a a |   104 bits

	if message.DownLinkFormat != 19 {
		return nil, errors.New("DF19 message must have a DownLinkFormat of 19")
	}
	if len(message.Payload) != 10 {
		return nil, errors.New("DF19 message must be 14 bytes long")
	}

	return &MessageDF19{
		MessageData:      message,
		ApplicationField: fields.ReadApplicationField(message),
	}, nil
}

// ToString returns a basic, but readable, representation of the field
func (message *MessageDF19) ToString() string {
	return fmt.Sprintf("ModeSMessage: %v\n"+
		"Downlink format:         %v\n"+
		"ApplicationField:        %v",
		message.GetName(),
		message.GetDownLinkFormat(),
		message.ApplicationField.ToString())
}
