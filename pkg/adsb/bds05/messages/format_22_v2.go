// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_formats_v2.go at 2024-06-02 22:43:43.3654964 +0300 EEST m=+0.007569401
package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/bds05/fields"
)

// ------------------------------------------------------------------------------------
//
//                                Format22V2
//
// ------------------------------------------------------------------------------------

// Format22V2 is a message at the format BDS 0,5 for ADSB V2
type Format22V2 struct {
	SurveillanceStatus                 fields.SurveillanceStatus
	NavigationIntegrityCodeSupplementB fields.NavigationIntegrityCodeSupplementB
	Altitude                           fields.Altitude
	Time                               fields.Time
	CPRFormat                          fields.CompactPositionReportingFormat
	EncodedLatitude                    fields.EncodedLatitude
	EncodedLongitude                   fields.EncodedLongitude
	HorizontalContainmentRadius        fields.HorizontalContainmentRadiusGNSSV2
	NavigationIntegrityCategory        byte
}

// GetMessageFormat returns the ADSB format of the message
func (message Format22V2) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format22
}

// GetSubtype returns the subtype of the message if any
func (message Format22V2) GetSubtype() adsb.Subtype {
	return nil
}

// GetMinimumADSBLevel returns the minimum ADSB Level for the message
func (message Format22V2) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel2
}

// GetMaximumADSBLevel returns the maximum ADSB Level for the message
func (message Format22V2) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel2
}

// GetSurveillanceStatus returns the Surveillance Status
func (message Format22V2) GetSurveillanceStatus() fields.SurveillanceStatus {
	return message.SurveillanceStatus
}

// GetNavigationIntegrityCodeSupplementB returns the NavigationIntegrityCodeSupplementB
func (message Format22V2) GetNavigationIntegrityCodeSupplementB() fields.NavigationIntegrityCodeSupplementB {
	return message.NavigationIntegrityCodeSupplementB
}

// GetAltitude returns the Altitude
func (message Format22V2) GetAltitude() fields.Altitude {
	return message.Altitude
}

// GetTime returns the Time
func (message Format22V2) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message Format22V2) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message Format22V2) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message Format22V2) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalContainmentRadius returns the HorizontalContainmentRadius
func (message Format22V2) GetHorizontalContainmentRadius() fields.HorizontalContainmentRadius {
	return message.HorizontalContainmentRadius
}

// GetNavigationIntegrityCategory returns the Navigation Integrity Category
func (message Format22V2) GetNavigationIntegrityCategory() byte {
	return message.NavigationIntegrityCategory
}

// ToString returns a basic, but readable, representation of the message
func (message Format22V2) ToString() string {
	return bds05v2ToString(message)
}

// ReadFormat22V2 reads a message at the format Format22V2
func ReadFormat22V2(data []byte) (*Format22V2, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format22.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format22V2", formatTypeCode)
	}

	nicSupplementB := fields.ReadNavigationIntegritySupplementB(data)

	hcr, nic := getHCRAndNICForV2GNSS(formatTypeCode)

	return &Format22V2{
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
