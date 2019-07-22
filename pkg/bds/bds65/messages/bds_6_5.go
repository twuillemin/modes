package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
)

// MessageBDS65 is the basic interface that ADSB messages at the format BDS 6,5 are expected to implement
type MessageBDS65 interface {
	adsb.Message

	// GetSubtype returns the subtype of the Operational Status Sub Type
	GetSubtype() fields.Subtype
}

// ReadBDS65 reads a message at the format BDS 6,5
func ReadBDS65(adsbLevel adsb.Level, data []byte) (MessageBDS65, adsb.Level, error) {

	if len(data) != 7 {
		return nil, adsbLevel, errors.New("the data for BDS message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	subType := data[0] & 0x07

	if formatTypeCode != 31 {
		return nil, adsbLevel, fmt.Errorf("the format type code %v can not be read as a BDS 6,5 format", formatTypeCode)
	}

	// Read the detectedADSBLevel of ADSB and the subtype
	detectedADSBLevel := fields.ReadVersionNumber(data)

	// If the version is fixed but different from the read one, return an error
	if (detectedADSBLevel == fields.ADSBVersion0 && (adsbLevel == adsb.Level1Exactly || adsbLevel == adsb.Level2)) ||
		(detectedADSBLevel == fields.ADSBVersion1 && (adsbLevel == adsb.Level0Exactly || adsbLevel == adsb.Level2)) ||
		(detectedADSBLevel == fields.ADSBVersion2 && (adsbLevel == adsb.Level0Exactly || adsbLevel == adsb.Level1Exactly)) {

		return nil, adsbLevel, fmt.Errorf("the request ADSB level (%v) is not coherent with the level detected in the message (%v)", adsbLevel, detectedADSBLevel)
	}

	// If the detected level is lower then the possible level
	if (detectedADSBLevel == fields.ADSBVersion0 && (adsbLevel == adsb.Level1OrMore || adsbLevel == adsb.Level2)) ||
		(detectedADSBLevel == fields.ADSBVersion1 && adsbLevel == adsb.Level2) {

		return nil, adsbLevel, fmt.Errorf("the request ADSB level (%v or more) is not higher with the level detected in the message (%v)", adsbLevel, detectedADSBLevel)
	}

	// As the level is not supposed to change, use Exact version
	resultingADSBLevel := adsb.Level0Exactly
	if detectedADSBLevel == fields.ADSBVersion1 {
		resultingADSBLevel = adsb.Level1Exactly
	} else if detectedADSBLevel == fields.ADSBVersion2 {
		resultingADSBLevel = adsb.Level2
	}

	switch resultingADSBLevel {

	case adsb.Level0Exactly:
		message, err := ReadFormat31V0(data)
		return message, resultingADSBLevel, err

	case adsb.Level1Exactly:
		switch subType {
		case 0:
			message, err := ReadFormat31V1Airborne(data)
			return message, resultingADSBLevel, err
		case 1:
			message, err := ReadFormat31V1Surface(data)
			return message, resultingADSBLevel, err
		default:
			return nil, resultingADSBLevel, fmt.Errorf("the subtype %v of Aircraft Operational Status is not supported", formatTypeCode)
		}

	case adsb.Level2:
		switch subType {
		case 0:
			message, err := ReadFormat31V2Airborne(data)
			return message, resultingADSBLevel, err
		case 1:
			message, err := ReadFormat31V2Surface(data)
			return message, resultingADSBLevel, err
		default:
			return nil, resultingADSBLevel, fmt.Errorf("the subtype %v of Aircraft Operational Status is not supported", formatTypeCode)
		}

	default:
		return nil, resultingADSBLevel, fmt.Errorf("the detectedADSBLevel of ADSB %v is not supported", formatTypeCode)
	}
}
