package bds65

import (
	"errors"
	"fmt"

	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
	"github.com/twuillemin/modes/pkg/common"
)

// AircraftOperationalStatusAirborneADSB2 is a message at the format BDS 6,5 the ADSB V2 / SubtypeAirborne
//
// Specified in Doc 9871 / C.2.3.10
type AircraftOperationalStatusAirborneADSB2 struct {
	FormatTypeCode                       byte
	Subtype                              fields.Subtype
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

func (message AircraftOperationalStatusAirborneADSB2) GetADSBLevel() byte {
	return 2
}

// GetRegister returns the Register the message
func (message AircraftOperationalStatusAirborneADSB2) GetRegister() register.Register {
	return register.BDS65
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message AircraftOperationalStatusAirborneADSB2) CheckCoherency() error {
	return nil
}

// ToString returns a basic, but readable, representation of the message
func (message AircraftOperationalStatusAirborneADSB2) ToString() string {
	return fmt.Sprintf(""+
		"Message:                                 %v\n"+
		"Subtype:                                 %v\n"+
		"SubtypeAirborne Capability Class:        \n%v\n"+
		"Operational Mode:                        \n%v\n"+
		"ADSB Version Number:                     %v\n"+
		"NIC Supplement:                          %v\n"+
		"Navigational Accuracy Category Position: %v\n"+
		"Geometric Vertical Accuracy:             %v\n"+
		"Source Integrity ReaderLevel:            %v\n"+
		"Source Integrity ReaderLevel Supplement: %v\n"+
		"NIC Baro:                                %v\n"+
		"Horizontal Reference Direction:          %v",
		message.GetRegister().ToString(),
		message.Subtype.ToString(),
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

// ReadAircraftOperationalStatusAirborneADSB2 reads a message at the format Format31 / subtype 0 (Airborne) for ADSB V2
func ReadAircraftOperationalStatusAirborneADSB2(data []byte) (*AircraftOperationalStatusAirborneADSB2, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != 31 {
		return nil, fmt.Errorf("the field FormatTypeCode must be 31, got %v", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeAirborne {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadAircraftOperationalStatusADSB0", subType)
	}

	// Check the ADSB ReaderLevel
	detectedADSBLevel := fields.ReadVersionNumber(data)
	if detectedADSBLevel != fields.ADSBVersion2 {
		return nil, fmt.Errorf("the data are given at at %v format and can not be read by ReadAircraftOperationalStatusAirborneADSB2", detectedADSBLevel.ToString())
	}

	capabilityContent := (data[1] & 0xC0) >> 6
	if capabilityContent != 0 {
		return nil, fmt.Errorf("the Capability Content (field Capability Class) must be 0 (%v given)", capabilityContent)
	}

	operationalModeFormat := (data[3] & 0xC0) >> 6
	if operationalModeFormat != 0 {
		return nil, fmt.Errorf("the Operational Mode Format (field Operational Mode) must be 0 (%v given)", operationalModeFormat)
	}

	if data[6]&0x03 != 0 {
		return nil, errors.New("the bits 55 and 56 must be 0")
	}

	return &AircraftOperationalStatusAirborneADSB2{
		FormatTypeCode:                       formatTypeCode,
		Subtype:                              subType,
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
