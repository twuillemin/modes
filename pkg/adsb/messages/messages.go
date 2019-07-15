package messages

// ADSBMessage is the basic interface that ADSB messages are expected to implement
type ADSBMessage interface {
	// GetName returns the name of the message
	GetName() string
	// GetBDS returns the binary data format
	GetBDS() string
	// GetFormatTypeCode returns the Format Type Code
	GetFormatTypeCode() byte
	// ToString returns a basic, but readable, representation of the field
	ToString() string
}
