package messages

import (
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/fields"
	"testing"
)

func TestReadFormatDF24Valid(t *testing.T) {

	msg, err := ParseDF24(buildValidDF24Message())
	if err != nil {
		t.Fatal(err)
	}

	if msg.ControlELM != fields.UplinkELMAcknowledgement {
		t.Errorf("Expected Control ELM \"%v\", got \"%v\"",
			fields.UplinkELMAcknowledgement.ToString(),
			msg.ControlELM.ToString())
	}

	if msg.NumberOfDSegment != 10 {
		t.Errorf("Expected Number Of D Segment to \"10\", got \"%v\"", msg.NumberOfDSegment)
	}

	if msg.MessageCommD[0] != 0xFF {
		t.Errorf("Expected MessageCommD[0] to be \"0xFF\", got \"%v\"", msg.MessageCommD[0])
	}

	if msg.MessageCommD[1] != 0x00 {
		t.Errorf("Expected MessageCommD[1] to be \"0x00\", got \"%v\"", msg.MessageCommD[1])
	}

	if msg.MessageCommD[2] != 0xAA {
		t.Errorf("Expected MessageCommD[2] to be \"0xAA\", got \"%v\"", msg.MessageCommD[2])
	}

	if msg.MessageCommD[3] != 0x01 {
		t.Errorf("Expected MessageCommD[0] to be \"0x01\", got \"%v\"", msg.MessageCommD[3])
	}

	if msg.MessageCommD[4] != 0x23 {
		t.Errorf("Expected MessageCommD[1] to be \"0x23\", got \"%v\"", msg.MessageCommD[4])
	}

	if msg.MessageCommD[5] != 0x45 {
		t.Errorf("Expected MessageCommD[2] to be \"0x45\", got \"%v\"", msg.MessageCommD[5])
	}

	if msg.MessageCommD[6] != 0x67 {
		t.Errorf("Expected MessageCommD[3] to be \"0x67\", got \"%v\"", msg.MessageCommD[6])
	}

	if msg.MessageCommD[7] != 0x89 {
		t.Errorf("Expected MessageCommD[4] to be \"0x89\", got \"%v\"", msg.MessageCommD[7])
	}

	if msg.MessageCommD[8] != 0xAB {
		t.Errorf("Expected MessageCommD[5] to be \"0xAB\", got \"%v\"", msg.MessageCommD[8])
	}

	if msg.MessageCommD[9] != 0xCD {
		t.Errorf("Expected MessageCommD[6] to be \"0xCD\", got \"%v\"", msg.MessageCommD[9])
	}

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func buildValidDF24Message() common.MessageData {

	// Format of the message is as follow:
	//
	//  DF _ KE   ND   |  Comm-D |   AP
	// 1 1 _ k n n n n | 80 bits | 24bits

	// downlinkFormat 1 1 + Reserved (0) + Control ELM UplinkELMAcknowledgement (1) + Number of D segment: 10 (1[010])
	// downlinkFormat 1 1011
	downlinkFormat := byte(27)

	// 0000 0101: Unused (00000) + Number of D segment: 10 ([1]010)
	firstField := byte(0x02)

	// Not used
	payload := make([]byte, 10)
	payload[0] = 0xFF
	payload[1] = 0x00
	payload[2] = 0xAA
	payload[3] = 0x01
	payload[4] = 0x23
	payload[5] = 0x45
	payload[6] = 0x67
	payload[7] = 0x89
	payload[8] = 0xAB
	payload[9] = 0xCD

	parity := make([]byte, 3)

	return common.MessageData{
		DownLinkFormat: downlinkFormat,
		FirstField:     firstField,
		Payload:        payload,
		Parity:         parity,
	}
}
