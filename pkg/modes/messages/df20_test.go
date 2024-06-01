package messages

import (
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/fields"
	"testing"
)

func TestReadFormatDF20Valid(t *testing.T) {

	msg, err := ParseDF20(buildValidDF20Message())
	if err != nil {
		t.Fatal(err)
	}

	if msg.FlightStatus != fields.FlightStatusAlertSPIOnTheGround {
		t.Errorf("Expected Flight Status \"%v\", got \"%v\"",
			fields.FlightStatusAlertSPIOnTheGround.ToString(),
			msg.FlightStatus.ToString())
	}

	if msg.DownlinkRequest != fields.DownlinkRequestELMAvailable6Segments {
		t.Errorf("Expected Downlink Request \"%v\", got \"%v\"",
			fields.DownlinkRequestELMAvailable6Segments.ToString(),
			msg.DownlinkRequest.ToString())
	}

	if msg.UtilityMessage.InterrogatorIdentifier != 5 {
		t.Errorf("Expected Utility Message Interrogator Identifier to be \"5\", got \"%v\"", msg.UtilityMessage.InterrogatorIdentifier)
	}

	if msg.UtilityMessage.IdentifierDesignator != fields.UtilityMessageIdentifierDesignatorCommB {
		t.Errorf("Expected Utility Message Interrogator Identifier \"%v\", got \"%v\"",
			fields.UtilityMessageIdentifierDesignatorCommB.ToString(),
			msg.UtilityMessage.IdentifierDesignator.ToString())
	}

	if msg.AltitudeCode.AltitudeInFeet != 33125 {
		t.Errorf("Expected Altitude In Feet \"33125\", got \"%v\"",
			msg.AltitudeCode.AltitudeInFeet)
	}

	if msg.AltitudeCode.ReportMethod != fields.AltitudeCodeReport25FootIncrements {
		t.Errorf("Expected Reply Information \"%v\", got \"%v\"",
			fields.AltitudeCodeReport25FootIncrements.ToString(),
			msg.AltitudeCode.ReportMethod.ToString())
	}

	if msg.MessageCommB[0] != 0x01 {
		t.Errorf("Expected MessageCommB[0] to be \"0x01\", got \"%v\"", msg.MessageCommB[0])
	}

	if msg.MessageCommB[1] != 0x23 {
		t.Errorf("Expected MessageCommB[1] to be \"0x23\", got \"%v\"", msg.MessageCommB[1])
	}

	if msg.MessageCommB[2] != 0x45 {
		t.Errorf("Expected MessageCommB[2] to be \"0x45\", got \"%v\"", msg.MessageCommB[2])
	}

	if msg.MessageCommB[3] != 0x67 {
		t.Errorf("Expected MessageCommB[3] to be \"0x67\", got \"%v\"", msg.MessageCommB[3])
	}

	if msg.MessageCommB[4] != 0x89 {
		t.Errorf("Expected MessageCommB[4] to be \"0x89\", got \"%v\"", msg.MessageCommB[4])
	}

	if msg.MessageCommB[5] != 0xAB {
		t.Errorf("Expected MessageCommB[5] to be \"0xAB\", got \"%v\"", msg.MessageCommB[5])
	}

	if msg.MessageCommB[6] != 0xCD {
		t.Errorf("Expected MessageCommB[6] to be \"0xCD\", got \"%v\"", msg.MessageCommB[6])
	}

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func buildValidDF20Message() common.MessageData {

	// Format of the message is as follows:
	//
	//     DF     FS   |      DR     UM  |   UM      AC    |        AC       |  Comm-B |  AP/DP
	// 1 0 1 0 0 f f f | d d d d d u u u | u u u a a a a a | a a a a a a a a | 56 bits | 24bits

	// Altitude code is a 13 bits fields, so read a uint16
	// bit         |17 18 19 20 21 22 23 24|25 26 27 28 29 30 31 32
	// message     |_  _  _  x  x  x  x  x |x  M  x  Q  x  x  x  x
	// 100 foot    |_  _  _  C1 A1 C2 A2 C4|A4 0  B1 0  B2 D2 B4 D4

	// M: 0
	// Q: 1 binary 25 feet
	// 1 -- 0 1 0 1 -- 0 0 1 1 -- 0 1 0 1 => 101 0101 0101 = 1365 => 33125 feet

	// 0000 0101: Unused (00000) + FS: On the ground + alert SPI (101)
	firstField := byte(0x05)

	payload := make([]byte, 10)

	// 1010 1010: DR : ELMAvailable6Segments (10101) + Utility message: IIS = 5 : 0101 + IDS CommB 01 (010[101])
	payload[0] = 0xAA

	// 1011 0101: Utility message: IIS : 0101 + IDS CommC 01 ([010]101) +  Altitude Code (1 0101 [0011 0101])
	payload[1] = 0xB5

	// 0011 0101: Altitude Code ([1 0101] 0011 0101)
	payload[2] = 0x35

	// CommB message
	payload[3] = 0x01
	payload[4] = 0x23
	payload[5] = 0x45
	payload[6] = 0x67
	payload[7] = 0x89
	payload[8] = 0xAB
	payload[9] = 0xCD

	parity := make([]byte, 3)

	return common.MessageData{
		DownLinkFormat: 20,
		FirstField:     firstField,
		Payload:        payload,
		Parity:         parity,
	}
}
