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
	FlightStatus         fields.FlightStatus
	DownlinkRequest      fields.DownlinkRequest
	UtilityMessage       fields.UtilityMessage
	AltitudeReportMethod fields.AltitudeReportMethod
	Altitude             int32
	MessageCommB         fields.MessageCommB
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

	// Format of the message is as follows:
	//
	//     DF     FS   |      DR     UM  |   UM      AC    |        AC       |  Comm-B |  AP/DP
	// 1 0 1 0 0 f f f | d d d d d u u u | u u u a a a a a | a a a a a a a a | 56 bits | 24bits

	if message.DownLinkFormat != 20 {
		return nil, errors.New("DF20 message must have a DownLinkFormat of 20")
	}

	if len(message.Payload) != 10 {
		return nil, errors.New("DF20 message must be 14 bytes long")
	}

	altitude, altitudeReportMethod, err := fields.ReadAltitude(message)
	if err != nil {
		return nil, errors.New("the field Altitude is not readable")
	}

	return &MessageDF20{
		MessageData:          message,
		FlightStatus:         fields.ReadFlightStatus(message),
		DownlinkRequest:      fields.ReadDownlinkRequest(message),
		UtilityMessage:       fields.ReadUtilityMessage(message),
		AltitudeReportMethod: altitudeReportMethod,
		Altitude:             altitude,
		MessageCommB:         fields.ReadMessageCommB(message),
	}, nil
}

// ToString returns a basic, but readable, representation of the field
func (message *MessageDF20) ToString() string {
	return fmt.Sprintf(""+
		"Downlink format:      %v - %v\n"+
		"FlightStatus:         %v\n"+
		"DownlinkRequest:      %v\n"+
		"UtilityMessage:       %v\n"+
		"AltitudeReportMethod: %v\n"+
		"Altitude:             %v\n"+
		"MessageCommB:         %v",
		message.GetDownLinkFormat(),
		message.GetName(),
		message.FlightStatus.ToString(),
		message.DownlinkRequest.ToString(),
		message.UtilityMessage.ToString(),
		message.AltitudeReportMethod.ToString(),
		message.Altitude,
		message.MessageCommB.ToString())
}
