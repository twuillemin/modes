package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds06/fields"
	"github.com/twuillemin/modes/pkg/bds/common"
)

// MessageBDS06 is the basic interface that ADSB messages at the format BDS 0,6 are expected to implement
type MessageBDS06 interface {
	common.BDSMessage
	// GetFormatTypeCode returns the Format Type Code
	GetFormatTypeCode() byte
	// GetMovement returns the Movement
	GetMovement() fields.Movement
	// GetGroundTrackStatus returns the GroundTrackStatus
	GetGroundTrackStatus() fields.GroundTrackStatus
	// GetGroundTrack returns the GroundTrack
	GetGroundTrack() fields.GroundTrack
	// GetTime returns the Time
	GetTime() fields.Time
	// GetCPRFormat returns the CompactPositionReportingFormat
	GetCPRFormat() fields.CompactPositionReportingFormat
	// GetEncodedLatitude returns the EncodedLatitude
	GetEncodedLatitude() fields.EncodedLatitude
	// GetEncodedLongitude returns the EncodedLongitude
	GetEncodedLongitude() fields.EncodedLongitude
	// GetHorizontalProtectionLimit returns the HorizontalProtectionLimit
	GetHorizontalProtectionLimit() fields.HorizontalProtectionLimit
	// GetContainmentRadius returns the ContainmentRadius
	GetContainmentRadius() fields.ContainmentRadius
}

var bds06Code = "BDS 0,6"
var bds06Name = "Extended squitter surface position"

func bds06ToString(message MessageBDS06) string {
	return fmt.Sprintf("Message:                           %v (%v)\n"+
		"FormatTypeCode:                    %v\n"+
		"Horizontal Protection Limit:       %v\n"+
		"Containment Radius:                %v\n"+
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

// ReadBDS06 reads a message at the format BDS 0,6
func ReadBDS06(data []byte) (MessageBDS06, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for BDS message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3

	switch formatTypeCode {
	case 5:
		return ReadFormat05(data)
	case 6:
		return ReadFormat06(data)
	case 7:
		return ReadFormat07(data)
	case 8:
		return ReadFormat08(data)
	}

	return nil, fmt.Errorf("the format type code %v can not be read as a BDS 0,6 format", formatTypeCode)
}
