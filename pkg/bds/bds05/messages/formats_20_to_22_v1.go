package messages

import (
	"github.com/twuillemin/modes/pkg/bds/bds05/fields"
)

// Format20To22V1 is a message at the format BDS 0,5 for ADSB V1
type Format20To22V1 struct {
	TypeCode                    byte
	SurveillanceStatus          fields.SurveillanceStatus
	SingleAntennaFlag           fields.SingleAntennaFlag
	Altitude                    fields.Altitude
	Time                        fields.Time
	CPRFormat                   fields.CompactPositionReportingFormat
	EncodedLatitude             fields.EncodedLatitude
	EncodedLongitude            fields.EncodedLongitude
	HorizontalContainmentRadius fields.HorizontalContainmentRadiusGNSSV1
	NavigationIntegrityCategory byte
}

// GetName returns the name of the message
func (message *Format20To22V1) GetName() string {
	return bds05Name
}

// GetRegister returns the binary data format
func (message *Format20To22V1) GetBDS() string {
	return bds05Code
}

// GetFormatTypeCode returns the Format Type Code
func (message *Format20To22V1) GetFormatTypeCode() byte {
	return message.TypeCode
}

// ToString returns a basic, but readable, representation of the message
func (message *Format20To22V1) ToString() string {
	return bds05v1ToString(message)
}

// GetSurveillanceStatus returns the Surveillance Status
func (message *Format20To22V1) GetSurveillanceStatus() fields.SurveillanceStatus {
	return message.SurveillanceStatus
}

// GetSingleAntennaFlag returns the SingleAntennaFlag
func (message *Format20To22V1) GetSingleAntennaFlag() fields.SingleAntennaFlag {
	return message.SingleAntennaFlag
}

// GetAltitude returns the Altitude
func (message *Format20To22V1) GetAltitude() fields.Altitude {
	return message.Altitude
}

// GetTime returns the Time
func (message *Format20To22V1) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message *Format20To22V1) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message *Format20To22V1) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message *Format20To22V1) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalContainmentRadius returns the HorizontalContainmentRadius
func (message *Format20To22V1) GetHorizontalContainmentRadius() fields.HorizontalContainmentRadius {
	return message.HorizontalContainmentRadius
}

// GetNavigationIntegrityCategory returns the Navigation Integrity Category
func (message *Format20To22V1) GetNavigationIntegrityCategory() byte {
	return 0
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

// readFormat20To22V1 reads a message at the format BDS 0,5
func readFormat20To22V1(data []byte) (*Format20To22V1, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3

	hcr, nic := getHCRAndNICForV1GNSS(formatTypeCode)

	return &Format20To22V1{
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
