package messages

import (
	"github.com/twuillemin/modes/pkg/adsb/fields"
)

// Format17 is a message at the format BDS 0,5
type Format17 struct {
	SurveillanceStatus fields.SurveillanceStatus
	SingleAntennaFlag  fields.SingleAntennaFlag
	Altitude           fields.Altitude
	Time               fields.Time
	CPRFormat          fields.CPRFormat
	EncodedLatitude    fields.EncodedLatitude
	EncodedLongitude   fields.EncodedLongitude
}

// GetName returns the name of the message
func (message *Format17) GetName() string {
	return bds05Name
}

// GetBDS returns the binary data format
func (message *Format17) GetBDS() string {
	return bds05Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format17) GetFormatTypeCode() byte {
	return 9
}

// ToString returns a basic, but readable, representation of the field
func (message *Format17) ToString() string {
	return bds05ToString(message)
}

// GetSurveillanceStatus returns the Surveillance Status
func (message *Format17) GetSurveillanceStatus() fields.SurveillanceStatus {
	return message.SurveillanceStatus
}

// GetSingleAntennaFlag returns the SingleAntennaFlag
func (message *Format17) GetSingleAntennaFlag() fields.SingleAntennaFlag {
	return message.SingleAntennaFlag
}

// GetAltitude returns the Altitude
func (message *Format17) GetAltitude() fields.Altitude {
	return message.Altitude
}

// GetTime returns the Time
func (message *Format17) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CPRFormat
func (message *Format17) GetCPRFormat() fields.CPRFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message *Format17) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message *Format17) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalProtectionLimit returns the HorizontalProtectionLimit
func (message *Format17) GetHorizontalProtectionLimit() fields.HPLAirborne {
	return fields.HPLABetween18Point52KmAnd37Point04Km
}

// GetContainmentRadius returns the ContainmentRadius
func (message *Format17) GetContainmentRadius() fields.ContainmentRadiusAirborne {
	return fields.CRABetween9Point26KmAnd18Point52
}

// ReadFormat17 reads a message at the format BDS 0,5
func ReadFormat17(data []byte) (*Format17, error) {

	return &Format17{
		SurveillanceStatus: fields.ReadSurveillanceStatus(data),
		SingleAntennaFlag:  fields.ReadSingleAntennaFlag(data),
		Altitude:           fields.ReadAltitude(data),
		Time:               fields.ReadTime(data),
		CPRFormat:          fields.ReadCPRFormat(data),
		EncodedLatitude:    fields.ReadEncodedLatitude(data),
		EncodedLongitude:   fields.ReadEncodedLongitude(data),
	}, nil
}
