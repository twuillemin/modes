// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_formats_v1.go at 2024-06-02 22:43:46.9205365 +0300 EEST m=+0.002581001
package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/bds06/fields"
)

// ------------------------------------------------------------------------------------
//
//                                Format06V1
//
// ------------------------------------------------------------------------------------

// Format06V1 is a message at the format BDS 0,6
type Format06V1 struct {
	Movement                    fields.Movement
	GroundTrackStatus           fields.GroundTrackStatus
	GroundTrack                 fields.GroundTrack
	Time                        fields.Time
	CPRFormat                   fields.CompactPositionReportingFormat
	EncodedLatitude             fields.EncodedLatitude
	EncodedLongitude            fields.EncodedLongitude
	HorizontalContainmentRadius fields.HorizontalContainmentRadiusV1
	NavigationIntegrityCategory byte
}

// GetMessageFormat returns the ADSB format of the message
func (message Format06V1) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format06
}

// GetSubtype returns the subtype of the message if any
func (message Format06V1) GetSubtype() adsb.Subtype {
	return nil
}

// GetMinimumADSBLevel returns the minimum ADSB Level for the message
func (message Format06V1) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel1
}

// GetMaximumADSBLevel returns the maximum ADSB Level for the message
func (message Format06V1) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel1
}

// GetMovement returns the Movement
func (message Format06V1) GetMovement() fields.Movement {
	return message.Movement
}

// GetGroundTrackStatus returns the GroundTrackStatus
func (message Format06V1) GetGroundTrackStatus() fields.GroundTrackStatus {
	return message.GroundTrackStatus
}

// GetGroundTrack returns the GroundTrack
func (message Format06V1) GetGroundTrack() fields.GroundTrack {
	return message.GroundTrack
}

// GetTime returns the Time
func (message Format06V1) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message Format06V1) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message Format06V1) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message Format06V1) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalContainmentRadius returns the HorizontalContainmentRadiusV1
func (message Format06V1) GetHorizontalContainmentRadius() fields.HorizontalContainmentRadiusV1 {
	return message.HorizontalContainmentRadius
}

// GetNavigationIntegrityCategory returns the Navigation Integrity Category
func (message Format06V1) GetNavigationIntegrityCategory() byte {
	return message.NavigationIntegrityCategory
}

// ToString returns a basic, but readable, representation of the message
func (message Format06V1) ToString() string {
	return messageBDS06V1ToString(message)
}

// ReadFormat06V1 reads a message at the format Format06V1 for ADSB V1
func ReadFormat06V1(nicSupplementA bool, data []byte) (*Format06V1, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format06.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format06V1", formatTypeCode)
	}

	hcr, nic := getHCRAndNICForV1(formatTypeCode, nicSupplementA)

	return &Format06V1{
		Movement:                    fields.ReadMovement(data),
		GroundTrackStatus:           fields.ReadGroundTrackStatus(data),
		GroundTrack:                 fields.ReadGroundTrack(data),
		Time:                        fields.ReadTime(data),
		CPRFormat:                   fields.ReadCompactPositionReportingFormat(data),
		EncodedLatitude:             fields.ReadEncodedLatitude(data),
		EncodedLongitude:            fields.ReadEncodedLongitude(data),
		HorizontalContainmentRadius: hcr,
		NavigationIntegrityCategory: nic,
	}, nil
}
