package messages

import (
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/fields"
	"testing"
)

func TestReadFormatDF19Valid(t *testing.T) {

	msg, err := ParseDF19(buildValidDF19Message())
	if err != nil {
		t.Fatal(err)
	}

	if msg.ApplicationField != fields.ApplicationFieldReserved5 {
		t.Errorf("Expected Application Field \"%v\", got \"%v\"",
			fields.ApplicationFieldReserved5.ToString(),
			msg.ApplicationField.ToString())
	}

	if len(msg.ToString()) <= 0 {
		t.Error("Expected a printable message, but get nothing")
	}
}

func buildValidDF19Message() common.MessageData {

	// Format of the message is as follows:
	//
	//     DF     AF   | Military use
	// 1 0 0 1 1 a a a |   104 bits

	// 0000 0101: Unused (00000) + Application Field: Reserved5 (101)
	firstField := byte(0x05)

	// Not used
	payload := make([]byte, 10)
	parity := make([]byte, 3)

	return common.MessageData{
		DownLinkFormat: 19,
		FirstField:     firstField,
		Payload:        payload,
		Parity:         parity,
	}
}
