package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	fields2 "github.com/twuillemin/modes/pkg/bds/bds08/fields"
)

// Format04 is a message at the format BDS 0,8
type Format04 struct {
	AircraftCategory       fields2.AircraftCategorySetA
	AircraftIdentification fields2.AircraftIdentification
}

// GetMessageFormat returns the ADSB format of the message
func (message Format04) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format04
}

// GetRegister returns the register of the message
func (message Format04) GetRegister() bds.Register {
	return adsb.Format04.GetRegister()
}

// GetSubtype returns the subtype of the message if any
func (message Format04) GetSubtype() adsb.Subtype {
	return nil
}

// GetMinimumADSBLevel returns the minimum ADSB ReaderLevel for the message
func (message Format04) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel0
}

// GetMaximumADSBLevel returns the maximum ADSB ReaderLevel for the message
func (message Format04) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel2
}

// GetAircraftCategory returns the category of the aircraft
func (message Format04) GetAircraftCategory() fields2.AircraftCategory {
	return message.AircraftCategory
}

// GetAircraftIdentification returns the identification of the aircraft
func (message Format04) GetAircraftIdentification() fields2.AircraftIdentification {
	return message.AircraftIdentification
}

// ToString returns a basic, but readable, representation of the message
func (message Format04) ToString() string {
	return bds08ToString(message)
}

// ReadFormat04 reads a message at the format Format04
func ReadFormat04(data []byte) (*Format04, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format04.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format04", formatTypeCode)
	}

	return &Format04{
		AircraftCategory:       fields2.ReadAircraftCategorySetA(data),
		AircraftIdentification: fields2.ReadAircraftIdentification(data),
	}, nil
}
