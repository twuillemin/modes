// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_formats_v0.go at 2024-06-02 22:43:46.3289943 +0300 EEST m=+0.003475301
package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/bds06/fields"
)

// ------------------------------------------------------------------------------------
//
//                                Format08V0
//
// ------------------------------------------------------------------------------------

// Format08V0 is a message at the format BDS 0,6
type Format08V0 struct {
	Movement                  fields.Movement
	GroundTrackStatus         fields.GroundTrackStatus
	GroundTrack               fields.GroundTrack
	Time                      fields.Time
	CPRFormat                 fields.CompactPositionReportingFormat
	EncodedLatitude           fields.EncodedLatitude
	EncodedLongitude          fields.EncodedLongitude
	HorizontalProtectionLimit fields.HorizontalProtectionLimit
	ContainmentRadius         fields.ContainmentRadius
}

// GetMessageFormat returns the ADSB format of the message
func (message Format08V0) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format08
}

// GetSubtype returns the subtype of the message if any
func (message Format08V0) GetSubtype() adsb.Subtype {
	return nil
}

// GetMinimumADSBLevel returns the minimum ADSB Level for the message
func (message Format08V0) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel0
}

// GetMaximumADSBLevel returns the maximum ADSB Level for the message
func (message Format08V0) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel0
}

// GetMovement returns the Movement
func (message Format08V0) GetMovement() fields.Movement {
	return message.Movement
}

// GetGroundTrackStatus returns the GroundTrackStatus
func (message Format08V0) GetGroundTrackStatus() fields.GroundTrackStatus {
	return message.GroundTrackStatus
}

// GetGroundTrack returns the GroundTrack
func (message Format08V0) GetGroundTrack() fields.GroundTrack {
	return message.GroundTrack
}

// GetTime returns the Time
func (message Format08V0) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message Format08V0) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message Format08V0) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message Format08V0) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalProtectionLimit returns the HorizontalProtectionLimit
func (message Format08V0) GetHorizontalProtectionLimit() fields.HorizontalProtectionLimit {
	return message.HorizontalProtectionLimit
}

// GetContainmentRadius returns the ContainmentRadius
func (message Format08V0) GetContainmentRadius() fields.ContainmentRadius {
	return message.ContainmentRadius
}

// ToString returns a basic, but readable, representation of the message
func (message Format08V0) ToString() string {
	return messageBDS06V0ToString(message)
}

// ReadFormat08V0 reads a message at the format Format08V0 for ADSB V0
func ReadFormat08V0(data []byte) (*Format08V0, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format08.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format Format08V0", formatTypeCode)
	}

	return &Format08V0{
		Movement:                  fields.ReadMovement(data),
		GroundTrackStatus:         fields.ReadGroundTrackStatus(data),
		GroundTrack:               fields.ReadGroundTrack(data),
		Time:                      fields.ReadTime(data),
		CPRFormat:                 fields.ReadCompactPositionReportingFormat(data),
		EncodedLatitude:           fields.ReadEncodedLatitude(data),
		EncodedLongitude:          fields.ReadEncodedLongitude(data),
		HorizontalProtectionLimit: hplByFormat[formatTypeCode],
		ContainmentRadius:         crByFormat[formatTypeCode],
	}, nil
}