// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_formats_v0.go at 2019-08-15 19:22:00.7194863 +0300 EEST m=+0.006981001
package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
)

// ------------------------------------------------------------------------------------
//
//                                Format09V0
//
// ------------------------------------------------------------------------------------

// Format09V0 is a message at the format BDS 0,5 for ADSB V0
type Format09V0 struct {
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
func (message Format09V0) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format09
}

// GetRegister returns the register of the message
func (message Format09V0) GetRegister() bds.Register {
	return adsb.Format09.GetRegister()
}

// GetSubtype returns the subtype of the message if any
func (message Format09V0) GetSubtype() adsb.Subtype {
	return nil
}

// GetMinimumADSBLevel returns the minimum ADSB Level for the message
func (message Format09V0) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel0
}

// GetMaximumADSBLevel returns the maximum ADSB Level for the message
func (message Format09V0) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel0
}

// GetSurveillanceStatus returns the Surveillance Status
func (message Format09V0) GetSurveillanceStatus() fields.SurveillanceStatus {
	return message.SurveillanceStatus
}

// GetSingleAntennaFlag returns the SingleAntennaFlag
func (message Format09V0) GetSingleAntennaFlag() fields.SingleAntennaFlag {
	return message.SingleAntennaFlag
}

// GetAltitude returns the Altitude
func (message Format09V0) GetAltitude() fields.Altitude {
	return message.Altitude
}

// GetTime returns the Time
func (message Format09V0) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message Format09V0) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message Format09V0) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message Format09V0) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalProtectionLimit returns the HorizontalProtectionLimit
func (message Format09V0) GetHorizontalProtectionLimit() fields.HorizontalProtectionLimit {
	return message.HorizontalProtectionLimit
}

// GetContainmentRadius returns the ContainmentRadius
func (message Format09V0) GetContainmentRadius() fields.ContainmentRadius {
	return message.ContainmentRadius
}

// ToString returns a basic, but readable, representation of the message
func (message Format09V0) ToString() string {
	return bds05v0ToString(message)
}

// ReadFormat09V0 reads a message at the format Format09V0
func ReadFormat09V0(data []byte) (*Format09V0, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format09.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format09V0", formatTypeCode)
	}

	horizontalProtectionLimit := hplBarometricByFormat[formatTypeCode]
	containmentRadius := crBarometricByFormat[formatTypeCode]

	return &Format09V0{
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
