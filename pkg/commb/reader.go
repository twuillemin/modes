package commb

import (
	"errors"
	"github.com/twuillemin/modes/pkg/bds/bds00"
	"github.com/twuillemin/modes/pkg/bds/bds40"

	"github.com/twuillemin/modes/pkg/bds"
	"github.com/twuillemin/modes/pkg/bds/bds07"
	"github.com/twuillemin/modes/pkg/bds/bds10"
	"github.com/twuillemin/modes/pkg/bds/bds17"
	"github.com/twuillemin/modes/pkg/bds/bds20"
	"github.com/twuillemin/modes/pkg/bds/bds30"
)

// ReadCommBMessage reads and parse a Comm-B message.
//
// params:
//   - data: The body of the message. The message must be 7 bytes long
//
// Return the parsed message, the detected ADSB ReaderLevel and an optional error. The detected ADSB ReaderLevel will generally be
// the same as the given one, except if the decoded message has information to change it.
func ReadCommBMessage(data []byte) (bds.Message, error) {
	if len(data) != 7 {
		return nil, errors.New("the data for Comm-B message must be 7 bytes long")
	}

	messages := make([]bds.Message, 0, 10)

	var message bds.Message
	var err error

	// Remove the all-zero messages directly
	message, err = bds00.ReadNoMessageAvailable(data)
	if err == nil {
		return message, nil
	}

	// Force analysis on all know types
	message, err = bds07.ReadStatus(data)
	if err == nil {
		messages = append(messages, message)
	}
	message, err = bds10.ReadDataLinkCapabilityReport(data)
	if err == nil {
		messages = append(messages, message)
	}
	message, err = bds17.ReadCommonUsageGICBCapabilityReport(data)
	if err == nil {
		messages = append(messages, message)
	}
	message, err = bds20.ReadAircraftIdentification(data)
	if err == nil {
		messages = append(messages, message)
	}
	message, err = bds30.ReadACASResolutionAdvisory(data)
	if err == nil {
		messages = append(messages, message)
	}
	message, err = bds40.ReadSelectedVerticalIntention(data)
	if err == nil {
		messages = append(messages, message)
	}

	// If no message is found, just return
	if len(messages) == 1 {
		return nil, errors.New("message can not be read")
	}

	// If only one message is found, it should be the good one
	if len(messages) == 1 {
		return messages[0], nil
	}

	coherentMessages := make([]bds.Message, 0, len(messages))
	for _, message := range messages {
		if message.CheckCoherency() == nil {
			coherentMessages = append(coherentMessages, message)
		}
	}

	switch len(coherentMessages) {
	case 0:
		return nil, errors.New("no coherent message can be read")
	case 1:
		return coherentMessages[0], nil
	default:
		return nil, errors.New("multiple formats match the message")
	}
}
