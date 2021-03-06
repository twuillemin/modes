package common

import "fmt"

// ICAOAddress is a 3 bytes ICAO Address, or Interrogator Identifier. The Most Significant Byte of the address
// is always 0.
type ICAOAddress uint32

// ToString returns a string representation of an address
func (address ICAOAddress) ToString() string {
	return fmt.Sprintf("%02X %02X %02X ",
		byte((address&0x00FF0000)>>16),
		byte((address&0x0000FF00)>>8),
		byte(address&0x000000FF))
}
