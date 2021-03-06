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
func (message Format02) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format02
}

// GetRegister returns the register of the message
func (message Format02) GetRegister() bds.Register {
	return adsb.Format02.GetRegister()
}

// GetSubtype returns the subtype of the message if any
func (message Format02) GetSubtype() adsb.Subtype {
	return nil
}

// GetMinimumADSBLevel returns the minimum ADSB ReaderLevel for the message
func (message Format02) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel0
}

// GetMaximumADSBLevel returns the maximum ADSB ReaderLevel for the message
func (message Format02) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel2
}

// GetAircraftCategory returns the category of the aircraft
func (message Format02) GetAircraftCategory() fields2.AircraftCategory {
	return message.AircraftCategory
}

// GetAircraftIdentification returns the identification of the aircraft
func (message Format02) GetAircraftIdentification() fields2.AircraftIdentification {
	return message.AircraftIdentification
}

// ToString returns a basic, but readable, representation of the message
func (message Format02) ToString() string {
	return bds08ToString(message)
}

// ReadFormat02 reads a message at the format Format02
func ReadFormat02(data []byte) (*Format02, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format02.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format02", formatTypeCode)
	}

	return &Format02{
		AircraftCategory:       fields2.ReadAircraftCategorySetC(data),
		AircraftIdentification: fields2.ReadAircraftIdentification(data),
	}, nil
}
