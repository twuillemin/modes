package messages

import (
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
)

// Format09To18V0 is a message at the format BDS 0,5 for ADSB V0
type Format09To18V0 struct {
	TypeCode                  byte
	SurveillanceStatus        fields.SurveillanceStatus
	SingleAntennaFlag         fields.SingleAntennaFlag
	Altitude                  fields.Altitude
	Time                      fields.Time
	CPRFormat                 fields.CompactPositionReportingFormat
	EncodedLatitude           fields.EncodedLatitude
	EncodedLongitude          fields.EncodedLongitude
	HorizontalProtectionLimit fields.HorizontalProtectionLimitBarometric
	ContainmentRadius         fields.ContainmentRadiusBarometric
}

// GetName returns the name of the message
func (message *Format09To18V0) GetName() string {
	return bds05Name
}

// GetBDS returns the binary data format
func (message *Format09To18V0) GetBDS() string {
	return bds05Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format09To18V0) GetFormatTypeCode() byte {
	return message.TypeCode
}

// ToString returns a basic, but readable, representation of the message
func (message *Format09To18V0) ToString() string {
	return bds05v0ToString(message)
}

// GetSurveillanceStatus returns the Surveillance Status
func (message *Format09To18V0) GetSurveillanceStatus() fields.SurveillanceStatus {
	return message.SurveillanceStatus
}

// GetSingleAntennaFlag returns the SingleAntennaFlag
func (message *Format09To18V0) GetSingleAntennaFlag() fields.SingleAntennaFlag {
	return message.SingleAntennaFlag
}

// GetAltitude returns the Altitude
func (message *Format09To18V0) GetAltitude() fields.Altitude {
	return message.Altitude
}

// GetTime returns the Time
func (message *Format09To18V0) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message *Format09To18V0) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message *Format09To18V0) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message *Format09To18V0) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalProtectionLimit returns the HorizontalProtectionLimit
func (message *Format09To18V0) GetHorizontalProtectionLimit() fields.HorizontalProtectionLimit {
	return message.HorizontalProtectionLimit
}

// GetContainmentRadius returns the ContainmentRadius
func (message *Format09To18V0) GetContainmentRadius() fields.ContainmentRadius {
	return message.ContainmentRadius
}

var hplBaroByFormat = map[byte]fields.HorizontalProtectionLimitBarometric{
	9:  fields.HPLBLowerThan7Dot5M,
	10: fields.HPLBBetween7Dot5MAnd25M,
	11: fields.HPLBBetween25MAnd185Dot2M,
	12: fields.HPLBBetween185Dot2MAnd370Dot4M,
	13: fields.HPLBBetween370Dot4MAnd926M,
	14: fields.HPLBBetween926MAnd1852M,
	15: fields.HPLBBetween1852MAnd3704M,
	16: fields.HPLBBetween3704MAnd18Point52Km,
	17: fields.HPLBBetween18Point52KmAnd37Point04Km,
	18: fields.HPLBLargerThan37Point04Km,
}

var crBaroByFormat = map[byte]fields.ContainmentRadiusBarometric{
	9:  fields.CRBLowerThan3M,
	10: fields.CRBBetween3MAnd10M,
	11: fields.CRBBetween10MAnd92Dot6M,
	12: fields.CRBBetween92Dot6MAnd185Dot2M,
	13: fields.CRBBetween185Dot2MAnd463M,
	14: fields.CRBBetween463MAnd926M,
	15: fields.CRBBetween926MAnd1852M,
	16: fields.CRBBetween1Point852KmAnd9Point26Km,
	17: fields.CRBBetween9Point26KmAnd18Point52Km,
	18: fields.CRBLargerThan18Point52Km,
}

// readFormat09To18V0 reads a message at the format BDS 0,5
func readFormat09To18V0(data []byte) (*Format09To18V0, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3

	return &Format09To18V0{
		TypeCode:                  formatTypeCode,
		SurveillanceStatus:        fields.ReadSurveillanceStatus(data),
		SingleAntennaFlag:         fields.ReadSingleAntennaFlag(data),
		Altitude:                  fields.ReadAltitude(data),
		Time:                      fields.ReadTime(data),
		CPRFormat:                 fields.ReadCompactPositionReportingFormat(data),
		EncodedLatitude:           fields.ReadEncodedLatitude(data),
		EncodedLongitude:          fields.ReadEncodedLongitude(data),
		HorizontalProtectionLimit: hplBaroByFormat[formatTypeCode],
		ContainmentRadius:         crBaroByFormat[formatTypeCode],
	}, nil
}
