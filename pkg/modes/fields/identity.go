package fields

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/modes/common"
)

// -----------------------------------------------------------------------------------------
//
//                                 Identity (ID)
//
// -----------------------------------------------------------------------------------------

// Identity (ID) field shall contain aircraft identity code, in accordance with the pattern for Mode A replies
//
// Defined at 3.1.2.6.7.1
type Identity struct {
	// Identity is the identity, which is always a 4 digit chars.
	Identity string
}

// IsSpecialEmergency informs that the identity is a special code for an aircraft in emergency
// Identity string is "7700".
func (identity Identity) IsSpecialEmergency() bool {
	return identity.Identity == "7700"
}

// IsSpecialRadiocommunicationFailure informs that the identity is a special code for an aircraft with
// radiocommunication failure. Identity string is "7600".
func (identity Identity) IsSpecialRadiocommunicationFailure() bool {
	return identity.Identity == "7600"
}

// IsSpecialUnlawfulInterference informs that the identity is a special code for an aircraft t which is being
// subjected to unlawful interference. Identity string is "7500".
func (identity Identity) IsSpecialUnlawfulInterference() bool {
	return identity.Identity == "7500"
}

// IsSpecialInstructionNotReceived informs that the identity is a special code for an aircraft which has not
// received any instructions from air traffic control units to operate the transponder. Identity string is "2000".
func (identity Identity) IsSpecialInstructionNotReceived() bool {
	return identity.Identity == "2000"
}

// ReadIdentity reads the identity from a message
func ReadIdentity(message common.MessageData) Identity {

	// Identity is a 13 bits fields, so read a uint16
	// bit         |17 18 19 20 21 22 23 24|25 26 27 28 29 30 31 32
	// id bits     |_  _  _  C1 A1 C2 A2 C4|A4 0  B1 D1  B2 D2 B4 D4

	// Get the raw identity code
	identity := (uint16(message.Payload[1])<<8 | uint16(message.Payload[2])) & 0x1fff

	// Starting with bit 20 the sequence shall be C1, A1, C2, A2, C4, A4, ZERO, B1, D1, B2, D2, B4, D4.
	c1 := (identity & 0x1000) >> 12
	a1 := (identity & 0x0800) >> 11
	c2 := (identity & 0x0400) >> 10
	a2 := (identity & 0x0200) >> 9
	c4 := (identity & 0x0100) >> 8
	a4 := (identity & 0x0080) >> 7
	b1 := (identity & 0x0020) >> 5
	d1 := (identity & 0x0010) >> 4
	b2 := (identity & 0x0008) >> 3
	d2 := (identity & 0x0004) >> 2
	b4 := (identity & 0x0002) >> 1
	d4 := identity & 0x0001

	// Defined at 3.1.1.6
	a := a4<<2 | a2<<1 | a1
	b := b4<<2 | b2<<1 | b1
	c := c4<<2 | c2<<1 | c1
	d := d4<<2 | d2<<1 | d1

	return Identity{
		Identity: fmt.Sprintf("%v%v%v%v", a, b, c, d),
	}
}

// ToString returns a basic, but readable, representation of the field
func (identity Identity) ToString() string {

	special := ""

	if identity.IsSpecialEmergency() {
		special = " [Emergency]"
	} else if identity.IsSpecialInstructionNotReceived() {
		special = " [Instruction Not Received]"
	} else if identity.IsSpecialRadiocommunicationFailure() {
		special = " [Radiocommunication Failure]"
	} else if identity.IsSpecialUnlawfulInterference() {
		special = " [Unlawful Interference]"
	}
	return identity.Identity + special
}

// ToExtendedString returns a complete representation of the field
func (identity Identity) ToExtendedString() string {

	special := ""

	if identity.IsSpecialEmergency() {
		special = " [Emergency: aircraft in emergency]"
	} else if identity.IsSpecialInstructionNotReceived() {
		special = " [Instruction Not Received: instructions from air traffic control units to operate the transponder not received]"
	} else if identity.IsSpecialRadiocommunicationFailure() {
		special = " [Radiocommunication Failure: aircraft with radiocommunication failure]"
	} else if identity.IsSpecialUnlawfulInterference() {
		special = " [Unlawful Interference: aircraft subjected to unlawful interference]"
	}
	return identity.Identity + special
}
