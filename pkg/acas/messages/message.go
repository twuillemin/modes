package messages

// ACASMessage is the basic interface that ACAS messages are expected to implement
type ACASMessage interface {
	// GetName returns the name of the message
	GetName() string
	// GetVDS1 returns the VDS1
	GetVDS1() byte
	// GetVDS2 returns the VDS1
	GetVDS2() byte
	// ToString returns a basic, but readable, representation of the field
	ToString() string
}
