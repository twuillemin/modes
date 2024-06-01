package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
)

// ReadBDS62 reads a message at the format BDS 6,2. This format was created in ADSB V1, and extended in ADSB V2
//
// Changes between version:
//   - ADSB V0 -> ADSB V1: Create message with Subtype 0
//   - ADSB V1 -> ADSB V2: Add Subtype 1
//
// Params:
//   - adsbLevel: The ADSB level request
//   - data: The data of the message must be 7 bytes
//
// Returns the message read or an error
func ReadBDS62(adsbLevel adsb.ReaderLevel, data []byte) (adsb.Message, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for BDS 6,2 message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	subType := (data[0] & 0x06) >> 1

	if formatTypeCode != 29 {
		return nil, fmt.Errorf("the Format Type %v of BSD 6,2 is not supported", formatTypeCode)
	}

	switch adsbLevel {

	case adsb.ReaderLevel1:
		switch subType {
		case 0:
			return ReadFormat29Subtype0(data)
		default:
			return nil, fmt.Errorf("the subtype %v of BSD 6,2 is not supported", subType)
		}

	case adsb.ReaderLevel2:
		switch subType {
		case 0:
			return ReadFormat29Subtype0(data)
		case 1:
			return ReadFormat29Subtype1(data)
		default:
			return nil, fmt.Errorf("the subtype %v of BSD 6,2 is not supported", subType)
		}

	default:
		return nil, fmt.Errorf("the BDS 6,2 format is not readable with the given reader level")
	}
}
