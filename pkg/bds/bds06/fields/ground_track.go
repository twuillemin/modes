package fields

import (
	"fmt"
	"math"
)

// GroundTrack is the Ground Track definition
//
// Specified in Doc 9871 / A.2.3.3.2
type GroundTrack byte

// ToString returns a basic, but readable, representation of the field
func (groundTrack GroundTrack) ToString() string {

	return fmt.Sprintf("%v degrees", groundTrack)
}

// ReadGroundTrack reads the GroundTrack from a 56 bits data field
func ReadGroundTrack(data []byte) GroundTrack {

	bits := (data[0]&0x07)<<3 + (data[1]&0xF0)>>4
	degrees := float64(bits) / 360.0

	return GroundTrack(int(math.Round(degrees)))
}
