package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb/messages"
	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
)

// MessageBDS65 is the basic interface that ADSB messages at the format BDS 6,5 are expected to implement
type MessageBDS65 interface {
	messages.ADSBMessage
	// GetOperationalStatusSubtypeCode returns the code of the Operational Status Sub Type
	GetOperationalStatusSubtypeCode() byte
}

var bds65Code = "BDS 6,5"
var bds65Name = "Extended squitter aircraft operational status"

// ReadBDS65 reads a message at the format BDS 6,5
func ReadBDS65(data []byte) (MessageBDS65, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for BDS message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3

	if formatTypeCode == 31 {
		return nil, fmt.Errorf("the format type code %v can not be read as a BDS 6,5 format", formatTypeCode)
	}

	// Read the version of ADSB and the subtype
	version := fields.ReadVersionNumber(data)
	subType := fields.ReadOperationalStatusSubtypeCode(data)

	switch version {

	case fields.ADSBVersion0:
		return ReadFormat31V0(data)

	case fields.ADSBVersion1, fields.ADSBVersion2:
		switch subType {
		case fields.OSSCAirborne:
			return ReadFormat31V1Airborne(data)
		case fields.OSSCSurface:
			return ReadFormat31V1Surface(data)
		default:
			return nil, fmt.Errorf("the subtype %v of Aircraft Operational Status is not supported", formatTypeCode)
		}

	default:
		return nil, fmt.Errorf("the version of ADSB %v is not supported", formatTypeCode)
	}
}
