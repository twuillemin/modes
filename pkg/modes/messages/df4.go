package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/fields"
)

// -------------------------------------------------------------------------------------
//                                         DF4
// -------------------------------------------------------------------------------------

// MessageDF4 is a message at the format DF4
type MessageDF4 struct {
	common.MessageData
	FlightStatus    fields.FlightStatus
	DownlinkRequest fields.DownlinkRequest
	UtilityMessage  fields.UtilityMessage
	AltitudeCode    fields.AltitudeCode
}

// GetName returns the name of the message
func (message *MessageDF4) GetName() string {
	return "Surveillance, altitude reply"
}

// GetDownLinkFormat returns the downlink format of the message
func (message *MessageDF4) GetDownLinkFormat() int {
	return 4
}

// ParseDF4 parses a message at the DF4 format
func ParseDF4(message common.MessageData) (*MessageDF4, error) {

	// Format of the message is as follow:
	//
	//     DF     FS   |      DR     UM  |   UM      AC    |        AC       |   AP
	// 0 0 1 0 0 f f f | d d d d d u u u | u u u a a a a a | a a a a a a a a | 24bits

	if message.DownLinkFormat != 4 {
		return nil, errors.New("DF4 message must have a DownLinkFormat of 4")
	}
	if len(message.Payload) != 3 {
		return nil, errors.New("DF4 message must be 7 bytes long")
	}

	return &MessageDF4{
		MessageData:     message,
		FlightStatus:    fields.ReadFlightStatus(message),
		DownlinkRequest: fields.ReadDownlinkRequest(message),
		UtilityMessage:  fields.ReadUtilityMessage(message),
		AltitudeCode:    fields.ReadAltitudeCode(message),
	}, nil
}

// ToString returns a basic, but readable, representation of the field
func (message *MessageDF4) ToString() string {
	return fmt.Sprintf("ModeSMessage: %v\n"+
		"Downlink format:  %v\n"+
		"FlightStatus:     %v\n"+
		"DownlinkRequest:  %v\n"+
		"UtilityMessage:   %v\n"+
		"AltitudeCode:     %v",
		message.GetName(),
		message.GetDownLinkFormat(),
		message.FlightStatus.ToString(),
		message.DownlinkRequest.ToString(),
		message.UtilityMessage.ToString(),
		message.AltitudeCode.ToString())
}
