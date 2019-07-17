package fields

import "fmt"

// NavigationalAccuracyCategoryPosition is the Navigational Accuracy Category Position definition
//
// Specified in Doc 9871 / B.2.3.10.7
type NavigationalAccuracyCategoryPosition byte

const (
	// NACPEPUGreaterThan18Point52Km indicates EPU >= 18.52 km (10 NM) - Unknown accuracy
	NACPEPUGreaterThan18Point52Km NavigationalAccuracyCategoryPosition = 0
	// NACPEPULowerThan18Point52Km indicates  EPU < 18.52 km (10 NM) - RNP-10 accuracy
	NACPEPULowerThan18Point52Km NavigationalAccuracyCategoryPosition = 1
	// NACPEPULowerThan7Point408Km indicates EPU < 7.408 km (4 NM) - RNP-4 accuracy
	NACPEPULowerThan7Point408Km NavigationalAccuracyCategoryPosition = 2
	// NACPEPULowerThan3Point704Km indicates EPU < 3.704 km (2 NM) - RNP-2 accuracy
	NACPEPULowerThan3Point704Km NavigationalAccuracyCategoryPosition = 3
	// NACPEPUGreaterThan1852M indicates EPU < 1 852 m (1 NM) - RNP-1 accuracy
	NACPEPUGreaterThan1852M NavigationalAccuracyCategoryPosition = 4
	// NACPEPULowerThan926M indicates EPU < 926 m (0.5 NM) - RNP-0.5 accuracy
	NACPEPULowerThan926M NavigationalAccuracyCategoryPosition = 5
	// NACPEPUGreaterThan555Point6M indicates EPU < 555.6 m ( 0.3 NM) - RNP-0.3 accuracy
	NACPEPUGreaterThan555Point6M NavigationalAccuracyCategoryPosition = 6
	// NACPEPULowerThan185Point2M indicates EPU < 185.2 m (0.1 NM) - RNP-0.1 accuracy
	NACPEPULowerThan185Point2M NavigationalAccuracyCategoryPosition = 7
	// NACPEPUGreaterThan92Point6M indicates EPU < 92.6 m (0.05 NM) - e.g. GPS (with SA)
	NACPEPUGreaterThan92Point6M NavigationalAccuracyCategoryPosition = 8
	// NACPEPULowerThan30MAndVEPULowerThan45M indicates EPU < 30 m and VEPU < 45 m - e.g. GPS (SA off)
	NACPEPULowerThan30MAndVEPULowerThan45M NavigationalAccuracyCategoryPosition = 9
	// NACPEPULowerThan10MAndVEPULowerThan15M indicates EPU < 10 m and VEPU < 15 m - e.g. WAAS
	NACPEPULowerThan10MAndVEPULowerThan15M NavigationalAccuracyCategoryPosition = 10
	// NACPEPULowerThan4MAndVEPULowerThan3M indicates EPU < 3 m and VEPU < 4 m - e.g. LAAS
	NACPEPULowerThan4MAndVEPULowerThan3M NavigationalAccuracyCategoryPosition = 11
	// NACPReserved12 is reserved
	NACPReserved12 NavigationalAccuracyCategoryPosition = 12
	// NACPReserved13 is reserved
	NACPReserved13 NavigationalAccuracyCategoryPosition = 13
	// NACPReserved14 is reserved
	NACPReserved14 NavigationalAccuracyCategoryPosition = 14
	// NACPReserved15 is reserved
	NACPReserved15 NavigationalAccuracyCategoryPosition = 15
)

// ToString returns a basic, but readable, representation of the field
func (category NavigationalAccuracyCategoryPosition) ToString() string {

	switch category {

	case NACPEPUGreaterThan18Point52Km:
		return "0 - EPU >= 18.52 km (10 NM) - Unknown accuracy"
	case NACPEPULowerThan18Point52Km:
		return "1 - EPU < 18.52 km (10 NM) - RNP-10 accuracy"
	case NACPEPULowerThan7Point408Km:
		return "2 - EPU < 7.408 km (4 NM) - RNP-4 accuracy"
	case NACPEPULowerThan3Point704Km:
		return "3 - EPU < 3.704 km (2 NM) - RNP-2 accuracy"
	case NACPEPUGreaterThan1852M:
		return "4 - EPU < 1 852 m (1 NM) - RNP-1 accuracy"
	case NACPEPULowerThan926M:
		return "5 - EPU < 926 m (0.5 NM) - RNP-0.5 accuracy"
	case NACPEPUGreaterThan555Point6M:
		return "6 - EPU < 555.6 m ( 0.3 NM) - RNP-0.3 accuracy"
	case NACPEPULowerThan185Point2M:
		return "7 - EPU < 185.2 m (0.1 NM) - RNP-0.1 accuracy"
	case NACPEPUGreaterThan92Point6M:
		return "8 - EPU < 92.6 m (0.05 NM) - e.g. GPS (with SA)"
	case NACPEPULowerThan30MAndVEPULowerThan45M:
		return "9 - EPU < 30 m and VEPU < 45 m - e.g. GPS (SA off)"
	case NACPEPULowerThan10MAndVEPULowerThan15M:
		return "10 - EPU < 10 m and VEPU < 15 m - e.g. WAAS"
	case NACPEPULowerThan4MAndVEPULowerThan3M:
		return "11 - EPU < 3 m and VEPU < 4 m - e.g. LAAS"
	case NACPReserved12:
		return "12 - Reserved"
	case NACPReserved13:
		return "13 - Reserved"
	case NACPReserved14:
		return "14 - Reserved"
	case NACPReserved15:
		return "15 - Reserved"
	default:
		return fmt.Sprintf("%v - Unknown code", category)
	}
}

// ReadNavigationalAccuracyCategoryPosition reads the NavigationalAccuracyCategoryPosition from a 56 bits data field
func ReadNavigationalAccuracyCategoryPosition(data []byte) NavigationalAccuracyCategoryPosition {
	bits := data[5] & 0x0F
	return NavigationalAccuracyCategoryPosition(bits)
}
