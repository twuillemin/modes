package messages

import (
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/fields"
	"testing"
)

func TestReadFormatDF17Valid(t *testing.T) {

	msg, err := ParseDF17(buildValidDF17Message())
	if err != nil {
		t.Fatal(err)
	}

	if msg.Capability != fields.CapabilityLevel2Airborne {
		t.Errorf("Expected Capability \"%v\", got \"%v\"",
			fields.CapabilityLevel2Airborne.ToString(),
			msg.Capability.ToString())
	}

	if msg.AddressAnnounced.Address != 6924272 {
		t.Errorf("Expected Address Announced \"0x00 0x69 0xA7 0xF0\", got \"%v\"",
			msg.AddressAnnounced.ToString())
	}

	if msg.MessageExtendedSquitter[0] != 0x01 {
		t.Errorf("Expected MessageExtendedSquitter[0] to be \"0x01\", got \"%v\"", msg.MessageExtendedSquitter[0])
	}

	if msg.MessageExtendedSquitter[1] != 0x23 {
		t.Errorf("Expected MessageExtendedSquitter[1] to be \"0x23\", got \"%v\"", msg.MessageExtendedSquitter[1])
	}

	if msg.MessageExtendedSquitter[2] != 0x45 {
		t.Errorf("Expected MessageExtendedSquitter[2] to be \"0x45\", got \"%v\"", msg.MessageExtendedSquitter[2])
	}

	if msg.MessageExtendedSquitter[3] != 0x67 {
		t.Errorf("Expected MessageExtendedSquitter[3] to be \"0x67\", got \"%v\"", msg.MessageExtendedSquitter[3])
	}

	if msg.MessageExtendedSquitter[4] != 0x89 {
		t.Errorf("Expected MessageExtendedSquitter[4] to be \"0x89\", got \"%v\"", msg.MessageExtendedSquitter[4])
	}

	if msg.MessageExtendedSquitter[5] != 0xAB {
		t.Errorf("Expected MessageExtendedSquitter[5] to be \"0xAB\", got \"%v\"", msg.MessageExtendedSquitter[5])
	}

	if msg.MessageExtendedSquitter[6] != 0xCD {
		t.Errorf("Expected MessageExtendedSquitter[6] to be \"0xCD\", got \"%v\"", msg.MessageExtendedSquitter[6])
	}

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func buildValidDF17Message() common.MessageData {

	// Format of the message is as follows:
	//
	//     DF     CA   |        AA       |        AA       |        AA       |    ME   |   PI
	// 1 0 0 0 1 c c c | a a a a a a a a | a a a a a a a a | a a a a a a a a | 56 bits | 24bits

	// AddressAnnounced: 0x00 0x69 0xA7 0xF0 = ‭6924272‬

	// 0000 0101: Unused (00000) + Capability: Level2Airborne (101)
	firstField := byte(0x05)

	payload := make([]byte, 10)

	// 0110 1001: Address Announced
	payload[0] = 0x69

	// 1010 0101: Address Announced
	payload[1] = 0xA7

	// 1111 0000: Address Announced
	payload[2] = 0xF0

	// Message Extended Squitter
	payload[3] = 0x01
	payload[4] = 0x23
	payload[5] = 0x45
	payload[6] = 0x67
	payload[7] = 0x89
	payload[8] = 0xAB
	payload[9] = 0xCD

	parity := make([]byte, 3)

	return common.MessageData{
		DownLinkFormat: 17,
		FirstField:     firstField,
		Payload:        payload,
		Parity:         parity,
	}
}
