package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds06/fields"
)

// MessageBDS06V1 is the basic interface that ADSB messages at the format BDS 0,6 Version 1 are expected to implement
type MessageBDS06V1 interface {
	MessageBDS06

	// GetHorizontalContainmentRadius returns the HorizontalContainmentRadius
	GetHorizontalContainmentRadius() fields.HorizontalContainmentRadiusV1
	// GetNavigationIntegrityCategory returns the Navigation Integrity Category
	GetNavigationIntegrityCategory() byte
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

// messageBDS06V1ToString returns a basic, but readable, representation of the message
func messageBDS06V1ToString(message MessageBDS06V1) string {
	return fmt.Sprintf("Message:                           %v\n"+
		"Horizontal Containment Radius:     %v\n"+
		"Navigation Integrity Category      %v\n"+
		"Movement:                          %v\n"+
		"Ground Track Status:               %v\n"+
		"Ground Track:                      %v\n"+
		"Time:                              %v\n"+
		"Compact Position Reporting Format: %v\n"+
		"Encoded Latitude:                  %v\n"+
		"Encoded Longitude:                 %v",
		adsb.GetMessageFormatInformation(message),
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
