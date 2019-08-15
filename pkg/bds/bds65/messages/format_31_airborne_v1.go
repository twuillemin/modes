package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
	"github.com/twuillemin/modes/pkg/common"
)

// Format31AirborneV1 is a message at the format BDS 6,5 the ADSB V1 / SubtypeAirborne
//
// Specified in Doc 9871 / B.2.3.10
type Format31AirborneV1 struct {
	AirborneCapabilityClass              fields.AirborneCapabilityClassV1
	OperationalMode                      fields.OperationalMode
	VersionNumber                        fields.VersionNumber
	NICSupplement                        fields.NICSupplementA
	NavigationalAccuracyCategoryPosition fields.NavigationalAccuracyCategoryPositionV1
	SurveillanceIntegrityLevel           fields.SurveillanceIntegrityLevel
	NICBaro                              fields.NICBaro
	HorizontalReferenceDirection         fields.HorizontalReferenceDirection
}

// GetMessageFormat returns the ADSB format of the message
func (message Format31AirborneV1) GetMessageFormat() adsb.MessageFormat {
	return adsb.Format31
}

// GetRegister returns the register of the message
func (message Format31AirborneV1) GetRegister() bds.Register {
	return adsb.Format31.GetRegister()
}

// GetSubtype returns the Subtype
func (message Format31AirborneV1) GetSubtype() adsb.Subtype {
	return fields.SubtypeAirborne
}

// GetMinimumADSBLevel returns the minimum ADSB ReaderLevel for the message
func (message Format31AirborneV1) GetMinimumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel1
}

// GetMaximumADSBLevel returns the maximum ADSB ReaderLevel for the message
func (message Format31AirborneV1) GetMaximumADSBLevel() adsb.MessageLevel {
	return adsb.MessageLevel1
}

// ToString returns a basic, but readable, representation of the message
func (message Format31AirborneV1) ToString() string {
	return fmt.Sprintf("Message:                                 %v\n"+
		"Subtype:                                 %v\n"+
		"SubtypeAirborne Capability Class:\n%v\n"+
		"Operational Mode:\n%v\n"+
		"ADSB Version Number:                     %v\n"+
		"NIC Supplement:                          %v\n"+
		"Navigational Accuracy Category Position: %v\n"+
		"Surveillance Integrity ReaderLevel:\n%v\n"+
		"NIC Baro:                                %v\n"+
		"Horizontal Reference Direction:          %v",
		adsb.GetMessageFormatInformation(&message),
		message.GetSubtype().ToString(),
		common.PrefixMultiLine(message.AirborneCapabilityClass.ToString(), "    "),
		common.PrefixMultiLine(message.OperationalMode.ToString(), "    "),
		message.VersionNumber.ToString(),
		message.NICSupplement.ToString(),
		message.NavigationalAccuracyCategoryPosition.ToString(),
		common.PrefixMultiLine(message.SurveillanceIntegrityLevel.ToString(), "    "),
		message.NICBaro.ToString(),
		message.HorizontalReferenceDirection.ToString())
}

// ReadFormat31AirborneV1 reads a message at the format Format31 / subtype 0 (Airborne) for ADSB V1
func ReadFormat31AirborneV1(data []byte) (*Format31AirborneV1, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != adsb.Format31.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read by ReadFormat31AirborneV1", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeAirborne {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadFormat31AirborneV1", subType.ToString())
	}

	// Check the ADSB ReaderLevel
	detectedADSBLevel := fields.ReadVersionNumber(data)
	if detectedADSBLevel != fields.ADSBVersion1 {
		return nil, fmt.Errorf("the data are given at %v format and can not be read by ReadFormat31AirborneV1", detectedADSBLevel.ToString())
	}

	serviceLevel := (data[1]&0xC0)>>4 + (data[1]&0x0C)>>2
	if serviceLevel != 0 {
		return nil, fmt.Errorf("the ServiceLevel (field Capability Class) must be 0 (%v given)", serviceLevel)
	}

	operationalModeFormat := (data[3] & 0xC0) >> 6
	if operationalModeFormat != 0 {
		return nil, fmt.Errorf("the Operational Mode Format (field Operational Mode) must be 0 (%v given)", operationalModeFormat)
	}

	return &Format31AirborneV1{
		AirborneCapabilityClass:              fields.ReadAirborneCapabilityClassV1(data),
		OperationalMode:                      fields.ReadOperationalMode(data),
		VersionNumber:                        fields.ReadVersionNumber(data),
		NICSupplement:                        fields.ReadNICSupplementA(data),
		NavigationalAccuracyCategoryPosition: fields.ReadNavigationalAccuracyCategoryPositionV1(data),
		SurveillanceIntegrityLevel:           fields.ReadSurveillanceIntegrityLevel(data),
		NICBaro:                              fields.ReadNICBaro(data),
		HorizontalReferenceDirection:         fields.ReadHorizontalReferenceDirection(data),
	}, nil
}
