package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
)

// Format28NoInformation is a message at the format BDS 6,1
type Format28NoInformation struct {
	Subtype fields.SubtypeV0
}

// GetMessageFormat returns the ADSB format of the message
func (message *Format28NoInformation) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format28V0OrMore
}

// GetRegister returns the register of the message
func (message *Format28NoInformation) GetRegister() bds.Register {
	return adsb.Format28V0OrMore.GetRegister()
}

// GetSubtype returns the Subtype
func (message *Format28NoInformation) GetSubtype() fields.Subtype {
	return message.Subtype
}

// ToString returns a basic, but readable, representation of the message
func (message *Format28NoInformation) ToString() string {
	return fmt.Sprintf("Message:                   %v\n"+
		"Subtype:                   %v - No Information",
		adsb.Format28V0OrMore.ToString(),
		message.GetSubtype().ToString())
}

// readFormat28NoInformation reads a message at the format BDS 6,1 / subtype 0
func readFormat28NoInformation(data []byte) (*Format28NoInformation, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format28V0.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format28", formatTypeCode)
	}

	return &Format28NoInformation{
		Subtype: fields.ReadSubtypeV0(data),
	}, nil
}
