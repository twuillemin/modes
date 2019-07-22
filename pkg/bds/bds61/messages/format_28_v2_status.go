package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// Format28V2Status is a message at the format BDS 6,1
type Format28V2Status struct {
	Subtype                 fields.SubtypeV2
	EmergencyPriorityStatus fields.EmergencyPriorityStatus
	ModeACode               uint16
}

// GetMessageFormat returns the ADSB format of the message
func (message *Format28V2Status) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format28V2
}

// GetRegister returns the register of the message
func (message *Format28V2Status) GetRegister() bds.Register {
	return adsb.Format28V2.GetRegister()
}

// GetSubtype returns the Subtype
func (message *Format28V2Status) GetSubtype() fields.Subtype {
	return message.Subtype
}

// GetEmergencyPriorityStatus returns the EmergencyPriorityStatus
func (message *Format28V2Status) GetEmergencyPriorityStatus() fields.EmergencyPriorityStatus {
	return message.EmergencyPriorityStatus
}

// ToString returns a basic, but readable, representation of the message
func (message *Format28V2Status) ToString() string {
	return fmt.Sprintf("Message:                   %v\n"+
		"Subtype:                   %v\n"+
		"Emergency Priority Status: %v\n"+
		"Mode A Code:               %v\n",
		adsb.Format28V2.ToString(),
		message.GetSubtype().ToString(),
		message.GetEmergencyPriorityStatus().ToString(),
		message.ModeACode)
}

// readFormat28V2Status reads a message at the format BDS 6,1
func readFormat28V2Status(data []byte) (*Format28V2Status, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format28V2.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format28", formatTypeCode)
	}

	byte1 := data[1] & 0x1F
	byte2 := data[2]

	return &Format28V2Status{
		Subtype:                 fields.ReadSubtypeV2(data),
		EmergencyPriorityStatus: fields.ReadEmergencyPriorityStatus(data),
		ModeACode:               bitutils.Pack2Bytes(byte1, byte2),
	}, nil
}
