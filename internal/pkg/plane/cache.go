package plane

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/modes/common"
)

var planes = make(map[common.ICAOAddress]*Plane)

var defaultADSBLevel = adsb.Level0OrMore

// SetDefaultADSBLevel defines the level used by default at place creation. By default, the conservative
// Level 0 or more is used
func SetDefaultADSBLevel(level adsb.Level) {
	defaultADSBLevel = level
}

// CheckoutPlane retrieves a plane from the cache. If no plane was present in the cache
// a new plane is created
func CheckoutPlane(timestamp uint32, address common.ICAOAddress) *Plane {

	if knownPlane, ok := planes[address]; ok {
		knownPlane.LastSeenTimestamp = timestamp
		return knownPlane
	}

	newPlane := &Plane{
		ICAOAddress:        address,
		ADSBLevel:          defaultADSBLevel,
		Identification:     "",
		FirstSeenTimestamp: timestamp,
		LastSeenTimestamp:  timestamp,
		Altitude:           0,
		EvenCPRLatitude:    0,
		EvenCPRLongitude:   0,
		EventCPRTimestamp:  0,
		OddCPRLatitude:     0,
		OddCPRLongitude:    0,
		OddCPRTimestamp:    0,
		NICSupplementA:     false,
		NICSupplementC:     false,
	}

	planes[address] = newPlane

	return newPlane

}
