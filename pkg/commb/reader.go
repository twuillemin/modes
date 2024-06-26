package commb

import (
	"errors"

	"github.com/twuillemin/modes/pkg/bds"
	"github.com/twuillemin/modes/pkg/bds/bds00"
	"github.com/twuillemin/modes/pkg/bds/bds07"
	"github.com/twuillemin/modes/pkg/bds/bds10"
	"github.com/twuillemin/modes/pkg/bds/bds17"
	"github.com/twuillemin/modes/pkg/bds/bds20"
	"github.com/twuillemin/modes/pkg/bds/bds30"
	"github.com/twuillemin/modes/pkg/bds/bds40"
	"github.com/twuillemin/modes/pkg/bds/bds44"
	"github.com/twuillemin/modes/pkg/bds/bds45"
	"github.com/twuillemin/modes/pkg/bds/bds50"
	"github.com/twuillemin/modes/pkg/bds/bds60"
)

// ReadCommBMessage reads and parse a Comm-B message.
//
// params:
//   - data: The body of the message. The message must be 7 bytes long
//
// Return the possible messages (which may be empty) and an optional error.
func ReadCommBMessage(data []byte) ([]bds.Message, error) {
	if len(data) != 7 {
		return nil, errors.New("the data for Comm-B message must be 7 bytes long")
	}

	var message bds.Message
	var err error

	//
	// Remove the all-zero messages directly
	//
	message, err = bds00.ReadNoMessageAvailable(data)
	if err == nil {
		return []bds.Message{message}, nil
	}

	//
	// Then the messages that leave no place to doubt
	//

	// Very short message (3 bits) rest must be blank
	message, err = bds07.ReadStatus(data)
	if err == nil && message.CheckCoherency() == nil {
		return []bds.Message{message}, nil
	}

	// Expect a well-formed string
	message, err = bds20.ReadAircraftIdentification(data)
	if err == nil && message.CheckCoherency() == nil {
		return []bds.Message{message}, nil
	}

	//
	// Then the message that could be interpreted in different formats
	//
	messages := make([]bds.Message, 0, 10)

	// Force analysis on all know types
	message, err = bds10.ReadDataLinkCapabilityReport(data)
	if err == nil {
		messages = append(messages, message)
	}
	message, err = bds17.ReadCommonUsageGICBCapabilityReport(data)
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
	message, err = bds44.ReadMeteorologicalRoutineAirReport(data)
	if err == nil {
		messages = append(messages, message)
	}
	message, err = bds45.ReadMeteorologicalHazardReport(data)
	if err == nil {
		messages = append(messages, message)
	}
	message, err = bds50.ReadTrackAndTurnReport(data)
	if err == nil {
		messages = append(messages, message)
	}
	message, err = bds60.ReadHeadingAndSpeedReport(data)
	if err == nil {
		messages = append(messages, message)
	}

	coherentMessages := make([]bds.Message, 0, len(messages))
	for _, message := range messages {
		if message.CheckCoherency() == nil {
			coherentMessages = append(coherentMessages, message)
		}
	}

	return coherentMessages, nil
}
