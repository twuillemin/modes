package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb/fields"
)

// MessageBDS05 is the basic interface that ADSB messages at the format BDS 0,5 are expected to implement
type MessageBDS05 interface {
	ADSBMessage
	GetSurveillanceStatus() fields.SurveillanceStatus
	GetSingleAntennaFlag() fields.SingleAntennaFlag
	GetAltitude() fields.Altitude
	GetTime() fields.Time
	GetCPRFormat() fields.CPRFormat
	GetEncodedLatitude() fields.EncodedLatitude
	GetEncodedLongitude() fields.EncodedLongitude
	GetHorizontalProtectionLimit() fields.HPL
	GetContainmentRadius() fields.ContainmentRadiusAirborne
}

var bds05Code = "BDS 0,5"
var bds05Name = "Extended squitter airborne position"

func bds05ToString(message MessageBDS05) string {
	return fmt.Sprintf("Message:                   %v (%v)\n"+
		"FormatTypeCode:            %v\n"+
		"SurveillanceStatus:        %v\n"+
		"HorizontalProtectionLimit: %v\n"+
		"ContainmentRadius:         %v\n"+
		"SingleAntenna:             %v\n"+
		"Time:                      %v\n"+
		"CPRFormat:                 %v\n"+
		"GetAltitude:               %v\n"+
		"EncodedLatitude:           %v\n"+
		"EncodedLongitude:          %v",
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
	}

	return nil, fmt.Errorf("the format type code %v can not be read as a BDS 0,5 format", formatTypeCode)
}
