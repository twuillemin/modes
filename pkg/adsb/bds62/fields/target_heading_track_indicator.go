package fields

import "fmt"

// TargetHeadingTrackIndicator is the Target Heading / Track Angle Indicator definition
//
// Specified in Doc 9871 / B.2.3.9.10
type TargetHeadingTrackIndicator byte

const (
	// THTIHeading indicates target heading angle is being reported
	THTIHeading TargetHeadingTrackIndicator = 0
	// THTITrack indicates track angle is being reported
	THTITrack TargetHeadingTrackIndicator = 1
)

// ToString returns a basic, but readable, representation of the field
func (targetType TargetHeadingTrackIndicator) ToString() string {

	switch targetType {
	case THTIHeading:
		return "0 - target heading angle is being reported"
	case THTITrack:
		return "1 - track angle is being reported"
	default:
		return fmt.Sprintf("%v - Unknown code", targetType)
	}
}

// ReadTargetHeadingTrackIndicator reads the TargetHeadingTrackIndicator from a 56 bits data field
func ReadTargetHeadingTrackIndicator(data []byte) TargetHeadingTrackIndicator {
	bits := (data[4] & 0x08) >> 3
	return TargetHeadingTrackIndicator(bits)
}
