package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	resolutionAdvisoryMessage "github.com/twuillemin/modes/pkg/acas/ra/messages"
	"github.com/twuillemin/modes/pkg/adsbspy"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	bds05Fields "github.com/twuillemin/modes/pkg/bds/bds05/fields"
	bds05Messages "github.com/twuillemin/modes/pkg/bds/bds05/messages"
	bds08Messages "github.com/twuillemin/modes/pkg/bds/bds08/messages"
	adsbReader "github.com/twuillemin/modes/pkg/bds/reader"
	modeSFields "github.com/twuillemin/modes/pkg/modes/fields"
	modeSMessages "github.com/twuillemin/modes/pkg/modes/messages"
	modeSReader "github.com/twuillemin/modes/pkg/modes/reader"
	"github.com/twuillemin/modes/pkg/plane"
)

func main() {

	// The example file
	for _, str := range readExampleFile() {
		readSingleLine(str)
	}
	/*
		conn, err := net.Dial("tcp", "localhost:47806")
		if err != nil {
			log.Fatal(err)
		}

		reader := bufio.NewReader(conn)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				log.Fatal(err)
			}

			readSingleLine(string(line))
		}

	*/
}

var planes = make(map[modeSReader.ICAOAddress]*plane.Plane)
var defaultADSBLevel = adsb.Level2

func readSingleLine(str string) {
	fmt.Printf("Reading: %s\n", str)

	// Read the line
	messageADSBSpy, err := adsbspy.ReadLine(str)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert to a message if possible
	messageModeS, err := modeSReader.ReadMessage(messageADSBSpy.Message)
	if err != nil {
		fmt.Println(err)
		return
	}

	timestamp := uint32(messageADSBSpy.Timestamp)

	// Check the CRC and get the Address or the Interrogator Identifier
	address, err := modeSReader.CheckCRC(messageModeS, messageADSBSpy.Message, nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// For message 11 (Reply to all call) the address is the address of the caller
	fmt.Printf(" -- Mode-S Information --\n")
	if messageModeS.GetDownLinkFormat() == 11 {
		fmt.Printf("In reply to interrogator: %v\n", address.ToString())
		// Update address
		message11 := messageModeS.(*modeSMessages.MessageDF11)
		address = modeSReader.ICAOAddress(message11.AddressAnnounced.Address)
	} else {
		fmt.Printf("From: %v\n", address.ToString())
	}

	// Print the content of the mode S message
	fmt.Printf("%v\n", messageModeS.ToString())

	// Get the plane
	viewedPlane := checkoutPlane(timestamp, address)

	// For message with additional content
	switch messageModeS.GetDownLinkFormat() {
	case 16:
		postProcessMessage16(messageModeS.(*modeSMessages.MessageDF16))
	case 17:
		postProcessMessage17(timestamp, viewedPlane, messageModeS.(*modeSMessages.MessageDF17))
	case 18:
		postProcessMessage18(timestamp, viewedPlane, messageModeS.(*modeSMessages.MessageDF18))
	}

	fmt.Printf("\n")
}

func checkoutPlane(timestamp uint32, address modeSReader.ICAOAddress) *plane.Plane {

	if knownPlane, ok := planes[address]; ok {
		knownPlane.LastSeenTimestamp = timestamp
		return knownPlane
	}

	newPlane := &plane.Plane{
		ICAOAddress:        address,
		ADSBLevel:          defaultADSBLevel,
		Identification:     "",
		FirstSeenTimestamp: timestamp,
		LastSeenTimestamp:  timestamp,
		EvenCPRLatitude:    0,
		EvenCPRLongitude:   0,
		EventCPRTimestamp:  0,
		OddCPRLatitude:     0,
		OddCPRLongitude:    0,
		OddCPRTimestamp:    0,
	}

	planes[address] = newPlane

	return newPlane

}

func postProcessMessage16(messageDF16 *modeSMessages.MessageDF16) {

	fmt.Printf(" -- ACAS Information --\n")

	// Extract the format
	vds1 := (messageDF16.MessageACAS[0] & 0xF0) >> 4
	vds2 := messageDF16.MessageACAS[0] & 0x0F

	if vds1 == 3 && vds2 == 0 {

		// Read the ACAS content
		messageACAS, errACAS := resolutionAdvisoryMessage.ReadResolutionAdvisory(messageDF16.MessageACAS[1:])
		if errACAS != nil {
			fmt.Println(errACAS)
			return
		}

		fmt.Printf("%v\n", messageACAS.ToString())
	} else {
		fmt.Printf("Unknown message type: [%X:%X]\n", vds1, vds2)
	}
}

func postProcessMessage17(timestamp uint32, plane *plane.Plane, messageDF17 *modeSMessages.MessageDF17) {

	processADSBMessage(timestamp, plane, messageDF17.MessageExtendedSquitter.Data)
}

func postProcessMessage18(timestamp uint32, plane *plane.Plane, messageDF18 *modeSMessages.MessageDF18) {

	if messageDF18.ControlField == modeSFields.ControlFieldADSB || messageDF18.ControlField == modeSFields.ControlFieldADSBReserved {
		processADSBMessage(timestamp, plane, messageDF18.MessageExtendedSquitter.Data)
	} else if messageDF18.ControlField == modeSFields.ControlFieldTISBFineFormat ||
		messageDF18.ControlField == modeSFields.ControlFieldTISBCoarseFormat ||
		messageDF18.ControlField == modeSFields.ControlFieldTISBReservedManagement ||
		messageDF18.ControlField == modeSFields.ControlFieldTISBRelayADSB {

		fmt.Printf(" -- TISB Information --\n")
		fmt.Printf("Not implemented\n")
	}
}

func processADSBMessage(timestamp uint32, plane *plane.Plane, data []byte) {

	fmt.Printf(" -- ADSB Information --\n")

	// Get the content
	messageADSB, _, errADSB := adsbReader.ReadADSBMessage(plane.ADSBLevel, false, false, data)
	if errADSB != nil {
		fmt.Println(errADSB)
		return
	}

	if messageADSB == nil {
		return
	}
	fmt.Printf("%v\n", messageADSB.ToString())

	planeUpdated := false

	// If message with position
	if message05, ok := messageADSB.(bds05Messages.MessageBDS05); ok {
		if message05.GetCPRFormat() == bds05Fields.CPRFormatEven {
			plane.EvenCPRLatitude = uint32(message05.GetEncodedLatitude())
			plane.EvenCPRLongitude = uint32(message05.GetEncodedLongitude())
			plane.EventCPRTimestamp = timestamp
		} else {
			plane.OddCPRLatitude = uint32(message05.GetEncodedLatitude())
			plane.OddCPRLongitude = uint32(message05.GetEncodedLongitude())
			plane.OddCPRTimestamp = timestamp
		}

		planeUpdated = true
	}

	// If message with identification
	if message08, ok := messageADSB.(bds08Messages.MessageBDS08); ok {
		if len(plane.Identification) == 0 {
			plane.Identification = string(message08.GetAircraftIdentification())
			planeUpdated = true
		}
	}

	if planeUpdated {
		fmt.Printf("==>%v\n", plane.ToString())
	}
}

// TODO Add management for ADSBLevel, nicSupplementA and nicSupplementC

func readExampleFile() []string {

	file, err := os.Open("example/example2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		errClose := file.Close()
		if errClose != nil {
			log.Fatal(errClose)
		}
	}()

	result := make([]string, 0, 1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
