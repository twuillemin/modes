package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb/messages"
	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
	"github.com/twuillemin/modes/pkg/bds/common"
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
func ReadBDS65(adsbLevel common.ADSBLevel, data []byte) (MessageBDS65, common.ADSBLevel, error) {

	if len(data) != 7 {
		return nil, adsbLevel, errors.New("the data for BDS message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3

	if formatTypeCode != 31 {
		return nil, adsbLevel, fmt.Errorf("the format type code %v can not be read as a BDS 6,5 format", formatTypeCode)
	}

	// Read the detectedADSBLevel of ADSB and the subtype
	detectedADSBLevel := fields.ReadVersionNumber(data)
	subType := fields.ReadOperationalStatusSubtypeCode(data)

	// If the version is fixed but different from the read one, return an error
	if (detectedADSBLevel == fields.ADSBVersion0 && (adsbLevel == common.Level1Exactly || adsbLevel == common.Level2)) ||
		(detectedADSBLevel == fields.ADSBVersion1 && (adsbLevel == common.Level0Exactly || adsbLevel == common.Level2)) ||
		(detectedADSBLevel == fields.ADSBVersion2 && (adsbLevel == common.Level0Exactly || adsbLevel == common.Level1Exactly)) {

		return nil, adsbLevel, fmt.Errorf("the request ADSB level (%v) is not coherent with the level detected in the message (%v)", adsbLevel, detectedADSBLevel)
	}

	// If the detected level is lower then the possible level
	if (detectedADSBLevel == fields.ADSBVersion0 && (adsbLevel == common.Level1OrMore || adsbLevel == common.Level2)) ||
		(detectedADSBLevel == fields.ADSBVersion1 && adsbLevel == common.Level2) {

		return nil, adsbLevel, fmt.Errorf("the request ADSB level (%v or more) is not higher with the level detected in the message (%v)", adsbLevel, detectedADSBLevel)
	}

	// As the level is not supposed to change, use Exact version
	resultingADSBLevel := common.Level0Exactly
	if detectedADSBLevel == fields.ADSBVersion1 {
		resultingADSBLevel = common.Level1Exactly
	} else if detectedADSBLevel == fields.ADSBVersion2 {
		resultingADSBLevel = common.Level2
	}

	switch detectedADSBLevel {

	case fields.ADSBVersion0:
		message, err := ReadFormat31V0(data)
		return message, resultingADSBLevel, err

	case fields.ADSBVersion1, fields.ADSBVersion2:
		switch subType {
		case fields.OSSCAirborne:
			message, err := ReadFormat31V1Airborne(data)
			return message, resultingADSBLevel, err
		case fields.OSSCSurface:
			message, err := ReadFormat31V1Surface(data)
			return message, resultingADSBLevel, err
		default:
			return nil, resultingADSBLevel, fmt.Errorf("the subtype %v of Aircraft Operational Status is not supported", formatTypeCode)
		}

	default:
		return nil, resultingADSBLevel, fmt.Errorf("the detectedADSBLevel of ADSB %v is not supported", formatTypeCode)
	}
}
