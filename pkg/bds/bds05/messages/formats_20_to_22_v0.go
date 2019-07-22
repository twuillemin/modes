package messages

import (
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
)

// Format20To22V0 is a message at the format BDS 0,5
type Format20To22V0 struct {
	TypeCode                  byte
	SurveillanceStatus        fields.SurveillanceStatus
	SingleAntennaFlag         fields.SingleAntennaFlag
	Altitude                  fields.Altitude
	Time                      fields.Time
	CPRFormat                 fields.CompactPositionReportingFormat
	EncodedLatitude           fields.EncodedLatitude
	EncodedLongitude          fields.EncodedLongitude
	HorizontalProtectionLimit fields.HorizontalProtectionLimitGNSS
	ContainmentRadius         fields.ContainmentRadiusGNSS
}

// GetName returns the name of the message
func (message *Format20To22V0) GetName() string {
	return bds05Name
}

// GetRegister returns the binary data format
func (message *Format20To22V0) GetBDS() string {
	return bds05Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format20To22V0) GetFormatTypeCode() byte {
	return message.TypeCode
}

// ToString returns a basic, but readable, representation of the message
func (message *Format20To22V0) ToString() string {
	return bds05v0ToString(message)
}

// GetSurveillanceStatus returns the Surveillance Status
func (message *Format20To22V0) GetSurveillanceStatus() fields.SurveillanceStatus {
	return message.SurveillanceStatus
}

// GetSingleAntennaFlag returns the SingleAntennaFlag
func (message *Format20To22V0) GetSingleAntennaFlag() fields.SingleAntennaFlag {
	return message.SingleAntennaFlag
}

// GetAltitude returns the Altitude
func (message *Format20To22V0) GetAltitude() fields.Altitude {
	return message.Altitude
}

// GetTime returns the Time
func (message *Format20To22V0) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat  for ADSB V0
func (message *Format20To22V0) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message *Format20To22V0) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message *Format20To22V0) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalProtectionLimit returns the HorizontalProtectionLimit
func (message *Format20To22V0) GetHorizontalProtectionLimit() fields.HorizontalProtectionLimit {
	return message.HorizontalProtectionLimit
}

// GetContainmentRadius returns the ContainmentRadius
func (message *Format20To22V0) GetContainmentRadius() fields.ContainmentRadius {
	return message.ContainmentRadius
}

var hplGNSSByFormat = map[byte]fields.HorizontalProtectionLimitGNSS{
	20: fields.HPLGLowerThan7Dot5M,
	21: fields.HPLGLowerThan25M,
	22: fields.HPLGGreaterThan25M,
}

var crGNSSByFormat = map[byte]fields.ContainmentRadiusGNSS{
	20: fields.CRGHorizontalLowerThan3MAndVerticalLowerThan4M,
	21: fields.CRGHorizontalLowerThan10MAndVerticalLowerThan15M,
	22: fields.CRGHorizontalGreaterThan10MOrVerticalGreaterThan15M,
}

// readFormat20To22V0 reads a message at the format BDS 0,5
func readFormat20To22V0(data []byte) (*Format20To22V0, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3

	return &Format20To22V0{
		TypeCode:                  formatTypeCode,
		SurveillanceStatus:        fields.ReadSurveillanceStatus(data),
		SingleAntennaFlag:         fields.ReadSingleAntennaFlag(data),
		Altitude:                  fields.ReadAltitude(data),
		Time:                      fields.ReadTime(data),
		CPRFormat:                 fields.ReadCompactPositionReportingFormat(data),
		EncodedLatitude:           fields.ReadEncodedLatitude(data),
		EncodedLongitude:          fields.ReadEncodedLongitude(data),
		HorizontalProtectionLimit: hplGNSSByFormat[formatTypeCode],
		ContainmentRadius:         crGNSSByFormat[formatTypeCode],
	}, nil
}
