package fields

import "fmt"

// HorizontalContainmentRadiusV2 is the Horizontal Containment Radius Barometric definition for ADSB V2
//
// Specified in Doc 9871 / B.2.3.1
type HorizontalContainmentRadiusV2 byte

const (
	// HCRBaroV2RcLowerThan7Point5M indicates Rc < 7.5 m
	HCRBaroV2RcLowerThan7Point5M HorizontalContainmentRadiusV2 = 9
	// HCRBaroV2RcLowerThan25M indicates Rc < 25 m
	HCRBaroV2RcLowerThan25M HorizontalContainmentRadiusV2 = 10
	// HCRBaroV2RcLowerThan0Point1NM indicates Rc < 0.1 NM (185.2 m)
	HCRBaroV2RcLowerThan0Point1NM HorizontalContainmentRadiusV2 = 11
	// HCRBaroV2RcLowerThan0Point2NM indicates Rc < 0.2 NM (370.4 m)
	HCRBaroV2RcLowerThan0Point2NM HorizontalContainmentRadiusV2 = 12
	// HCRBaroV2RcLowerThan0Point5NM indicates Rc < 0.5 NM (926 m)
	HCRBaroV2RcLowerThan0Point5NM HorizontalContainmentRadiusV2 = 13
	// HCRBaroV2RcLowerThan1Point0NM indicates Rc < 1.0 NM (1852 m)
	HCRBaroV2RcLowerThan1Point0NM HorizontalContainmentRadiusV2 = 14
	// HCRBaroV2RcLowerThan2NM indicates Rc < 2 NM (3.704 km)
	HCRBaroV2RcLowerThan2NM HorizontalContainmentRadiusV2 = 15
	// HCRBaroV2RcLowerThan8NM indicates Rc < 8 NM (14.816 km)
	HCRBaroV2RcLowerThan8NM HorizontalContainmentRadiusV2 = 16
	// HCRBaroV2RcLowerThan20NM indicates Rc < 20 NM (37.04 km)
	HCRBaroV2RcLowerThan20NM HorizontalContainmentRadiusV2 = 17
	// HCRBaroV2RcGreaterThan20NM indicates Rc >= 20 NM (37.04 km) or unknown
	HCRBaroV2RcGreaterThan20NM HorizontalContainmentRadiusV2 = 18

	// HCRBaroV2RcLowerThan75M indicates Rc < 75 m
	HCRBaroV2RcLowerThan75M HorizontalContainmentRadiusV2 = 111
	// HCRBaroV2RcLowerThan0Point6NM indicates Rc < 0.6 NM (1111.2 m)
	HCRBaroV2RcLowerThan0Point6NM HorizontalContainmentRadiusV2 = 113
	// HCRBaroV2RcLowerThan4NM indicates Rc < 4 NM (7.408 km)
	HCRBaroV2RcLowerThan4NM HorizontalContainmentRadiusV2 = 116

	// HCRBaroV2RcLowerThan0Point3NM indicates Rc < 0.3 NM (555.6 m)
	HCRBaroV2RcLowerThan0Point3NM HorizontalContainmentRadiusV2 = 213

	// HCRGNSSV2RcLowerThan7Point5M indicates Rc < 7.5 m
	HCRGNSSV2RcLowerThan7Point5M HorizontalContainmentRadiusV2 = 20
	// HCRGNSSV2RcLowerThan25M indicates Rc < 25 m
	HCRGNSSV2RcLowerThan25M HorizontalContainmentRadiusV2 = 21
	// HCRGNSSV2RcGreaterThan25MOrUnknown indicates Rc >= 25 m or unknown
	HCRGNSSV2RcGreaterThan25MOrUnknown HorizontalContainmentRadiusV2 = 22
)

// ToString returns a basic, but readable, representation of the field
func (hcr HorizontalContainmentRadiusV2) ToString() string {
	switch hcr {
	case HCRBaroV2RcLowerThan7Point5M:
		return "(Baro) Rc < 7.5 m"
	case HCRBaroV2RcLowerThan25M:
		return "(Baro) Rc < 25 m"
	case HCRBaroV2RcLowerThan0Point1NM:
		return "(Baro) Rc < 0.1 NM (185.2 m)"
	case HCRBaroV2RcLowerThan0Point2NM:
		return "(Baro) Rc < 0.2 NM (370.4 m)"
	case HCRBaroV2RcLowerThan0Point5NM:
		return "(Baro) Rc < 0.5 NM (926 m)"
	case HCRBaroV2RcLowerThan1Point0NM:
		return "(Baro) Rc < 1.0 NM (1852 m)"
	case HCRBaroV2RcLowerThan2NM:
		return "(Baro) Rc < 2 NM (3.704 km)"
	case HCRBaroV2RcLowerThan8NM:
		return "(Baro) Rc < 8 NM (14.816 km)"
	case HCRBaroV2RcLowerThan20NM:
		return "(Baro) Rc < 20 NM (37.04 km)"
	case HCRBaroV2RcGreaterThan20NM:
		return "(Baro) Rc >= 20 NM (37.04 km) or unknown"

	case HCRBaroV2RcLowerThan75M:
		return "(Baro) Rc < 75 m"
	case HCRBaroV2RcLowerThan0Point6NM:
		return "(Baro) Rc < 0.6 NM (1111.2 m)"
	case HCRBaroV2RcLowerThan4NM:
		return "(Baro) Rc < 4 NM (7.408 km)"

	case HCRBaroV2RcLowerThan0Point3NM:
		return "(Baro) Rc < 0.3 NM (555.6 m)"

	case HCRGNSSV2RcLowerThan7Point5M:
		return "(GNSS) Rc < 7.5 m"
	case HCRGNSSV2RcLowerThan25M:
		return "(GNSS) Rc < 25 m"
	case HCRGNSSV2RcGreaterThan25MOrUnknown:
		return "(GNSS) Rc >= 25 m or unknown"

	default:
		return fmt.Sprintf("%v - Unknown code", hcr)
	}
}
