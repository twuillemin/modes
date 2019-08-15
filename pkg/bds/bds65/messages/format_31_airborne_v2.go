package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
	"github.com/twuillemin/modes/pkg/common"
)

// Format31AirborneV2 is a message at the format BDS 6,5 the ADSB V2 / SubtypeAirborne
//
// Specified in Doc 9871 / C.2.3.10
type Format31AirborneV2 struct {
	AirborneCapabilityClass              fields.AirborneCapabilityClassV2
	OperationalMode                      fields.AirborneOperationalMode
	VersionNumber                        fields.VersionNumber
	NICSupplementA                       fields.NICSupplementA
	NavigationalAccuracyCategoryPosition fields.NavigationalAccuracyCategoryPositionV2
	GeometricVerticalAccuracy            fields.GeometricVerticalAccuracy
	SourceIntegrityLevel                 fields.SourceIntegrityLevel
	NICBaro                              fields.NICBaro
	HorizontalReferenceDirection         fields.HorizontalReferenceDirection
	SourceIntegrityLevelSupplement       fields.SourceIntegrityLevelSupplement
}

// GetMessageFormat returns the ADSB format of the message
func (message Format31AirborneV2) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format31
}

// GetRegister returns the register of the message
func (message Format31AirborneV2) GetRegister() bds.Register {
	return adsb.Format31.GetRegister()
}

// GetSubtype returns the Subtype
func (message Format31AirborneV2) GetSubtype() adsb.Subtype {
	return fields.SubtypeAirborne
}

// GetMinimumADSBLevel returns the minimum ADSB ReaderLevel for the message
func (message Format31AirborneV2) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel2
}

// GetMaximumADSBLevel returns the maximum ADSB ReaderLevel for the message
func (message Format31AirborneV2) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel2
}

// ToString returns a basic, but readable, representation of the message
func (message Format31AirborneV2) ToString() string {
	return fmt.Sprintf("Message:                                 %v\n"+
		"Subtype:                                 %v\n"+
		"SubtypeAirborne Capability Class:\n%v\n"+
		"Operational Mode:\n%v\n"+
		"ADSB Version Number:                     %v\n"+
		"NIC Supplement:                          %v\n"+
		"Navigational Accuracy Category Position: %v\n"+
		"Geometric Vertical Accuracy:             %v\n"+
		"Source Integrity ReaderLevel:                  %v\n"+
		"Source Integrity ReaderLevel Supplement:       %v\n"+
		"NIC Baro:                                %v\n"+
		"Horizontal Reference Direction:          %v",
		adsb.GetMessageFormatInformation(&message),
		message.GetSubtype().ToString(),
		common.PrefixMultiLine(message.AirborneCapabilityClass.ToString(), "    "),
		common.PrefixMultiLine(message.OperationalMode.ToString(), "    "),
		message.VersionNumber.ToString(),
		message.NICSupplementA.ToString(),
		message.NavigationalAccuracyCategoryPosition.ToString(),
		message.GeometricVerticalAccuracy.ToString(),
		message.SourceIntegrityLevel.ToString(),
		message.SourceIntegrityLevelSupplement.ToString(),
		message.NICBaro.ToString(),
		message.HorizontalReferenceDirection.ToString())
}

// ReadFormat31AirborneV2 reads a message at the format Format31 / subtype 0 (Airborne) for ADSB V2
func ReadFormat31AirborneV2(data []byte) (*Format31AirborneV2, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format31.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read by ReadFormat31AirborneV2", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeAirborne {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadFormat31AirborneV2", subType.ToString())
	}

	// Check the ADSB ReaderLevel
	detectedADSBLevel := fields.ReadVersionNumber(data)
	if detectedADSBLevel != fields.ADSBVersion2 {
		return nil, fmt.Errorf("the data are given at at %v format and can not be read by ReadFormat31AirborneV2", detectedADSBLevel.ToString())
	}

	capabilityContent := (data[1] & 0xC0) >> 6
	if capabilityContent != 0 {
		return nil, fmt.Errorf("the Capability Content (field Capability Class) must be 0 (%v given)", capabilityContent)
	}

	operationalModeFormat := (data[3] & 0xC0) >> 6
	if operationalModeFormat != 0 {
		return nil, fmt.Errorf("the Operational Mode Format (field Operational Mode) must be 0 (%v given)", operationalModeFormat)
	}

	return &Format31AirborneV2{
		AirborneCapabilityClass:              fields.ReadAirborneCapabilityClassV2(data),
		OperationalMode:                      fields.ReadAirborneOperationalMode(data),
		VersionNumber:                        fields.ReadVersionNumber(data),
		NICSupplementA:                       fields.ReadNICSupplementA(data),
		NavigationalAccuracyCategoryPosition: fields.ReadNavigationalAccuracyCategoryPositionV2(data),
		GeometricVerticalAccuracy:            fields.ReadGeometricVerticalAccuracy(data),
		SourceIntegrityLevel:                 fields.ReadSourceIntegrityLevel(data),
		NICBaro:                              fields.ReadNICBaro(data),
		HorizontalReferenceDirection:         fields.ReadHorizontalReferenceDirection(data),
		SourceIntegrityLevelSupplement:       fields.ReadSourceIntegrityLevelSupplement(data),
	}, nil
}
