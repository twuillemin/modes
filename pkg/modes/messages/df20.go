package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/fields"
)

// -------------------------------------------------------------------------------------
//                                         DF20
// -------------------------------------------------------------------------------------

// MessageDF20 is a message at the format DF20
type MessageDF20 struct {
	common.MessageData
	FlightStatus    fields.FlightStatus
	DownlinkRequest fields.DownlinkRequest
	UtilityMessage  fields.UtilityMessage
	AltitudeCode    fields.AltitudeCode
	MessageCommB    fields.MessageCommB
}

// GetName returns the name of the message
func (message *MessageDF20) GetName() string {
	return "Comm-B altitude reply"
}

// GetDownLinkFormat returns the downlink format of the message
func (message *MessageDF20) GetDownLinkFormat() int {
	return 20
}

// ParseDF20 parses a message at the DF4 format
func ParseDF20(message common.MessageData) (*MessageDF20, error) {

	// Format of the message is as follow:
	//
	//     DF     FS   |      DR     UM  |   UM      AC    |        AC       |  Comm-B |  AP/DP
	// 1 0 1 0 0 f f f | d d d d d u u u | u u u a a a a a | a a a a a a a a | 56 bits | 24bits

	if message.DownLinkFormat != 20 {
		return nil, errors.New("DF20 message must have a DownLinkFormat of 20")
	}
	if len(message.Payload) != 10 {
		return nil, errors.New("DF20 message must be 14 bytes long")
	}

	return &MessageDF20{
		MessageData:     message,
		FlightStatus:    fields.ReadFlightStatus(message),
		DownlinkRequest: fields.ReadDownlinkRequest(message),
		UtilityMessage:  fields.ReadUtilityMessage(message),
		AltitudeCode:    fields.ReadAltitudeCode(message),
		MessageCommB:    fields.ReadMessageCommB(message),
	}, nil
}

// ToString returns a basic, but readable, representation of the field
func (message *MessageDF20) ToString() {
	fmt.Printf("Message: %v\n", message.GetName())
	fmt.Printf("Downlink format:  %v\n", message.GetDownLinkFormat())
	fmt.Printf("FlightStatus:     %v\n", message.FlightStatus.ToString())
	fmt.Printf("DownlinkRequest:  %v\n", message.DownlinkRequest.ToString())
	fmt.Printf("UtilityMessage:   %v\n", message.UtilityMessage.ToString())
	fmt.Printf("AltitudeCode:     %v\n", message.AltitudeCode.ToString())
	fmt.Printf("MessageCommB:     %v\n", message.MessageCommB.ToString())
}
