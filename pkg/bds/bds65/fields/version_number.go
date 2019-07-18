package fields

import "fmt"

// VersionNumber is the ADSB version number
//
// Specified in Doc 9871 / B.2.3.10.5
type VersionNumber byte

const (
	// ADSBV0 indicates ADSB V0: Conformant to Doc 9871, 1st Edition, Appendix A
	ADSBV0 VersionNumber = 0
	// ADSBV1 indicates ADSB V1: Conformant to Doc 9871, 1st Edition, Appendix B
	ADSBV1 VersionNumber = 1
	// ADSBV2 indicates ADSB V2: Conformant to Doc 9871, 2nd Edition
	ADSBV2 VersionNumber = 2
)

// ToString returns a basic, but readable, representation of the field
func (versionNumber VersionNumber) ToString() string {

	switch versionNumber {
	case ADSBV0:
		return "0 - ADSB V0: Conformant to Doc 9871, 1st Edition, Appendix A"
	case ADSBV1:
		return "1 - ADSB V1: Conformant to Doc 9871, 1st Edition, Appendix B"
	case ADSBV2:
		return "2 - ADSB V2: Conformant to Doc 9871, 2nd Edition"
	}
	return fmt.Sprintf("%v", versionNumber)
}

// ReadVersionNumber reads the VersionNumber from a 56 bits data field
func ReadVersionNumber(data []byte) VersionNumber {

	bits := (data[5] % 0xE0) >> 5

	return VersionNumber(bits)
}
