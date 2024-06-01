package fields

import "fmt"

// UplinkELMThroughputCapability is the Uplink ELM average throughput capability definition
//
// Specified in Doc 9871 / D.2.4.1
type UplinkELMThroughputCapability byte

const (
	// UplinkELMNoCapability indicates No UELM Capability.
	UplinkELMNoCapability UplinkELMThroughputCapability = 0
	// UplinkELM16SegmentsIn1Second indicates 16 UELM segments in 1 second.
	UplinkELM16SegmentsIn1Second UplinkELMThroughputCapability = 1
	// UplinkELM16SegmentsIn500Milliseconds indicates 16 UELM segments in 500ms.
	UplinkELM16SegmentsIn500Milliseconds UplinkELMThroughputCapability = 2
	// UplinkELM16SegmentsIn250Milliseconds indicates 16 UELM segments in 250 ms.
	UplinkELM16SegmentsIn250Milliseconds UplinkELMThroughputCapability = 3
	// UplinkELM16SegmentsIn125Milliseconds indicates 16 UELM segments in 125 ms.
	UplinkELM16SegmentsIn125Milliseconds UplinkELMThroughputCapability = 4
	// UplinkELM16SegmentsIn60Milliseconds indicates 16 UELM segments in 60ms.
	UplinkELM16SegmentsIn60Milliseconds UplinkELMThroughputCapability = 5
	// UplinkELM16SegmentsIn30Milliseconds indicates 16 UELM segments in 30ms.
	UplinkELM16SegmentsIn30Milliseconds UplinkELMThroughputCapability = 6
)

// ToString returns a basic, but readable, representation of the field
func (uetc UplinkELMThroughputCapability) ToString() string {

	switch uetc {
	case UplinkELMNoCapability:
		return "0 - No UELM Capability"
	case UplinkELM16SegmentsIn1Second:
		return "1 - 16 UELM segments in 1 second"
	case UplinkELM16SegmentsIn500Milliseconds:
		return "2 - 16 UELM segments in 500 milliseconds"
	case UplinkELM16SegmentsIn250Milliseconds:
		return "3 - 16 UELM segments in 250 milliseconds"
	case UplinkELM16SegmentsIn125Milliseconds:
		return "4 - 16 UELM segments in 125 milliseconds"
	case UplinkELM16SegmentsIn60Milliseconds:
		return "5 - 16 UELM segments in 60 milliseconds"
	case UplinkELM16SegmentsIn30Milliseconds:
		return "6 - 16 UELM segments in 30 milliseconds"
	default:
		return fmt.Sprintf("%v - Unknown code", uetc)
	}
}

// ReadUplinkELMThroughputCapability reads the UplinkELMThroughputCapability from a 56 bits data field
func ReadUplinkELMThroughputCapability(data []byte) UplinkELMThroughputCapability {
	bits := (data[3] & 0x70) >> 4
	return UplinkELMThroughputCapability(bits)
}
