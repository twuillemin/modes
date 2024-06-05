package bds17

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/register"
)

type CommonUsageGICBCapabilityReport struct {
	ExtendedSquitterAirbornePosition            bool
	ExtendedSquitterSurfacePosition             bool
	ExtendedSquitterStatus                      bool
	ExtendedSquitterIdentificationAndCategory   bool
	ExtendedSquitterAirborneVelocityInformation bool
	ExtendedSquitterEventDrivenInformation      bool
	AircraftIdentification                      bool
	AircraftRegistrationNumber                  bool
	SelectedVerticalIntention                   bool
	NextWaypointIdentifier                      bool
	NextWaypointPosition                        bool
	NextWaypointInformation                     bool
	MeteorologicalRoutineReport                 bool
	MeteorologicalHazardReport                  bool
	VHFChannelReport                            bool
	TrackAndTurnReport                          bool
	PositionCoarse                              bool
	PositionFine                                bool
	AirReferencedStateVector                    bool
	Waypoint1                                   bool
	Waypoint2                                   bool
	Waypoint3                                   bool
	QuasiStaticParameterMonitoring              bool
	HeadingAndSpeedReport                       bool
	MilitaryApplications                        bool
}

// GetRegister returns the Register the message
func (message CommonUsageGICBCapabilityReport) GetRegister() register.Register {
	return register.BDS17
}

// ToString returns a basic, but readable, representation of the message
func (message CommonUsageGICBCapabilityReport) ToString() string {
	return fmt.Sprintf(""+
		"Message:                                             %v\n"+
		"0,5 Extended squitter airborne position:             %v\n"+
		"0,6 Extended squitter surface position:              %v\n"+
		"0,7 Extended squitter status:                        %v\n"+
		"0,8 Extended squitter identification and category:   %v\n"+
		"0,9 Extended squitter airborne velocity information: %v\n"+
		"0,A Extended squitter event-driven information:      %v\n"+
		"2,0 Aircraft identification:                         %v\n"+
		"2,l Aircraft registration number:                    %v\n"+
		"4,0 Selected vertical intention:                     %v\n"+
		"4,l Next waypoint identifier:                        %v\n"+
		"4,2 Next waypoint position:                          %v\n"+
		"4,3 Next waypoint information:                       %v\n"+
		"4,4 Meteorological routine report:                   %v\n"+
		"4,5 Meteorological hazard report:                    %v\n"+
		"4.8 VHF channel report:                              %v\n"+
		"5,0 Track and turn report:                           %v\n"+
		"5,1 Position coarse:                                 %v\n"+
		"5,2 Position fine:                                   %v\n"+
		"5,3 Air-referenced state vector:                     %v\n"+
		"5,4 Waypoint 1:                                      %v\n"+
		"5,5 Waypoint 2:                                      %v\n"+
		"5,6 Waypoint 3:                                      %v\n"+
		"5,F Quasi-static parameter monitoring:               %v\n"+
		"6,0 Heading and speed report:                        %v\n"+
		"F,1 Military applications:                           %v\n",
		message.GetRegister().ToString(),
		message.ExtendedSquitterAirbornePosition,
		message.ExtendedSquitterSurfacePosition,
		message.ExtendedSquitterStatus,
		message.ExtendedSquitterIdentificationAndCategory,
		message.ExtendedSquitterAirborneVelocityInformation,
		message.ExtendedSquitterEventDrivenInformation,
		message.AircraftIdentification,
		message.AircraftRegistrationNumber,
		message.SelectedVerticalIntention,
		message.NextWaypointIdentifier,
		message.NextWaypointPosition,
		message.NextWaypointInformation,
		message.MeteorologicalRoutineReport,
		message.MeteorologicalHazardReport,
		message.VHFChannelReport,
		message.TrackAndTurnReport,
		message.PositionCoarse,
		message.PositionFine,
		message.AirReferencedStateVector,
		message.Waypoint1,
		message.Waypoint2,
		message.Waypoint3,
		message.QuasiStaticParameterMonitoring,
		message.HeadingAndSpeedReport,
		message.MilitaryApplications)
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message CommonUsageGICBCapabilityReport) CheckCoherency() error {
	return nil
}

// ReadCommonUsageGICBCapabilityReport reads a message as a CommonUsageGICBCapabilityReport
func ReadCommonUsageGICBCapabilityReport(data []byte) (*CommonUsageGICBCapabilityReport, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for Comm-B CommonUsageGICBCapabilityReport message must be 7 bytes long")
	}

	// Bits 30 to 32 are reserved and must be 0
	if data[3]&0x07 != 0 {
		return nil, errors.New("the bits 30 to 32 are reserved and must be 0")
	}

	for i := uint(4); i < 7; i++ {
		if data[i] != 0 {
			return nil, errors.New("the bits 33 to 56 are reserved and must be 0")
		}
	}

	return &CommonUsageGICBCapabilityReport{
		ExtendedSquitterAirbornePosition:            (data[0]&0x80)>>7 != 0,
		ExtendedSquitterSurfacePosition:             (data[0]&0x40)>>6 != 0,
		ExtendedSquitterStatus:                      (data[0]&0x20)>>5 != 0,
		ExtendedSquitterIdentificationAndCategory:   (data[0]&0x10)>>4 != 0,
		ExtendedSquitterAirborneVelocityInformation: (data[0]&0x08)>>3 != 0,
		ExtendedSquitterEventDrivenInformation:      (data[0]&0x04)>>2 != 0,
		AircraftIdentification:                      (data[0]&0x02)>>1 != 0,
		AircraftRegistrationNumber:                  (data[0]&0x01)>>0 != 0,
		SelectedVerticalIntention:                   (data[1]&0x80)>>7 != 0,
		NextWaypointIdentifier:                      (data[1]&0x40)>>6 != 0,
		NextWaypointPosition:                        (data[1]&0x20)>>5 != 0,
		NextWaypointInformation:                     (data[1]&0x10)>>4 != 0,
		MeteorologicalRoutineReport:                 (data[1]&0x08)>>3 != 0,
		MeteorologicalHazardReport:                  (data[1]&0x04)>>2 != 0,
		VHFChannelReport:                            (data[1]&0x02)>>1 != 0,
		TrackAndTurnReport:                          (data[1]&0x01)>>0 != 0,
		PositionCoarse:                              (data[2]&0x80)>>7 != 0,
		PositionFine:                                (data[2]&0x40)>>6 != 0,
		AirReferencedStateVector:                    (data[2]&0x20)>>5 != 0,
		Waypoint1:                                   (data[2]&0x10)>>4 != 0,
		Waypoint2:                                   (data[2]&0x08)>>3 != 0,
		Waypoint3:                                   (data[2]&0x04)>>2 != 0,
		QuasiStaticParameterMonitoring:              (data[2]&0x02)>>1 != 0,
		HeadingAndSpeedReport:                       (data[2]&0x01)>>0 != 0,
		MilitaryApplications:                        (data[3]&0x08)>>3 != 0,
	}, nil
}
