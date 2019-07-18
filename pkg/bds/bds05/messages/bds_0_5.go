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
	// GetSingleAntennaFlag returns the SingleAntennaFlag
	GetSingleAntennaFlag() fields.SingleAntennaFlag
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
	// GetHorizontalProtectionLimit returns the HorizontalProtectionLimit
	GetHorizontalProtectionLimit() fields.HPL
	// GetContainmentRadius returns the ContainmentRadius
	GetContainmentRadius() fields.ContainmentRadius
}

var bds05Code = "BDS 0,5"
var bds05Name = "Extended squitter airborne position"

func bds05ToString(message MessageBDS05) string {
	return fmt.Sprintf("Message:                        %v (%v)\n"+
		"FormatTypeCode:                 %v\n"+
		"SurveillanceStatus:             %v\n"+
		"HorizontalProtectionLimit:      %v\n"+
		"ContainmentRadius:              %v\n"+
		"SingleAntenna:                  %v\n"+
		"Time:                           %v\n"+
		"CompactPositionReportingFormat: %v\n"+
		"GetAltitude:                    %v\n"+
		"EncodedLatitude:                %v\n"+
		"EncodedLongitude:               %v",
		message.GetBDS(),
		message.GetName(),
		message.GetFormatTypeCode(),
		message.GetSurveillanceStatus().ToString(),
		message.GetHorizontalProtectionLimit().ToString(),
		message.GetContainmentRadius().ToString(),
		message.GetSingleAntennaFlag().ToString(),
		message.GetTime().ToString(),
		message.GetCPRFormat().ToString(),
		message.GetAltitude().ToString(),
		message.GetEncodedLatitude(),
		message.GetEncodedLongitude())
}

// ReadBDS05 reads a message at the format BDS 0,5
func ReadBDS05(data []byte) (MessageBDS05, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for BDS message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3

	switch formatTypeCode {
	case 9:
		return ReadFormat09(data)
	case 10:
		return ReadFormat10(data)
	case 11:
		return ReadFormat11(data)
	case 12:
		return ReadFormat12(data)
	case 13:
		return ReadFormat13(data)
	case 14:
		return ReadFormat14(data)
	case 15:
		return ReadFormat15(data)
	case 16:
		return ReadFormat16(data)
	case 17:
		return ReadFormat17(data)
	case 18:
		return ReadFormat18(data)
	case 20:
		return ReadFormat20(data)
	case 21:
		return ReadFormat21(data)
	case 22:
		return ReadFormat22(data)
	}

	return nil, fmt.Errorf("the format type code %v can not be read as a BDS 0,5 format", formatTypeCode)
}
