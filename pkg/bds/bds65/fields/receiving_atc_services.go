package fields

import "fmt"

// ReceivingATCServices is the Receiving ATC Services definition
//
// Specified in Doc 9871 / B.2.3.10.4
type ReceivingATCServices byte

const (
	// RASNotReceivingATC indicates Aircraft not receiving ATC services
	RASNotReceivingATC ReceivingATCServices = 0
	// RASReceivingATC indicates Aircraft receiving ATC services
	RASReceivingATC ReceivingATCServices = 1
)

// ToString returns a basic, but readable, representation of the field
func (status ReceivingATCServices) ToString() string {

	switch status {
	case RASNotReceivingATC:
		return "0 - Aircraft not receiving ATC services"
	case RASReceivingATC:
		return "1 - Aircraft receiving ATC services"
	default:
		return fmt.Sprintf("%v - Unknown code", status)
	}
}

// ReadReceivingATCServices reads the ReceivingATCServices from a 56 bits data field
func ReadReceivingATCServices(data []byte) ReceivingATCServices {
	bits := (data[3] & 0x08) >> 3
	return ReceivingATCServices(bits)
}
