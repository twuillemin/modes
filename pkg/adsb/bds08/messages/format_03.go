package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/bds08/fields"
	"github.com/twuillemin/modes/pkg/bds"
)

// Format03 is a message at the format BDS 0,8
type Format03 struct {
	AircraftCategory       fields.AircraftCategorySetB
	AircraftIdentification fields.AircraftIdentification
}

// GetMessageFormat returns the ADSB format of the message
func (message Format03) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format03
}

// GetRegister returns the register of the message
func (message Format03) GetRegister() bds.Register {
	return adsb.Format03.GetRegister()
}

// GetSubtype returns the subtype of the message if any
func (message Format03) GetSubtype() adsb.Subtype {
	return nil
}

// GetMinimumADSBLevel returns the minimum ADSB ReaderLevel for the message
func (message Format03) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel0
}

// GetMaximumADSBLevel returns the maximum ADSB ReaderLevel for the message
func (message Format03) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel2
}

// GetAircraftCategory returns the category of the aircraft
func (message Format03) GetAircraftCategory() fields.AircraftCategory {
	return message.AircraftCategory
}

// GetAircraftIdentification returns the identification of the aircraft
func (message Format03) GetAircraftIdentification() fields.AircraftIdentification {
	return message.AircraftIdentification
}

// ToString returns a basic, but readable, representation of the message
func (message Format03) ToString() string {
	return bds08ToString(message)
}

// ReadFormat03 reads a message at the format Format03
func ReadFormat03(data []byte) (*Format03, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format03.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format03", formatTypeCode)
	}

	return &Format03{
		AircraftCategory:       fields.ReadAircraftCategorySetB(data),
		AircraftIdentification: fields.ReadAircraftIdentification(data),
	}, nil
}
