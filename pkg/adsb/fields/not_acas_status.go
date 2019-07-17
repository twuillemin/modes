package fields

import "fmt"

// NotACASStatus is the Not ACAS flag definition
//
// Specified in Doc 9871 / B.2.3.10.3
type NotACASStatus byte

const (
	// NotACASStatusOperationalUnknown indicates ACAS operational or unknown
	NotACASStatusOperationalUnknown NotACASStatus = 0
	// NotACASStatusNotInstalledNotOperational indicates ACAS not installed or not operational
	NotACASStatusNotInstalledNotOperational NotACASStatus = 1
)

// ToString returns a basic, but readable, representation of the field
func (status NotACASStatus) ToString() string {

	switch status {
	case NotACASStatusOperationalUnknown:
		return "0 - ACAS operational or unknown"
	case NotACASStatusNotInstalledNotOperational:
		return "1 - ACAS not installed or not operational"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadNotACASStatus reads the NotACASStatus from a 56 bits data field
func ReadNotACASStatus(data []byte) NotACASStatus {
	bits := (data[1] & 0x20) >> 5
	return NotACASStatus(bits)
}
