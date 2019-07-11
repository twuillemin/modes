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

func (message *MessageDF18) GetName() string {
	return "1090 Extended Squitter, supplementary"
}

func (message *MessageDF18) GetDownLinkFormat() int {
	return 18
}

// ParseDF18 parses a message at the DF17 format
func ParseDF18(message common.MessageData) (*MessageDF18, error) {

	// Format of the message is as follow:
	//
	//     DF     CF   |        AA       |        AA       |        AA       |    ME   |   PI
	// 1 0 0 1 0 c c c | a a a a a a a a | a a a a a a a a | a a a a a a a a | 56 bits | 24bits

	if message.DownLinkFormat != 17 {
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

func (message *MessageDF18) PrettyPrint() {
	fmt.Printf("Message: %v\n", message.GetName())
	fmt.Printf("Downlink format:          %v\n", message.GetDownLinkFormat())
	fmt.Printf("ControlField:             %v\n", message.ControlField.PrettyPrint())
	fmt.Printf("AddressAnnounced:         %v\n", message.AddressAnnounced.PrettyPrint())
	fmt.Printf("MessageExtendedSquitter:  %v\n", message.MessageExtendedSquitter.PrettyPrint())
}
