package reader

import (
	"errors"
	"github.com/twuillemin/modes/pkg/adsb/messages"
	messages05 "github.com/twuillemin/modes/pkg/bds/bds05/messages"
	messages06 "github.com/twuillemin/modes/pkg/bds/bds06/messages"
	messages08 "github.com/twuillemin/modes/pkg/bds/bds08/messages"
	messages09 "github.com/twuillemin/modes/pkg/bds/bds09/messages"
	messages61 "github.com/twuillemin/modes/pkg/bds/bds61/messages"
	messages62 "github.com/twuillemin/modes/pkg/bds/bds62/messages"
	messages65 "github.com/twuillemin/modes/pkg/bds/bds65/messages"
	"github.com/twuillemin/modes/pkg/bds/common"
)

// ReadMessage reads and parse an ADSB message.
//
// params:
//    - adsbLevel: The ADSB level request (not used, but present for coherency)
//    - nicSupplementA: The NIC Supplement-A comes from the Aircraft  Operational  Status - Message Type Format 31 (see
//                      C.2.3.10.20). If no previous Type Format 31 message was received before calling this function, a
//                      default value of 0 can be used.
//                      Note: This value is name simply NIC Supplement in ADSB V1
//    - nicSupplementC: The NIC Supplement-C comes from the Surface Capability Class (CC) Code  Subfield  of  the
//                      Aircraft  Operational  Status - Message Type Format 31 (see  C.2.3.10.20). If no previous Type
//                      Format 31 message was received before calling this function, a default value of 0 can be used.
//                      Note: This value does is only present since ADSB V2
//    - message: The body of the message. The message must be 7 bytes long
//
// Return the parsed message or an error
func ReadMessage(
	adsbLevel common.ADSBLevel,
	nicSupplementA bool,
	nicSupplementC bool,
	data []byte) (messages.ADSBMessage, common.ADSBLevel, error) {

	if len(data) != 7 {
		return nil, adsbLevel, errors.New("the data for ADSB message must be 7 bytes long")
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
	case 1, 2, 3, 4:
		return messages08.ReadBDS08(adsbLevel, data)
	case 5, 6, 7, 8:
		return messages06.ReadBDS06(adsbLevel, nicSupplementA, nicSupplementC, data)
	case 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22:
		return messages05.ReadBDS05(adsbLevel, nicSupplementA, data)
	case 19:
		return messages09.ReadBDS09(adsbLevel, data)
	case 28:
		return messages61.ReadBDS61(adsbLevel, data)
	case 29:
		return messages62.ReadBDS62(adsbLevel, data)
	case 31:
		return messages65.ReadBDS65(adsbLevel, data)
	}

	return nil, adsbLevel, nil
}
