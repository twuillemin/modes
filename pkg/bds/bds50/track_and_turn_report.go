package bds50

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/bds50/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// TrackAndTurnReport is a message at the format BDS 5,0
//
// Specified in Doc 9871 / Table A-2-48
type TrackAndTurnReport struct {
	RollAngleStatus           bool
	RollAngleWingDown         fields.WingDown
	RollAngle                 float32
	TrueTrackAngleStatus      bool
	TrueTrackAngleOrientation fields.TrueTrackOrientation
	TrueTrackAngle            float32
	GroundSpeedStatus         bool
	GroundSpeed               float32
	TrackAngleRateStatus      bool
	TrackAngleRate            float32
	TrueAirSpeedStatus        bool
	TrueAirSpeed              float32
}

// GetRegister returns the Register the message
func (message TrackAndTurnReport) GetRegister() register.Register {
	return register.BDS50
}

// ToString returns a basic, but readable, representation of the message
func (message TrackAndTurnReport) ToString() string {
	return fmt.Sprintf(""+
		"Message:                               %v\n"+
		"Roll Angle Status:                     %v\n"+
		"Roll Angle Wing:                       %v\n"+
		"Roll Angle (degrees):                  %v\n"+
		"True Track Angle Status:               %v\n"+
		"True Track Angle Orientation:          %v\n"+
		"True Track (degrees):                  %v\n"+
		"Ground Speed Status:                   %v\n"+
		"Ground Speed (knots):                  %v\n"+
		"Track Angle Rate Status:               %v\n"+
		"Track Angle Rate (degrees/second):     %v\n"+
		"True Air Speed Status:                 %v\n"+
		"True Air Speed (knots):                %v",
		message.GetRegister().ToString(),
		message.RollAngleStatus,
		message.RollAngleWingDown,
		message.RollAngle,
		message.TrueTrackAngleStatus,
		message.TrueTrackAngleOrientation,
		message.TrueTrackAngle,
		message.GroundSpeedStatus,
		message.GroundSpeed,
		message.TrackAngleRateStatus,
		message.TrackAngleRate,
		message.TrueAirSpeedStatus,
		message.TrueAirSpeed)
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message TrackAndTurnReport) CheckCoherency() error {
	// If no data available, it is probably not coherent
	if !message.RollAngleStatus && !message.TrueTrackAngleStatus && !message.GroundSpeedStatus && !message.TrackAngleRateStatus && !message.TrueAirSpeedStatus {
		return errors.New("the message does not convey any information")
	}

	if message.RollAngle > 50 {
		return errors.New("the roll angle is too high (above 50 degrees)")
	}

	if message.GroundSpeed > 600 {
		return errors.New("the ground speed is too high (above 600 knots)")
	}

	if message.TrueAirSpeed > 500 {
		return errors.New("the true air speed is too high (above 500 knots)")
	}

	speedDifference := message.GroundSpeed - message.TrueAirSpeed
	if (speedDifference < -200) || (speedDifference > 200) {
		return errors.New("the difference between true air speed and ground speed is too high (above 200 knots)")
	}

	return nil
}

// ReadTrackAndTurnReport reads a message as a TrackAndTurnReport
func ReadTrackAndTurnReport(data []byte) (*TrackAndTurnReport, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for Comm-B TrackAndTurnReport message must be 7 bytes long")
	}

	rollAngleStatus, rollAngleWingDown, rollAngle := fields.ReadRollAngle(data)
	trueTrackAngleStatus, trueTrackAngleOrientation, trueTrackAngle := fields.ReadTrueTrackAngle(data)
	groundSpeedStatus, groundSpeed := fields.ReadGroundSpeed(data)
	trackAngleRateStatus, trackAngleRate := fields.ReadTrackAngleRate(data)
	trueAirSpeedStatus, trueAirSpeed := fields.ReadTrueAirSpeed(data)

	return &TrackAndTurnReport{
		RollAngleStatus:           rollAngleStatus,
		RollAngleWingDown:         rollAngleWingDown,
		RollAngle:                 rollAngle,
		TrueTrackAngleStatus:      trueTrackAngleStatus,
		TrueTrackAngleOrientation: trueTrackAngleOrientation,
		TrueTrackAngle:            trueTrackAngle,
		GroundSpeedStatus:         groundSpeedStatus,
		GroundSpeed:               groundSpeed,
		TrackAngleRateStatus:      trackAngleRateStatus,
		TrackAngleRate:            trackAngleRate,
		TrueAirSpeedStatus:        trueAirSpeedStatus,
		TrueAirSpeed:              trueAirSpeed,
	}, nil
}
