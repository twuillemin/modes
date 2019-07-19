package messages

import (
	"github.com/twuillemin/modes/pkg/bds/bds06/fields"
)

// Format05 is a message at the format BDS 0,6
type Format05 struct {
	Movement          fields.Movement
	GroundTrackStatus fields.GroundTrackStatus
	GroundTrack       fields.GroundTrack
	Time              fields.Time
	CPRFormat         fields.CompactPositionReportingFormat
	EncodedLatitude   fields.EncodedLatitude
	EncodedLongitude  fields.EncodedLongitude
}

// GetName returns the name of the message
func (message *Format05) GetName() string {
	return bds06Name
}

// GetBDS returns the binary data format
func (message *Format05) GetBDS() string {
	return bds06Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format05) GetFormatTypeCode() byte {
	return 5
}

// ToString returns a basic, but readable, representation of the message
func (message *Format05) ToString() string {
	return bds06ToString(message)
}

// GetMovement returns the Movement
func (message *Format05) GetMovement() fields.Movement {
	return message.Movement
}

// GetGroundTrackStatus returns the GroundTrackStatus
func (message *Format05) GetGroundTrackStatus() fields.GroundTrackStatus {
	return message.GroundTrackStatus
}

// GetGroundTrack returns the GroundTrack
func (message *Format05) GetGroundTrack() fields.GroundTrack {
	return message.GroundTrack
}

// GetTime returns the Time
func (message *Format05) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message *Format05) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message *Format05) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message *Format05) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalProtectionLimit returns the HorizontalProtectionLimit
func (message *Format05) GetHorizontalProtectionLimit() fields.HorizontalProtectionLimit {
	return fields.HPLLowerThan7Point5M
}

// GetContainmentRadius returns the ContainmentRadius
func (message *Format05) GetContainmentRadius() fields.ContainmentRadius {
	return fields.CRLowerThan3M
}

// ReadFormat05 reads a message at the format BDS 0,6
func ReadFormat05(data []byte) (*Format05, error) {

	return &Format05{
		Movement:          fields.ReadMovement(data),
		GroundTrackStatus: fields.ReadGroundTrackStatus(data),
		GroundTrack:       fields.ReadGroundTrack(data),
		Time:              fields.ReadTime(data),
		CPRFormat:         fields.ReadCompactPositionReportingFormat(data),
		EncodedLatitude:   fields.ReadEncodedLatitude(data),
		EncodedLongitude:  fields.ReadEncodedLongitude(data),
	}, nil
}
