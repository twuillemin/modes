package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
)

// MessageBDS05V2 is the basic interface that ADSB messages at the format BDS 0,5  for ADSB v2 are expected to implement
type MessageBDS05V2 interface {
	MessageBDS05
	// GetNavigationIntegrityCodeSupplementB returns the NavigationIntegrityCodeSupplementB
	GetNavigationIntegrityCodeSupplementB() fields.NavigationIntegrityCodeSupplementB
	// GetHorizontalContainmentRadius returns the HorizontalContainmentRadius
	GetHorizontalContainmentRadius() fields.HorizontalContainmentRadius
	// GetNavigationIntegrityCategory returns the Navigation Integrity Category
	GetNavigationIntegrityCategory() byte
}

func bds05v2ToString(message MessageBDS05V2) string {
	return fmt.Sprintf("Message:                           %v\n"+
		"Surveillance Status:               %v\n"+
		"Horizontal Containment Radius      %v\n"+
		"Navigation Integrity Category      %v\n"+
		"NIC Supplement B:                  %v\n"+
		"Time:                              %v\n"+
		"Compact Position Reporting Format: %v\n"+
		"Altitude:                          %v\n"+
		"Encoded Latitude:                  %v\n"+
		"Encoded Longitude:                 %v",
		adsb.GetMessageFormatInformation(message),
		message.GetSurveillanceStatus().ToString(),
		message.GetHorizontalContainmentRadius().ToString(),
		message.GetNavigationIntegrityCategory(),
		message.GetNavigationIntegrityCodeSupplementB().ToString(),
		message.GetTime().ToString(),
		message.GetCPRFormat().ToString(),
		message.GetAltitude().ToString(),
		message.GetEncodedLatitude(),
		message.GetEncodedLongitude())
}

func getHCRAndNICForV2Barometric(
	formatTypeCode byte,
	nicSupplementA bool,
	nicSupplementB fields.NavigationIntegrityCodeSupplementB) (fields.HorizontalContainmentRadiusBarometricV2, byte) {

	switch formatTypeCode {
	case 9:
		return fields.HCRBV2RcLowerThan7Point5M, 11
	case 10:
		return fields.HCRBV2RcLowerThan25M, 10
	case 11:
		if nicSupplementA {
			return fields.HCRBV2RcLowerThan75M, 9
		} else {
			return fields.HCRBV2RcLowerThan0Point1NM, 8
		}
	case 12:
		return fields.HCRBV2RcLowerThan7Point5M, 7
	case 13:
		if !nicSupplementA {
			if nicSupplementB == fields.NICBOne {
				return fields.HCRBV2RcLowerThan0Point3NM, 6
			} else {
				return fields.HCRBV2RcLowerThan0Point5NM, 6
			}
		} else {
			return fields.HCRBV2RcLowerThan0Point6NM, 6
		}
	case 14:
		return fields.HCRBV2RcLowerThan1Point0NM, 5
	case 15:
		return fields.HCRBV2RcLowerThan2NM, 4
	case 16:
		if nicSupplementA {
			return fields.HCRBV2RcLowerThan4NM, 3
		} else {
			return fields.HCRBV2RcLowerThan8NM, 2
		}
	case 17:
		return fields.HCRBV2RcLowerThan20NM, 1
	default:
		return fields.HCRBV2RcGreaterThan20NM, 0
	}
}

func getHCRAndNICForV2GNSS(formatTypeCode byte) (fields.HorizontalContainmentRadiusGNSSV2, byte) {
	switch formatTypeCode {
	case 20:
		return fields.HCRGV2RcLowerThan7Point5M, 11
	case 21:
		return fields.HCRGV2RcLowerThan25M, 10
	default:
		return fields.HCRGV2RcGreaterThan25MOrUnknown, 0
	}
}
