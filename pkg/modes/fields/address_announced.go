package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
)

// -----------------------------------------------------------------------------------------
//
//                                 Address Announced (AA)
//
// -----------------------------------------------------------------------------------------

// AddressAnnounced (AA) field shall contain the aircraft address which provides unambiguous identification of
// the aircraft.
//
// Defined at 3.1.2.5.2.2.2
type AddressAnnounced struct {
	Address uint32
}

// ReadAddressAnnounced reads the AA field from a message
func ReadAddressAnnounced(message common.MessageData) AddressAnnounced {

	return AddressAnnounced{
		Address: uint32(message.Payload[0])<<16 | uint32(message.Payload[1])<<8 | uint32(message.Payload[2]),
	}
}

// ToString returns a basic, but readable, representation of the field
func (addressAnnounced AddressAnnounced) ToString() string {
	return fmt.Sprintf("%X %X %X",
		uint8((addressAnnounced.Address&0x00FF0000)>>16),
		uint8((addressAnnounced.Address&0x0000FF00)>>8),
		uint8(addressAnnounced.Address&0x000000FF))
}
