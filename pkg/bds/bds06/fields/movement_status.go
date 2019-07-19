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
	// MSReserved indicates that the value is reserved
	MSReserved MovementStatus = 3
)
