package bds65

import (
	"errors"
	"fmt"

	"github.com/twuillemin/modes/pkg/bds/bds65/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
	"github.com/twuillemin/modes/pkg/common"
)

// AircraftOperationalStatusSurfaceADSB1 is a message at the format BDS 6,5 the ADSB V1 / SubtypeSurface
//
// Specified in Doc 9871 / B.2.3.10
type AircraftOperationalStatusSurfaceADSB1 struct {
	FormatTypeCode                       byte
	Subtype                              fields.Subtype
	SurfaceCapabilityClass               fields.SurfaceCapabilityClassV1
	LengthAndWidth                       fields.LengthWidth
	OperationalMode                      fields.OperationalMode
	VersionNumber                        fields.VersionNumber
	NICSupplement                        fields.NICSupplementA
	NavigationalAccuracyCategoryPosition fields.NavigationalAccuracyCategoryPositionV1
	SurveillanceIntegrityLevel           fields.SurveillanceIntegrityLevel
	TrackAngleHeading                    fields.TrackAngleHeading
	HorizontalReferenceDirection         fields.HorizontalReferenceDirection
}

func (message AircraftOperationalStatusSurfaceADSB1) GetADSBLevel() byte {
	return 1
}

// GetRegister returns the Register the message
func (message AircraftOperationalStatusSurfaceADSB1) GetRegister() register.Register {
	return register.BDS65
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message AircraftOperationalStatusSurfaceADSB1) CheckCoherency() error {
	return nil
}

// ToString returns a basic, but readable, representation of the message
func (message AircraftOperationalStatusSurfaceADSB1) ToString() string {
	return fmt.Sprintf(""+
		"Message:                                 %v\n"+
		"Subtype:                                 %v\n"+
		"SubtypeAirborne Capability Class:        \n%v\n"+
		"Aircraft Length And Width                %v\n"+
		"Operational Mode:                        \n%v\n"+
		"ADSB Version Number:                     %v\n"+
		"NIC Supplement:                          %v\n"+
		"Navigational Accuracy Category Position: %v\n"+
		"Surveillance Integrity ReaderLevel:      \n%v\n"+
		"NIC Baro:                                %v\n"+
		"Horizontal Reference Direction:          %v",
		message.GetRegister().ToString(),
		message.Subtype.ToString(),
		common.PrefixMultiLine(message.SurfaceCapabilityClass.ToString(), "    "),
		message.LengthAndWidth.ToString(),
		common.PrefixMultiLine(message.OperationalMode.ToString(), "    "),
		message.VersionNumber.ToString(),
		message.NICSupplement.ToString(),
		message.NavigationalAccuracyCategoryPosition.ToString(),
		common.PrefixMultiLine(message.SurveillanceIntegrityLevel.ToString(), "    "),
		message.TrackAngleHeading.ToString(),
		message.HorizontalReferenceDirection.ToString())
}

// ReadAircraftOperationalStatusSurfaceADSB1 reads a message at the format Format31 / subtype 1 (Surface) for ADSB V1
func ReadAircraftOperationalStatusSurfaceADSB1(data []byte) (*AircraftOperationalStatusSurfaceADSB1, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != 31 {
		return nil, fmt.Errorf("the field FormatTypeCode must be 31, got %v", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeSurface {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadAircraftOperationalStatusSurfaceADSB1", subType.ToString())
	}

	// Check the ADSB ReaderLevel
	detectedADSBLevel := fields.ReadVersionNumber(data)
	if detectedADSBLevel != fields.ADSBVersion1 {
		return nil, fmt.Errorf("the data are given at %v format and can not be read by ReadAircraftOperationalStatusSurfaceADSB1", detectedADSBLevel.ToString())
	}

	serviceLevel := (data[1]&0xC0)>>4 + (data[1]&0x0C)>>2
	if serviceLevel != 0 {
		return nil, fmt.Errorf("the ServiceLevel (field Capability Class) must be 0 (%v given)", serviceLevel)
	}

	operationalModeFormat := (data[3] & 0xC0) >> 6
	if operationalModeFormat != 0 {
		return nil, fmt.Errorf("the Operational Mode Format (field Operational Mode) must be 0 (%v given)", operationalModeFormat)
	}

	if data[6]&0xC0 != 0 {
		return nil, errors.New("the bits 49 and 50 must be 0")
	}

	if data[6]&0x03 != 0 {
		return nil, errors.New("the bits 55 and 56 must be 0")
	}

	return &AircraftOperationalStatusSurfaceADSB1{
		FormatTypeCode:                       formatTypeCode,
		Subtype:                              subType,
		SurfaceCapabilityClass:               fields.ReadSurfaceCapabilityClassV1(data),
		LengthAndWidth:                       fields.ReadAircraftLengthAndWidth(data),
		OperationalMode:                      fields.ReadOperationalMode(data),
		VersionNumber:                        fields.ReadVersionNumber(data),
		NICSupplement:                        fields.ReadNICSupplementA(data),
		NavigationalAccuracyCategoryPosition: fields.ReadNavigationalAccuracyCategoryPositionV1(data),
		SurveillanceIntegrityLevel:           fields.ReadSurveillanceIntegrityLevel(data),
		TrackAngleHeading:                    fields.ReadTrackAngleHeading(data),
		HorizontalReferenceDirection:         fields.ReadHorizontalReferenceDirection(data),
	}, nil
}
