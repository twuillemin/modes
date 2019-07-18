package fields

// TargetAltitudeStatus is the status of the Target Altitude information
type TargetAltitudeStatus int

const (
	// TASValid indicates that the Target Altitude is valid
	TASValid TargetAltitudeStatus = 0
	// TASInvalid indicates that the Target Altitude is invalid (out of range)
	TASInvalid TargetAltitudeStatus = 1
)
