package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/bds61/fields"
)

// Format28StatusV1 is a message at the format BDS 6,1
type Format28StatusV1 struct {
	EmergencyPriorityStatus fields.EmergencyPriorityStatus
}

// GetMessageFormat returns the ADSB format of the message
func (message Format28StatusV1) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format28
}

// GetSubtype returns the Subtype
func (message Format28StatusV1) GetSubtype() adsb.Subtype {
	return fields.SubtypeEmergencyPriorityStatus
}

// GetMinimumADSBLevel returns the minimum ADSB ReaderLevel for the message
func (message Format28StatusV1) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel1
}

// GetMaximumADSBLevel returns the maximum ADSB ReaderLevel for the message
func (message Format28StatusV1) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel1
}

// GetEmergencyPriorityStatus returns the EmergencyPriorityStatus
func (message Format28StatusV1) GetEmergencyPriorityStatus() fields.EmergencyPriorityStatus {
	return message.EmergencyPriorityStatus
}

// ToString returns a basic, but readable, representation of the message
func (message Format28StatusV1) ToString() string {
	return fmt.Sprintf("Message:                   %v\n"+
		"Subtype:                   %v\n"+
		"Emergency Priority Status: %v",
		adsb.GetMessageFormatInformation(&message),
		message.GetSubtype().ToString(),
		message.GetEmergencyPriorityStatus().ToString())
}

// ReadFormat28StatusV1 reads a message at the format Format 28 / Subtype 1 (Emergency/priority status) for ADSB V1
func ReadFormat28StatusV1(data []byte) (*Format28StatusV1, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format28.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read by ReadFormat28StatusV1", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeEmergencyPriorityStatus {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadFormat28StatusV1", subType.ToString())
	}

	return &Format28StatusV1{
		EmergencyPriorityStatus: fields.ReadEmergencyPriorityStatus(data),
	}, nil
}
