package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds06/fields"
)

// MessageBDS06V2 is the basic interface that ADSB messages at the format BDS 0,6 Version 2 are expected to implement
type MessageBDS06V2 interface {
	MessageBDS06

	// GetHorizontalContainmentRadius returns the HorizontalContainmentRadius
	GetHorizontalContainmentRadius() fields.HorizontalContainmentRadiusV2
	// GetNavigationIntegrityCategory returns the Navigation Integrity Category
	GetNavigationIntegrityCategory() byte
}

func getHCRAndNICForV2(formatTypeCode byte, nicSupplementA bool, nicSupplementC bool) (fields.HorizontalContainmentRadiusV2, byte) {
	switch formatTypeCode {
	case 5:
		return fields.HCRV2RcLowerThan7Point5M, 11
	case 6:
		return fields.HCRV2RcLowerThan25M, 10
	case 7:
		if nicSupplementA {
			return fields.HCRV2RcLowerThan75M, 9
		} else {
			return fields.HCRV2RcLowerThan0Point1NM, 8
		}
	default:
		if nicSupplementA {
			if nicSupplementC {
				return fields.HCRV2RcLowerThan0Point2NM, 7
			} else {
				return fields.HCRV2RcLowerThan0Point3NM, 6
			}
		} else {
			if nicSupplementA {
				return fields.HCRV2RcLowerThan0Point6NM, 6
			} else {
				return fields.HCRV2RcGreaterThan0Point6NM, 0
			}
		}
	}
}

// messageBDS06V2ToString returns a basic, but readable, representation of the message
func messageBDS06V2ToString(message MessageBDS06V2) string {
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
