package messages

import (
	"github.com/twuillemin/modes/pkg/bds/bds06/fields"
)

// Format07 is a message at the format BDS 0,6
type Format07 struct {
	Movement          fields.Movement
	GroundTrackStatus fields.GroundTrackStatus
	GroundTrack       fields.GroundTrack
	Time              fields.Time
	CPRFormat         fields.CompactPositionReportingFormat
	EncodedLatitude   fields.EncodedLatitude
	EncodedLongitude  fields.EncodedLongitude
}

// GetName returns the name of the message
func (message *Format07) GetName() string {
	return bds06Name
}

// GetBDS returns the binary data format
func (message *Format07) GetBDS() string {
	return bds06Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format07) GetFormatTypeCode() byte {
	return 5
}

// ToString returns a basic, but readable, representation of the message
func (message *Format07) ToString() string {
	return bds06ToString(message)
}

// GetMovement returns the Movement
func (message *Format07) GetMovement() fields.Movement {
	return message.Movement
}

// GetGroundTrackStatus returns the GroundTrackStatus
func (message *Format07) GetGroundTrackStatus() fields.GroundTrackStatus {
	return message.GroundTrackStatus
}

// GetGroundTrack returns the GroundTrack
func (message *Format07) GetGroundTrack() fields.GroundTrack {
	return message.GroundTrack
}

// GetTime returns the Time
func (message *Format07) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message *Format07) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message *Format07) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message *Format07) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalProtectionLimit returns the HorizontalProtectionLimit
func (message *Format07) GetHorizontalProtectionLimit() fields.HorizontalProtectionLimit {
	return fields.HPLLowerThan25M
}

// GetContainmentRadius returns the ContainmentRadius
func (message *Format07) GetContainmentRadius() fields.ContainmentRadius {
	return fields.CRBetween3MAnd10M
}

// ReadFormat07 reads a message at the format BDS 0,6
func ReadFormat07(data []byte) (*Format07, error) {

	return &Format07{
		Movement:          fields.ReadMovement(data),
		GroundTrackStatus: fields.ReadGroundTrackStatus(data),
		GroundTrack:       fields.ReadGroundTrack(data),
		Time:              fields.ReadTime(data),
		CPRFormat:         fields.ReadCompactPositionReportingFormat(data),
		EncodedLatitude:   fields.ReadEncodedLatitude(data),
		EncodedLongitude:  fields.ReadEncodedLongitude(data),
	}, nil
}
