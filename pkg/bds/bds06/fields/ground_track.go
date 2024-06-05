package fields

import (
	"fmt"
)

// GroundTrack is the Ground Track definition
//
// Specified in Doc 9871 / A.2.3.3.2
type GroundTrack byte

// ToString returns a basic, but readable, representation of the field
func (groundTrack GroundTrack) ToString() string {

	return fmt.Sprintf("%v degrees", groundTrack.GetGroundTrack())
}

// GetGroundTrack returns the GroundTrack denoted in degrees.
func (groundTrack GroundTrack) GetGroundTrack() float64 {

	return float64(groundTrack) * 360.0 / 128.0
}

// ReadGroundTrack reads the GroundTrack from a 56 bits data field
func ReadGroundTrack(data []byte) GroundTrack {

	bits := (data[1]&0x07)<<4 + (data[2]&0xF0)>>4

	return GroundTrack(bits)
}
