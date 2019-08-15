package adsb

import "fmt"

// ReaderLevel is the definition of the ADSB ReaderLevel used to communicate with reader function. As a parameter
// of the reader function it determine the level that the client want to read. As a returned value, it allows
// the function to provide more details after the data has been read. For example, sending Level0OrBetter is sent to
// the function for reading type code 31. As the type code 31 contains the exact level, the read function may return
// ReaderLevel1Exactly.
type ReaderLevel byte

const (
	// ReaderLevel0OrMore indicates that the message could be level ADSB level 0 or more
	ReaderLevel0OrMore ReaderLevel = 0
	// ReaderLevel0Exactly indicates that the message must be read as ADSB 0 or has been determined as being level 0 only
	ReaderLevel0Exactly ReaderLevel = 1
	// ReaderLevel1OrMore indicates that the message could be level ADSB level 1 or more
	ReaderLevel1OrMore ReaderLevel = 2
	// ReaderLevel1Exactly indicates that the message must be read as ADSB 1 or has been determined as being level 1 only
	ReaderLevel1Exactly ReaderLevel = 3
	// ReaderLevel2 indicates that the message must be read as ADSB 2 or has been determined as being level 1 only
	ReaderLevel2 ReaderLevel = 4
)

// ToString returns a basic, but readable, representation of the message
func (level ReaderLevel) ToString() string {
	switch level {
	case ReaderLevel0OrMore:
		return "ADSB V0, V1 and V2"
	case ReaderLevel0Exactly:
		return "ADSB V0"
	case ReaderLevel1OrMore:
		return "ADSB V1 and V2"
	case ReaderLevel1Exactly:
		return "ADSB V1"
	case ReaderLevel2:
		return "ADSB V2"
	default:
		return fmt.Sprintf("there is no ADSB ReaderLevel associated with value %v", level)
	}
}
