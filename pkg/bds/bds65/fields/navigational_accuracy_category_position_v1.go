package fields

import "fmt"

// NavigationalAccuracyCategoryPositionV1 is the Navigational Accuracy Category Position definition
//
// Specified in Doc 9871 / B.2.3.10.7
type NavigationalAccuracyCategoryPositionV1 byte

const (
	// NACV1PEPUGreaterThan18Point52Km indicates EPU >= 18.52 km (10 NM) - Unknown accuracy
	NACV1PEPUGreaterThan18Point52Km NavigationalAccuracyCategoryPositionV1 = 0
	// NACV1PEPULowerThan18Point52Km indicates  EPU < 18.52 km (10 NM) - RNP-10 accuracy
	NACV1PEPULowerThan18Point52Km NavigationalAccuracyCategoryPositionV1 = 1
	// NACV1PEPULowerThan7Point408Km indicates EPU < 7.408 km (4 NM) - RNP-4 accuracy
	NACV1PEPULowerThan7Point408Km NavigationalAccuracyCategoryPositionV1 = 2
	// NACV1PEPULowerThan3Point704Km indicates EPU < 3.704 km (2 NM) - RNP-2 accuracy
	NACV1PEPULowerThan3Point704Km NavigationalAccuracyCategoryPositionV1 = 3
	// NACV1PEPUGreaterThan1852M indicates EPU < 1 852 m (1 NM) - RNP-1 accuracy
	NACV1PEPUGreaterThan1852M NavigationalAccuracyCategoryPositionV1 = 4
	// NACV1PEPULowerThan926M indicates EPU < 926 m (0.5 NM) - RNP-0.5 accuracy
	NACV1PEPULowerThan926M NavigationalAccuracyCategoryPositionV1 = 5
	// NACV1PEPUGreaterThan555Point6M indicates EPU < 555.6 m ( 0.3 NM) - RNP-0.3 accuracy
	NACV1PEPUGreaterThan555Point6M NavigationalAccuracyCategoryPositionV1 = 6
	// NACV1PEPULowerThan185Point2M indicates EPU < 185.2 m (0.1 NM) - RNP-0.1 accuracy
	NACV1PEPULowerThan185Point2M NavigationalAccuracyCategoryPositionV1 = 7
	// NACV1PEPUGreaterThan92Point6M indicates EPU < 92.6 m (0.05 NM) - e.g. GPS (with SA)
	NACV1PEPUGreaterThan92Point6M NavigationalAccuracyCategoryPositionV1 = 8
	// NACV1PEPULowerThan30MAndVEPULowerThan45M indicates EPU < 30 m and VEPU < 45 m - e.g. GPS (SA off)
	NACV1PEPULowerThan30MAndVEPULowerThan45M NavigationalAccuracyCategoryPositionV1 = 9
	// NACV1PEPULowerThan10MAndVEPULowerThan15M indicates EPU < 10 m and VEPU < 15 m - e.g. WAAS
	NACV1PEPULowerThan10MAndVEPULowerThan15M NavigationalAccuracyCategoryPositionV1 = 10
	// NACV1PEPULowerThan3MAndVEPULowerThan4M indicates EPU < 3 m and VEPU < 4 m - e.g. LAAS
	NACV1PEPULowerThan3MAndVEPULowerThan4M NavigationalAccuracyCategoryPositionV1 = 11
	// NACV1PReserved12 is reserved
	NACV1PReserved12 NavigationalAccuracyCategoryPositionV1 = 12
	// NACV1PReserved13 is reserved
	NACV1PReserved13 NavigationalAccuracyCategoryPositionV1 = 13
	// NACV1PReserved14 is reserved
	NACV1PReserved14 NavigationalAccuracyCategoryPositionV1 = 14
	// NACV1PReserved15 is reserved
	NACV1PReserved15 NavigationalAccuracyCategoryPositionV1 = 15
)

// ToString returns a basic, but readable, representation of the field
func (category NavigationalAccuracyCategoryPositionV1) ToString() string {

	switch category {

	case NACV1PEPUGreaterThan18Point52Km:
		return "0 - EPU >= 18.52 km (10 NM) - Unknown accuracy"
	case NACV1PEPULowerThan18Point52Km:
		return "1 - EPU < 18.52 km (10 NM) - RNP-10 accuracy"
	case NACV1PEPULowerThan7Point408Km:
		return "2 - EPU < 7.408 km (4 NM) - RNP-4 accuracy"
	case NACV1PEPULowerThan3Point704Km:
		return "3 - EPU < 3.704 km (2 NM) - RNP-2 accuracy"
	case NACV1PEPUGreaterThan1852M:
		return "4 - EPU < 1 852 m (1 NM) - RNP-1 accuracy"
	case NACV1PEPULowerThan926M:
		return "5 - EPU < 926 m (0.5 NM) - RNP-0.5 accuracy"
	case NACV1PEPUGreaterThan555Point6M:
		return "6 - EPU < 555.6 m ( 0.3 NM) - RNP-0.3 accuracy"
	case NACV1PEPULowerThan185Point2M:
		return "7 - EPU < 185.2 m (0.1 NM) - RNP-0.1 accuracy"
	case NACV1PEPUGreaterThan92Point6M:
		return "8 - EPU < 92.6 m (0.05 NM) - e.g. GPS (with SA)"
	case NACV1PEPULowerThan30MAndVEPULowerThan45M:
		return "9 - EPU < 30 m and VEPU < 45 m - e.g. GPS (SA off)"
	case NACV1PEPULowerThan10MAndVEPULowerThan15M:
		return "10 - EPU < 10 m and VEPU < 15 m - e.g. WAAS"
	case NACV1PEPULowerThan3MAndVEPULowerThan4M:
		return "11 - EPU < 3 m and VEPU < 4 m - e.g. LAAS"
	case NACV1PReserved12:
		return "12 - Reserved"
	case NACV1PReserved13:
		return "13 - Reserved"
	case NACV1PReserved14:
		return "14 - Reserved"
	case NACV1PReserved15:
		return "15 - Reserved"
	default:
		return fmt.Sprintf("%v - Unknown code", category)
	}
}

// ReadNavigationalAccuracyCategoryPositionV1 reads the NavigationalAccuracyCategoryPositionV1 from a 56 bits data field
func ReadNavigationalAccuracyCategoryPositionV1(data []byte) NavigationalAccuracyCategoryPositionV1 {
	bits := data[5] & 0x0F
	return NavigationalAccuracyCategoryPositionV1(bits)
}
