package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
	"github.com/twuillemin/modes/pkg/bds/common"
)

// MessageBDS05 is the basic interface that ADSB messages at the format BDS 0,5 are expected to implement
type MessageBDS05 interface {
	common.BDSMessage
	// GetFormatTypeCode returns the Format Type Code
	GetFormatTypeCode() byte
	// GetSurveillanceStatus returns the Surveillance Status
	GetSurveillanceStatus() fields.SurveillanceStatus
	// GetAltitude returns the Altitude
	GetAltitude() fields.Altitude
	// GetTime returns the Time
	GetTime() fields.Time
	// GetCPRFormat returns the CompactPositionReportingFormat
	GetCPRFormat() fields.CompactPositionReportingFormat
	// GetEncodedLatitude returns the EncodedLatitude
	GetEncodedLatitude() fields.EncodedLatitude
	// GetEncodedLongitude returns the EncodedLongitude
	GetEncodedLongitude() fields.EncodedLongitude
}

var bds05Code = "BDS 0,5"
var bds05Name = "Extended squitter airborne position"

// ReadBDS05 reads a message at the format BDS 0,5. As there is no information in this message for guessing the
// correct ADSB version, the lowest adsbLevel given is used and returned.
//
// Changes between version:
//    - ADSB V0 -> ADSB V1: Add NIC Supplement A bit coming from a previous message type 31
//    - ADSB V1 -> ADSB V2: Precision on the movement values 125, 126 and 127, from simply Reserved to
//                          Reserved with details.
//    - ADSB V1 -> ADSB V2: Add the replace the SingleAntennaFlag bit by the NIC B bit in first byte of data
//
// Params:
//    - adsbLevel: The ADSB level request (not used, but present for coherency)
//    - data: The data of the message must be 7 bytes
//    - nicSupplementA: The nic supplement A bit coming from previous Format Code 31 message if any. If none, 0 is fine
//
// Returns the message read, the given ADSBLevel or an error
func ReadBDS05(adsbLevel common.ADSBLevel, nicSupplementA bool, data []byte) (MessageBDS05, common.ADSBLevel, error) {

	if len(data) != 7 {
		return nil, adsbLevel, errors.New("the data for BDS message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3

	switch adsbLevel {

	case common.Level0Exactly, common.Level0OrMore:

		if 9 <= formatTypeCode && formatTypeCode <= 18 {
			message, err := readFormat09To18V0(data)
			return message, adsbLevel, err
		} else if 20 <= formatTypeCode && formatTypeCode <= 22 {
			message, err := readFormat20To22V0(data)
			return message, adsbLevel, err
		}

	case common.Level1Exactly, common.Level1OrMore:

		if 9 <= formatTypeCode && formatTypeCode <= 18 {
			message, err := readFormat09To18V1(nicSupplementA, data)
			return message, adsbLevel, err
		} else if 20 <= formatTypeCode && formatTypeCode <= 22 {
			message, err := readFormat20To22V1(data)
			return message, adsbLevel, err
		}

	case common.Level2:

		if 9 <= formatTypeCode && formatTypeCode <= 18 {
			message, err := readFormat09To18V2(nicSupplementA, data)
			return message, adsbLevel, err
		} else if 20 <= formatTypeCode && formatTypeCode <= 22 {
			message, err := readFormat20To22V2(data)
			return message, adsbLevel, err
		}
	}

	return nil, adsbLevel, fmt.Errorf("the format type code %v can not be read as a BDS 0,5 format", formatTypeCode)
}
