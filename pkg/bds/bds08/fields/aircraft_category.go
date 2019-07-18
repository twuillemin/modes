package fields

// AircraftCategory is the base type that all Resolution Advisory should implement
type AircraftCategory interface {
	// GetCategorySetName returns the name of the category set
	GetCategorySetName() string
	// ToString returns a basic, but readable, representation of the field
	ToString() string
}
