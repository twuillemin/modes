package bds60

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds60/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// HeadingAndTrackReport is a message at the format BDS 5,0
//
// Specified in Doc 9871 / Table A-2-48
type HeadingAndTrackReport struct {
	MagneticHeadingStatus         bool
	MagneticHeadingOrientation    fields.MagneticHeading
	MagneticHeading               float32
	IndicatedAirSpeedStatus       bool
	IndicatedAirSpeed             uint32
	MachStatus                    bool
	Mach                          float32
	BarometricAltitudeRateStatus  bool
	BarometricAltitudeRate        int32
	InitialVerticalVelocityStatus bool
	InitialVerticalVelocity       int32
}

// GetRegister returns the Register the message
func (message HeadingAndTrackReport) GetRegister() register.Register {
	return register.BDS60
}

// ToString returns a basic, but readable, representation of the message
func (message HeadingAndTrackReport) ToString() string {
	return fmt.Sprintf(""+
		"Message:                                  %v\n"+
		"Magnetic Heading Status:                  %v\n"+
		"Magnetic Heading Orientation:             %v\n"+
		"Magnetic Heading (degrees):               %v\n"+
		"Indicated Air Speed Status:               %v\n"+
		"Indicated Air Speed (knots):              %v\n"+
		"Mach Status:                              %v\n"+
		"Mach:                                     %v\n"+
		"Barometric Altitude Rate Status:          %v\n"+
		"Barometric Altitude Rate (feet/minute):   %v\n"+
		"Inertial Vertical Velocity Status:        %v\n"+
		"Inertial Vertical Velocity (feet/minute): %v",
		message.GetRegister().ToString(),
		message.MagneticHeadingStatus,
		message.MagneticHeadingOrientation,
		message.MagneticHeading,
		message.IndicatedAirSpeedStatus,
		message.IndicatedAirSpeed,
		message.MachStatus,
		message.Mach,
		message.BarometricAltitudeRateStatus,
		message.BarometricAltitudeRate,
		message.InitialVerticalVelocityStatus,
		message.InitialVerticalVelocity)
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message HeadingAndTrackReport) CheckCoherency() error {
	// If no data available, it is probably not coherent
	if !message.MagneticHeadingStatus && !message.IndicatedAirSpeedStatus && !message.MachStatus && !message.BarometricAltitudeRateStatus && !message.InitialVerticalVelocityStatus {
		return errors.New("the message does not convey any information")
	}

	if !message.MagneticHeadingStatus && message.MagneticHeading != 0 {
		return errors.New("the magnetic heading status is set to false, but a magnetic heading value is given")
	}

	if !message.IndicatedAirSpeedStatus && message.IndicatedAirSpeed != 0 {
		return errors.New("the indicated air speed status is set to false, but a indicated air speed value is given")
	}

	if !message.MachStatus && message.Mach != 0 {
		return errors.New("the mach status is set to false, but a mach value is given")
	}

	if !message.BarometricAltitudeRateStatus && message.BarometricAltitudeRate != 0 {
		return errors.New("the barometric altitude rate status is set to false, but a barometric altitude rate value is given")
	}

	if !message.InitialVerticalVelocityStatus && message.InitialVerticalVelocity != 0 {
		return errors.New("the initial vertical velocity status is set to false, but a initial vertical velocity value is given")
	}

	if message.IndicatedAirSpeed > 500 {
		return errors.New("the indicated air speed is too high (above 500 knots)")
	}

	if message.Mach > 1.0 {
		return errors.New("the mach is too high (above 1.0)")
	}

	if (message.BarometricAltitudeRate < -6000) || (message.BarometricAltitudeRate > 600) {
		return errors.New("the barometric altitude rate is to high or to low (above or under 6000fpm)")
	}

	if (message.InitialVerticalVelocity < -6000) || (message.InitialVerticalVelocity > 600) {
		return errors.New("the initial vertical velocity is to high or to low (above or under 6000fpm)")
	}

	return nil
}

// ReadHeadingAndTrackReport reads a message as a HeadingAndTrackReport
func ReadHeadingAndTrackReport(data []byte) (*HeadingAndTrackReport, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for Comm-B HeadingAndTrackReport message must be 7 bytes long")
	}

	magneticHeadingStatus, magneticHeadingOrientation, magneticHeading := fields.ReadMagneticHeading(data)
	indicatedAirSpeedStatus, indicatedAirSpeed := fields.ReadIndicatedAirSpeed(data)
	machStatus, mach := fields.ReadMach(data)
	barometricAltitudeRateStatus, barometricAltitudeRate := fields.ReadBarometricAltitudeRate(data)
	initialVerticalVelocityStatus, initialVerticalVelocity := fields.ReadInertialVerticalVelocity(data)

	return &HeadingAndTrackReport{
		MagneticHeadingStatus:         magneticHeadingStatus,
		MagneticHeadingOrientation:    magneticHeadingOrientation,
		MagneticHeading:               magneticHeading,
		IndicatedAirSpeedStatus:       indicatedAirSpeedStatus,
		IndicatedAirSpeed:             indicatedAirSpeed,
		MachStatus:                    machStatus,
		Mach:                          mach,
		BarometricAltitudeRateStatus:  barometricAltitudeRateStatus,
		BarometricAltitudeRate:        barometricAltitudeRate,
		InitialVerticalVelocityStatus: initialVerticalVelocityStatus,
		InitialVerticalVelocity:       initialVerticalVelocity,
	}, nil
}
