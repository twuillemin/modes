package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds06/fields"
)

// Format05To08V1 is a message at the format BDS 0,6
type Format05To08V1 struct {
	TypeCode                      byte
	Movement                      fields.Movement
	GroundTrackStatus             fields.GroundTrackStatus
	GroundTrack                   fields.GroundTrack
	Time                          fields.Time
	CPRFormat                     fields.CompactPositionReportingFormat
	EncodedLatitude               fields.EncodedLatitude
	EncodedLongitude              fields.EncodedLongitude
	HorizontalContainmentRadiusV1 fields.HorizontalContainmentRadiusV1
	NavigationIntegrityCategory   byte
}

// GetName returns the name of the message
func (message *Format05To08V1) GetName() string {
	return bds06Name
}

// GetBDS returns the binary data format
func (message *Format05To08V1) GetBDS() string {
	return bds06Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format05To08V1) GetFormatTypeCode() byte {
	return message.TypeCode
}

// GetMovement returns the Movement
func (message *Format05To08V1) GetMovement() fields.Movement {
	return message.Movement
}

// GetGroundTrackStatus returns the GroundTrackStatus
func (message *Format05To08V1) GetGroundTrackStatus() fields.GroundTrackStatus {
	return message.GroundTrackStatus
}

// GetGroundTrack returns the GroundTrack
func (message *Format05To08V1) GetGroundTrack() fields.GroundTrack {
	return message.GroundTrack
}

// GetTime returns the Time
func (message *Format05To08V1) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message *Format05To08V1) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message *Format05To08V1) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message *Format05To08V1) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalContainmentRadius returns the HorizontalContainmentRadius
func (message *Format05To08V1) GetHorizontalContainmentRadius() fields.HorizontalContainmentRadius {
	return message.HorizontalContainmentRadiusV1
}

// GetNavigationIntegrityCategory returns the Navigation Integrity Category
func (message *Format05To08V1) GetNavigationIntegrityCategory() byte {
	return 0
}

// ToString returns a basic, but readable, representation of the message
func (message *Format05To08V1) ToString() string {
	return fmt.Sprintf("Message:                           %v (%v)\n"+
		"FormatTypeCode:                    %v\n"+
		"Horizontal Containment Radius:     %v\n"+
		"Navigation Integrity Category      %v\n"+
		"Movement:                          %v\n"+
		"Ground Track Status:               %v\n"+
		"Ground Track:                      %v\n"+
		"Time:                              %v\n"+
		"Compact Position Reporting Format: %v\n"+
		"Encoded Latitude:                  %v\n"+
		"Encoded Longitude:                 %v",
		message.GetBDS(),
		message.GetName(),
		message.GetFormatTypeCode(),
		message.GetHorizontalContainmentRadius().ToString(),
		message.GetNavigationIntegrityCategory(),
		message.GetMovement().ToString(),
		message.GetGroundTrackStatus(),
		message.GetGroundTrack(),
		message.GetTime().ToString(),
		message.GetCPRFormat().ToString(),
		message.GetEncodedLatitude(),
		message.GetEncodedLongitude())
}

func getHCRAndNICForV1(formatTypeCode byte, nicSupplementA bool) (fields.HorizontalContainmentRadiusV1, byte) {
	switch formatTypeCode {
	case 5:
		return fields.HCRV1RcLowerThan7Point5M, 11
	case 6:
		return fields.HCRV1RcLowerThan25M, 10
	case 7:
		if nicSupplementA {
			return fields.HCRV1RcLowerThan75M, 9
		} else {
			return fields.HCRV1RcGreaterThan0Point1NM, 8
		}
	default:
		return fields.HCRV1RcGreaterThan0Point1NM, 0
	}
}

// readFormat05To08V1 reads a message at the format BDS 0,6
func readFormat05To08V1(nicSupplementA bool, data []byte) (*Format05To08V1, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3

	hcr, nic := getHCRAndNICForV1(formatTypeCode, nicSupplementA)

	return &Format05To08V1{
		TypeCode:                      formatTypeCode,
		Movement:                      fields.ReadMovement(data),
		GroundTrackStatus:             fields.ReadGroundTrackStatus(data),
		GroundTrack:                   fields.ReadGroundTrack(data),
		Time:                          fields.ReadTime(data),
		CPRFormat:                     fields.ReadCompactPositionReportingFormat(data),
		EncodedLatitude:               fields.ReadEncodedLatitude(data),
		EncodedLongitude:              fields.ReadEncodedLongitude(data),
		HorizontalContainmentRadiusV1: hcr,
		NavigationIntegrityCategory:   nic,
	}, nil
}
