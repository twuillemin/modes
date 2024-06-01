package fields

import "fmt"

// CompactPositionReportingFormat is the CPR (Compact Position Reporting) format definition
//
// Specified in Doc 9871 / A.2.3.3.3
type CompactPositionReportingFormat byte

const (
	// CPRFormatEven denotes an even format coding
	CPRFormatEven CompactPositionReportingFormat = 0
	// CPRFormatOdd indicates an odd format coding
	CPRFormatOdd CompactPositionReportingFormat = 1
)

// ToString returns a basic, but readable, representation of the field
func (format CompactPositionReportingFormat) ToString() string {

	switch format {
	case CPRFormatEven:
		return "0 - even format coding"
	case CPRFormatOdd:
		return "1 - odd format coding"
	default:
		return fmt.Sprintf("%v - Unknown code", format)
	}
}

// ReadCompactPositionReportingFormat reads the CompactPositionReportingFormat from a 56 bits data field
func ReadCompactPositionReportingFormat(data []byte) CompactPositionReportingFormat {
	bits := (data[2] & 0x04) >> 2
	return CompactPositionReportingFormat(bits)
}
