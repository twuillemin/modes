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

// GetName returns the name of the message
func (message *MessageDF5) GetName() string {
	return "Surveillance, identify reply"
}

// GetDownLinkFormat returns the downlink format of the message
func (message *MessageDF5) GetDownLinkFormat() int {
	return 5
}

// ParseDF5 parses a message at the DF5 format
func ParseDF5(message common.MessageData) (*MessageDF5, error) {

	// Format of the message is as follows:
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

// ToString returns a basic, but readable, representation of the field
func (message *MessageDF5) ToString() string {
	return fmt.Sprintf("Downlink format:  %v - %v\n"+
		"FlightStatus:     %v\n"+
		"DownlinkRequest:  %v\n"+
		"UtilityMessage:   %v\n"+
		"Identity:         %v",
		message.GetDownLinkFormat(),
		message.GetName(),
		message.FlightStatus.ToString(),
		message.DownlinkRequest.ToString(),
		message.UtilityMessage.ToString(),
		message.Identity.ToString())
}
