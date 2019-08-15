package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	fields2 "github.com/twuillemin/modes/pkg/bds/bds08/fields"
)

// Format03 is a message at the format BDS 0,8
type Format03 struct {
	AircraftCategory       fields2.AircraftCategorySetB
	AircraftIdentification fields2.AircraftIdentification
}

// GetMessageFormat returns the ADSB format of the message
func (message *Format03) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format03V0OrMore
}

// GetRegister returns the register of the message
func (message *Format03) GetRegister() bds.Register {
	return adsb.Format03V0OrMore.GetRegister()
}

// ToString returns a basic, but readable, representation of the message
func (message *Format03) ToString() string {
	return bds08ToString(message)
}

// GetAircraftCategory returns the category of the aircraft
func (message *Format03) GetAircraftCategory() fields2.AircraftCategory {
	return message.AircraftCategory
}

// GetAircraftIdentification returns the identification of the aircraft
func (message *Format03) GetAircraftIdentification() fields2.AircraftIdentification {
	return message.AircraftIdentification
}

// ReadFormat03 reads a message at the format Format03
func ReadFormat03(data []byte) (*Format03, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format03V0OrMore.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format03", formatTypeCode)
	}

	return &Format03{
		AircraftCategory:       fields2.ReadAircraftCategorySetB(data),
		AircraftIdentification: fields2.ReadAircraftIdentification(data),
	}, nil
}
