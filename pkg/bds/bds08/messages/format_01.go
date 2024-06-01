package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds08/fields"
)

// Format01 is a message at the format BDS 0,8
type Format01 struct {
	AircraftCategory       fields.AircraftCategorySetD
	AircraftIdentification fields.AircraftIdentification
}

// GetMessageFormat returns the ADSB format of the message
func (message Format01) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format01
}

// GetRegister returns the register of the message
func (message Format01) GetRegister() bds.Register {
	return adsb.Format01.GetRegister()
}

// GetSubtype returns the subtype of the message if any
func (message Format01) GetSubtype() adsb.Subtype {
	return nil
}

// GetMinimumADSBLevel returns the minimum ADSB ReaderLevel for the message
func (message Format01) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel0
}

// GetMaximumADSBLevel returns the maximum ADSB ReaderLevel for the message
func (message Format01) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel2
}

// GetAircraftCategory returns the category of the aircraft
func (message Format01) GetAircraftCategory() fields.AircraftCategory {
	return message.AircraftCategory
}

// GetAircraftIdentification returns the identification of the aircraft
func (message Format01) GetAircraftIdentification() fields.AircraftIdentification {
	return message.AircraftIdentification
}

// ToString returns a basic, but readable, representation of the message
func (message Format01) ToString() string {
	return bds08ToString(message)
}

// ReadFormat01 reads a message at the format Format01
func ReadFormat01(data []byte) (*Format01, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format01.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format01", formatTypeCode)
	}

	return &Format01{
		AircraftCategory:       fields.ReadAircraftCategorySetD(data),
		AircraftIdentification: fields.ReadAircraftIdentification(data),
	}, nil
}
