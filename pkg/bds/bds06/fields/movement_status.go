package fields

// MovementStatus is the status of the Movement information
type MovementStatus int

const (
	// MSNoInformation indicates no information
	MSNoInformation MovementStatus = 0
	// MSValid indicates that Movement is valid
	MSValid MovementStatus = 1
	// MSAboveMaximum indicates that the Movement is above the maximum
	MSAboveMaximum MovementStatus = 2
	// MSReservedDecelerating indicates that the value is reserved
	MSReservedDecelerating MovementStatus = 3
	// MSReservedAccelerating indicates that the value is reserved
	MSReservedAccelerating MovementStatus = 4
	// MSReservedBackingUp indicates that the value is reserved
	MSReservedBackingUp MovementStatus = 5
)
