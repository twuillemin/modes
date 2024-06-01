package fields

import "fmt"

// NavigationUncertaintyCategory is the Navigation Uncertainty Category definition
//
// Specified in Doc 9871 / Table A-2-9
type NavigationUncertaintyCategory byte

const (
	// NUCPUnknown indicates Unknown
	NUCPUnknown NavigationUncertaintyCategory = 0
	// NUCPHorizontalLowerThan10VerticalLowerThan15Point2 indicates Horizontal < 10m/s and Vertical < 15.2m/s
	NUCPHorizontalLowerThan10VerticalLowerThan15Point2 NavigationUncertaintyCategory = 1
	// NUCPHorizontalLowerThan3VerticalLowerThan4Point6 indicates Horizontal < 3m/s and Vertical < 4.6m/s
	NUCPHorizontalLowerThan3VerticalLowerThan4Point6 NavigationUncertaintyCategory = 2
	// NUCPHorizontalLowerThan1VerticalLowerThan1Point5 indicates Horizontal < 1m/s and Vertical < 1.5m/s
	NUCPHorizontalLowerThan1VerticalLowerThan1Point5 NavigationUncertaintyCategory = 3
	// NUCPHorizontalLowerThan0Point3VerticalLowerThan0Point46 indicates Horizontal < 0.3m/s and Vertical < 0.46m/s
	NUCPHorizontalLowerThan0Point3VerticalLowerThan0Point46 NavigationUncertaintyCategory = 4
)

var nucpHorizontalLabel = "Horizontal Velocity Error (95%):"
var nucpVerticalLabel = "Vertical Velocity Error (95%):"

// ToString returns a basic, but readable, representation of the field
func (category NavigationUncertaintyCategory) ToString() string {

	switch category {
	case NUCPUnknown:
		return "0 - Unknown "
	case NUCPHorizontalLowerThan10VerticalLowerThan15Point2:
		return fmt.Sprintf("1 - %v %v / %v %v",
			nucpHorizontalLabel, "< 10 m/s",
			nucpVerticalLabel, "< 15.2 m/s (50 fps)")
	case NUCPHorizontalLowerThan3VerticalLowerThan4Point6:
		return fmt.Sprintf("1 - %v %v / %v %v",
			nucpHorizontalLabel, "< 3 m/s",
			nucpVerticalLabel, "< 4.6 m/s (15 fps)")
	case NUCPHorizontalLowerThan1VerticalLowerThan1Point5:
		return fmt.Sprintf("1 - %v %v / %v %v",
			nucpHorizontalLabel, "< 2 m/s",
			nucpVerticalLabel, "< 1.5 m/s (5 fps)")
	case NUCPHorizontalLowerThan0Point3VerticalLowerThan0Point46:
		return fmt.Sprintf("1 - %v %v / %v %v",
			nucpHorizontalLabel, "< 0.3 m/s",
			nucpVerticalLabel, "< 0.46 m/s (1.5 fps)")

	default:
		return fmt.Sprintf("%v - Unknown code", category)
	}
}

// ReadNavigationUncertaintyCategory reads the NavigationUncertaintyCategory from a 56 bits data field
func ReadNavigationUncertaintyCategory(data []byte) NavigationUncertaintyCategory {
	bits := (data[1] & 0x38) >> 3
	return NavigationUncertaintyCategory(bits)
}
