package fields

import "fmt"

// NavigationalAccuracyCategoryPositionV1 is the Navigational Accuracy Category Position definition
//
// Specified in Doc 9871 / B.2.3.9.12
type NavigationalAccuracyCategoryPositionV1 byte

const (
	// NACPV1EPUGreaterThan18Point52Km indicates EPU >= 18.52 km (10 NM) - Unknown accuracy
	NACPV1EPUGreaterThan18Point52Km NavigationalAccuracyCategoryPositionV1 = 0
	// NACPV1EPULowerThan18Point52Km indicates  EPU < 18.52 km (10 NM) - RNP-10 accuracy
	NACPV1EPULowerThan18Point52Km NavigationalAccuracyCategoryPositionV1 = 1
	// NACPV1EPULowerThan7Point408Km indicates EPU < 7.408 km (4 NM) - RNP-4 accuracy
	NACPV1EPULowerThan7Point408Km NavigationalAccuracyCategoryPositionV1 = 2
	// NACPV1EPULowerThan3Point704Km indicates EPU < 3.704 km (2 NM) - RNP-2 accuracy
	NACPV1EPULowerThan3Point704Km NavigationalAccuracyCategoryPositionV1 = 3
	// NACPV1EPUGreaterThan1852M indicates EPU < 1 852 m (1 NM) - RNP-1 accuracy
	NACPV1EPUGreaterThan1852M NavigationalAccuracyCategoryPositionV1 = 4
	// NACPV1EPULowerThan926M indicates EPU < 926 m (0.5 NM) - RNP-0.5 accuracy
	NACPV1EPULowerThan926M NavigationalAccuracyCategoryPositionV1 = 5
	// NACPV1EPUGreaterThan555Point6M indicates EPU < 555.6 m ( 0.3 NM) - RNP-0.3 accuracy
	NACPV1EPUGreaterThan555Point6M NavigationalAccuracyCategoryPositionV1 = 6
	// NACPV1EPULowerThan185Point2M indicates EPU < 185.2 m (0.1 NM) - RNP-0.1 accuracy
	NACPV1EPULowerThan185Point2M NavigationalAccuracyCategoryPositionV1 = 7
	// NACPV1EPUGreaterThan92Point6M indicates EPU < 92.6 m (0.05 NM) - e.g. GPS (with SA)
	NACPV1EPUGreaterThan92Point6M NavigationalAccuracyCategoryPositionV1 = 8
	// NACPV1EPULowerThan30MAndVEPULowerThan45M indicates EPU < 30 m and VEPU < 45 m - e.g. GPS (SA off)
	NACPV1EPULowerThan30MAndVEPULowerThan45M NavigationalAccuracyCategoryPositionV1 = 9
	// NACPV1EPULowerThan10MAndVEPULowerThan15M indicates EPU < 10 m and VEPU < 15 m - e.g. WAAS
	NACPV1EPULowerThan10MAndVEPULowerThan15M NavigationalAccuracyCategoryPositionV1 = 10
	// NACPV1EPULowerThan4MAndVEPULowerThan3M indicates EPU < 3 m and VEPU < 4 m - e.g. LAAS
	NACPV1EPULowerThan4MAndVEPULowerThan3M NavigationalAccuracyCategoryPositionV1 = 11
	// NACPV1Reserved12 is reserved
	NACPV1Reserved12 NavigationalAccuracyCategoryPositionV1 = 12
	// NACPV1Reserved13 is reserved
	NACPV1Reserved13 NavigationalAccuracyCategoryPositionV1 = 13
	// NACPV1Reserved14 is reserved
	NACPV1Reserved14 NavigationalAccuracyCategoryPositionV1 = 14
	// NACPV1Reserved15 is reserved
	NACPV1Reserved15 NavigationalAccuracyCategoryPositionV1 = 15
)

// ToString returns a basic, but readable, representation of the field
func (category NavigationalAccuracyCategoryPositionV1) ToString() string {

	switch category {

	case NACPV1EPUGreaterThan18Point52Km:
		return "0 - EPU >= 18.52 km (10 NM) - Unknown accuracy"
	case NACPV1EPULowerThan18Point52Km:
		return "1 - EPU < 18.52 km (10 NM) - RNP-10 accuracy"
	case NACPV1EPULowerThan7Point408Km:
		return "2 - EPU < 7.408 km (4 NM) - RNP-4 accuracy"
	case NACPV1EPULowerThan3Point704Km:
		return "3 - EPU < 3.704 km (2 NM) - RNP-2 accuracy"
	case NACPV1EPUGreaterThan1852M:
		return "4 - EPU < 1 852 m (1 NM) - RNP-1 accuracy"
	case NACPV1EPULowerThan926M:
		return "5 - EPU < 926 m (0.5 NM) - RNP-0.5 accuracy"
	case NACPV1EPUGreaterThan555Point6M:
		return "6 - EPU < 555.6 m ( 0.3 NM) - RNP-0.3 accuracy"
	case NACPV1EPULowerThan185Point2M:
		return "7 - EPU < 185.2 m (0.1 NM) - RNP-0.1 accuracy"
	case NACPV1EPUGreaterThan92Point6M:
		return "8 - EPU < 92.6 m (0.05 NM) - e.g. GPS (with SA)"
	case NACPV1EPULowerThan30MAndVEPULowerThan45M:
		return "9 - EPU < 30 m and VEPU < 45 m - e.g. GPS (SA off)"
	case NACPV1EPULowerThan10MAndVEPULowerThan15M:
		return "10 - EPU < 10 m and VEPU < 15 m - e.g. WAAS"
	case NACPV1EPULowerThan4MAndVEPULowerThan3M:
		return "11 - EPU < 3 m and VEPU < 4 m - e.g. LAAS"
	case NACPV1Reserved12:
		return "12 - Reserved"
	case NACPV1Reserved13:
		return "13 - Reserved"
	case NACPV1Reserved14:
		return "14 - Reserved"
	case NACPV1Reserved15:
		return "15 - Reserved"
	default:
		return fmt.Sprintf("%v - Unknown code", category)
	}
}

// ReadNavigationalAccuracyCategoryPositionV1 reads the NavigationalAccuracyCategoryPositionV1 from a 56 bits data field
func ReadNavigationalAccuracyCategoryPositionV1(data []byte) NavigationalAccuracyCategoryPositionV1 {
	bits := (data[4]&0x01)<<3 + (data[5]&0xE)>>5
	return NavigationalAccuracyCategoryPositionV1(bits)
}
