package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/fields"
)

// -------------------------------------------------------------------------------------
//                                         DF0
// -------------------------------------------------------------------------------------

// MessageDF0 is a message at the format DF0
//
// Specified in Annex 10, Volume IV, 3.1.2.8.2
type MessageDF0 struct {
	common.MessageData
	VerticalStatus      fields.VerticalStatus
	CrossLinkCapability fields.CrossLinkCompatibility
	SensitivityLevel    fields.SensitivityLevelReport
	ReplyInformation    fields.ReplyInformationAirAir
	AltitudeCode        fields.AltitudeCode
}

// GetName returns the name of the message
func (message *MessageDF0) GetName() string {
	return "Short air-air surveillance (ACAS)"
}

// GetDownLinkFormat returns the downlink format of the message
func (message *MessageDF0) GetDownLinkFormat() int {
	return 0
}

// ParseDF0 parses a message at the DF0 format
func ParseDF0(message common.MessageData) (*MessageDF0, error) {

	// Format of the message is as follow:
	//
	//     DF  VS CC _ |   SL  _ _   RI  |RI _ _     AC    |        AC       |   AP
	// 0 0 0 0 0 x x _ | x x x _ _ x x x | x _ _ x x x x x | x x x x x x x x | 24bits

	if message.DownLinkFormat != 0 {
		return nil, errors.New("DF0 message must have a DownLinkFormat of 0")
	}

	if len(message.Payload) != 3 {
		return nil, errors.New("DF0 message must be 7 bytes long")
	}

	return &MessageDF0{
		MessageData:         message,
		VerticalStatus:      fields.ReadVerticalStatus(message),
		CrossLinkCapability: fields.ReadCrossLinkCompatibility(message),
		SensitivityLevel:    fields.ReadSensitivityLevelReport(message),
		ReplyInformation:    fields.ReadReplyInformationAirAir(message),
		AltitudeCode:        fields.ReadAltitudeCode(message),
	}, nil
}

// ToString returns a basic, but readable, representation of the field
func (message *MessageDF0) ToString() string {
	return fmt.Sprintf("Downlink format:     %v - %v\n"+
		"VerticalStatus:      %v\n"+
		"CrossLinkCapability: %v\n"+
		"SensitivityLevel:    %v\n"+
		"ReplyInformation:    %v\n"+
		"AltitudeCode:        %v",
		message.GetDownLinkFormat(),
		message.GetName(),
		message.VerticalStatus.ToString(),
		message.CrossLinkCapability.ToString(),
		message.SensitivityLevel.ToString(),
		message.ReplyInformation.ToString(),
		message.AltitudeCode.ToString())
}
