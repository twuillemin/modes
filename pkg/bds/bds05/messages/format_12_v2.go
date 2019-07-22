// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_formats_v2.go at 2019-07-22 22:14:35.3309519 +0300 EEST m=+0.010969601
package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
)

// ------------------------------------------------------------------------------------
//
//                                Format12V2
//
// ------------------------------------------------------------------------------------

// Format12V2 is a message at the format BDS 0,5 for ADSB V2
type Format12V2 struct {
	SurveillanceStatus                 fields.SurveillanceStatus
	NavigationIntegrityCodeSupplementB fields.NavigationIntegrityCodeSupplementB
	Altitude                           fields.Altitude
	Time                               fields.Time
	CPRFormat                          fields.CompactPositionReportingFormat
	EncodedLatitude                    fields.EncodedLatitude
	EncodedLongitude                   fields.EncodedLongitude
	HorizontalContainmentRadius        fields.HorizontalContainmentRadiusBarometricV2
	NavigationIntegrityCategory        byte
}

// GetMessageFormat returns the ADSB format of the message
func (message *Format12V2) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format12V2
}

// GetRegister returns the register of the message
func (message *Format12V2) GetRegister() bds.Register {
	return adsb.Format12V2.GetRegister()
}

// ToString returns a basic, but readable, representation of the message
func (message *Format12V2) ToString() string {
	return bds05v2ToString(message)
}

// GetSurveillanceStatus returns the Surveillance Status
func (message *Format12V2) GetSurveillanceStatus() fields.SurveillanceStatus {
	return message.SurveillanceStatus
}

// GetNavigationIntegrityCodeSupplementB returns the NavigationIntegrityCodeSupplementB
func (message *Format12V2) GetNavigationIntegrityCodeSupplementB() fields.NavigationIntegrityCodeSupplementB {
	return message.NavigationIntegrityCodeSupplementB
}

// GetAltitude returns the Altitude
func (message *Format12V2) GetAltitude() fields.Altitude {
	return message.Altitude
}

// GetTime returns the Time
func (message *Format12V2) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message *Format12V2) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message *Format12V2) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message *Format12V2) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalContainmentRadius returns the HorizontalContainmentRadius
func (message *Format12V2) GetHorizontalContainmentRadius() fields.HorizontalContainmentRadius {
	return message.HorizontalContainmentRadius
}

// GetNavigationIntegrityCategory returns the Navigation Integrity Category
func (message *Format12V2) GetNavigationIntegrityCategory() byte {
	return message.NavigationIntegrityCategory
}

// readFormat12V2 reads a message at the format BDS 0,5
func readFormat12V2(nicSupplementA bool, data []byte) (*Format12V2, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3

	if formatTypeCode != adsb.Format12V2.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format12V2", formatTypeCode)
	}

	nicSupplementB := fields.ReadNavigationIntegritySupplementB(data)

	hcr, nic := getHCRAndNICForV2Barometric(formatTypeCode, nicSupplementA, nicSupplementB)

	return &Format12V2{
		SurveillanceStatus:                 fields.ReadSurveillanceStatus(data),
		NavigationIntegrityCodeSupplementB: nicSupplementB,
		Altitude:                           fields.ReadAltitude(data),
		Time:                               fields.ReadTime(data),
		CPRFormat:                          fields.ReadCompactPositionReportingFormat(data),
		EncodedLatitude:                    fields.ReadEncodedLatitude(data),
		EncodedLongitude:                   fields.ReadEncodedLongitude(data),
		HorizontalContainmentRadius:        hcr,
		NavigationIntegrityCategory:        nic,
	}, nil
}
