package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
)

// ReadBDS65 reads a message at the format BDS 6,5
//
// Contrary to the other readers which rely on the given ReaderLevel to be decoded, the BDS 6,5 includes the ADSB
// level. So, the ReaderLevel is ignored.
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

	switch detectedADSBLevel {

	case fields.ADSBVersion0:
		message, err := ReadFormat31Reserved(data)
		return message, adsb.ReaderLevel0, err

	case fields.ADSBVersion1:
		switch subType {
		case 0:
			message, err := ReadFormat31AirborneV1(data)
			return message, adsb.ReaderLevel1, err
		case 1:
			message, err := ReadFormat31SurfaceV1(data)
			return message, adsb.ReaderLevel1, err
		default:
			return nil, adsb.ReaderLevel1, fmt.Errorf("the subtype %v of Aircraft Operational Status is not supported", subType)
		}

	case fields.ADSBVersion2:
		switch subType {
		case 0:
			message, err := ReadFormat31AirborneV2(data)
			return message, adsb.ReaderLevel2, err
		case 1:
			message, err := ReadFormat31SurfaceV2(data)
			return message, adsb.ReaderLevel2, err
		default:
			return nil, adsb.ReaderLevel2, fmt.Errorf("the subtype %v of Aircraft Operational Status is not supported", subType)
		}

	default:
		return nil, adsbLevel, fmt.Errorf("the detectedADSBLevel of ADSB %v is not supported", detectedADSBLevel)
	}
}
