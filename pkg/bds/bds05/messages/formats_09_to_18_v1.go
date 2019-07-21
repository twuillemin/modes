package messages

import (
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
)

// Format09To18V1 is a message at the format BDS 0,5 for ADSB V1
type Format09To18V1 struct {
	TypeCode                    byte
	SurveillanceStatus          fields.SurveillanceStatus
	SingleAntennaFlag           fields.SingleAntennaFlag
	Altitude                    fields.Altitude
	Time                        fields.Time
	CPRFormat                   fields.CompactPositionReportingFormat
	EncodedLatitude             fields.EncodedLatitude
	EncodedLongitude            fields.EncodedLongitude
	HorizontalContainmentRadius fields.HorizontalContainmentRadiusBarometricV1
	NavigationIntegrityCategory byte
}

// GetName returns the name of the message
func (message *Format09To18V1) GetName() string {
	return bds05Name
}

// GetBDS returns the binary data format
func (message *Format09To18V1) GetBDS() string {
	return bds05Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format09To18V1) GetFormatTypeCode() byte {
	return message.TypeCode
}

// ToString returns a basic, but readable, representation of the message
func (message *Format09To18V1) ToString() string {
	return bds05v1ToString(message)
}

// GetSurveillanceStatus returns the Surveillance Status
func (message *Format09To18V1) GetSurveillanceStatus() fields.SurveillanceStatus {
	return message.SurveillanceStatus
}

// GetSingleAntennaFlag returns the SingleAntennaFlag
func (message *Format09To18V1) GetSingleAntennaFlag() fields.SingleAntennaFlag {
	return message.SingleAntennaFlag
}

// GetAltitude returns the Altitude
func (message *Format09To18V1) GetAltitude() fields.Altitude {
	return message.Altitude
}

// GetTime returns the Time
func (message *Format09To18V1) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message *Format09To18V1) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message *Format09To18V1) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message *Format09To18V1) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalContainmentRadius returns the HorizontalContainmentRadius
func (message *Format09To18V1) GetHorizontalContainmentRadius() fields.HorizontalContainmentRadius {
	return message.HorizontalContainmentRadius
}

// GetNavigationIntegrityCategory returns the Navigation Integrity Category
func (message *Format09To18V1) GetNavigationIntegrityCategory() byte {
	return 0
}

func getHCRAndNICForV1Baro(formatTypeCode byte, nicSupplementA bool) (fields.HorizontalContainmentRadiusBarometricV1, byte) {
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

// readFormat09To18V1 reads a message at the format BDS 0,5
func readFormat09To18V1(nicSupplementA bool, data []byte) (*Format09To18V1, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3

	hcr, nic := getHCRAndNICForV1Baro(formatTypeCode, nicSupplementA)

	return &Format09To18V1{
		TypeCode:                    formatTypeCode,
		SurveillanceStatus:          fields.ReadSurveillanceStatus(data),
		SingleAntennaFlag:           fields.ReadSingleAntennaFlag(data),
		Altitude:                    fields.ReadAltitude(data),
		Time:                        fields.ReadTime(data),
		CPRFormat:                   fields.ReadCompactPositionReportingFormat(data),
		EncodedLatitude:             fields.ReadEncodedLatitude(data),
		EncodedLongitude:            fields.ReadEncodedLongitude(data),
		HorizontalContainmentRadius: hcr,
		NavigationIntegrityCategory: nic,
	}, nil
}
