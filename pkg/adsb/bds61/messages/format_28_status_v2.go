package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/bds61/fields"
	"github.com/twuillemin/modes/pkg/bds"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// Format28StatusV2 is a message at the format BDS 6,1
type Format28StatusV2 struct {
	EmergencyPriorityStatus fields.EmergencyPriorityStatus
	ModeACode               uint16
}

// GetMessageFormat returns the ADSB format of the message
func (message Format28StatusV2) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format28
}

// GetRegister returns the register of the message
func (message Format28StatusV2) GetRegister() bds.Register {
	return adsb.Format28.GetRegister()
}

// GetSubtype returns the Subtype
func (message Format28StatusV2) GetSubtype() adsb.Subtype {
	return fields.SubtypeEmergencyPriorityStatus
}

// GetMinimumADSBLevel returns the minimum ADSB ReaderLevel for the message
func (message Format28StatusV2) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel1
}

// GetMaximumADSBLevel returns the maximum ADSB ReaderLevel for the message
func (message Format28StatusV2) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel1
}

// GetEmergencyPriorityStatus returns the EmergencyPriorityStatus
func (message Format28StatusV2) GetEmergencyPriorityStatus() fields.EmergencyPriorityStatus {
	return message.EmergencyPriorityStatus
}

// ToString returns a basic, but readable, representation of the message
func (message Format28StatusV2) ToString() string {
	return fmt.Sprintf("Message:                   %v\n"+
		"Subtype:                   %v\n"+
		"Emergency Priority Status: %v\n"+
		"Mode A Code:               %v\n",
		adsb.GetMessageFormatInformation(&message),
		message.GetSubtype().ToString(),
		message.GetEmergencyPriorityStatus().ToString(),
		message.ModeACode)
}

// ReadFormat28StatusV2 reads a message at the format Format 28 / Subtype 1 (Emergency/priority status) for ADSB V2
func ReadFormat28StatusV2(data []byte) (*Format28StatusV2, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format28.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read by ReadFormat28StatusV2", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeEmergencyPriorityStatus {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadFormat28StatusV2", subType.ToString())
	}

	byte1 := data[1] & 0x1F
	byte2 := data[2]

	return &Format28StatusV2{
		EmergencyPriorityStatus: fields.ReadEmergencyPriorityStatus(data),
		ModeACode:               bitutils.Pack2Bytes(byte1, byte2),
	}, nil
}
