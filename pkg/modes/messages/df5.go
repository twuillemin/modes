package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/fields"
)

// -------------------------------------------------------------------------------------
//                                         DF5
// -------------------------------------------------------------------------------------

// MessageDF5 is a message at the format DF5
type MessageDF5 struct {
	common.MessageData
	FlightStatus    fields.FlightStatus
	DownlinkRequest fields.DownlinkRequest
	UtilityMessage  fields.UtilityMessage
	Identity        fields.Identity
}

func (message *MessageDF5) GetName() string {
	return "Surveillance (roll call) IDENT Reply"
}

func (message *MessageDF5) GetDownLinkFormat() int {
	return 5
}

// ParseDF5 parses a message at the DF5 format
func ParseDF5(message common.MessageData) (*MessageDF5, error) {

	// Format of the message is as follow:
	//
	//     DF     FS   |      DR     UM  |   UM      ID    |        ID       |   AP
	// 0 0 1 0 1 f f f | d d d d d u u u | u u u i i i i i | i i i i i i i i | 24bits

	if message.DownLinkFormat != 5 {
		return nil, errors.New("DF5 message must have a DownLinkFormat of 5")
	}
	if len(message.Payload) != 3 {
		return nil, errors.New("DF5 message must be 7 bytes long")
	}

	return &MessageDF5{
		MessageData:     message,
		FlightStatus:    fields.ReadFlightStatus(message),
		DownlinkRequest: fields.ReadDownlinkRequest(message),
		UtilityMessage:  fields.ReadUtilityMessage(message),
		Identity:        fields.ReadIdentity(message),
	}, nil
}

func (message *MessageDF5) PrettyPrint() {
	fmt.Printf("Message: %v\n", message.GetName())
	fmt.Printf("Downlink format:  %v\n", message.GetDownLinkFormat())
	fmt.Printf("FlightStatus:     %v\n", message.FlightStatus.PrettyPrint())
	fmt.Printf("DownlinkRequest:  %v\n", message.DownlinkRequest.PrettyPrint())
	fmt.Printf("UtilityMessage:   %v\n", message.UtilityMessage.PrettyPrint())
	fmt.Printf("Identity:         %v\n", message.Identity.PrettyPrint())
}
