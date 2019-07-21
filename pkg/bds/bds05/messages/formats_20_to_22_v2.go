package messages

import (
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
)

// Format20To22V2 is a message at the format BDS 0,5 for ADSB v2
type Format20To22V2 struct {
	TypeCode                           byte
	SurveillanceStatus                 fields.SurveillanceStatus
	NavigationIntegrityCodeSupplementB fields.NavigationIntegrityCodeSupplementB
	Altitude                           fields.Altitude
	Time                               fields.Time
	CPRFormat                          fields.CompactPositionReportingFormat
	EncodedLatitude                    fields.EncodedLatitude
	EncodedLongitude                   fields.EncodedLongitude
	HorizontalContainmentRadius        fields.HorizontalContainmentRadiusGNSSV2
	NavigationIntegrityCategory        byte
}

// GetName returns the name of the message
func (message *Format20To22V2) GetName() string {
	return bds05Name
}

// GetBDS returns the binary data format
func (message *Format20To22V2) GetBDS() string {
	return bds05Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format20To22V2) GetFormatTypeCode() byte {
	return message.TypeCode
}

// ToString returns a basic, but readable, representation of the message
func (message *Format20To22V2) ToString() string {
	return bds05v2ToString(message)
}

// GetSurveillanceStatus returns the Surveillance Status
func (message *Format20To22V2) GetSurveillanceStatus() fields.SurveillanceStatus {
	return message.SurveillanceStatus
}

// GetNavigationIntegrityCodeSupplementB returns the NavigationIntegrityCodeSupplementB
func (message *Format20To22V2) GetNavigationIntegrityCodeSupplementB() fields.NavigationIntegrityCodeSupplementB {
	return message.NavigationIntegrityCodeSupplementB
}

// GetAltitude returns the Altitude
func (message *Format20To22V2) GetAltitude() fields.Altitude {
	return message.Altitude
}

// GetTime returns the Time
func (message *Format20To22V2) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message *Format20To22V2) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message *Format20To22V2) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message *Format20To22V2) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalContainmentRadius returns the HorizontalContainmentRadius
func (message *Format20To22V2) GetHorizontalContainmentRadius() fields.HorizontalContainmentRadius {
	return message.HorizontalContainmentRadius
}

// GetNavigationIntegrityCategory returns the Navigation Integrity Category
func (message *Format20To22V2) GetNavigationIntegrityCategory() byte {
	return 0
}

func getHCRAndNICForv2GNSS(formatTypeCode byte) (fields.HorizontalContainmentRadiusGNSSV2, byte) {
	switch formatTypeCode {
	case 20:
		return fields.HCRGV2RcLowerThan7Point5M, 11
	case 21:
		return fields.HCRGV2RcLowerThan25M, 10
	default:
		return fields.HCRGV2RcGreaterThan25MOrUnknown, 0
	}
}

// readFormat20To22V2 reads a message at the format BDS 0,5
func readFormat20To22V2(data []byte) (*Format20To22V2, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3

	nicb := fields.ReadNavigationIntegritySupplementB(data)
	hcr, nic := getHCRAndNICForv2GNSS(formatTypeCode)

	return &Format20To22V2{
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
