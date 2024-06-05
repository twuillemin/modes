package fields

import "fmt"

// HorizontalContainmentRadiusV1 is the Horizontal Containment Radius Barometric definition for ADSB V1
//
// Specified in Doc 9871 / B.2.3.1
type HorizontalContainmentRadiusV1 byte

const (
	// HCRBaroV1RcLowerThan7Point5MAndVPLLowerThan11M indicates Rc < 7.5 m and VPL < 11 m
	HCRBaroV1RcLowerThan7Point5MAndVPLLowerThan11M HorizontalContainmentRadiusV1 = 9
	// HCRBaroV1RcLowerThan25MAndVPLLowerThan37Point5M indicates Rc < 25 m and VPL < 37.5 m
	HCRBaroV1RcLowerThan25MAndVPLLowerThan37Point5M HorizontalContainmentRadiusV1 = 10
	// HCRBaroV1RcLowerThan0Point1NM indicates Rc < 0.1 NM (185.2 m)
	HCRBaroV1RcLowerThan0Point1NM HorizontalContainmentRadiusV1 = 11
	// HCRBaroV1RcLowerThan0Point2NM indicates Rc < 0.2 NM (370.4 m)
	HCRBaroV1RcLowerThan0Point2NM HorizontalContainmentRadiusV1 = 12
	// HCRBaroV1RcLowerThan0Point5NM indicates Rc < 0.5 NM (926 m)
	HCRBaroV1RcLowerThan0Point5NM HorizontalContainmentRadiusV1 = 13
	// HCRBaroV1RcLowerThan1Point0NM indicates Rc < 1.0 NM (1852 m)
	HCRBaroV1RcLowerThan1Point0NM HorizontalContainmentRadiusV1 = 14
	// HCRBaroV1RcLowerThan2NM indicates Rc < 2 NM (3.704 km)
	HCRBaroV1RcLowerThan2NM HorizontalContainmentRadiusV1 = 15
	// HCRBaroV1RcLowerThan8NM indicates Rc < 8 NM (14.816 km)
	HCRBaroV1RcLowerThan8NM HorizontalContainmentRadiusV1 = 16
	// HCRBaroV1RcLowerThan20NM indicates Rc < 20 NM (37.04 km)
	HCRBaroV1RcLowerThan20NM HorizontalContainmentRadiusV1 = 17
	// HCRBaroV1RcGreaterThan20NM indicates Rc >= 20 NM (37.04 km) or unknown
	HCRBaroV1RcGreaterThan20NM HorizontalContainmentRadiusV1 = 18

	// HCRBaroV1RcLowerThan75MAndVPLLowerThan112M indicates Rc < 75 m and VPL < 112 m
	HCRBaroV1RcLowerThan75MAndVPLLowerThan112M HorizontalContainmentRadiusV1 = 111
	// HCRBaroV1RcLowerThan0Point6NM indicates Rc < 0.6 NM (1111.2 m)
	HCRBaroV1RcLowerThan0Point6NM HorizontalContainmentRadiusV1 = 113
	// HCRBaroV1RcLowerThan4NM indicates Rc < 4 NM (7.408 km)
	HCRBaroV1RcLowerThan4NM HorizontalContainmentRadiusV1 = 116

	// HCRGNSSV1RcLowerThan7Point5MAndVPLLowerThan11M indicates Rc < 7.5 m and VPL < 11 m
	HCRGNSSV1RcLowerThan7Point5MAndVPLLowerThan11M HorizontalContainmentRadiusV1 = 0
	// HCRGNSSV1RcLowerThan25MAndVPLLowerThan37Point5M indicates Rc < 25 m and VPL < 37.5 m
	HCRGNSSV1RcLowerThan25MAndVPLLowerThan37Point5M HorizontalContainmentRadiusV1 = 1
	// HCRGNSSV1RcGreaterThan25MOrVPLGreaterThan37Point5MOrUnknown indicates Rc >= 25 m or Pc >= 37.5 m or unknown
	HCRGNSSV1RcGreaterThan25MOrVPLGreaterThan37Point5MOrUnknown HorizontalContainmentRadiusV1 = 2
)

// ToString returns a basic, but readable, representation of the field
func (hcr HorizontalContainmentRadiusV1) ToString() string {
	switch hcr {
	case HCRBaroV1RcLowerThan7Point5MAndVPLLowerThan11M:
		return "(Baro) Rc < 7.5 m and VPL < 11 m"
	case HCRBaroV1RcLowerThan25MAndVPLLowerThan37Point5M:
		return "(Baro) Rc < 25 m and VPL < 37.5 m"
	case HCRBaroV1RcLowerThan0Point1NM:
		return "(Baro) Rc < 0.1 NM (185.2 m)"
	case HCRBaroV1RcLowerThan0Point2NM:
		return "(Baro) Rc < 0.2 NM (370.4 m)"
	case HCRBaroV1RcLowerThan0Point5NM:
		return "(Baro) Rc < 0.5 NM (926 m)"
	case HCRBaroV1RcLowerThan1Point0NM:
		return "(Baro) Rc < 1.0 NM (1852 m)"
	case HCRBaroV1RcLowerThan2NM:
		return "(Baro) Rc < 2 NM (3.704 km)"
	case HCRBaroV1RcLowerThan8NM:
		return "(Baro) Rc < 8 NM (14.816 km)"
	case HCRBaroV1RcLowerThan20NM:
		return "(Baro) Rc < 20 NM (37.04 km)"
	case HCRBaroV1RcGreaterThan20NM:
		return "(Baro) Rc >= 20 NM (37.04 km) or unknown"

	case HCRBaroV1RcLowerThan75MAndVPLLowerThan112M:
		return "(Baro) Rc < 75 m and VPL < 112 m"
	case HCRBaroV1RcLowerThan0Point6NM:
		return "(Baro) Rc < 0.6 NM (1111.2 m)"
	case HCRBaroV1RcLowerThan4NM:
		return "(Baro) Rc < 4 NM (7.408 km)"

	case HCRGNSSV1RcLowerThan7Point5MAndVPLLowerThan11M:
		return "(GNSS) Rc < 7.5 m and VPL < 11 m"
	case HCRGNSSV1RcLowerThan25MAndVPLLowerThan37Point5M:
		return "(GNSS) Rc < 25 m and VPL < 37.5 m"
	case HCRGNSSV1RcGreaterThan25MOrVPLGreaterThan37Point5MOrUnknown:
		return "(GNSS) Rc >= 25 m or VPL >= 37.5 m or unknown"

	default:
		return fmt.Sprintf("%v - Unknown code", hcr)
	}
}
