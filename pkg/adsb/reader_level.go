package adsb

import "fmt"

// ReaderLevel is the definition of the ADSB ReaderLevel used to communicate with reader function. The ReaderLevel is
// normally present given by the query (which is not available) or is given by the Aircraft Status Operational Message
// BDS 6,5.
//
// As stated by official documentation:
// An ADS-B Version Two (2) Receiving Subsystem shall, as a default, assume the received messages are
// using ADS-B Version Zero (0) ADS-B Message format unless, or until, an Aircraft Operational Status
// Message is received and the ADS-B Version Number is confirmed to be other than Zero (0)
type ReaderLevel byte

const (
	// ReaderLevel0 indicates that the message must be read as ADSB 0 or has been determined as being level 0 only
	ReaderLevel0 ReaderLevel = 1
	// ReaderLevel1 indicates that the message must be read as ADSB 1 or has been determined as being level 1 only
	ReaderLevel1 ReaderLevel = 3
	// ReaderLevel2 indicates that the message must be read as ADSB 2 or has been determined as being level 1 only
	ReaderLevel2 ReaderLevel = 4
)

// ToString returns a basic, but readable, representation of the message
func (level ReaderLevel) ToString() string {
	switch level {
	case ReaderLevel0:
		return "ADSB V0"
	case ReaderLevel1:
		return "ADSB V1"
	case ReaderLevel2:
		return "ADSB V2"
	default:
		return fmt.Sprintf("there is no ADSB ReaderLevel associated with value %v", level)
	}
}
