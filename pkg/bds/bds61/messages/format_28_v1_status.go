package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
)

// Format28V1Status is a message at the format BDS 6,1
type Format28V1Status struct {
	Subtype                 fields.SubtypeV1
	EmergencyPriorityStatus fields.EmergencyPriorityStatus
}

// GetMessageFormat returns the ADSB format of the message
func (message *Format28V1Status) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format28V1
}

// GetRegister returns the register of the message
func (message *Format28V1Status) GetRegister() bds.Register {
	return adsb.Format28V1.GetRegister()
}

// GetSubtype returns the Subtype
func (message *Format28V1Status) GetSubtype() fields.Subtype {
	return message.Subtype
}

// GetEmergencyPriorityStatus returns the EmergencyPriorityStatus
func (message *Format28V1Status) GetEmergencyPriorityStatus() fields.EmergencyPriorityStatus {
	return message.EmergencyPriorityStatus
}

// ToString returns a basic, but readable, representation of the message
func (message *Format28V1Status) ToString() string {
	return fmt.Sprintf("Message:                   %v\n"+
		"Subtype:                   %v\n"+
		"Emergency Priority Status: %v",
		adsb.Format28V1.ToString(),
		message.GetSubtype().ToString(),
		message.GetEmergencyPriorityStatus().ToString())
}

// readFormat28V1Status reads a message at the format BDS 6,1
func readFormat28V1Status(data []byte) (*Format28V1Status, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format28V1.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format28", formatTypeCode)
	}

	return &Format28V1Status{
		Subtype:                 fields.ReadSubtypeV1(data),
		EmergencyPriorityStatus: fields.ReadEmergencyPriorityStatus(data),
	}, nil
}
