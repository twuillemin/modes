package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// MagneticHeading is the Velocity Magnetic Heading definition
//
// Specified in Doc 9871 / A.2.3.5.6
type MagneticHeading float64

// ToString returns a basic, but readable, representation of the field
func (heading MagneticHeading) ToString() string {

	return fmt.Sprintf("%v", heading)
}

// ReadMagneticHeading reads the MagneticHeading from a 56 bits data field
func ReadMagneticHeading(data []byte) MagneticHeading {
	bit1 := data[1] & 0x03
	bit2 := data[2]

	value := float64(bitutils.Pack2Bytes(bit1, bit2)) * 360 / 1024.0

	return MagneticHeading(value)
}
