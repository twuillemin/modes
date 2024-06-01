package fields

import "fmt"

// ModeSSubnetworkVersion is the Mode S Subnetwork Version Number definition
//
// Specified in Doc 9871 / D.2.4.1
type ModeSSubnetworkVersion byte

const (
	// ModeSSubnetworkNotAvailable indicates that Mode S subnetwork is not available.
	ModeSSubnetworkNotAvailable ModeSSubnetworkVersion = 0
	// ModeSSubnetworkVersion1 indicates that the version is 1.
	ModeSSubnetworkVersion1 ModeSSubnetworkVersion = 1
	// ModeSSubnetworkVersion2 indicates that the version is 2.
	ModeSSubnetworkVersion2 = 2
	// ModeSSubnetworkVersion3 indicates that the version is 3.
	ModeSSubnetworkVersion3 = 3
	// ModeSSubnetworkVersion4 indicates that the version is 4.
	ModeSSubnetworkVersion4 = 4
	// ModeSSubnetworkVersion5 indicates that the version is 5.
	ModeSSubnetworkVersion5 = 5
)

// ToString returns a basic, but readable, representation of the field
func (mssv ModeSSubnetworkVersion) ToString() string {

	switch mssv {
	case ModeSSubnetworkNotAvailable:
		return "0 - Mode S subnetwork not available"
	case ModeSSubnetworkVersion1:
		return "1 - Mode S subnetwork Version 1 - ICAO Doc 9688 (1996)"
	case ModeSSubnetworkVersion2:
		return "1 - Mode S subnetwork Version 2 - ICAO Doc 9688 (1997)"
	case ModeSSubnetworkVersion3:
		return "1 - Mode S subnetwork Version 3 - ICAO Annex 10, Vol III, Amdt 77"
	case ModeSSubnetworkVersion4:
		return "1 - Mode S subnetwork Version 4 - ICAO Doc 9871, Edition 1"
	case ModeSSubnetworkVersion5:
		return "1 - Mode S subnetwork Version 5 - ICAO Doc 9871, Edition 2"
	default:
		return fmt.Sprintf("%v - Unknown code", mssv)
	}
}

// ReadModeSSubnetworkVersion reads the ModeSSubnetworkVersion from a 56 bits data field
func ReadModeSSubnetworkVersion(data []byte) ModeSSubnetworkVersion {
	bits := (data[2] & 0xFE) >> 1
	return ModeSSubnetworkVersion(bits)
}
