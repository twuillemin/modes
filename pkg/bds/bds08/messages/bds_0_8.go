package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds08/fields"
	"github.com/twuillemin/modes/pkg/bds/common"
)

// MessageBDS08 is the basic interface that ADSB messages at the format BDS 0,8 are expected to implement
type MessageBDS08 interface {
	common.BDSMessage
	// GetFormatTypeCode returns the Format Type Code
	GetFormatTypeCode() byte
	// GetAircraftCategory returns the aircraft category
	GetAircraftCategory() fields.AircraftCategory
	// GetAircraftIdentification returns the identity of the aircraft
	GetAircraftIdentification() fields.AircraftIdentification
}

var bds08Code = "BDS 0,8"
var bds08Name = "Extended squitter aircraft identification and category"

func bds08ToString(message MessageBDS08) string {
	return fmt.Sprintf("Message:                 %v (%v)\n"+
		"Format Type Code:        %v\n"+
		"Aircraft Category:       %v (%v)\n"+
		"Aircraft Identification: %v",
		message.GetBDS(),
		message.GetName(),
		message.GetFormatTypeCode(),
		message.GetAircraftCategory().ToString(),
		message.GetAircraftCategory().GetCategorySetName(),
		message.GetAircraftIdentification())
}

// ReadBDS08 reads a message at the format BDS 0,8
func ReadBDS08(data []byte) (MessageBDS08, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for BDS message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3

	switch formatTypeCode {
	case 1:
		return ReadFormat01(data)
	case 2:
		return ReadFormat02(data)
	case 3:
		return ReadFormat03(data)
	case 4:
		return ReadFormat04(data)
	}

	return nil, fmt.Errorf("the format type code %v can not be read as a BDS 0,8 format", formatTypeCode)
}
