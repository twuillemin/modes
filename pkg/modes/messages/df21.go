package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/fields"
)

// -------------------------------------------------------------------------------------
//                                         DF21
// -------------------------------------------------------------------------------------

// MessageDF21 is a message at the format DF21
type MessageDF21 struct {
	common.MessageData
	FlightStatus    fields.FlightStatus
	DownlinkRequest fields.DownlinkRequest
	UtilityMessage  fields.UtilityMessage
	Identity        fields.Identity
	MessageCommB    fields.MessageCommB
}

func (message *MessageDF21) GetName() string {
	return "Comm-B identify reply"
}

func (message *MessageDF21) GetDownLinkFormat() int {
	return 21
}

// ParseDF21 parses a message at the DF21 format
func ParseDF21(message common.MessageData) (*MessageDF21, error) {

	// Format of the message is as follow:
	//
	//     DF     FS   |      DR     UM  |   UM      ID    |        ID       |  Comm-B |  AP/DP
	// 1 0 1 0 1 f f f | d d d d d u u u | u u u i i i i i | i i i i i i i i | 56 bits | 24bits

	if message.DownLinkFormat != 21 {
		return nil, errors.New("DF21 message must have a DownLinkFormat of 21")
	}
	if len(message.Payload) != 10 {
		return nil, errors.New("DF21 message must be 14 bytes long")
	}

	return &MessageDF21{
		MessageData:     message,
		FlightStatus:    fields.ReadFlightStatus(message),
		DownlinkRequest: fields.ReadDownlinkRequest(message),
		UtilityMessage:  fields.ReadUtilityMessage(message),
		Identity:        fields.ReadIdentity(message),
		MessageCommB:    fields.ReadMessageCommB(message),
	}, nil
}

func (message *MessageDF21) PrettyPrint() {
	fmt.Printf("Message: %v\n", message.GetName())
	fmt.Printf("Downlink format:  %v\n", message.GetDownLinkFormat())
	fmt.Printf("FlightStatus:     %v\n", message.FlightStatus.PrettyPrint())
	fmt.Printf("DownlinkRequest:  %v\n", message.DownlinkRequest.PrettyPrint())
	fmt.Printf("UtilityMessage:   %v\n", message.UtilityMessage.PrettyPrint())
	fmt.Printf("Identity:         %v\n", message.Identity.PrettyPrint())
	fmt.Printf("MessageCommB:     %v\n", message.MessageCommB.PrettyPrint())
}
