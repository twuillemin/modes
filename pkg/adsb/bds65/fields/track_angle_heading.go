package fields

import "fmt"

// TrackAngleHeading is the Track Angle / Heading definition
//
// Specified in Doc 9871 / B.2.3.10.12
type TrackAngleHeading byte

const (
	// TAHTargetHeadingReported indicates Target heading angle is being reported
	TAHTargetHeadingReported TrackAngleHeading = 0
	// TAHTrackAngleReported indicates Track angle is being reported
	TAHTrackAngleReported TrackAngleHeading = 1
)

// ToString returns a basic, but readable, representation of the field
func (headingAngle TrackAngleHeading) ToString() string {

	switch headingAngle {
	case TAHTargetHeadingReported:
		return "0 - Target heading angle is being reported"
	case TAHTrackAngleReported:
		return "1 - Track angle is being reported"
	default:
		return fmt.Sprintf("%v - Unknown code", headingAngle)
	}
}

// ReadTrackAngleHeading reads the TrackAngleHeading from a 56 bits data field
func ReadTrackAngleHeading(data []byte) TrackAngleHeading {
	bits := (data[6] & 0x08) >> 3
	return TrackAngleHeading(bits)
}
