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
//
// Params:
//   - adsbLevel: The ADSB level request (currently unused)
//   - data: The data of the message must be 7 bytes
//
// Returns the message read or an error
func ReadBDS65(_ adsb.ReaderLevel, data []byte) (adsb.Message, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for BDS 6,5 message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	subType := data[0] & 0x07

	if formatTypeCode != 31 {
		return nil, fmt.Errorf("the Format Type %v of BSD 6,5 is not supported", formatTypeCode)
	}

	// Read the detectedADSBLevel of ADSB and the subtype
	detectedADSBLevel := fields.ReadVersionNumber(data)

	switch detectedADSBLevel {

	case fields.ADSBVersion0:
		return ReadFormat31Reserved(data)

	case fields.ADSBVersion1:
		switch subType {
		case 0:
			return ReadFormat31AirborneV1(data)
		case 1:
			return ReadFormat31SurfaceV1(data)
		default:
			return nil, fmt.Errorf("the subtype %v of BSD 6,5 is not supported", subType)
		}

	case fields.ADSBVersion2:
		switch subType {
		case 0:
			return ReadFormat31AirborneV2(data)
		case 1:
			return ReadFormat31SurfaceV2(data)
		default:
			return nil, fmt.Errorf("the subtype %v of BSD 6,5 is not supported", subType)
		}

	default:
		return nil, fmt.Errorf("the BDS 6,5 format is not readable with the given reader level")
	}
}
