package fields

// VelocityStatus is the status of the velocity / airspeed information
type VelocityStatus int

const (
	// VelocityStatusNoInformation indicates no velocity information
	VelocityStatusNoInformation VelocityStatus = 0
	// VelocityStatusRegular indicates that the Velocity is computed on the linear scale value of field * factor
	VelocityStatusRegular VelocityStatus = 1
	// VelocityStatusMaximum indicates that the Velocity field value indicates velocity greater the maximum of the scale
	VelocityStatusMaximum VelocityStatus = 2
)
