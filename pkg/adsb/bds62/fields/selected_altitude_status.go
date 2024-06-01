package fields

// SelectedAltitudeStatus is the status of the Selected Altitude information
type SelectedAltitudeStatus int

const (
	// SASInvalid indicates that the Selected Altitude is invalid of absent
	SASInvalid SelectedAltitudeStatus = 0
	// SASValid indicates that the Selected Altitude is valid
	SASValid SelectedAltitudeStatus = 1
)
