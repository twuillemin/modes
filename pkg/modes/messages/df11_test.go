package messages

import (
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/fields"
	"testing"
)

func TestReadFormatDF11Valid(t *testing.T) {

	msg, err := ParseDF11(buildValidDF11Message())
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

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func buildValidDF11Message() common.MessageData {

	// Format of the message is as follows:
	//
	//     DF     CA   |        AA       |        AA       |        AA       |   PI
	// 0 1 0 1 1 c c c | a a a a a a a a | a a a a a a a a | a a a a a a a a | 24bits

	// AddressAnnounced: 0x00 0x69 0xA7 0xF0 = ‭6924272‬

	// 0000 0101: Unused (00000) + Capability: Level2Airborne (101)
	firstField := byte(0x05)

	payload := make([]byte, 3)

	// 0110 1001: Address Announced
	payload[0] = 0x69

	// 1010 0101: Address Announced
	payload[1] = 0xA7

	// 1111 0000: Address Announced
	payload[2] = 0xF0

	parity := make([]byte, 3)

	return common.MessageData{
		DownLinkFormat: 11,
		FirstField:     firstField,
		Payload:        payload,
		Parity:         parity,
	}
}
