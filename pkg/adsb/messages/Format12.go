package messages

import (
	"github.com/twuillemin/modes/pkg/adsb/fields"
)

// Format12 is a message at the format BDS 0,5
type Format12 struct {
	SurveillanceStatus fields.SurveillanceStatus
	SingleAntennaFlag  fields.SingleAntennaFlag
	Altitude           fields.Altitude
	Time               fields.Time
	CPRFormat          fields.CompactPositionReportingFormat
	EncodedLatitude    fields.EncodedLatitude
	EncodedLongitude   fields.EncodedLongitude
}

// GetName returns the name of the message
func (message *Format12) GetName() string {
	return bds05Name
}

// GetBDS returns the binary data format
func (message *Format12) GetBDS() string {
	return bds05Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format12) GetFormatTypeCode() byte {
	return 12
}

// ToString returns a basic, but readable, representation of the message
func (message *Format12) ToString() string {
	return bds05ToString(message)
}

// GetSurveillanceStatus returns the Surveillance Status
func (message *Format12) GetSurveillanceStatus() fields.SurveillanceStatus {
	return message.SurveillanceStatus
}

// GetSingleAntennaFlag returns the SingleAntennaFlag
func (message *Format12) GetSingleAntennaFlag() fields.SingleAntennaFlag {
	return message.SingleAntennaFlag
}

// GetAltitude returns the Altitude
func (message *Format12) GetAltitude() fields.Altitude {
	return message.Altitude
}

// GetTime returns the Time
func (message *Format12) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message *Format12) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message *Format12) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message *Format12) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalProtectionLimit returns the HorizontalProtectionLimit
func (message *Format12) GetHorizontalProtectionLimit() fields.HPL {
	return fields.HPLABBetween185Dot2MAnd370Dot4M
}

// GetContainmentRadius returns the ContainmentRadius
func (message *Format12) GetContainmentRadius() fields.ContainmentRadius {
	return fields.CRABBetween92Dot6MAnd185Dot2M
}

// ReadFormat12 reads a message at the format BDS 0,5
func ReadFormat12(data []byte) (*Format12, error) {

	return &Format12{
		SurveillanceStatus: fields.ReadSurveillanceStatus(data),
		SingleAntennaFlag:  fields.ReadSingleAntennaFlag(data),
		Altitude:           fields.ReadAltitude(data),
		Time:               fields.ReadTime(data),
		CPRFormat:          fields.ReadCompactPositionReportingFormat(data),
		EncodedLatitude:    fields.ReadEncodedLatitude(data),
		EncodedLongitude:   fields.ReadEncodedLongitude(data),
	}, nil
}
