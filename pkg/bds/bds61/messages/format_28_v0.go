package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
)

// Format28V0 is a message at the format BDS 6,1
type Format28V0 struct {
	Subtype                 fields.SubtypeV0
	EmergencyPriorityStatus fields.EmergencyPriorityStatus
}

// GetMessageFormat returns the ADSB format of the message
func (message *Format28V0) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format28V0
}

// GetRegister returns the register of the message
func (message *Format28V0) GetRegister() bds.Register {
	return adsb.Format28V0.GetRegister()
}

// GetSubtype returns the Subtype
func (message *Format28V0) GetSubtype() fields.Subtype {
	return message.Subtype
}

// GetEmergencyPriorityStatus returns the EmergencyPriorityStatus
func (message *Format28V0) GetEmergencyPriorityStatus() fields.EmergencyPriorityStatus {
	return message.EmergencyPriorityStatus
}

// ToString returns a basic, but readable, representation of the message
func (message *Format28V0) ToString() string {
	return fmt.Sprintf("Message:                   %v\n"+
		"Subtype:                   %v\n"+
		"Emergency Priority Status: %v",
		adsb.Format28V0.ToString(),
		message.GetSubtype().ToString(),
		message.GetEmergencyPriorityStatus().ToString())
}

// ReadFormat28V0 reads a message at the format BDS 6,1
func ReadFormat28V0(data []byte) (*Format28V0, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format28V0.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format28", formatTypeCode)
	}

	return &Format28V0{
		Subtype:                 fields.ReadSubtypeV0(data),
		EmergencyPriorityStatus: fields.ReadEmergencyPriorityStatus(data),
	}, nil
}
