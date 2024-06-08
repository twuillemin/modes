package bds09

import (
	"errors"
	"fmt"
	"math"

	"github.com/twuillemin/modes/pkg/bds/bds09/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// AirborneVelocityGroundSpeedSupersonic is a message at the format BDS 9,0
type AirborneVelocityGroundSpeedSupersonic struct {
	FormatTypeCode                   byte
	Subtype                          fields.Subtype
	IntentChange                     fields.IntentChange
	IFRCapability                    fields.IFRCapability
	NavigationUncertaintyCategory    fields.NavigationUncertaintyCategory
	DirectionEastWest                fields.DirectionEastWest
	VelocityEWStatus                 fields.NumericValueStatus
	VelocityEW                       uint16
	DirectionNorthSouth              fields.DirectionNorthSouth
	VelocityNSStatus                 fields.NumericValueStatus
	VelocityNS                       uint16
	VerticalRateSource               fields.VerticalRateSource
	VerticalRateStatus               fields.NumericValueStatus
	VerticalRate                     int16
	DifferenceAltitudeGNSSBaroStatus fields.NumericValueStatus
	DifferenceAltitudeGNSSBaro       int16
}

func (message AirborneVelocityGroundSpeedSupersonic) GetSubtype() fields.Subtype {
	return message.Subtype
}

// GetRegister returns the Register the message
func (message AirborneVelocityGroundSpeedSupersonic) GetRegister() register.Register {
	return register.BDS09
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message AirborneVelocityGroundSpeedSupersonic) CheckCoherency() error {
	return nil
}

// ToString returns a basic, but readable, representation of the message
func (message AirborneVelocityGroundSpeedSupersonic) ToString() string {
	return fmt.Sprintf(""+
		"Message:                           %v\n"+
		"Subtype:                           %v\n"+
		"Intent Change:                     %v\n"+
		"IFR Capability:                    %v\n"+
		"Navigation Uncertainty Category:   %v\n"+
		"Direction EW:                      %v\n"+
		"Velocity EW Status:                %v\n"+
		"Velocity EW (knots):               %v\n"+
		"Direction NS:                      %v\n"+
		"Velocity NS Status:                %v\n"+
		"Velocity NS (knots):               %v\n"+
		"Vertical Rate Source:              %v\n"+
		"Vertical Rate Status:              %v\n"+
		"Vertical Rate (ft/min):            %v\n"+
		"Difference Alt. GNSS Baro Status:  %v\n"+
		"Difference Alt. GNSS Baro (ft):    %v\n"+
		"Computed Speed:                    %v\n"+
		"Computed Track:                    %v",
		message.GetRegister().ToString(),
		message.Subtype.ToString(),
		message.IntentChange.ToString(),
		message.IFRCapability.ToString(),
		message.NavigationUncertaintyCategory.ToString(),
		message.DirectionEastWest.ToString(),
		message.VelocityEWStatus.ToString(),
		message.VelocityEW,
		message.DirectionNorthSouth.ToString(),
		message.VelocityNSStatus.ToString(),
		message.VelocityNS,
		message.VerticalRateSource.ToString(),
		message.VerticalRateStatus.ToString(),
		message.VerticalRate,
		message.DifferenceAltitudeGNSSBaroStatus.ToString(),
		message.DifferenceAltitudeGNSSBaro,
		message.toSpeedString(),
		message.toTrackString())
}

// ReadAirborneVelocityGroundSpeedSupersonic reads a message AirborneVelocity (Ground speed supersonic)
func ReadAirborneVelocityGroundSpeedSupersonic(data []byte) (*AirborneVelocityGroundSpeedSupersonic, error) {

	if len(data) != 7 {
		return nil, fmt.Errorf("the data must be 7 bytes long (%v given)", len(data))
	}

	formatTypeCode := (data[0] & 0xF8) >> 3
	if formatTypeCode != 19 {
		return nil, fmt.Errorf("the field FormatTypeCode must be 19, got %v", formatTypeCode)
	}

	subType := fields.ReadSubtype(data)
	if subType != fields.SubtypeGroundSpeedSupersonic {
		return nil, fmt.Errorf("the data are given for subtype %v format and can not be read by ReadAirborneVelocityGroundSpeedSupersonic", subType.ToString())
	}

	velocityEW, velocityEWStatus := fields.ReadVelocityEWSupersonic(data)
	velocityNS, velocityNSStatus := fields.ReadVelocityNSSupersonic(data)
	verticalRate, verticalRateStatus := fields.ReadVerticalRate(data)
	diffBaro, diffBaroStatus := fields.ReadHeightDifference(data)

	return &AirborneVelocityGroundSpeedSupersonic{
		FormatTypeCode:                   formatTypeCode,
		Subtype:                          subType,
		IntentChange:                     fields.ReadIntentChange(data),
		IFRCapability:                    fields.ReadIFRCapability(data),
		NavigationUncertaintyCategory:    fields.ReadNavigationUncertaintyCategory(data),
		DirectionEastWest:                fields.ReadDirectionEastWest(data),
		VelocityEWStatus:                 velocityEWStatus,
		VelocityEW:                       velocityEW,
		DirectionNorthSouth:              fields.ReadDirectionNorthSouth(data),
		VelocityNSStatus:                 velocityNSStatus,
		VelocityNS:                       velocityNS,
		VerticalRateSource:               fields.ReadVerticalRateSource(data),
		VerticalRateStatus:               verticalRateStatus,
		VerticalRate:                     verticalRate,
		DifferenceAltitudeGNSSBaroStatus: diffBaroStatus,
		DifferenceAltitudeGNSSBaro:       diffBaro,
	}, nil
}

func (message AirborneVelocityGroundSpeedSupersonic) GetSpeed() (uint32, error) {
	if message.VelocityEWStatus != fields.NVSRegular {
		return 0, errors.New("velocity EW is not regular")
	}

	if message.VelocityNSStatus != fields.NVSRegular {
		return 0, errors.New("velocity NS is not regular")
	}

	velocityEW := float64(message.VelocityEW)
	velocityNS := float64(message.VelocityNS)

	speed := math.Sqrt(velocityEW*velocityEW + velocityNS*velocityNS)

	return uint32(speed), nil
}

func (message AirborneVelocityGroundSpeedSupersonic) toSpeedString() string {
	speed, err := message.GetSpeed()
	if err != nil {
		return fmt.Sprintf("not computable: %v", err)
	}

	return fmt.Sprintf("%v (knots)", speed)
}

func (message AirborneVelocityGroundSpeedSupersonic) GetTrack() (float32, error) {
	if message.VelocityEWStatus != fields.NVSRegular {
		return 0, errors.New("velocity EW is not regular")
	}

	if message.VelocityNSStatus != fields.NVSRegular {
		return 0, errors.New("velocity NS is not regular")
	}

	velocityEW := float64(message.VelocityEW)
	velocityNS := float64(message.VelocityNS)

	if message.DirectionEastWest == fields.DEWWest {
		velocityEW = -velocityEW
	}

	if message.DirectionNorthSouth == fields.DNSSouth {
		velocityNS = -velocityNS
	}

	trackRadian := math.Atan2(velocityEW, velocityNS)
	trackDegree := trackRadian * (180 / math.Pi)
	if trackDegree < 0 {
		trackDegree = trackDegree + 360
	}

	return float32(trackDegree), nil
}

func (message AirborneVelocityGroundSpeedSupersonic) toTrackString() string {
	track, err := message.GetTrack()
	if err != nil {
		return fmt.Sprintf("not computable: %v", err)
	}

	return fmt.Sprintf("%v (degrees)", track)
}
