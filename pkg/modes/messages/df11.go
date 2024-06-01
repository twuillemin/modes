package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/fields"
)

// -------------------------------------------------------------------------------------
//                                         DF11
// -------------------------------------------------------------------------------------

// MessageDF11 is a message at the format DF11
type MessageDF11 struct {
	common.MessageData
	Capability       fields.Capability
	AddressAnnounced fields.AddressAnnounced
}

// GetName returns the name of the message
func (message *MessageDF11) GetName() string {
	return "All-call reply"
}

// GetDownLinkFormat returns the downlink format of the message
func (message *MessageDF11) GetDownLinkFormat() int {
	return 11
}

// ParseDF11 parses a message at the DF11 format
func ParseDF11(message common.MessageData) (*MessageDF11, error) {

	// Format of the message is as follows:
	//
	//     DF     CA   |        AA       |        AA       |        AA       |   PI
	// 0 1 0 1 1 c c c | a a a a a a a a | a a a a a a a a | a a a a a a a a | 24bits

	if message.DownLinkFormat != 11 {
		return nil, errors.New("DF11 message must have a DownLinkFormat of 11")
	}
	if len(message.Payload) != 3 {
		return nil, errors.New("DF11 message must be 7 bytes long")
	}

	return &MessageDF11{
		MessageData:      message,
		Capability:       fields.ReadCapability(message),
		AddressAnnounced: fields.ReadAddressAnnounced(message),
	}, nil
}

// ToString returns a basic, but readable, representation of the field
func (message *MessageDF11) ToString() string {
	return fmt.Sprintf("Downlink format:  %v - %v\n"+
		"Capability:       %v\n"+
		"AddressAnnounced: %v",
		message.GetDownLinkFormat(),
		message.GetName(),
		message.Capability.ToString(),
		message.AddressAnnounced.ToString())
}
