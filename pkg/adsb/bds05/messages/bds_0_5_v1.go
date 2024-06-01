package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/bds05/fields"
)

// MessageBDS05V1 is the basic interface that ADSB messages at the format BDS 0,5  for ADSB v1 are expected to implement
type MessageBDS05V1 interface {
	MessageBDS05
	GetSingleAntennaFlag() fields.SingleAntennaFlag
	// GetHorizontalContainmentRadius returns the HorizontalContainmentRadius
	GetHorizontalContainmentRadius() fields.HorizontalContainmentRadius
	// GetNavigationIntegrityCategory returns the Navigation Integrity Category
	GetNavigationIntegrityCategory() byte
}

func bds05v1ToString(message MessageBDS05V1) string {
	return fmt.Sprintf("Message:                           %v\n"+
		"Surveillance Status:               %v\n"+
		"Horizontal Containment Radius      %v\n"+
		"Navigation Integrity Category      %v\n"+
		"Single Antenna:                    %v\n"+
		"Time:                              %v\n"+
		"Compact Position Reporting Format: %v\n"+
		"Altitude:                          %v\n"+
		"Encoded Latitude:                  %v\n"+
		"Encoded Longitude:                 %v",
		adsb.GetMessageFormatInformation(message),
		message.GetSurveillanceStatus().ToString(),
		message.GetHorizontalContainmentRadius().ToString(),
		message.GetNavigationIntegrityCategory(),
		message.GetSingleAntennaFlag().ToString(),
		message.GetTime().ToString(),
		message.GetCPRFormat().ToString(),
		message.GetAltitude().ToString(),
		message.GetEncodedLatitude(),
		message.GetEncodedLongitude())
}

func getHCRAndNICForV1Barometric(formatTypeCode byte, nicSupplementA bool) (fields.HorizontalContainmentRadiusBarometricV1, byte) {
	switch formatTypeCode {
	case 9:
		return fields.HCRBV1RcLowerThan7Point5MAndVPLLowerThan11M, 11
	case 10:
		return fields.HCRBV1RcLowerThan25MAndVPLLowerThan37Point5M, 10
	case 11:
		if nicSupplementA {
			return fields.HCRBV1RcLowerThan75MAndVPLLowerThan112M, 9
		} else {
			return fields.HCRBV1RcLowerThan0Point1NM, 8
		}

	case 12:
		return fields.HCRBV1RcLowerThan7Point5MAndVPLLowerThan11M, 7
	case 13:
		if nicSupplementA {
			return fields.HCRBV1RcLowerThan0Point6NM, 6
		} else {
			return fields.HCRBV1RcLowerThan0Point5NM, 6
		}
	case 14:
		return fields.HCRBV1RcLowerThan1Point0NM, 5
	case 15:
		return fields.HCRBV1RcLowerThan2NM, 4
	case 16:
		if nicSupplementA {
			return fields.HCRBV1RcLowerThan4NM, 3
		} else {
			return fields.HCRBV1RcLowerThan8NM, 2
		}
	case 17:
		return fields.HCRBV1RcLowerThan20NM, 1
	default:
		return fields.HCRBV1RcGreaterThan20NM, 0
	}
}

func getHCRAndNICForV1GNSS(formatTypeCode byte) (fields.HorizontalContainmentRadiusGNSSV1, byte) {
	switch formatTypeCode {
	case 20:
		return fields.HCRGV1RcLowerThan7Point5MAndVPLLowerThan11M, 11
	case 21:
		return fields.HCRGV1RcLowerThan25MAndVPLLowerThan37Point5M, 10
	default:
		return fields.HCRGV1RcGreaterThan25MOrVPLGreaterThan37Point5MOrUnknown, 0
	}
}
