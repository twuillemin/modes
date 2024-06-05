package fields

import "fmt"

// CommonUsageGICBCapability is the Common usage GICB capability definition
//
// Specified in Doc 9871 / D.2.4.1
type CommonUsageGICBCapability byte

const (
	// CommonUsageGICBReportUnChanged indicates that the common usage GICB capability report (register 17) is unchanged.
	CommonUsageGICBReportUnChanged CommonUsageGICBCapability = 0
	// CommonUsageGICBReportChanged indicates that the common usage GICB capability report (register 17) changed.
	CommonUsageGICBReportChanged CommonUsageGICBCapability = 1
)

// ToString returns a basic, but readable, representation of the field
func (cugc CommonUsageGICBCapability) ToString() string {

	switch cugc {
	case CommonUsageGICBReportUnChanged:
		return "0 - The common usage GICB capability report (register 17) is unchange"
	case CommonUsageGICBReportChanged:
		return "1 - The common usage GICB capability report (register 17) changed"
	default:
		return fmt.Sprintf("%v - Unknown code", cugc)
	}
}

// ReadCommonUsageGICBCapability reads the CommonUsageGICBCapability from a 56 bits data field
func ReadCommonUsageGICBCapability(data []byte) CommonUsageGICBCapability {
	bits := (data[4] & 0x10) >> 4
	return CommonUsageGICBCapability(bits)
}
