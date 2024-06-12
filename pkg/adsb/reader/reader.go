package reader

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/airborne_position"
	"github.com/twuillemin/modes/pkg/adsb/airborne_velocity"
	"github.com/twuillemin/modes/pkg/adsb/aircraft_identification_and_category"
	"github.com/twuillemin/modes/pkg/adsb/aircraft_operational_status"
	"github.com/twuillemin/modes/pkg/adsb/aircraft_status"
	"github.com/twuillemin/modes/pkg/adsb/no_position_information"
	"github.com/twuillemin/modes/pkg/adsb/surface_position"
	"github.com/twuillemin/modes/pkg/adsb/target_state_and_status"
)

// ReadADSBMessage reads and parse an ADSB message.
//
// params:
//   - adsbLevel: The ADSB level request (not used, but present for coherency)
//   - nicSupplementA: The NIC Supplement-A comes from the Aircraft  Operational  Status - Message Type Format 31 (see
//     C.2.3.10.20). If no previous Type Format 31 message was received before calling this function, a
//     default value of 0 can be used.
//     Note: This value is name simply NIC Supplement in ADSB V1
//   - nicSupplementC: The NIC Supplement-C comes from the Surface Capability Class (CC) Code  Subfield  of  the
//     Aircraft  Operational  Status - Message Type Format 31 (see  C.2.3.10.20). If no previous Type
//     Format 31 message was received before calling this function, a default value of 0 can be used.
//     Note: This value does is only present since ADSB V2
//   - data: The body of the message. The message must be 7 bytes long
//
// Return the parsed message and an optional error. The detected ADSB ReaderLevel will generally be
// the same as the given one, except if the decoded message has information to change it.
func ReadADSBMessage(
	adsbLevel adsb.ADSBVersion,
	nicSupplementA bool,
	nicSupplementC bool,
	data []byte) (adsb.Message, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for ADSB message must be 7 bytes long")
	}

	// -----------------------------------------------------
	// Code   BDS  Description           V0      V1      V2
	//  0      ?   No Position
	//  1     0,8  Aircraft Id           OK      OK      OK
	//  2     0,8  Aircraft Id           OK      OK      OK
	//  3     0,8  Aircraft Id           OK      OK      OK
	//  4     0,8  Aircraft Id           OK      OK      OK
	//  5     0,6  Surface position      OK      OK      OK
	//  6     0,6  Surface position      OK      OK      OK
	//  7     0,6  Surface position      OK      OK      OK
	//  8     0,6  Surface position      OK      OK      OK
	//  9     0,5  Airborne position     OK      OK      OK
	// 10     0,5  Airborne position     OK      OK      OK
	// 14     0,5  Airborne position     OK      OK      OK
	// 12     0,5  Airborne position     OK      OK      OK
	// 13     0,5  Airborne position     OK      OK      OK
	// 14     0,5  Airborne position     OK      OK      OK
	// 15     0,5  Airborne position     OK      OK      OK
	// 16     0,5  Airborne position     OK      OK      OK
	// 17     0,5  Airborne position     OK      OK      OK
	// 18     0,5  Airborne position     OK      OK      OK
	// 19     0,9  Airborne velocity     OK      OK      OK
	// 20     0,5  Airborne position     OK      OK      OK
	// 21     0,5  Airborne position     OK      OK      OK
	// 22     0,5  Airborne position     OK      OK      OK
	// 23          Reserved
	// 24          Reserved
	// 25          Reserved
	// 26          Reserved
	// 27          Reserved
	// 28     6,1  Emergency report      OK      OK      OK
	// 29     6,2  Target and status     __      OK      OK
	// 30          Reserved
	// 31     6,5  Operational status    OK      OK      OK

	// Get the type
	formatTypeCode := (data[0] & 0xF8) >> 3

	switch formatTypeCode {
	case 0:
		if data[2]&0x0F == 0 && data[3] == 0 && data[4] == 0 && data[5] == 0 && data[6] == 0 {
			return no_position_information.ReadNoPositionInformation(adsbLevel, data)
		} else {
			return airborne_position.ReadAirbornePositionType0(adsbLevel, data)
		}
	case 1, 2, 3, 4:
		return aircraft_identification_and_category.ReadAircraftIdentificationAndCategory(adsbLevel, data)
	case 5, 6, 7, 8:
		return surface_position.ReadSurfacePosition(adsbLevel, data, nicSupplementA, nicSupplementC)
	case 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22:
		return airborne_position.ReadAirbornePosition(adsbLevel, data, nicSupplementA)
	case 19:
		return airborne_velocity.ReadAirborneVelocity(adsbLevel, data)
	case 28:
		return aircraft_status.ReadAircraftStatus(adsbLevel, data)
	case 29:
		return target_state_and_status.ReadTargetStateAndStatus(adsbLevel, data)
	case 31:
		return aircraft_operational_status.ReadAircraftOperationalStatus(adsbLevel, data)
	}

	return nil, fmt.Errorf("the formatTypeCode %v is not supported", formatTypeCode)
}
