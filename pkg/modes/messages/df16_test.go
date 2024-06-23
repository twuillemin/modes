package messages

import (
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/fields"
	"testing"
)

func TestReadFormatDF16Valid(t *testing.T) {

	msg, err := ParseDF16(buildValidDF16Message())
	if err != nil {
		t.Fatal(err)
	}

	if msg.VerticalStatus != fields.VerticalStatusOnTheGround {
		t.Errorf("Expected Vertical Status \"%v\", got \"%v\"",
			fields.VerticalStatusOnTheGround.ToString(),
			msg.VerticalStatus.ToString())
	}

	if msg.ReplyInformation != fields.ReplyInformationMaximumAirSpeed140To280 {
		t.Errorf("Expected Reply Information \"%v\", got \"%v\"",
			fields.ReplyInformationMaximumAirSpeed140To280.ToString(),
			msg.ReplyInformation.ToString())
	}

	if msg.Altitude != 33125 {
		t.Errorf("Expected Altitude In Feet \"33125\", got \"%v\"",
			msg.Altitude)
	}

	if msg.AltitudeReportMethod != fields.AltitudeCodeReport25FootIncrements {
		t.Errorf("Expected Reply Information \"%v\", got \"%v\"",
			fields.AltitudeCodeReport25FootIncrements.ToString(),
			msg.AltitudeReportMethod.ToString())
	}

	if msg.MessageACAS[0] != 0x01 {
		t.Errorf("Expected MessageACAS[0] to be \"0x01\", got \"%v\"", msg.MessageACAS[0])
	}

	if msg.MessageACAS[1] != 0x23 {
		t.Errorf("Expected MessageACAS[1] to be \"0x23\", got \"%v\"", msg.MessageACAS[1])
	}

	if msg.MessageACAS[2] != 0x45 {
		t.Errorf("Expected MessageACAS[2] to be \"0x45\", got \"%v\"", msg.MessageACAS[2])
	}

	if msg.MessageACAS[3] != 0x67 {
		t.Errorf("Expected MessageACAS[3] to be \"0x67\", got \"%v\"", msg.MessageACAS[3])
	}

	if msg.MessageACAS[4] != 0x89 {
		t.Errorf("Expected MessageACAS[4] to be \"0x89\", got \"%v\"", msg.MessageACAS[4])
	}

	if msg.MessageACAS[5] != 0xAB {
		t.Errorf("Expected MessageACAS[5] to be \"0xAB\", got \"%v\"", msg.MessageACAS[5])
	}

	if msg.MessageACAS[6] != 0xCD {
		t.Errorf("Expected MessageACAS[6] to be \"0xCD\", got \"%v\"", msg.MessageACAS[6])
	}

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func buildValidDF16Message() common.MessageData {

	// Format of the message is as follows:
	//
	//     DF   VS _ _ |   SL  _ _   RI  |RI _ _     AC    |        AC       |    MV   |  AP
	// 1 0 0 0 0 x _ _ | x x x _ _ x x x | x _ _ x x x x x | x x x x x x x x | 56 bits |24bits

	// Altitude code is a 13 bits fields, so read a uint16
	// bit         |17 18 19 20 21 22 23 24|25 26 27 28 29 30 31 32
	// message     |_  _  _  x  x  x  x  x |x  M  x  Q  x  x  x  x
	// 100 foot    |_  _  _  C1 A1 C2 A2 C4|A4 0  B1 0  B2 D2 B4 D4

	// M: 0
	// Q: 1 binary 25 feet
	// 1 -- 0 1 0 1 -- 0 0 1 1 -- 0 1 0 1 => 101 0101 0101 = 1365 => 33125 feet

	// 0000 0100: Unused (00000) + VS: On the ground (1) +  Reserved (00)
	firstField := byte(0x04)

	payload := make([]byte, 10)

	// 1010 0101: Sensitivity Level 5 (101) + Reserved (00) + RI: Max speed 140 to 280 10 (101 [0])
	payload[0] = 0xA5

	// 0001 0101: RI: Max speed 140 to 280 10 ([101] 0) + Reserved (00) +  Altitude Code (1 0101 [0011 0101])
	payload[1] = 0x15

	// 0011 0101: Altitude Code ([1 0101] 0011 0101)
	payload[2] = 0x35

	// Message ACAS
	payload[3] = 0x01
	payload[4] = 0x23
	payload[5] = 0x45
	payload[6] = 0x67
	payload[7] = 0x89
	payload[8] = 0xAB
	payload[9] = 0xCD

	parity := make([]byte, 16)

	return common.MessageData{
		DownLinkFormat: 16,
		FirstField:     firstField,
		Payload:        payload,
		Parity:         parity,
	}
}
