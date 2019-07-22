package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds06/fields"
)

// MessageBDS06V0 is the basic interface that ADSB messages at the format BDS 0,6 Version 0 are expected to implement
type MessageBDS06V0 interface {
	MessageBDS06

	// GetHorizontalProtectionLimit returns the HorizontalProtectionLimit
	GetHorizontalProtectionLimit() fields.HorizontalProtectionLimit
	// GetContainmentRadius returns the ContainmentRadius
	GetContainmentRadius() fields.ContainmentRadius
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

// MessageBDS06V0ToString returns a basic, but readable, representation of the message
func messageBDS06V0ToString(message MessageBDS06V0) string {
	return fmt.Sprintf("Message:                           %v\n"+
		"Horizontal Protection Limit:       %v\n"+
		"Containment Radius:                %v\n"+
		"Movement:                          %v\n"+
		"Ground Track Status:               %v\n"+
		"Ground Track:                      %v\n"+
		"Time:                              %v\n"+
		"Compact Position Reporting Format: %v\n"+
		"Encoded Latitude:                  %v\n"+
		"Encoded Longitude:                 %v",
		message.GetMessageFormat().ToString(),
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
