package messages

import (
	"errors"
	"fmt"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/bds08/fields"
)

// MessageBDS08 is the basic interface that ADSB messages at the format BDS 0,8 are expected to implement
type MessageBDS08 interface {
	adsb.Message

	// GetAircraftCategory returns the aircraft category
	GetAircraftCategory() fields.AircraftCategory
	// GetAircraftIdentification returns the identity of the aircraft
	GetAircraftIdentification() fields.AircraftIdentification
}

func bds08ToString(message MessageBDS08) string {
	return fmt.Sprintf("Message:                 %v\n"+
		"Aircraft Category:       %v (%v)\n"+
		"Aircraft Identification: %v",
		adsb.GetMessageFormatInformation(message),
		message.GetAircraftCategory().ToString(),
		message.GetAircraftCategory().GetCategorySetName(),
		message.GetAircraftIdentification())
}

// ReadBDS08 reads a message at the format BDS 0,8. As this format does not have changes from ADSB V0 to
// ADSB V2, the returned ADSBLevel is always the given one.
//
// Params:
//   - adsbLevel: The ADSB level request (currently unused)
//   - data: The data of the message must be 7 bytes
//
// Returns the message read or an error
func ReadBDS08(_ adsb.ReaderLevel, data []byte) (MessageBDS08, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for BDS 0,8 message must be 7 bytes long")
	}

	formatTypeCode := (data[0] & 0xF8) >> 3

	switch formatTypeCode {
	case 1:
		message, err := ReadFormat01(data)
		return message, err
	case 2:
		message, err := ReadFormat02(data)
		return message, err
	case 3:
		message, err := ReadFormat03(data)
		return message, err
	case 4:
		message, err := ReadFormat04(data)
		return message, err
	default:
		return nil, fmt.Errorf("the Format Type %v of BSD 0,8 is not supported", formatTypeCode)
	}
}
