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
type MessageDF0 struct {
	common.MessageData
	VerticalStatus      fields.VerticalStatus
	CrossLinkCapability fields.CrossLinkCompatibility
	SensitivityLevel    fields.SensitivityLevelReport
	ReplyInformation    fields.ReplyInformationAirAir
	AltitudeCode        fields.AltitudeCode
}

func (message *MessageDF0) GetName() string {
	return "Short air-air surveillance (ACAS)"
}

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

func (message *MessageDF0) PrettyPrint() {
	fmt.Printf("Message: %v\n", message.GetName())
	fmt.Printf("Downlink format:     %v\n", message.GetDownLinkFormat())
	fmt.Printf("VerticalStatus:      %v\n", message.VerticalStatus.PrettyPrint())
	fmt.Printf("CrossLinkCapability: %v\n", message.CrossLinkCapability.PrettyPrint())
	fmt.Printf("SensitivityLevel:    %v\n", message.SensitivityLevel.PrettyPrint())
	fmt.Printf("ReplyInformation:    %v\n", message.ReplyInformation.PrettyPrint())
	fmt.Printf("AltitudeCode:        %v\n", message.AltitudeCode.PrettyPrint())
}
