package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds62/fields"
	"github.com/twuillemin/modes/pkg/bds/common"
)

// MessageBDS62 is the basic interface that ADSB messages at the format BDS 6,2 are expected to implement
type MessageBDS62 interface {
	common.BDSMessage
	// GetFormatTypeCode returns the Format Type Code
	GetFormatTypeCode() byte
	// GetSubtype returns the Subtype
	GetSubtype() fields.Subtype
}

var bds62Code = "BDS 6,2"
var bds62Name = "Target state and status information"

// ReadBDS62 reads a message at the format BDS 6,2. This format was created in ADSB V1, and extended in ADSB V2
//
// Changes between version:
//    - ADSB V0 -> ADSB V1: Create message with Subtype 0
//    - ADSB V1 -> ADSB V2: Add Subtype 1
//
// Params:
//    - adsbLevel: The ADSB level request
//    - data: The data of the message must be 7 bytes
//
// Returns the message read, the given ADSBLevel or an error
func ReadBDS62(adsbLevel common.ADSBLevel, data []byte) (MessageBDS62, common.ADSBLevel, error) {

	if len(data) != 7 {
		return nil, adsbLevel, errors.New("the data for BDS message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	subType := data[0] & 0x07

	if formatTypeCode != 29 {
		return nil, adsbLevel, fmt.Errorf("the format type code %v can not be read as a BDS 6,2 format", formatTypeCode)
	}

	// Default is given code
	adsbLevelToUse := adsbLevel

	switch adsbLevel {

	case common.Level0OrMore:
		switch subType {
		case 0:
			// Subtype 0 only exists from ADSB V1
			adsbLevelToUse = common.Level1OrMore
		case 1:
			// Subtype 1 only exists from ADSB V2
			adsbLevelToUse = common.Level2
		}

	case common.Level1OrMore:
		// Mode A Data for subtype 1 are only provided for ADSB Level 2
		if subType == 1 {
			adsbLevelToUse = common.Level2
		}
	}

	switch adsbLevelToUse {

	case common.Level1OrMore, common.Level1Exactly:
		if subType == 0 {
			message, err := readFormat29Subtype0(data)
			return message, adsbLevelToUse, err
		}

	case common.Level2:
		if subType == 0 {
			message, err := readFormat29Subtype0(data)
			return message, adsbLevelToUse, err
		} else if subType == 1 {
			// TODO implement for subtype 2
			return nil, adsbLevelToUse, fmt.Errorf("the subtype %v of Target state and status information is not implement", formatTypeCode)
		}
	}

	return nil, adsbLevelToUse, fmt.Errorf("the subtype %v of Target state and status information is not supported", subType)
}
