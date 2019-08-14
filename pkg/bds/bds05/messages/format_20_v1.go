// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_formats_v1.go at 2019-08-15 00:36:41.2683773 +0300 EEST m=+0.015953101
package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
)

// ------------------------------------------------------------------------------------
//
//                                Format20V1
//
// ------------------------------------------------------------------------------------

// Format20V1 is a message at the format BDS 0,5 for ADSB V1
type Format20V1 struct {
	SurveillanceStatus          fields.SurveillanceStatus
	SingleAntennaFlag           fields.SingleAntennaFlag
	Altitude                    fields.Altitude
	Time                        fields.Time
	CPRFormat                   fields.CompactPositionReportingFormat
	EncodedLatitude             fields.EncodedLatitude
	EncodedLongitude            fields.EncodedLongitude
	HorizontalContainmentRadius fields.HorizontalContainmentRadiusGNSSV1
	NavigationIntegrityCategory byte
}

// GetMessageFormat returns the ADSB format of the message
func (message *Format20V1) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format20V1
}

// GetRegister returns the register of the message
func (message *Format20V1) GetRegister() bds.Register {
	return adsb.Format20V1.GetRegister()
}

// ToString returns a basic, but readable, representation of the message
func (message *Format20V1) ToString() string {
	return bds05v1ToString(message)
}

// GetSurveillanceStatus returns the Surveillance Status
func (message *Format20V1) GetSurveillanceStatus() fields.SurveillanceStatus {
	return message.SurveillanceStatus
}

// GetSingleAntennaFlag returns the SingleAntennaFlag
func (message *Format20V1) GetSingleAntennaFlag() fields.SingleAntennaFlag {
	return message.SingleAntennaFlag
}

// GetAltitude returns the Altitude
func (message *Format20V1) GetAltitude() fields.Altitude {
	return message.Altitude
}

// GetTime returns the Time
func (message *Format20V1) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message *Format20V1) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message *Format20V1) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message *Format20V1) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalContainmentRadius returns the HorizontalContainmentRadius
func (message *Format20V1) GetHorizontalContainmentRadius() fields.HorizontalContainmentRadius {
	return message.HorizontalContainmentRadius
}

// GetNavigationIntegrityCategory returns the Navigation Integrity Category
func (message *Format20V1) GetNavigationIntegrityCategory() byte {
	return message.NavigationIntegrityCategory
}

// readFormat20V1 reads a message at the format BDS 0,5
func readFormat20V1(data []byte) (*Format20V1, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format20V1.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format20V1", formatTypeCode)
	}

	hcr, nic := getHCRAndNICForV1GNSS(formatTypeCode)

	return &Format20V1{
		SurveillanceStatus:          fields.ReadSurveillanceStatus(data),
		SingleAntennaFlag:           fields.ReadSingleAntennaFlag(data),
		Altitude:                    fields.ReadAltitude(data),
		Time:                        fields.ReadTime(data),
		CPRFormat:                   fields.ReadCompactPositionReportingFormat(data),
		EncodedLatitude:             fields.ReadEncodedLatitude(data),
		EncodedLongitude:            fields.ReadEncodedLongitude(data),
		HorizontalContainmentRadius: hcr,
		NavigationIntegrityCategory: nic,
	}, nil
}
