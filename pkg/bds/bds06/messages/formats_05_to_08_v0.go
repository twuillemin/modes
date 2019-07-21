package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds06/fields"
)

// Format05To08V0 is a message at the format BDS 0,6
type Format05To08V0 struct {
	TypeCode                  byte
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

// GetName returns the name of the message
func (message *Format05To08V0) GetName() string {
	return bds06Name
}

// GetBDS returns the binary data format
func (message *Format05To08V0) GetBDS() string {
	return bds06Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format05To08V0) GetFormatTypeCode() byte {
	return message.TypeCode
}

// GetMovement returns the Movement
func (message *Format05To08V0) GetMovement() fields.Movement {
	return message.Movement
}

// GetGroundTrackStatus returns the GroundTrackStatus
func (message *Format05To08V0) GetGroundTrackStatus() fields.GroundTrackStatus {
	return message.GroundTrackStatus
}

// GetGroundTrack returns the GroundTrack
func (message *Format05To08V0) GetGroundTrack() fields.GroundTrack {
	return message.GroundTrack
}

// GetTime returns the Time
func (message *Format05To08V0) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message *Format05To08V0) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message *Format05To08V0) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message *Format05To08V0) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalProtectionLimit returns the HorizontalProtectionLimit
func (message *Format05To08V0) GetHorizontalProtectionLimit() fields.HorizontalProtectionLimit {
	return message.HorizontalProtectionLimit
}

// GetContainmentRadius returns the ContainmentRadius
func (message *Format05To08V0) GetContainmentRadius() fields.ContainmentRadius {
	return message.ContainmentRadius
}

var hplByFormat = map[byte]fields.HorizontalProtectionLimit{
	5: fields.HPLLowerThan7Point5M,
	6: fields.HPLLowerThan25M,
	7: fields.HPLLowerThan185Point2M,
	8: fields.HPLGreaterThan185Point2M,
}

var crByFormat = map[byte]fields.ContainmentRadius{
	5: fields.CRLowerThan3M,
	6: fields.CRBetween3MAnd10M,
	7: fields.CRBetween10MAnd92Point6M,
	8: fields.CRGreaterThan92Point6M,
}

// ToString returns a basic, but readable, representation of the message
func (message *Format05To08V0) ToString() string {
	return fmt.Sprintf("Message:                           %v - %v (%v)\n"+
		"Horizontal Protection Limit:       %v\n"+
		"Containment Radius:                %v\n"+
		"Movement:                          %v\n"+
		"Ground Track Status:               %v\n"+
		"Ground Track:                      %v\n"+
		"Time:                              %v\n"+
		"Compact Position Reporting Format: %v\n"+
		"Encoded Latitude:                  %v\n"+
		"Encoded Longitude:                 %v",
		message.GetFormatTypeCode(),
		message.GetName(),
		message.GetBDS(),
		message.GetHorizontalProtectionLimit().ToString(),
		message.GetContainmentRadius().ToString(),
		message.GetMovement().ToString(),
		message.GetGroundTrackStatus(),
		message.GetGroundTrack(),
		message.GetTime().ToString(),
		message.GetCPRFormat().ToString(),
		message.GetEncodedLatitude(),
		message.GetEncodedLongitude())
}

// readFormat05To08V0 reads a message at the format BDS 0,6
func readFormat05To08V0(data []byte) (*Format05To08V0, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3

	return &Format05To08V0{
		TypeCode:                  formatTypeCode,
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
