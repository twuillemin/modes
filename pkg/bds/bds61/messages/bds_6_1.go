package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds61/fields"
	"github.com/twuillemin/modes/pkg/bds/common"
)

// MessageBDS61 is the basic interface that ADSB messages at the format BDS 6,1 are expected to implement
type MessageBDS61 interface {
	common.BDSMessage
	// GetFormatTypeCode returns the Format Type Code
	GetFormatTypeCode() byte
	// GetSubtype returns the Subtype
	GetSubtype() fields.Subtype
}

var bds61Code = "BDS 6,1"
var bds61Name = "Extended squitter emergency/priority status"

// ReadBDS61 reads a message at the format BDS 6,1. This format was extended from ADSB V1 to ADSB V2
//
// Changes between version:
//    - ADSB V0 -> ADSB V1: Add Subtype 2: Add ACAS Resolution Advisory Message
//    - ADSB V1 -> ADSB V2: Change Subtype 1: Add Mode A Code
//
// Params:
//    - adsbLevel: The ADSB level request
//    - data: The data of the message must be 7 bytes
//
// Returns the message read, the given ADSBLevel or an error
func ReadBDS61(adsbLevel common.ADSBLevel, data []byte) (MessageBDS61, common.ADSBLevel, error) {

	if len(data) != 7 {
		return nil, adsbLevel, errors.New("the data for BDS message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	subType := data[0] & 0x07
	possibleModeAData := (data[1]&0x1F) != 0 || data[2] != 0

	if formatTypeCode != 28 {
		return nil, adsbLevel, fmt.Errorf("the format type code %v can not be read as a BDS 6,1 format", formatTypeCode)
	}

	// Default is given code
	adsbLevelToUse := adsbLevel

	switch adsbLevel {

	case common.Level0OrMore:
		switch subType {
		case 1:
			// Mode A Data for subtype 1 are only provided for ADSB Level 2
			if possibleModeAData {
				adsbLevelToUse = common.Level2
			}
		case 2:
			// Subtype 2 only exists from ADSB V1
			adsbLevelToUse = common.Level1OrMore
		}

	case common.Level1OrMore:
		// Mode A Data for subtype 1 are only provided for ADSB Level 2
		if subType == 1 && possibleModeAData {
			adsbLevelToUse = common.Level2
		}
	}

	switch adsbLevelToUse {

	case common.Level0OrMore, common.Level0Exactly:
		if subType == 0 {
			message, err := readFormat28NoInformation(data)
			return message, adsbLevelToUse, err
		} else if subType == 1 {
			message, err := ReadFormat28V0(data)
			return message, adsbLevelToUse, err
		}

	case common.Level1OrMore, common.Level1Exactly:
		if subType == 0 {
			message, err := readFormat28NoInformation(data)
			return message, adsbLevelToUse, err
		} else if subType == 1 {
			message, err := readFormat28V1Status(data)
			return message, adsbLevelToUse, err
		} else if subType == 2 {
			message, err := readFormat28V1ACAS(data)
			return message, adsbLevelToUse, err
		}

	case common.Level2:
		if subType == 0 {
			message, err := readFormat28NoInformation(data)
			return message, adsbLevelToUse, err
		} else if subType == 1 {
			message, err := readFormat28V2Status(data)
			return message, adsbLevelToUse, err
		} else if subType == 2 {
			message, err := readFormat28V2ACAS(data)
			return message, adsbLevelToUse, err
		}
	}

	return nil, adsbLevelToUse, fmt.Errorf("the subtype %v of Emergency/Priority Status is not supported", formatTypeCode)

}
