package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/fields"
)

// -------------------------------------------------------------------------------------
//                                         DF17
// -------------------------------------------------------------------------------------

// MessageDF17 is a message at the format DF17
type MessageDF17 struct {
	common.MessageData
	Capability              fields.Capability
	AddressAnnounced        fields.AddressAnnounced
	MessageExtendedSquitter fields.MessageExtendedSquitter
}

// GetName returns the name of the message
func (message *MessageDF17) GetName() string {
	return "Extended squitter"
}

// GetDownLinkFormat returns the downlink format of the message
func (message *MessageDF17) GetDownLinkFormat() int {
	return 17
}

// ParseDF17 parses a message at the DF17 format
func ParseDF17(message common.MessageData) (*MessageDF17, error) {

	// Format of the message is as follow:
	//
	//     DF     CA   |        AA       |        AA       |        AA       |    ME   |   PI
	// 1 0 0 0 1 c c c | a a a a a a a a | a a a a a a a a | a a a a a a a a | 56 bits | 24bits

	if message.DownLinkFormat != 17 {
		return nil, errors.New("DF17 message must have a DownLinkFormat of 17")
	}
	if len(message.Payload) != 10 {
		return nil, errors.New("DF17 message must be 14 bytes long")
	}

	return &MessageDF17{
		MessageData:             message,
		Capability:              fields.ReadCapability(message),
		AddressAnnounced:        fields.ReadAddressAnnounced(message),
		MessageExtendedSquitter: fields.ReadMessageExtendedSquitter(message),
	}, nil
}

// ToString returns a basic, but readable, representation of the field
func (message *MessageDF17) ToString() string {
	return fmt.Sprintf("ModeSMessage: %v\n"+
		"Downlink format:         %v\n"+
		"Capability:              %v\n"+
		"AddressAnnounced:        %v\n"+
		"MessageExtendedSquitter: %v",
		message.GetName(),
		message.GetDownLinkFormat(),
		message.Capability.ToString(),
		message.AddressAnnounced.ToString(),
		message.MessageExtendedSquitter.ToString())
}
