package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/fields"
)

// -------------------------------------------------------------------------------------
//                                         DF24
// -------------------------------------------------------------------------------------

// MessageDF24 is a message at the format DF24
type MessageDF24 struct {
	common.MessageData
	ControlELM       fields.ControlELM
	NumberOfDSegment fields.NumberOfDSegment
	MessageCommD     fields.MessageCommD
}

func (message *MessageDF24) GetName() string {
	return "Comm. D Extended Length MessageModeS (ELM)"
}

func (message *MessageDF24) GetDownLinkFormat() int {
	return 24
}

// ParseDF24 parses a message at the DF24 format
func ParseDF24(message common.MessageData) (*MessageDF24, error) {

	// Format of the message is as follow:
	//
	//  DF _ KE   ND   |  Comm-B |   AP
	// 1 0 _ k n n n n | 80 bits | 24bits

	if message.DownLinkFormat&0x18 == 0x18 {
		return nil, errors.New("DF24 message must have a DownLinkFormat of 24")
	}
	if len(message.Payload) != 10 {
		return nil, errors.New("DF24 message must be 14 bytes long")
	}

	return &MessageDF24{
		MessageData:      message,
		ControlELM:       fields.ReadControlELM(message),
		NumberOfDSegment: fields.ReadNumberOfDSegment(message),
		MessageCommD:     fields.ReadMessageCommD(message),
	}, nil
}

func (message *MessageDF24) PrettyPrint() {
	fmt.Printf("MessageModeS: %v\n", message.GetName())
	fmt.Printf("Downlink format:   %v\n", message.GetDownLinkFormat())
	fmt.Printf("ControlELM:        %v\n", message.ControlELM.PrettyPrint())
	fmt.Printf("NumberOfDSegment:  %v\n", message.NumberOfDSegment.PrettyPrint())
	fmt.Printf("MessageCommD:      %v\n", message.MessageCommD.PrettyPrint())
}
