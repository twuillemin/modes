package fields

import "fmt"

// ActiveRANoVerticalRAGenerated is one of the possible resolution advisory contained in a MV field
type ActiveRANoVerticalRAGenerated struct{}

// GetType returns the type of Resolution Advisory
func (noVerticalRAGenerated ActiveRANoVerticalRAGenerated) GetType() RAType {
	return NoVerticalRAGenerated
}

// ToString returns a basic, but readable, representation of the field
func (noVerticalRAGenerated ActiveRANoVerticalRAGenerated) ToString() string {
	return fmt.Sprintf("Type: no Resolution Advisory has been generated")
}
