package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
)

// Format28StatusV0 is a message at the format BDS 6,1
type Format28StatusV0 struct {
	EmergencyPriorityStatus fields.EmergencyPriorityStatus
}

// GetMessageFormat returns the ADSB format of the message
func (message Format28StatusV0) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format28
}

// GetRegister returns the register of the message
func (message Format28StatusV0) GetRegister() bds.Register {
	return adsb.Format28.GetRegister()
}

// GetSubtype returns the Subtype
func (message Format28StatusV0) GetSubtype() adsb.Subtype {
	return fields.SubtypeEmergencyPriorityStatus
}

// GetMinimumADSBLevel returns the minimum ADSB ReaderLevel for the message
func (message Format28StatusV0) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel0
}

// GetMaximumADSBLevel returns the maximum ADSB ReaderLevel for the message
func (message Format28StatusV0) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel0
}

// GetEmergencyPriorityStatus returns the EmergencyPriorityStatus
func (message Format28StatusV0) GetEmergencyPriorityStatus() fields.EmergencyPriorityStatus {
	return message.EmergencyPriorityStatus
}

// ToString returns a basic, but readable, representation of the message
func (message Format28StatusV0) ToString() string {
	return fmt.Sprintf("Message:                   %v\n"+
		"Subtype:                   %v\n"+
		"Emergency Priority Status: %v",
		adsb.GetMessageFormatInformation(&message),
		message.GetSubtype().ToString(),
		message.GetEmergencyPriorityStatus().ToString())
}

// ReadFormat28StatusV0 reads a message at the format Format 28 / Subtype 1 (Emergency/priority status) for ADSB V0
func ReadFormat28StatusV0(data []byte) (*Format28StatusV0, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format28.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read by ReadFormat28StatusV0", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeEmergencyPriorityStatus {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadFormat28StatusV0", subType.ToString())
	}

	return &Format28StatusV0{
		EmergencyPriorityStatus: fields.ReadEmergencyPriorityStatus(data),
	}, nil
}
