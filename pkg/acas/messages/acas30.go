package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/acas/fields"
)

// MessageACAS (MV) field shall contain the aircraft address which provides the long ACAS message
//
// Defined at 3.1.2.8.3.1
type ACAS30 struct {
	VDS1                    byte
	VDS2                    byte
	ActiveRA                fields.ActiveResolutionAdvisory
	RAComplement            fields.RAComplement
	RATerminatedIndicator   fields.RATerminatedIndicator
	MultipleThreatEncounter fields.MultipleThreatEncounter
}

// ReadMessageACAS30 reads a message for VDS1=3 and VDS2=0
//
// Params:
//    - message: The content of the message including the field VDS. This is for example the full content
//               content of the MV field from Mode S message
//
// Returns a properly formatted ACASMessage
func ReadMessageACAS30(message []byte) ACASMessage {

	// Format of the message is as follow:
	//        0                 1                 2                 3           4,5,6
	//  VDS1     VDS2  |       ARA       |   ARA       RAC | RAC RAT MTE Res | Reserved |
	// 0 0 1 1 0 0 0 0 | a a a a a a a a | a a a a a a c c | c c t m _ _ _ _ | 18 bits  |

	// Extract the raw values
	vds1 := (message[0] & 0xF0) >> 4
	vds2 := message[0] & 0x0F

	ara1 := (message[1] & 0xFC) >> 2
	ara2 := (message[1] & 0x03) << 6
	ara3 := (message[2] & 0xFC) >> 2
	ara := uint16(ara1)<<8 + uint16(ara2) + uint16(ara3)

	rac1 := (message[2] & 0x03) << 2
	rac2 := (message[3] & 0xC0) >> 6
	rac := rac1 + rac2

	rat := (message[3] & 0x20) >> 5

	mte := (message[3] & 0x10) >> 4

	araFirstBit := (message[1] & 0x80) >> 7

	var activeRA fields.ActiveResolutionAdvisory
	if araFirstBit == 1 {
		activeRA = fields.ReadARAOneThreatOrSameSeparation(ara)
	} else if mte != 0 {
		activeRA = fields.ReadARAMultipleThreatsDifferentSeparation(ara)
	} else {
		activeRA = fields.ActiveRANoVerticalRAGenerated{}
	}

	return ACAS30{
		VDS1:                    vds1,
		VDS2:                    vds2,
		ActiveRA:                activeRA,
		RAComplement:            fields.ReadRAComplement(rac),
		RATerminatedIndicator:   fields.ReadRATerminatedIndicator(rat),
		MultipleThreatEncounter: fields.ReadMultipleThreatEncounter(mte),
	}
}

// GetName returns the name of the message
func (messageACAS ACAS30) GetName() string {
	return "Coordination reply"
}

// GetVDS1 returns the VDS1
func (messageACAS ACAS30) GetVDS1() byte {
	return messageACAS.VDS1
}

// GetVDS2 returns the VDS1
func (messageACAS ACAS30) GetVDS2() byte {
	return messageACAS.VDS2
}

// ToString returns a basic, but readable, representation of the field
func (messageACAS ACAS30) ToString() string {
	return fmt.Sprintf("VDS 1:                                           %02X\n"+
		"VDS 2:                                           %02X\n"+
		"Active Resolution Advisory:\n%v\n"+
		"Active Resolution Advisory Complement:\n%v\n"+
		"Active Resolution Advisory Terminated Indicator: %v\n"+
		"Multiple Threat Encounter:                       %v",
		messageACAS.VDS1,
		messageACAS.VDS2,
		messageACAS.ActiveRA.ToString(),
		messageACAS.RAComplement.ToString(),
		messageACAS.RATerminatedIndicator.ToString(),
		messageACAS.MultipleThreatEncounter.ToString())
}
