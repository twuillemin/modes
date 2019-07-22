package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	fields2 "github.com/twuillemin/modes/pkg/bds/bds08/fields"
)

// Format02 is a message at the format BDS 0,8
type Format02 struct {
	AircraftCategory       fields2.AircraftCategorySetC
	AircraftIdentification fields2.AircraftIdentification
}

// GetMessageFormat returns the ADSB format of the message
func (message *Format02) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format02V0OrMore
}

// GetRegister returns the register of the message
func (message *Format02) GetRegister() bds.Register {
	return adsb.Format02V0OrMore.GetRegister()
}

// ToString returns a basic, but readable, representation of the message
func (message *Format02) ToString() string {
	return bds08ToString(message)
}

// GetAircraftCategory returns the category of the aircraft
func (message *Format02) GetAircraftCategory() fields2.AircraftCategory {
	return message.AircraftCategory
}

// GetAircraftIdentification returns the identification of the aircraft
func (message *Format02) GetAircraftIdentification() fields2.AircraftIdentification {
	return message.AircraftIdentification
}

// readFormat02 reads a message at the format BDS 0,8
func readFormat02(data []byte) (*Format02, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format02V0OrMore.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format02", formatTypeCode)
	}

	return &Format02{
		AircraftCategory:       fields2.ReadAircraftCategorySetC(data),
		AircraftIdentification: fields2.ReadAircraftIdentification(data),
	}, nil
}
