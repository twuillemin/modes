package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
)

// ReadBDS61 reads a message at the format BDS 6,1. This format was extended from ADSB V1 to ADSB V2
//
// Changes between version:
//   - ADSB V0 -> ADSB V1: Add Subtype 2: Add ACAS Resolution Advisory Message
//   - ADSB V1 -> ADSB V2: Change Subtype 1: Add Mode A Code
//
// Params:
//   - adsbLevel: The ADSB level request
//   - data: The data of the message must be 7 bytes
//
// Returns the message read or an error
func ReadBDS61(adsbLevel adsb.ReaderLevel, data []byte) (adsb.Message, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for BDS 6,1 message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	subType := data[0] & 0x07

	if formatTypeCode != 28 {
		return nil, fmt.Errorf("the Format Type %v of BSD 6,1 is not supported", formatTypeCode)
	}

	switch adsbLevel {

	case adsb.ReaderLevel0:
		switch subType {
		case 0:
			return ReadFormat28NoInformation(data)
		case 1:
			return ReadFormat28StatusV0(data)
		default:
			return nil, fmt.Errorf("the subtype %v of BSD 6,1 is not supported", subType)
		}

	case adsb.ReaderLevel1:
		switch subType {
		case 0:
			return ReadFormat28NoInformation(data)
		case 1:
			return ReadFormat28StatusV1(data)
		case 2:
			return ReadFormat28ACAS(data)
		default:
			return nil, fmt.Errorf("the subtype %v of BSD 6,1 is not supported", subType)
		}

	case adsb.ReaderLevel2:
		switch subType {
		case 0:
			return ReadFormat28NoInformation(data)
		case 1:
			return ReadFormat28StatusV2(data)
		case 2:
			return ReadFormat28ACAS(data)
		default:
			return nil, fmt.Errorf("the subtype %v of BSD 6,1 is not supported", subType)
		}

	default:
		return nil, fmt.Errorf("the BDS 6,1 format is not readable with the given reader level")
	}
}
