package messages

//go:generate go run gen/gen_formats_v0.go
//go:generate go run gen/gen_formats_v1.go
//go:generate go run gen/gen_formats_v2.go

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
)

// MessageBDS05 is the basic interface that ADSB messages at the format BDS 0,5 are expected to implement
type MessageBDS05 interface {
	adsb.Message

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
func ReadBDS05(adsbLevel adsb.Level, nicSupplementA bool, data []byte) (MessageBDS05, adsb.Level, error) {

	if len(data) != 7 {
		return nil, adsbLevel, errors.New("the data for BDS message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3

	switch adsbLevel {

	case adsb.Level0Exactly, adsb.Level0OrMore:
		switch formatTypeCode {
		case 9:
			message, err := readFormat09V0(data)
			return message, adsbLevel, err
		case 10:
			message, err := readFormat10V0(data)
			return message, adsbLevel, err
		case 11:
			message, err := readFormat11V0(data)
			return message, adsbLevel, err
		case 12:
			message, err := readFormat12V0(data)
			return message, adsbLevel, err
		case 13:
			message, err := readFormat13V0(data)
			return message, adsbLevel, err
		case 14:
			message, err := readFormat14V0(data)
			return message, adsbLevel, err
		case 15:
			message, err := readFormat15V0(data)
			return message, adsbLevel, err
		case 16:
			message, err := readFormat16V0(data)
			return message, adsbLevel, err
		case 17:
			message, err := readFormat17V0(data)
			return message, adsbLevel, err
		case 18:
			message, err := readFormat18V0(data)
			return message, adsbLevel, err
		case 20:
			message, err := readFormat20V0(data)
			return message, adsbLevel, err
		case 21:
			message, err := readFormat21V0(data)
			return message, adsbLevel, err
		case 22:
			message, err := readFormat22V0(data)
			return message, adsbLevel, err
		}

	case adsb.Level1Exactly, adsb.Level1OrMore:
		switch formatTypeCode {
		case 9:
			message, err := readFormat09V1(nicSupplementA, data)
			return message, adsbLevel, err
		case 10:
			message, err := readFormat10V1(nicSupplementA, data)
			return message, adsbLevel, err
		case 11:
			message, err := readFormat11V1(nicSupplementA, data)
			return message, adsbLevel, err
		case 12:
			message, err := readFormat12V1(nicSupplementA, data)
			return message, adsbLevel, err
		case 13:
			message, err := readFormat13V1(nicSupplementA, data)
			return message, adsbLevel, err
		case 14:
			message, err := readFormat14V1(nicSupplementA, data)
			return message, adsbLevel, err
		case 15:
			message, err := readFormat15V1(nicSupplementA, data)
			return message, adsbLevel, err
		case 16:
			message, err := readFormat16V1(nicSupplementA, data)
			return message, adsbLevel, err
		case 17:
			message, err := readFormat17V1(nicSupplementA, data)
			return message, adsbLevel, err
		case 18:
			message, err := readFormat18V1(nicSupplementA, data)
			return message, adsbLevel, err
		case 20:
			message, err := readFormat20V1(data)
			return message, adsbLevel, err
		case 21:
			message, err := readFormat21V1(data)
			return message, adsbLevel, err
		case 22:
			message, err := readFormat22V1(data)
			return message, adsbLevel, err
		}

	case adsb.Level2:
		switch formatTypeCode {
		case 9:
			message, err := readFormat09V2(nicSupplementA, data)
			return message, adsbLevel, err
		case 10:
			message, err := readFormat10V2(nicSupplementA, data)
			return message, adsbLevel, err
		case 11:
			message, err := readFormat11V2(nicSupplementA, data)
			return message, adsbLevel, err
		case 12:
			message, err := readFormat12V2(nicSupplementA, data)
			return message, adsbLevel, err
		case 13:
			message, err := readFormat13V2(nicSupplementA, data)
			return message, adsbLevel, err
		case 14:
			message, err := readFormat14V2(nicSupplementA, data)
			return message, adsbLevel, err
		case 15:
			message, err := readFormat15V2(nicSupplementA, data)
			return message, adsbLevel, err
		case 16:
			message, err := readFormat16V2(nicSupplementA, data)
			return message, adsbLevel, err
		case 17:
			message, err := readFormat17V2(nicSupplementA, data)
			return message, adsbLevel, err
		case 18:
			message, err := readFormat18V2(nicSupplementA, data)
			return message, adsbLevel, err
		case 20:
			message, err := readFormat20V2(data)
			return message, adsbLevel, err
		case 21:
			message, err := readFormat21V2(data)
			return message, adsbLevel, err
		case 22:
			message, err := readFormat22V2(data)
			return message, adsbLevel, err
		}
	}

	return nil, adsbLevel, fmt.Errorf("the format type code %v can not be read as a BDS 0,5 format", formatTypeCode)
}
