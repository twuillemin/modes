package adsb

import "fmt"

// Level is the definition of the ADSB Level used to communicate with reader function. As a parameter
// of the reader function it determine the level that the client want to read. As a returned value, it allows
// the function to provide more details after the data has been read. For example, sending Level0OrBetter is sent to
// the function for reading type code 31. As the type code 31 contains the exact level, the read function may return
// Level1Exactly.
type Level byte

const (
	// Level0OrMore indicates that the message could be level ADSB level 0 or more
	Level0OrMore Level = 0
	// Level0Exactly indicates that the message must be read as ADSB 0 or has been determined as being level 0 only
	Level0Exactly Level = 1
	// Level1OrMore indicates that the message could be level ADSB level 1 or more
	Level1OrMore Level = 2
	// Level1Exactly indicates that the message must be read as ADSB 1 or has been determined as being level 1 only
	Level1Exactly Level = 3
	// Level2 indicates that the message must be read as ADSB 2 or has been determined as being level 1 only
	Level2 Level = 4
)

func (level Level) ToString() string {
	switch level {
	case Level0OrMore:
		return "ADSB V0, V1 and V2"
	case Level0Exactly:
		return "ADSB V0"
	case Level1OrMore:
		return "ADSB V1 and V2"
	case Level1Exactly:
		return "ADSB V1"
	case Level2:
		return "ADSB V2"
	default:
		return fmt.Sprintf("there is no ADSB Level associated with value %v", level)
	}
}
