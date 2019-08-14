// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_formats_v0.go at 2019-08-15 00:36:40.7148265 +0300 EEST m=+0.014954901
package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
)

// ------------------------------------------------------------------------------------
//
//                                Format18V0
//
// ------------------------------------------------------------------------------------

// Format18V0 is a message at the format BDS 0,5 for ADSB V0
type Format18V0 struct {
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

// GetMessageFormat returns the ADSB format of the message
func (message *Format18V0) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format18V0
}

// GetRegister returns the register of the message
func (message *Format18V0) GetRegister() bds.Register {
	return adsb.Format18V0.GetRegister()
}

// ToString returns a basic, but readable, representation of the message
func (message *Format18V0) ToString() string {
	return bds05v0ToString(message)
}

// GetSurveillanceStatus returns the Surveillance Status
func (message *Format18V0) GetSurveillanceStatus() fields.SurveillanceStatus {
	return message.SurveillanceStatus
}

// GetSingleAntennaFlag returns the SingleAntennaFlag
func (message *Format18V0) GetSingleAntennaFlag() fields.SingleAntennaFlag {
	return message.SingleAntennaFlag
}

// GetAltitude returns the Altitude
func (message *Format18V0) GetAltitude() fields.Altitude {
	return message.Altitude
}

// GetTime returns the Time
func (message *Format18V0) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message *Format18V0) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message *Format18V0) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message *Format18V0) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalProtectionLimit returns the HorizontalProtectionLimit
func (message *Format18V0) GetHorizontalProtectionLimit() fields.HorizontalProtectionLimit {
	return message.HorizontalProtectionLimit
}

// GetContainmentRadius returns the ContainmentRadius
func (message *Format18V0) GetContainmentRadius() fields.ContainmentRadius {
	return message.ContainmentRadius
}

// readFormat18V0 reads a message at the format BDS 0,5
func readFormat18V0(data []byte) (*Format18V0, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format18V0.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format18V0", formatTypeCode)
	}

	horizontalProtectionLimit := hplBarometricByFormat[formatTypeCode]
	containmentRadius := crBarometricByFormat[formatTypeCode]

	return &Format18V0{
		SurveillanceStatus:        fields.ReadSurveillanceStatus(data),
		SingleAntennaFlag:         fields.ReadSingleAntennaFlag(data),
		Altitude:                  fields.ReadAltitude(data),
		Time:                      fields.ReadTime(data),
		CPRFormat:                 fields.ReadCompactPositionReportingFormat(data),
		EncodedLatitude:           fields.ReadEncodedLatitude(data),
		EncodedLongitude:          fields.ReadEncodedLongitude(data),
		HorizontalProtectionLimit: horizontalProtectionLimit,
		ContainmentRadius:         containmentRadius,
	}, nil
}
