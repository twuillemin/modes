package fields

import "fmt"

// NavigationAccuracyCategory is the NACv (Navigation Accuracy Category for Velocity) definition
//
// Specified in Doc 9871 / C.2.3.10.3
type NavigationAccuracyCategory byte

// ToString returns a basic, but readable, representation of the field
func (category NavigationAccuracyCategory) ToString() string {
	return fmt.Sprintf("%v", category)
}

// ReadNavigationAccuracyCategory reads the NavigationAccuracyCategory from a 56 bits data field
func ReadNavigationAccuracyCategory(data []byte) NavigationAccuracyCategory {
	bits := (data[2] & 0xE) >> 5
	return NavigationAccuracyCategory(bits)
}
