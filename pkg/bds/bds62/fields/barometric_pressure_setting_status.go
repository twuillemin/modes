package fields

// BarometricPressureSettingStatus is the status of the Barometric Pressure Setting information
type BarometricPressureSettingStatus int

const (
	// BPSInvalid indicates that the Barometric Pressure Setting is invalid of absent
	BPSInvalid BarometricPressureSettingStatus = 0
	// BPSValid indicates that the Barometric Pressure Setting is valid
	BPSValid BarometricPressureSettingStatus = 1
)
