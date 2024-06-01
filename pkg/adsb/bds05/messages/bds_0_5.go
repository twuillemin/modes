package messages

//go:generate go run gen/gen_formats_v0.go
//go:generate go run gen/gen_formats_v1.go
//go:generate go run gen/gen_formats_v2.go
//go:generate go run gen/gen_tests_bds_0_5.go
//go:generate go run gen/gen_tests_v0.go
//go:generate go run gen/gen_tests_v1.go
//go:generate go run gen/gen_tests_v2.go

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/bds05/fields"
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

// ReadBDS05 reads a message at the format BDS 0,5.
//
// Changes between version:
//   - ADSB V0 -> ADSB V1: Add NIC Supplement A bit coming from a previous message type 31
//   - ADSB V1 -> ADSB V2: Precision on the movement values 125, 126 and 127, from simply Reserved to
//     Reserved with details.
//   - ADSB V1 -> ADSB V2: Replace the SingleAntennaFlag bit by the NIC B bit in first byte of data
//
// Params:
//   - adsbLevel: The ADSB level request
//   - data: The data of the message must be 7 bytes
//   - nicSupplementA: The nic supplement A bit coming from previous Format Code 31 message if any. If none, 0 is fine
//
// Returns the message read or an error
func ReadBDS05(adsbLevel adsb.ReaderLevel, nicSupplementA bool, data []byte) (MessageBDS05, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for BDS 0,5 message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3

	switch adsbLevel {

	case adsb.ReaderLevel0:
		switch formatTypeCode {
		case 9:
			message, err := ReadFormat09V0(data)
			return message, err
		case 10:
			message, err := ReadFormat10V0(data)
			return message, err
		case 11:
			message, err := ReadFormat11V0(data)
			return message, err
		case 12:
			message, err := ReadFormat12V0(data)
			return message, err
		case 13:
			message, err := ReadFormat13V0(data)
			return message, err
		case 14:
			message, err := ReadFormat14V0(data)
			return message, err
		case 15:
			message, err := ReadFormat15V0(data)
			return message, err
		case 16:
			message, err := ReadFormat16V0(data)
			return message, err
		case 17:
			message, err := ReadFormat17V0(data)
			return message, err
		case 18:
			message, err := ReadFormat18V0(data)
			return message, err
		case 20:
			message, err := ReadFormat20V0(data)
			return message, err
		case 21:
			message, err := ReadFormat21V0(data)
			return message, err
		case 22:
			message, err := ReadFormat22V0(data)
			return message, err
		default:
			return nil, fmt.Errorf("the Format Type %v of BSD 0,5 is not supported", formatTypeCode)
		}

	case adsb.ReaderLevel1:
		switch formatTypeCode {
		case 9:
			message, err := ReadFormat09V1(nicSupplementA, data)
			return message, err
		case 10:
			message, err := ReadFormat10V1(nicSupplementA, data)
			return message, err
		case 11:
			message, err := ReadFormat11V1(nicSupplementA, data)
			return message, err
		case 12:
			message, err := ReadFormat12V1(nicSupplementA, data)
			return message, err
		case 13:
			message, err := ReadFormat13V1(nicSupplementA, data)
			return message, err
		case 14:
			message, err := ReadFormat14V1(nicSupplementA, data)
			return message, err
		case 15:
			message, err := ReadFormat15V1(nicSupplementA, data)
			return message, err
		case 16:
			message, err := ReadFormat16V1(nicSupplementA, data)
			return message, err
		case 17:
			message, err := ReadFormat17V1(nicSupplementA, data)
			return message, err
		case 18:
			message, err := ReadFormat18V1(nicSupplementA, data)
			return message, err
		case 20:
			message, err := ReadFormat20V1(data)
			return message, err
		case 21:
			message, err := ReadFormat21V1(data)
			return message, err
		case 22:
			message, err := ReadFormat22V1(data)
			return message, err
		default:
			return nil, fmt.Errorf("the Format Type %v of BSD 0,5 is not supported", formatTypeCode)
		}

	case adsb.ReaderLevel2:
		switch formatTypeCode {
		case 9:
			message, err := ReadFormat09V2(nicSupplementA, data)
			return message, err
		case 10:
			message, err := ReadFormat10V2(nicSupplementA, data)
			return message, err
		case 11:
			message, err := ReadFormat11V2(nicSupplementA, data)
			return message, err
		case 12:
			message, err := ReadFormat12V2(nicSupplementA, data)
			return message, err
		case 13:
			message, err := ReadFormat13V2(nicSupplementA, data)
			return message, err
		case 14:
			message, err := ReadFormat14V2(nicSupplementA, data)
			return message, err
		case 15:
			message, err := ReadFormat15V2(nicSupplementA, data)
			return message, err
		case 16:
			message, err := ReadFormat16V2(nicSupplementA, data)
			return message, err
		case 17:
			message, err := ReadFormat17V2(nicSupplementA, data)
			return message, err
		case 18:
			message, err := ReadFormat18V2(nicSupplementA, data)
			return message, err
		case 20:
			message, err := ReadFormat20V2(data)
			return message, err
		case 21:
			message, err := ReadFormat21V2(data)
			return message, err
		case 22:
			message, err := ReadFormat22V2(data)
			return message, err
		default:
			return nil, fmt.Errorf("the Format Type %v of BSD 0,5 is not supported", formatTypeCode)
		}

	default:
		return nil, fmt.Errorf("the BDS 0,5 format is not readable with the given reader level")
	}
}
