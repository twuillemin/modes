package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds06/fields"
	"github.com/twuillemin/modes/pkg/bds/common"
)

// MessageBDS06 is the basic interface that ADSB messages at the format BDS 0,6 are expected to implement
type MessageBDS06 interface {
	common.BDSMessage
	// GetFormatTypeCode returns the Format Type Code
	GetFormatTypeCode() byte
	// GetMovement returns the Movement
	GetMovement() fields.Movement
	// GetGroundTrackStatus returns the GroundTrackStatus
	GetGroundTrackStatus() fields.GroundTrackStatus
	// GetGroundTrack returns the GroundTrack
	GetGroundTrack() fields.GroundTrack
	// GetTime returns the Time
	GetTime() fields.Time
	// GetCPRFormat returns the CompactPositionReportingFormat
	GetCPRFormat() fields.CompactPositionReportingFormat
	// GetEncodedLatitude returns the EncodedLatitude
	GetEncodedLatitude() fields.EncodedLatitude
	// GetEncodedLongitude returns the EncodedLongitude
	GetEncodedLongitude() fields.EncodedLongitude
}

var bds06Code = "BDS 0,6"
var bds06Name = "Extended squitter surface position"

// ReadBDS06 reads a message at the format BDS 0,6. As there is no information in this message for guessing the
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
//    - nicSupplementA: The NIC Supplement-A comes from the Aircraft  Operational  Status - Message Type Format 31 (see
//                      C.2.3.10.20). If no previous Type Format 31 message was received before calling this function, a
//                      default value of 0 can be used.
//    - nicSupplementC: The NIC Supplement-C comes from the Surface Capability Class (CC) Code  Subfield  of  the
//                      Aircraft  Operational  Status - Message Type Format 31 (see  C.2.3.10.20). If no previous Type
//                      Format 31 message was received before calling this function, a default value of 0 can be used.
//    - data: The data of the message must be 7 bytes
//
// Returns the message read, the given ADSBLevel or an error
func ReadBDS06(
	adsbLevel common.ADSBLevel,
	nicSupplementA bool,
	nicSupplementC bool,
	data []byte) (MessageBDS06, common.ADSBLevel, error) {

	if len(data) != 7 {
		return nil, adsbLevel, errors.New("the data for BDS message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3

	if formatTypeCode < 5 || formatTypeCode > 8 {
		return nil, adsbLevel, fmt.Errorf("the format type code %v can not be read as a BDS 0,6 format", formatTypeCode)
	}

	switch adsbLevel {
	case common.Level0Exactly, common.Level0OrMore:
		message, err := readFormat05To08V0(data)
		return message, adsbLevel, err

	case common.Level1Exactly, common.Level1OrMore:
		message, err := readFormat05To08V1(nicSupplementA, data)
		return message, adsbLevel, err
	// case common.Level2Exactly, common.Level2OrMore:
	default:
		message, err := readFormat05To08V2(nicSupplementA, nicSupplementC, data)
		return message, adsbLevel, err
	}
}
