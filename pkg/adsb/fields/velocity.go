package fields

// VelocityStatus is the status of the velocity
type VelocityStatus int

const (
	// VelocityStatusNoInformation indicates no velocity information
	VelocityStatusNoInformation VelocityStatus = 0
	// VelocityStatusRegular indicates that the Velocity is computed on the linear scale value of field * factor
	VelocityStatusRegular VelocityStatus = 1
	// VelocityStatusMaximum indicates that the Velocity field value indicates velocity greater the maximum of the scale
	VelocityStatusMaximum VelocityStatus = 2
)

type Velocity interface {

	// GetStatus returns the status of the velocity
	GetStatus() VelocityStatus
	// GetScaleFactor returns the scale factor between the message field and the value
	GetScaleFactor() int
	// GetVelocityValue returns the velocity. Note that the returned value will be 0 for VelocityStatusNoInformation and
	// the maximum for VelocityMaximum
	GetVelocityValue() int
	// ToString a representation of the Velocity
	ToString() string
}
