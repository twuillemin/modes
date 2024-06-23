package messages

import (
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/fields"
	"testing"
)

func TestReadFormatDF0Valid(t *testing.T) {

	msg, err := ParseDF0(buildValidDF0Message())
	if err != nil {
		t.Fatal(err)
	}

	if msg.VerticalStatus != fields.VerticalStatusOnTheGround {
		t.Errorf("Expected Vertical Status \"%v\", got \"%v\"",
			fields.VerticalStatusOnTheGround.ToString(),
			msg.VerticalStatus.ToString())
	}

	if msg.CrossLinkCapability != fields.CrossLinkCompatibilitySupported {
		t.Errorf("Expected Cross Link Capability \"%v\", got \"%v\"",
			fields.CrossLinkCompatibilitySupported.ToString(),
			msg.CrossLinkCapability.ToString())
	}

	if msg.ReplyInformation != fields.ReplyInformationMaximumAirSpeed140To280 {
		t.Errorf("Expected Reply Information \"%v\", got \"%v\"",
			fields.ReplyInformationMaximumAirSpeed140To280.ToString(),
			msg.ReplyInformation.ToString())
	}

	if msg.Altitude != 33125 {
		t.Errorf("Expected Altitude In Feet \"33125\", got \"%v\"",
			msg)
	}

	if msg.AltitudeReportMethod != fields.AltitudeCodeReport25FootIncrements {
		t.Errorf("Expected Reply Information \"%v\", got \"%v\"",
			fields.AltitudeCodeReport25FootIncrements.ToString(),
			msg.AltitudeReportMethod.ToString())
	}

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func buildValidDF0Message() common.MessageData {

	// Format of the message is as follows:
	//
	//     DF  VS CC _ |   SL  _ _   RI  |RI _ _     AC    |        AC       |   AP
	// 0 0 0 0 0 x x _ | x x x _ _ x x x | x _ _ x x x x x | x x x x x x x x | 24bits

	// Altitude code is a 13 bits fields, so read a uint16
	// bit         |17 18 19 20 21 22 23 24|25 26 27 28 29 30 31 32
	// message     |_  _  _  x  x  x  x  x |x  M  x  Q  x  x  x  x
	// 100 foot    |_  _  _  C1 A1 C2 A2 C4|A4 0  B1 0  B2 D2 B4 D4

	// M: 0
	// Q: 1 binary 25 feet
	// 1 -- 0 1 0 1 -- 0 0 1 1 -- 0 1 0 1 => 101 0101 0101 = 1365 => 33125 feet

	// 0000 0110: Unused (00000) + VS: On the ground (1) + CC: Supported (1) + Reserved (0)
	firstField := byte(0x06)

	payload := make([]byte, 3)

	// 1010 0101: Sensitivity Level 5 (101) + Reserved (00) + RI: Max speed 140 to 280 10 (101 [0])
	payload[0] = 0xA5

	// 0001 0101: RI: Max speed 140 to 280 10 ([101] 0) + Reserved (00) +  Altitude Code (1 0101 [0011 0101])
	payload[1] = 0x15

	// 0011 0101: Altitude Code ([1 0101] 0011 0101)
	payload[2] = 0x35

	parity := make([]byte, 3)

	return common.MessageData{
		DownLinkFormat: 0,
		FirstField:     firstField,
		Payload:        payload,
		Parity:         parity,
	}
}
