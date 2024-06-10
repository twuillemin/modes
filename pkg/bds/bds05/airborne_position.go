package bds05

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// AirbornePosition is a message at the format BDS 0,5
type AirbornePosition struct {
	FormatTypeCode       byte
	SurveillanceStatus   fields.SurveillanceStatus
	SingleAntennaFlag    fields.SingleAntennaFlag
	AltitudeSource       fields.AltitudeSource
	AltitudeReportMethod fields.AltitudeReportMethod
	AltitudeInFeet       int
	Time                 fields.Time
	CPRFormat            fields.CompactPositionReportingFormat
	EncodedLatitude      fields.EncodedLatitude
	EncodedLongitude     fields.EncodedLongitude
}

// GetRegister returns the Register the message
func (message AirbornePosition) GetRegister() register.Register {
	return register.BDS05
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message AirbornePosition) CheckCoherency() error {
	return nil
}

func (message AirbornePosition) ToString() string {
	return fmt.Sprintf(""+
		"Message:                           %v\n"+
		"Format Type Code:                  %v\n"+
		"Surveillance Status:               %v\n"+
		"Single Antenna Flag:               %v\n"+
		"Altitude Source:                   %v\n"+
		"Altitude Report Method:            %v\n"+
		"Altitude (feet):                   %v\n"+
		"Time:                              %v\n"+
		"Compact Position Reporting Format: %v\n"+
		"Encoded Latitude:                  %v\n"+
		"Encoded Longitude:                 %v",
		message.GetRegister().ToString(),
		message.FormatTypeCode,
		message.SurveillanceStatus.ToString(),
		message.SingleAntennaFlag.ToString(),
		message.AltitudeSource.ToString(),
		message.AltitudeReportMethod.ToString(),
		message.AltitudeInFeet,
		message.Time.ToString(),
		message.CPRFormat.ToString(),
		message.EncodedLatitude,
		message.EncodedLongitude)
}

// ReadAirbornePosition reads a message at the format Format09V1
func ReadAirbornePosition(data []byte) (*AirbornePosition, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3

	if formatTypeCode < 9 || formatTypeCode > 22 || formatTypeCode == 19 {
		return nil, fmt.Errorf("the field FormatTypeCode must be comprised between 9 and 18 included or 20 and 22 included, got %v", formatTypeCode)
	}

	altitude, altitudeSource, altitudeReportMethod := fields.ReadAltitude(data)

	return &AirbornePosition{
		FormatTypeCode:       formatTypeCode,
		SurveillanceStatus:   fields.ReadSurveillanceStatus(data),
		SingleAntennaFlag:    fields.ReadSingleAntennaFlag(data),
		AltitudeSource:       altitudeSource,
		AltitudeReportMethod: altitudeReportMethod,
		AltitudeInFeet:       altitude,
		Time:                 fields.ReadTime(data),
		CPRFormat:            fields.ReadCompactPositionReportingFormat(data),
		EncodedLatitude:      fields.ReadEncodedLatitude(data),
		EncodedLongitude:     fields.ReadEncodedLongitude(data),
	}, nil
}
