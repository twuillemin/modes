package messages

// ModeSMessage is the basic interface that Mode S messages are expected to implement
type ModeSMessage interface {
	// GetName returns the name of the message
	GetName() string
	// GetDownLinkFormat returns the downlink format of the message
	GetDownLinkFormat() int
	// ToString returns a basic, but readable, representation of the field
	ToString() string
}
