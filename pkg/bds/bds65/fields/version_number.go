package fields

import "fmt"

// VersionNumber is the ADSB version number
//
// Specified in Doc 9871 / B.2.3.10.5
type VersionNumber byte

const (
	// ADSBVersion0 indicates ADSB V0: Conformant to Doc 9871, Appendix A
	ADSBVersion0 VersionNumber = 0
	// ADSBVersion1 indicates ADSB V1: Conformant to Doc 9871, Appendix B
	ADSBVersion1 VersionNumber = 1
	// ADSBVersion2 indicates ADSB V2: Conformant to Doc 9871, Appendix C
	ADSBVersion2 VersionNumber = 2
)

// ToString returns a basic, but readable, representation of the field
func (versionNumber VersionNumber) ToString() string {

	switch versionNumber {
	case ADSBVersion0:
		return "0 - ADSB V0: Conformant to Doc 9871, Appendix A"
	case ADSBVersion1:
		return "1 - ADSB V1: Conformant to Doc 9871, Appendix B"
	case ADSBVersion2:
		return "2 - ADSB V2: Conformant to Doc 9871, Appendix C"
	}
	return fmt.Sprintf("%v", versionNumber)
}

// ReadVersionNumber reads the VersionNumber from a 56 bits data field
func ReadVersionNumber(data []byte) VersionNumber {

	bits := (data[5] % 0xE0) >> 5

	return VersionNumber(bits)
}
