package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds06/fields"
)

//go:generate go run gen/gen_formats_v0.go
//go:generate go run gen/gen_formats_v1.go
//go:generate go run gen/gen_formats_v2.go
//go:generate go run gen/gen_tests_bds_0_6.go
//go:generate go run gen/gen_tests_v0.go
//go:generate go run gen/gen_tests_v1.go
//go:generate go run gen/gen_tests_v2.go

// MessageBDS06 is the basic interface that ADSB messages at the format BDS 0,6 are expected to implement
type MessageBDS06 interface {
	adsb.Message

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

// ReadBDS06 reads a message at the format BDS 0,6. As there is no information in this message for guessing the
// correct ADSB version, the lowest adsbLevel given is used and returned.
//
// Changes between version:
//   - ADSB V0 -> ADSB V1: Add NIC Supplement A bit coming from a previous message type 31
//   - ADSB V1 -> ADSB V2: Precision on the movement values 125, 126 and 127, from simply Reserved to
//     Reserved with details.
//   - ADSB V1 -> ADSB V2: Replace the SingleAntennaFlag bit by the NIC B bit in first byte of data
//
// Params:
//   - adsbLevel: The ADSB level request (not used, but present for coherency)
//   - nicSupplementA: The NIC Supplement-A comes from the Aircraft  Operational  Status - Message Type Format 31 (see
//     C.2.3.10.20). If no previous Type Format 31 message was received before calling this function, a
//     default value of 0 can be used.
//   - nicSupplementC: The NIC Supplement-C comes from the SubtypeSurface Capability Class (CC) Code  Subfield  of  the
//     Aircraft  Operational  Status - Message Type Format 31 (see  C.2.3.10.20). If no previous Type
//     Format 31 message was received before calling this function, a default value of 0 can be used.
//   - data: The data of the message must be 7 bytes
//
// Returns the message read, the given ADSBLevel or an error
func ReadBDS06(
	adsbLevel adsb.ReaderLevel,
	nicSupplementA bool,
	nicSupplementC bool,
	data []byte) (MessageBDS06, adsb.ReaderLevel, error) {

	if len(data) != 7 {
		return nil, adsbLevel, errors.New("the data for BDS message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3

	if formatTypeCode < 5 || formatTypeCode > 8 {
		return nil, adsbLevel, fmt.Errorf("the format type code %v can not be read as a BDS 0,6 format", formatTypeCode)
	}

	switch adsbLevel {
	case adsb.ReaderLevel0Exactly, adsb.ReaderLevel0OrMore:
		switch formatTypeCode {
		case 5:
			message, err := ReadFormat05V0(data)
			return message, adsbLevel, err
		case 6:
			message, err := ReadFormat06V0(data)
			return message, adsbLevel, err
		case 7:
			message, err := ReadFormat07V0(data)
			return message, adsbLevel, err
		case 8:
			message, err := ReadFormat08V0(data)
			return message, adsbLevel, err
		}

	case adsb.ReaderLevel1Exactly, adsb.ReaderLevel1OrMore:
		switch formatTypeCode {
		case 5:
			message, err := ReadFormat05V1(nicSupplementA, data)
			return message, adsbLevel, err
		case 6:
			message, err := ReadFormat06V1(nicSupplementA, data)
			return message, adsbLevel, err
		case 7:
			message, err := ReadFormat07V1(nicSupplementA, data)
			return message, adsbLevel, err
		case 8:
			message, err := ReadFormat08V1(nicSupplementA, data)
			return message, adsbLevel, err
		}

	case adsb.ReaderLevel2:
		switch formatTypeCode {
		case 5:
			message, err := ReadFormat05V2(nicSupplementA, nicSupplementC, data)
			return message, adsbLevel, err
		case 6:
			message, err := ReadFormat06V2(nicSupplementA, nicSupplementC, data)
			return message, adsbLevel, err
		case 7:
			message, err := ReadFormat07V2(nicSupplementA, nicSupplementC, data)
			return message, adsbLevel, err
		case 8:
			message, err := ReadFormat08V2(nicSupplementA, nicSupplementC, data)
			return message, adsbLevel, err
		}
	}

	return nil, adsbLevel, fmt.Errorf("the format type code %v can not be read as a BDS 0,6 format", formatTypeCode)
}
