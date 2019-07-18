package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb/fields"
	"github.com/twuillemin/modes/pkg/adsb/messages"
)

// MessageBDS65 is the basic interface that ADSB messages at the format BDS 6,5 are expected to implement
type MessageBDS65 interface {
	messages.ADSBMessage
	// GetOperationalStatusSubTypeCode returns the code of the Operational Status Sub Type
	GetOperationalStatusSubTypeCode() byte
}

var bds65Code = "BDS 6,5"
var bds65Name = "Extended squitter aircraft operational status"

// ReadBDS65 reads a message at the format BDS 6,5
func ReadBDS65(data []byte) (MessageBDS65, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3

	if formatTypeCode == 31 {

		// Read the version of ADSB and the subtype
		version := fields.ReadVersionNumber(data)
		subType := fields.ReadSubTypeCode(data)

		switch version {

		case fields.ADSBV0:
			return messages.ReadFormat31V0(data)

		case fields.ADSBV1, fields.ADSBV2:
			switch subType {
			case fields.STCAirborne:
				return messages.ReadFormat31V1Airborne(data)
			case fields.STCSurface:
				return messages.ReadFormat31V1Surface(data)
			default:
				return nil, fmt.Errorf("the subtype %v of Aircraft Operational Status is not supported", formatTypeCode)
			}

		default:
			return nil, fmt.Errorf("the version of ADSB %v is not supported", formatTypeCode)
		}
	}

	return nil, fmt.Errorf("the format type code %v can not be read as a BDS 6,5 format", formatTypeCode)
}
