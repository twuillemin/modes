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
// Returns the message read, the given ADSBLevel or an error
func ReadBDS62(adsbLevel adsb.ReaderLevel, data []byte) (adsb.Message, adsb.ReaderLevel, error) {

	if len(data) != 7 {
		return nil, adsbLevel, errors.New("the data for BDS message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	subType := (data[0] & 0x06) >> 1

	if formatTypeCode != 29 {
		return nil, adsbLevel, fmt.Errorf("the format type code %v can not be read as a BDS 6,2 format", formatTypeCode)
	}

	switch adsbLevel {

	case adsb.ReaderLevel0:
		return nil, adsbLevel, fmt.Errorf("the BDS 6,2 format is not readable as an ADSB V0 message")

	case adsb.ReaderLevel1:
		if subType == 0 {
			message, err := ReadFormat29Subtype0(data)
			return message, adsbLevel, err
		} else {
			return nil, adsbLevel, fmt.Errorf("the subtype %v of Target state and status information is not supported", subType)
		}

	case adsb.ReaderLevel2:
		if subType == 0 {
			message, err := ReadFormat29Subtype0(data)
			return message, adsbLevel, err
		} else if subType == 1 {
			message, err := ReadFormat29Subtype1(data)
			return message, adsbLevel, err
		} else {
			return nil, adsbLevel, fmt.Errorf("the format type code %v can not be read as a BDS 6,2 format", formatTypeCode)
		}
	}

	return nil, adsbLevel, fmt.Errorf("the BDS 6,2 format is not readable with the given reader level")
}
