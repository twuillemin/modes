package fields

import "fmt"

// NavigationalAccuracyCategoryPositionV2 is the Navigational Accuracy Category Position definition
//
// Specified in Doc 9871 / C.2.3.9.9
type NavigationalAccuracyCategoryPositionV2 byte

const (
	// NACV2PEPUGreaterThan18Point52Km indicates EPU >= 18.52 km (10 NM) - Unknown accuracy
	NACV2PEPUGreaterThan18Point52Km NavigationalAccuracyCategoryPositionV2 = 0
	// NACV2PEPULowerThan18Point52Km indicates  EPU < 18.52 km (10 NM) - RNP-10 accuracy
	NACV2PEPULowerThan18Point52Km NavigationalAccuracyCategoryPositionV2 = 1
	// NACV2PEPULowerThan7Point408Km indicates EPU < 7.408 km (4 NM) - RNP-4 accuracy
	NACV2PEPULowerThan7Point408Km NavigationalAccuracyCategoryPositionV2 = 2
	// NACV2PEPULowerThan3Point704Km indicates EPU < 3.704 km (2 NM) - RNP-2 accuracy
	NACV2PEPULowerThan3Point704Km NavigationalAccuracyCategoryPositionV2 = 3
	// NACV2PEPUGreaterThan1852M indicates EPU < 1 852 m (1 NM) - RNP-1 accuracy
	NACV2PEPUGreaterThan1852M NavigationalAccuracyCategoryPositionV2 = 4
	// NACV2PEPULowerThan926M indicates EPU < 926 m (0.5 NM) - RNP-0.5 accuracy
	NACV2PEPULowerThan926M NavigationalAccuracyCategoryPositionV2 = 5
	// NACV2PEPUGreaterThan555Point6M indicates EPU < 555.6 m ( 0.3 NM) - RNP-0.3 accuracy
	NACV2PEPUGreaterThan555Point6M NavigationalAccuracyCategoryPositionV2 = 6
	// NACV2PEPULowerThan185Point2M indicates EPU < 185.2 m (0.1 NM) - RNP-0.1 accuracy
	NACV2PEPULowerThan185Point2M NavigationalAccuracyCategoryPositionV2 = 7
	// NACV2PEPUGreaterThan92Point6M indicates EPU < 92.6 m (0.05 NM) - e.g. GPS (with SA)
	NACV2PEPUGreaterThan92Point6M NavigationalAccuracyCategoryPositionV2 = 8
	// NACV2PEPULowerThan30M indicates EPU < 30 m - e.g. GPS (SA off)
	NACV2PEPULowerThan30M NavigationalAccuracyCategoryPositionV2 = 9
	// NACV2PEPULowerThan10M indicates EPU < 10 m - e.g. WAAS
	NACV2PEPULowerThan10M NavigationalAccuracyCategoryPositionV2 = 10
	// NACV2PEPULowerThan3M indicates EPU < 3 m - e.g. LAAS
	NACV2PEPULowerThan3M NavigationalAccuracyCategoryPositionV2 = 11
	// NACV2PReserved12 is reserved
	NACV2PReserved12 NavigationalAccuracyCategoryPositionV2 = 12
	// NACV2PReserved13 is reserved
	NACV2PReserved13 NavigationalAccuracyCategoryPositionV2 = 13
	// NACV2PReserved14 is reserved
	NACV2PReserved14 NavigationalAccuracyCategoryPositionV2 = 14
	// NACV2PReserved15 is reserved
	NACV2PReserved15 NavigationalAccuracyCategoryPositionV2 = 15
)

// ToString returns a basic, but readable, representation of the field
func (category NavigationalAccuracyCategoryPositionV2) ToString() string {

	switch category {

	case NACV2PEPUGreaterThan18Point52Km:
		return "0 - EPU >= 18.52 km (10 NM) - Unknown accuracy"
	case NACV2PEPULowerThan18Point52Km:
		return "1 - EPU < 18.52 km (10 NM) - RNP-10 accuracy"
	case NACV2PEPULowerThan7Point408Km:
		return "2 - EPU < 7.408 km (4 NM) - RNP-4 accuracy"
	case NACV2PEPULowerThan3Point704Km:
		return "3 - EPU < 3.704 km (2 NM) - RNP-2 accuracy"
	case NACV2PEPUGreaterThan1852M:
		return "4 - EPU < 1 852 m (1 NM) - RNP-1 accuracy"
	case NACV2PEPULowerThan926M:
		return "5 - EPU < 926 m (0.5 NM) - RNP-0.5 accuracy"
	case NACV2PEPUGreaterThan555Point6M:
		return "6 - EPU < 555.6 m ( 0.3 NM) - RNP-0.3 accuracy"
	case NACV2PEPULowerThan185Point2M:
		return "7 - EPU < 185.2 m (0.1 NM) - RNP-0.1 accuracy"
	case NACV2PEPUGreaterThan92Point6M:
		return "8 - EPU < 92.6 m (0.05 NM) - e.g. GPS (with SA)"
	case NACV2PEPULowerThan30M:
		return "9 - EPU < 30 m - e.g. GPS (SA off)"
	case NACV2PEPULowerThan10M:
		return "10 - EPU < 10 m - e.g. WAAS"
	case NACV2PEPULowerThan3M:
		return "11 - EPU < 3 m - e.g. LAAS"
	case NACV2PReserved12:
		return "12 - Reserved"
	case NACV2PReserved13:
		return "13 - Reserved"
	case NACV2PReserved14:
		return "14 - Reserved"
	case NACV2PReserved15:
		return "15 - Reserved"
	default:
		return fmt.Sprintf("%v - Unknown code", category)
	}
}

// ReadNavigationalAccuracyCategoryPositionV2 reads the NavigationalAccuracyCategoryPositionV2 from a 56 bits data field
func ReadNavigationalAccuracyCategoryPositionV2(data []byte) NavigationalAccuracyCategoryPositionV2 {
	bits := (data[4]&0x01)<<3 + (data[5]&0xE0)>>5
	return NavigationalAccuracyCategoryPositionV2(bits)
}
