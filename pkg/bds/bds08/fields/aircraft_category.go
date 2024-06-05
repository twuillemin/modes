package fields

// AircraftCategory is the base type that all Resolution Advisory should implement
//
// Specified in Doc 9871 / Table A-2-8
type AircraftCategory interface {
	// GetCategorySetName returns the name of the category set
	GetCategorySetName() string
	// ToString returns a basic, but readable, representation of the field
	ToString() string
	// CheckCoherency checks that the Category is coherent
	CheckCoherency() error
}
