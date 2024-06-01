package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
)

// ReadBDS65 reads a message at the format BDS 6,5
func ReadBDS65(adsbLevel adsb.ReaderLevel, data []byte) (adsb.Message, adsb.ReaderLevel, error) {

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
	if (detectedADSBLevel == fields.ADSBVersion0 && (adsbLevel == adsb.ReaderLevel1Exactly || adsbLevel == adsb.ReaderLevel2)) ||
		(detectedADSBLevel == fields.ADSBVersion1 && (adsbLevel == adsb.ReaderLevel0Exactly || adsbLevel == adsb.ReaderLevel2)) ||
		(detectedADSBLevel == fields.ADSBVersion2 && (adsbLevel == adsb.ReaderLevel0Exactly || adsbLevel == adsb.ReaderLevel1Exactly)) {

		return nil, adsbLevel, fmt.Errorf("the request ADSB level (%v) is not coherent with the level detected in the message (%v)", adsbLevel, detectedADSBLevel)
	}

	// If the detected level is lower than the possible level
	if (detectedADSBLevel == fields.ADSBVersion0 && (adsbLevel == adsb.ReaderLevel1OrMore || adsbLevel == adsb.ReaderLevel2)) ||
		(detectedADSBLevel == fields.ADSBVersion1 && adsbLevel == adsb.ReaderLevel2) {

		return nil, adsbLevel, fmt.Errorf("the request ADSB level (%v or more) is not higher with the level detected in the message (%v)", adsbLevel, detectedADSBLevel)
	}

	// As the level is not supposed to change, use Exact version
	resultingADSBLevel := adsb.ReaderLevel0Exactly
	if detectedADSBLevel == fields.ADSBVersion1 {
		resultingADSBLevel = adsb.ReaderLevel1Exactly
	} else if detectedADSBLevel == fields.ADSBVersion2 {
		resultingADSBLevel = adsb.ReaderLevel2
	}

	switch resultingADSBLevel {

	case adsb.ReaderLevel0Exactly:
		message, err := ReadFormat31Reserved(data)
		return message, resultingADSBLevel, err

	case adsb.ReaderLevel1Exactly:
		switch subType {
		case 0:
			message, err := ReadFormat31AirborneV1(data)
			return message, resultingADSBLevel, err
		case 1:
			message, err := ReadFormat31SurfaceV1(data)
			return message, resultingADSBLevel, err
		default:
			return nil, resultingADSBLevel, fmt.Errorf("the subtype %v of Aircraft Operational Status is not supported", formatTypeCode)
		}

	case adsb.ReaderLevel2:
		switch subType {
		case 0:
			message, err := ReadFormat31AirborneV2(data)
			return message, resultingADSBLevel, err
		case 1:
			message, err := ReadFormat31SurfaceV2(data)
			return message, resultingADSBLevel, err
		default:
			return nil, resultingADSBLevel, fmt.Errorf("the subtype %v of Aircraft Operational Status is not supported", formatTypeCode)
		}

	default:
		return nil, resultingADSBLevel, fmt.Errorf("the detectedADSBLevel of ADSB %v is not supported", formatTypeCode)
	}
}
