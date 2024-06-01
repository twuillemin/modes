package fields

import "fmt"

// DownlinkELMThroughputCapability is the Downlink ELM average throughput capability definition
//
// Specified in Doc 9871 / D.2.4.1
type DownlinkELMThroughputCapability byte

const (
	// DownlinkELMNoCapability indicates No DEML Capability.
	DownlinkELMNoCapability DownlinkELMThroughputCapability = 0
	// DownlinkELM4SegmentsEverySecond indicates One 4 segments DELM every second.
	DownlinkELM4SegmentsEverySecond DownlinkELMThroughputCapability = 1
	// DownlinkELM8SegmentsEverySecond indicates One 8 segments DELM every second.
	DownlinkELM8SegmentsEverySecond DownlinkELMThroughputCapability = 2
	// DownlinkELM16SegmentsEverySecond indicates One 16 segments DELM every second.
	DownlinkELM16SegmentsEverySecond DownlinkELMThroughputCapability = 3
	// DownlinkELM16SegmentsEvery500Milliseconds indicates One 16 segments DELM every 500 ms.
	DownlinkELM16SegmentsEvery500Milliseconds DownlinkELMThroughputCapability = 4
	// DownlinkELM16SegmentsEvery250Milliseconds indicates One 16 segments DELM every 250 ms.
	DownlinkELM16SegmentsEvery250Milliseconds DownlinkELMThroughputCapability = 5
	// DownlinkELM16SegmentsEvery125Milliseconds indicates One 16 segments DELM every 125 ms.
	DownlinkELM16SegmentsEvery125Milliseconds DownlinkELMThroughputCapability = 6
)

// ToString returns a basic, but readable, representation of the field
func (uetc DownlinkELMThroughputCapability) ToString() string {

	switch uetc {
	case DownlinkELMNoCapability:
		return "0 - No DEML Capability"
	case DownlinkELM4SegmentsEverySecond:
		return "1 - One 4 segments DELM every second"
	case DownlinkELM8SegmentsEverySecond:
		return "2 - One 8 segments DELM every second"
	case DownlinkELM16SegmentsEverySecond:
		return "3 - One 16 segments DELM every second"
	case DownlinkELM16SegmentsEvery500Milliseconds:
		return "4 - One 16 segments DELM every 500 ms"
	case DownlinkELM16SegmentsEvery250Milliseconds:
		return "5 - One 16 segments DELM every 250 ms"
	case DownlinkELM16SegmentsEvery125Milliseconds:
		return "6 - One 16 segments DELM every 125 ms"
	default:
		return fmt.Sprintf("%v - Unknown code", uetc)
	}
}

// ReadDownlinkELMThroughputCapability reads the DownlinkELMThroughputCapability from a 56 bits data field
func ReadDownlinkELMThroughputCapability(data []byte) DownlinkELMThroughputCapability {
	bits := data[3] & 0x0F
	return DownlinkELMThroughputCapability(bits)
}
