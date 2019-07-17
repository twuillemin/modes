package common

// Printable is the basic interface that most of fields and messages should implement
type Printable interface {
	// ToString returns a basic, but readable, representation of the field
	ToString() string
}
