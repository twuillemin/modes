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

func (message *MessageDF17) GetName() string {
	return "1090 Extended Squitter"
}

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

func (message *MessageDF17) PrettyPrint() {
	fmt.Printf("MessageModeS: %v\n", message.GetName())
	fmt.Printf("Downlink format:          %v\n", message.GetDownLinkFormat())
	fmt.Printf("Capability:               %v\n", message.Capability.PrettyPrint())
	fmt.Printf("AddressAnnounced:         %v\n", message.AddressAnnounced.PrettyPrint())
	fmt.Printf("MessageExtendedSquitter:  %v\n", message.MessageExtendedSquitter.PrettyPrint())
}
