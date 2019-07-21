package messages

import (
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
)

// Format09To18V2 is a message at the format BDS 0,5 for ADSB V2
type Format09To18V2 struct {
	TypeCode                           byte
	SurveillanceStatus                 fields.SurveillanceStatus
	NavigationIntegrityCodeSupplementB fields.NavigationIntegrityCodeSupplementB
	Altitude                           fields.Altitude
	Time                               fields.Time
	CPRFormat                          fields.CompactPositionReportingFormat
	EncodedLatitude                    fields.EncodedLatitude
	EncodedLongitude                   fields.EncodedLongitude
	HorizontalContainmentRadius        fields.HorizontalContainmentRadiusBarometricV2
	NavigationIntegrityCategory        byte
}

// GetName returns the name of the message
func (message *Format09To18V2) GetName() string {
	return bds05Name
}

// GetBDS returns the binary data format
func (message *Format09To18V2) GetBDS() string {
	return bds05Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format09To18V2) GetFormatTypeCode() byte {
	return message.TypeCode
}

// ToString returns a basic, but readable, representation of the message
func (message *Format09To18V2) ToString() string {
	return bds05v2ToString(message)
}

// GetSurveillanceStatus returns the Surveillance Status
func (message *Format09To18V2) GetSurveillanceStatus() fields.SurveillanceStatus {
	return message.SurveillanceStatus
}

// GetNavigationIntegrityCodeSupplementB returns the NavigationIntegrityCodeSupplementB
func (message *Format09To18V2) GetNavigationIntegrityCodeSupplementB() fields.NavigationIntegrityCodeSupplementB {
	return message.NavigationIntegrityCodeSupplementB
}

// GetAltitude returns the Altitude
func (message *Format09To18V2) GetAltitude() fields.Altitude {
	return message.Altitude
}

// GetTime returns the Time
func (message *Format09To18V2) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message *Format09To18V2) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message *Format09To18V2) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message *Format09To18V2) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalContainmentRadius returns the HorizontalContainmentRadius
func (message *Format09To18V2) GetHorizontalContainmentRadius() fields.HorizontalContainmentRadius {
	return message.HorizontalContainmentRadius
}

// GetNavigationIntegrityCategory returns the Navigation Integrity Category
func (message *Format09To18V2) GetNavigationIntegrityCategory() byte {
	return 0
}

func getHCRAndNICForV2Baro(
	formatTypeCode byte,
	nicSupplementA bool,
	nicb fields.NavigationIntegrityCodeSupplementB) (fields.HorizontalContainmentRadiusBarometricV2, byte) {

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
			if nicb == fields.NICBOne {
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

// readFormat09To18V2 reads a message at the format BDS 0,5
func readFormat09To18V2(nicSupplementA bool, data []byte) (*Format09To18V2, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3

	nicb := fields.ReadNavigationIntegritySupplementB(data)
	hcr, nic := getHCRAndNICForV2Baro(formatTypeCode, nicSupplementA, nicb)

	return &Format09To18V2{
		TypeCode:                           formatTypeCode,
		SurveillanceStatus:                 fields.ReadSurveillanceStatus(data),
		NavigationIntegrityCodeSupplementB: nicb,
		Altitude:                           fields.ReadAltitude(data),
		Time:                               fields.ReadTime(data),
		CPRFormat:                          fields.ReadCompactPositionReportingFormat(data),
		EncodedLatitude:                    fields.ReadEncodedLatitude(data),
		EncodedLongitude:                   fields.ReadEncodedLongitude(data),
		HorizontalContainmentRadius:        hcr,
		NavigationIntegrityCategory:        nic,
	}, nil
}
