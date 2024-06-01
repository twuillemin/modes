package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/bds61/fields"
	"github.com/twuillemin/modes/pkg/bds"
)

// Format28NoInformation is a message at the format BDS 6,1
type Format28NoInformation struct {
}

// GetMessageFormat returns the ADSB format of the message
func (message Format28NoInformation) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format28
}

// GetRegister returns the register of the message
func (message Format28NoInformation) GetRegister() bds.Register {
	return adsb.Format28.GetRegister()
}

// GetSubtype returns the Subtype
func (message Format28NoInformation) GetSubtype() adsb.Subtype {
	return fields.SubtypeNoInformation
}

// GetMinimumADSBLevel returns the minimum ADSB ReaderLevel for the message
func (message Format28NoInformation) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel0
}

// GetMaximumADSBLevel returns the maximum ADSB ReaderLevel for the message
func (message Format28NoInformation) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel2
}

// ToString returns a basic, but readable, representation of the message
func (message Format28NoInformation) ToString() string {
	return fmt.Sprintf("Message:                   %v\n"+
		"Subtype:                   %v - No Information",
		adsb.GetMessageFormatInformation(&message),
		message.GetSubtype().ToString())
}

// ReadFormat28NoInformation reads a message at the format Format28ACAS / subtype 0 (No information)
func ReadFormat28NoInformation(data []byte) (*Format28NoInformation, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format28.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format28ACAS", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeNoInformation {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadFormat28NoInformation", subType.ToString())
	}

	return &Format28NoInformation{}, nil
}
