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
	VerticalStatus       fields.VerticalStatus
	SensitivityLevel     fields.SensitivityLevelReport
	ReplyInformation     fields.ReplyInformationAirAir
	AltitudeReportMethod fields.AltitudeReportMethod
	Altitude             int32
	MessageACAS          fields.MessageACAS
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

	// Format of the message is as follows:
	//
	//     DF   VS _ _ |   SL  _ _   RI  |RI _ _     AC    |        AC       |    MV   |  AP
	// 1 0 0 0 0 x _ _ | x x x _ _ x x x | x _ _ x x x x x | x x x x x x x x | 56 bits |24bits

	if message.DownLinkFormat != 16 {
		return nil, errors.New("DF16 message must have a DownLinkFormat of 16")
	}

	if len(message.Payload) != 10 {
		return nil, errors.New("DF16 message must be 14 bytes long")
	}

	altitude, altitudeReportMethod, err := fields.ReadAltitude(message)
	if err != nil {
		return nil, errors.New("the field Altitude is not readable")
	}

	return &MessageDF16{
		MessageData:          message,
		VerticalStatus:       fields.ReadVerticalStatus(message),
		SensitivityLevel:     fields.ReadSensitivityLevelReport(message),
		ReplyInformation:     fields.ReadReplyInformationAirAir(message),
		AltitudeReportMethod: altitudeReportMethod,
		Altitude:             altitude,
		MessageACAS:          fields.ReadMessageACAS(message),
	}, nil
}

// ToString returns a basic, but readable, representation of the field
func (message *MessageDF16) ToString() string {
	return fmt.Sprintf(""+
		"Downlink format:      %v - %v\n"+
		"VerticalStatus:       %v\n"+
		"SensitivityLevel:     %v\n"+
		"ReplyInformation:     %v\n"+
		"AltitudeReportMethod: %v\n"+
		"Altitude:             %v\n"+
		"Message ACAS:         %v",
		message.GetDownLinkFormat(),
		message.GetName(),
		message.VerticalStatus.ToString(),
		message.SensitivityLevel.ToString(),
		message.ReplyInformation.ToString(),
		message.AltitudeReportMethod.ToString(),
		message.Altitude,
		message.MessageACAS.ToString())
}
