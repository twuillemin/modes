package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
)

// ThreatIdentityAddress is a 3 bytes ICAO Address. The Most Significant Byte of the address
// is always 0.
type ThreatIdentityAddress uint32

// ToString returns a string representation of an address
func (address ThreatIdentityAddress) ToString() string {
	return fmt.Sprintf("%02X %02X %02X ",
		byte((address&0x00FF0000)>>16),
		byte((address&0x0000FF00)>>8),
		byte(address&0x000000FF))
}

// ReadThreatIdentityAddress reads the ThreatIdentityAddress
func ReadThreatIdentityAddress(data []byte) ThreatIdentityAddress {

	bits := bitutils.Pack4Bytes(data[2], data[3], data[4], data[5])
	bits = (bits >> 2) & 0x00FFFFFF
	return ThreatIdentityAddress(bits)
}
