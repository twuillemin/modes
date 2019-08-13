package fields

import "fmt"

// NavigationAccuracyCategoryVelocity is the NACv (Navigation Accuracy Category for Velocity) definition
//
// Specified in Doc 9871 / C.2.3.10.3
type NavigationAccuracyCategoryVelocity byte

// ToString returns a basic, but readable, representation of the field
func (category NavigationAccuracyCategoryVelocity) ToString() string {
	return fmt.Sprintf("%v", category)
}

// ReadNavigationAccuracyCategoryVelocity reads the NavigationAccuracyCategoryVelocity from a 56 bits data field
func ReadNavigationAccuracyCategoryVelocity(data []byte) NavigationAccuracyCategoryVelocity {
	bits := (data[2] & 0xE0) >> 5
	return NavigationAccuracyCategoryVelocity(bits)
}
