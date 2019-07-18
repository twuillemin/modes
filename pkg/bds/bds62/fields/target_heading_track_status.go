package fields

// TargetHeadingTrackStatus is the status of the Target Heading / Track Indicator information
type TargetHeadingTrackStatus int

const (
	// THTSValid indicates that the Target Heading/Track is valid
	THTSValid TargetHeadingTrackStatus = 0
	// THTSInvalid indicates that the Target Heading/Track is invalid (out of range)
	THTSInvalid TargetHeadingTrackStatus = 1
)
