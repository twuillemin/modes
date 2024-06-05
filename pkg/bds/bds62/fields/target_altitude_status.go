package fields

// TargetAltitudeStatus is the status of the Target Altitude information
type TargetAltitudeStatus int

const (
	// TASInvalid indicates that the Target Altitude is invalid (out of range)
	TASInvalid TargetAltitudeStatus = 0
	// TASValid indicates that the Target Altitude is valid
	TASValid TargetAltitudeStatus = 1
)
