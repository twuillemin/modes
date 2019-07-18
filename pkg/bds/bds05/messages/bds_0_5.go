package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb/messages"
	fields2 "github.com/twuillemin/modes/pkg/bds/bds05/fields"
)

// MessageBDS05 is the basic interface that ADSB messages at the format BDS 0,5 are expected to implement
type MessageBDS05 interface {
	messages.ADSBMessage
	GetSurveillanceStatus() fields2.SurveillanceStatus
	GetSingleAntennaFlag() fields2.SingleAntennaFlag
	GetAltitude() fields2.Altitude
	GetTime() fields2.Time
	GetCPRFormat() fields2.CompactPositionReportingFormat
	GetEncodedLatitude() fields2.EncodedLatitude
	GetEncodedLongitude() fields2.EncodedLongitude
	GetHorizontalProtectionLimit() fields2.HPL
	GetContainmentRadius() fields2.ContainmentRadius
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

	formatTypeCode := (data[0] & 0xF8) >> 3

	switch formatTypeCode {
	case 9:
		return messages2.ReadFormat09(data)
	case 10:
		return messages2.ReadFormat10(data)
	case 11:
		return messages2.ReadFormat11(data)
	case 12:
		return messages2.ReadFormat12(data)
	case 13:
		return messages2.ReadFormat13(data)
	case 14:
		return messages2.ReadFormat14(data)
	case 15:
		return messages2.ReadFormat15(data)
	case 16:
		return messages2.ReadFormat16(data)
	case 17:
		return messages2.ReadFormat17(data)
	case 18:
		return messages2.ReadFormat18(data)
	case 20:
		return messages2.ReadFormat20(data)
	case 21:
		return messages2.ReadFormat21(data)
	case 22:
		return messages2.ReadFormat22(data)
	}

	return nil, fmt.Errorf("the format type code %v can not be read as a BDS 0,5 format", formatTypeCode)
}
