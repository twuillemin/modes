package adsb

import "fmt"

// ADSBVersion is the definition of the ADSB version used to communicate with reader function. The ADSBVersion is
// normally given by the query (which is not available) or is given by the Aircraft Status Operational Message
// BDS 6,5.
//
// As stated by official documentation:
// An ADS-B Version Two (2) Receiving Subsystem shall, as a default, assume the received messages are
// using ADS-B Version Zero (0) ADS-B Message format unless, or until, an Aircraft Operational Status
// Message is received and the ADS-B Version Number is confirmed to be other than Zero (0)
type ADSBVersion byte

const (
	// ADSBV0 indicates that the message must be read as ADSB 0 or has been determined as being level 0 only
	ADSBV0 ADSBVersion = 0
	// ADSBV1 indicates that the message must be read as ADSB 1 or has been determined as being level 1 only
	ADSBV1 ADSBVersion = 1
	// ADSBV2 indicates that the message must be read as ADSB 2 or has been determined as being level 1 only
	ADSBV2 ADSBVersion = 2
)

// ToString returns a basic, but readable, representation of the message
func (level ADSBVersion) ToString() string {
	switch level {
	case ADSBV0:
		return "ADSB V0"
	case ADSBV1:
		return "ADSB V1"
	case ADSBV2:
		return "ADSB V2"
	default:
		return fmt.Sprintf("there is no ADSB ReaderLevel associated with value %v", level)
	}
}
