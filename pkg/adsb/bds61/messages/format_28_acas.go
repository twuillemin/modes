package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/bds61/fields"
	"github.com/twuillemin/modes/pkg/bds"
)

// Format28ACAS is a message at the format BDS 6,1
type Format28ACAS struct {
	ACASData []byte
}

// GetMessageFormat returns the ADSB format of the message
func (message Format28ACAS) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format28
}

// GetRegister returns the register of the message
func (message Format28ACAS) GetRegister() bds.Register {
	return adsb.Format28.GetRegister()
}

// GetSubtype returns the Subtype
func (message Format28ACAS) GetSubtype() adsb.Subtype {
	return fields.SubtypeRABroadcast
}

// GetMinimumADSBLevel returns the minimum ADSB ReaderLevel for the message
func (message Format28ACAS) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel1
}

// GetMaximumADSBLevel returns the maximum ADSB ReaderLevel for the message
func (message Format28ACAS) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel2
}

// GetACASData returns the ACAS Data
func (message Format28ACAS) GetACASData() []byte {
	return message.ACASData
}

// ToString returns a basic, but readable, representation of the message
func (message Format28ACAS) ToString() string {
	return fmt.Sprintf("Message:                   %v\n"+
		"Subtype:                   %v\n"+
		"ACAS Data                  : %02X %02X %02X %02X %02X %02X",
		adsb.GetMessageFormatInformation(&message),
		message.GetSubtype().ToString(),
		message.ACASData[0], message.ACASData[1], message.ACASData[2], message.ACASData[3], message.ACASData[4], message.ACASData[5])
}

// ReadFormat28ACAS reads a message at the format Format 28 / Subtype 2 (ACAS RA Broadcast)
func ReadFormat28ACAS(data []byte) (*Format28ACAS, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format28.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format28ACAS", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeRABroadcast {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadFormat28ACAS", subType.ToString())
	}

	// Copy data to not keep a reference to the given data
	acasData := make([]byte, 6)
	for i := 0; i < 6; i++ {
		acasData[i] = data[i+1]
	}

	return &Format28ACAS{
		ACASData: acasData,
	}, nil
}
