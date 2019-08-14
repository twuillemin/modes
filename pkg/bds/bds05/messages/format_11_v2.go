// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_formats_v2.go at 2019-08-15 00:36:41.8108698 +0300 EEST m=+0.008971701
package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
)

// ------------------------------------------------------------------------------------
//
//                                Format11V2
//
// ------------------------------------------------------------------------------------

// Format11V2 is a message at the format BDS 0,5 for ADSB V2
type Format11V2 struct {
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
func (message *Format11V2) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format11V2
}

// GetRegister returns the register of the message
func (message *Format11V2) GetRegister() bds.Register {
	return adsb.Format11V2.GetRegister()
}

// ToString returns a basic, but readable, representation of the message
func (message *Format11V2) ToString() string {
	return bds05v2ToString(message)
}

// GetSurveillanceStatus returns the Surveillance Status
func (message *Format11V2) GetSurveillanceStatus() fields.SurveillanceStatus {
	return message.SurveillanceStatus
}

// GetNavigationIntegrityCodeSupplementB returns the NavigationIntegrityCodeSupplementB
func (message *Format11V2) GetNavigationIntegrityCodeSupplementB() fields.NavigationIntegrityCodeSupplementB {
	return message.NavigationIntegrityCodeSupplementB
}

// GetAltitude returns the Altitude
func (message *Format11V2) GetAltitude() fields.Altitude {
	return message.Altitude
}

// GetTime returns the Time
func (message *Format11V2) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message *Format11V2) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message *Format11V2) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message *Format11V2) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalContainmentRadius returns the HorizontalContainmentRadius
func (message *Format11V2) GetHorizontalContainmentRadius() fields.HorizontalContainmentRadius {
	return message.HorizontalContainmentRadius
}

// GetNavigationIntegrityCategory returns the Navigation Integrity Category
func (message *Format11V2) GetNavigationIntegrityCategory() byte {
	return message.NavigationIntegrityCategory
}

// readFormat11V2 reads a message at the format BDS 0,5
func readFormat11V2(nicSupplementA bool, data []byte) (*Format11V2, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format11V2.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format11V2", formatTypeCode)
	}

	nicSupplementB := fields.ReadNavigationIntegritySupplementB(data)

	hcr, nic := getHCRAndNICForV2Barometric(formatTypeCode, nicSupplementA, nicSupplementB)

	return &Format11V2{
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
