package fields

import "fmt"

// IntentChangeFlag is the Intent Change Flag definition
//
// Specified in Doc 9871 / A.2.3.5.3
type IntentChangeFlag byte

const (
	// ICRFNoChange indicates no change in intent
	ICRFNoChange IntentChangeFlag = 0
	// ICRFChangeInIntent indicates intent change
	ICRFChangeInIntent IntentChangeFlag = 1
)

// ToString returns a basic, but readable, representation of the field
func (flag IntentChangeFlag) ToString() string {

	switch flag {
	case ICRFNoChange:
		return "0 - no change in intent"
	case ICRFChangeInIntent:
		return "1 - intent change"
	default:
		return fmt.Sprintf("%v - Unknown code", flag)
	}
}

// ReadIntentChangeFlag reads the IntentChangeFlag from a 56 bits data field
func ReadIntentChangeFlag(data []byte) IntentChangeFlag {
	bits := (data[1] & 0x80) >> 7
	return IntentChangeFlag(bits)
}
