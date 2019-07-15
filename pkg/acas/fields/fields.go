package fields

// Field is the basic interface that Mode S messages fields are expected to implement
type Field interface {
	// ToString returns a basic, but readable, representation of the field
	ToString() string
}
