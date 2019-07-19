package fields

import "fmt"

// IntentChange is the Intent Change Flag definition
//
// Specified in Doc 9871 / A.2.3.5.3
type IntentChange byte

const (
	// ICNoChange indicates no change in intent
	ICNoChange IntentChange = 0
	// ICChangeInIntent indicates intent change
	ICChangeInIntent IntentChange = 1
)

// ToString returns a basic, but readable, representation of the field
func (flag IntentChange) ToString() string {

	switch flag {
	case ICNoChange:
		return "0 - no change in intent"
	case ICChangeInIntent:
		return "1 - intent change"
	default:
		return fmt.Sprintf("%v - Unknown code", flag)
	}
}

// ReadIntentChange reads the IntentChange from a 56 bits data field
func ReadIntentChange(data []byte) IntentChange {
	bits := (data[1] & 0x80) >> 7
	return IntentChange(bits)
}
