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

func (message *MessageDF16) GetName() string {
	return "Long Air to Air ACAS"
}

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

func (message *MessageDF16) PrettyPrint() {
	fmt.Printf("MessageModeS: %v\n", message.GetName())
	fmt.Printf("Downlink format:   %v\n", message.GetDownLinkFormat())
	fmt.Printf("VerticalStatus:    %v\n", message.VerticalStatus.PrettyPrint())
	fmt.Printf("SensitivityLevel:  %v\n", message.SensitivityLevel.PrettyPrint())
	fmt.Printf("ReplyInformation:  %v\n", message.ReplyInformation.PrettyPrint())
	fmt.Printf("AltitudeCode:      %v\n", message.AltitudeCode.PrettyPrint())
	fmt.Printf("MessageACAS:       %v\n", message.MessageACAS.PrettyPrint())
}
