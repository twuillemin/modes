package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
)

// Format28V1ACAS is a message at the format BDS 6,1
type Format28V1ACAS struct {
	Subtype  fields.SubtypeV1
	ACASData []byte
}

// GetMessageFormat returns the ADSB format of the message
func (message *Format28V1ACAS) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format28V1
}

// GetRegister returns the register of the message
func (message *Format28V1ACAS) GetRegister() bds.Register {
	return adsb.Format28V1.GetRegister()
}

// GetSubtype returns the Subtype
func (message *Format28V1ACAS) GetSubtype() fields.Subtype {
	return message.Subtype
}

// GetACASData returns the ACAS Data
func (message *Format28V1ACAS) GetACASData() []byte {
	return message.ACASData
}

// ToString returns a basic, but readable, representation of the message
func (message *Format28V1ACAS) ToString() string {
	return fmt.Sprintf("Message:                   %v\n"+
		"Subtype:                   %v\n"+
		"ACAS Data                  : %02X %02X %02X %02X %02X %02X",
		adsb.Format28V1.ToString(),
		message.GetSubtype().ToString(),
		message.ACASData[0], message.ACASData[1], message.ACASData[2], message.ACASData[3], message.ACASData[4], message.ACASData[5])
}

// readFormat28V1ACAS reads a message at the format BDS 6,1
func readFormat28V1ACAS(data []byte) (*Format28V1ACAS, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format28V1.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format28", formatTypeCode)
	}

	// Copy data to not keep a reference to the given data
	acasData := make([]byte, 6)
	for i := 0; i < 6; i++ {
		acasData[i] = data[i+1]
	}

	return &Format28V1ACAS{
		Subtype:  fields.ReadSubtypeV1(data),
		ACASData: acasData,
	}, nil
}
