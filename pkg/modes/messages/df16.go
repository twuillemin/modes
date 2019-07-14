package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/fields"
)

// -------------------------------------------------------------------------------------
//                                         DF16
// -------------------------------------------------------------------------------------

// MessageDF16 is a message at the format DF16
type MessageDF16 struct {
	common.MessageData
	VerticalStatus   fields.VerticalStatus
	SensitivityLevel fields.SensitivityLevelReport
	ReplyInformation fields.ReplyInformationAirAir
	AltitudeCode     fields.AltitudeCode
	MessageACAS      fields.MessageACAS
}

// GetName returns the name of the message
func (message *MessageDF16) GetName() string {
	return "Long air-air surveillance (ACAS)"
}

// GetDownLinkFormat returns the downlink format of the message
func (message *MessageDF16) GetDownLinkFormat() int {
	return 16
}

// ParseDF16 parses a message at the DF5 format
func ParseDF16(message common.MessageData) (*MessageDF16, error) {

	// Format of the message is as follow:
	//
	//     DF   VS _ _ |   SL  _ _   RI  |RI _ _     AC    |        AC       |    MV   |  AP
	// 1 0 0 0 0 x _ _ | x x x _ _ x x x | x _ _ x x x x x | x x x x x x x x | 56 bits |24bits

	if message.DownLinkFormat != 16 {
		return nil, errors.New("DF16 message must have a DownLinkFormat of 16")
	}
	if len(message.Payload) != 10 {
		return nil, errors.New("DF16 message must be 14 bytes long")
	}

	return &MessageDF16{
		MessageData:      message,
		VerticalStatus:   fields.ReadVerticalStatus(message),
		SensitivityLevel: fields.ReadSensitivityLevelReport(message),
		ReplyInformation: fields.ReadReplyInformationAirAir(message),
		AltitudeCode:     fields.ReadAltitudeCode(message),
		MessageACAS:      fields.ReadMessageACAS(message),
	}, nil
}

// ToString returns a basic, but readable, representation of the field
func (message *MessageDF16) ToString() string {
	return fmt.Sprintf("Message: %v\n"+
		"Downlink format:  %v\n"+
		"VerticalStatus:   %v\n"+
		"SensitivityLevel: %v\n"+
		"ReplyInformation: %v\n"+
		"AltitudeCode:     %v\n"+
		"MessageACAS:      %v",
		message.GetName(),
		message.GetDownLinkFormat(),
		message.VerticalStatus.ToString(),
		message.SensitivityLevel.ToString(),
		message.ReplyInformation.ToString(),
		message.AltitudeCode.ToString(),
		message.MessageACAS.ToString())
}
