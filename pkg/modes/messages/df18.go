package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/fields"
)

// -------------------------------------------------------------------------------------
//                                         DF18
// -------------------------------------------------------------------------------------

// MessageDF18 is a message at the format DF18
type MessageDF18 struct {
	common.MessageData
	ControlField            fields.ControlField
	AddressAnnounced        fields.AddressAnnounced
	MessageExtendedSquitter fields.MessageExtendedSquitter
}

// GetName returns the name of the message
func (message *MessageDF18) GetName() string {
	return "Extended squitter/non transponder"
}

// GetDownLinkFormat returns the downlink format of the message
func (message *MessageDF18) GetDownLinkFormat() int {
	return 18
}

// ParseDF18 parses a message at the DF17 format
func ParseDF18(message common.MessageData) (*MessageDF18, error) {

	// Format of the message is as follow:
	//
	//     DF     CF   |        AA       |        AA       |        AA       |    ME   |   PI
	// 1 0 0 1 0 c c c | a a a a a a a a | a a a a a a a a | a a a a a a a a | 56 bits | 24bits

	if message.DownLinkFormat != 18 {
		return nil, errors.New("DF18 message must have a DownLinkFormat of 18")
	}
	if len(message.Payload) != 10 {
		return nil, errors.New("DF18 message must be 14 bytes long")
	}

	return &MessageDF18{
		MessageData:             message,
		ControlField:            fields.ReadControlField(message),
		AddressAnnounced:        fields.ReadAddressAnnounced(message),
		MessageExtendedSquitter: fields.ReadMessageExtendedSquitter(message),
	}, nil
}

// ToString returns a basic, but readable, representation of the field
func (message *MessageDF18) ToString() string {
	return fmt.Sprintf("Downlink format:         %v - %v\n"+
		"ControlField:            %v\n"+
		"AddressAnnounced:        %v\n"+
		"MessageExtendedSquitter: %v",
		message.GetDownLinkFormat(),
		message.GetName(),
		message.ControlField.ToString(),
		message.AddressAnnounced.ToString(),
		message.MessageExtendedSquitter.ToString())
}
