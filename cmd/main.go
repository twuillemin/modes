package main

import (
	"bufio"
	"fmt"
	resolutionAdvisoryMessage "github.com/twuillemin/modes/pkg/acas/ra/messages"
	"github.com/twuillemin/modes/pkg/adsbspy"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	adsbReader "github.com/twuillemin/modes/pkg/bds/reader"
	modeSFields "github.com/twuillemin/modes/pkg/modes/fields"
	modeSMessages "github.com/twuillemin/modes/pkg/modes/messages"
	modeSReader "github.com/twuillemin/modes/pkg/modes/reader"
	"log"
	"os"
)

func main() {

	for _, str := range readExampleFile() {

		fmt.Printf("Reading: %s\n", str)

		// Read the line
		messageADSBSpy, err := adsbspy.ReadLine(str)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Convert to a message if possible
		messageModeS, err := modeSReader.ReadMessage(messageADSBSpy.Message)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Check the CRC and get the Address or the Interrogator Identifier
		address, err := modeSReader.CheckCRC(messageModeS, messageADSBSpy.Message, nil, nil)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// For message 11 (Reply to all call) the address is the address of the caller
		fmt.Printf(" -- Mode-S Information --\n")
		if messageModeS.GetDownLinkFormat() == 11 {
			fmt.Printf("In reply to interrogator: %v\n", address.ToString())
		} else {
			fmt.Printf("From: %v\n", address.ToString())
		}

		// Print the content of the mode S message
		fmt.Printf("%v\n", messageModeS.ToString())

		// For message with additional content
		switch messageModeS.GetDownLinkFormat() {
		case 16:
			postProcessMessage16(messageModeS.(*modeSMessages.MessageDF16))
		case 17:
			postProcessMessage17(messageModeS.(*modeSMessages.MessageDF17))
		case 18:
			postProcessMessage18(messageModeS.(*modeSMessages.MessageDF18))
		}

		fmt.Printf("\n")
	}
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

func postProcessMessage17(messageDF17 *modeSMessages.MessageDF17) {

	fmt.Printf(" -- ADSB Information --\n")

	// Get the content
	messageADSB, _, errADSB := adsbReader.ReadADSBMessage(adsb.Level0OrMore, false, false, messageDF17.MessageExtendedSquitter.Data)
	if errADSB != nil {
		fmt.Println(errADSB)
		return
	}
	if messageADSB != nil {
		fmt.Printf("%v\n", messageADSB.ToString())
	}
}

func postProcessMessage18(messageDF18 *modeSMessages.MessageDF18) {

	if messageDF18.ControlField == modeSFields.ControlFieldADSB || messageDF18.ControlField == modeSFields.ControlFieldADSBReserved {

		fmt.Printf(" -- ADSB Information --\n")

		// Get the content
		messageADSB, _, errADSB := adsbReader.ReadADSBMessage(adsb.Level0OrMore, false, false, messageDF18.MessageExtendedSquitter.Data)
		if errADSB != nil {
			fmt.Println(errADSB)
			return
		}
		if messageADSB != nil {
			fmt.Printf("%v\n", messageADSB.ToString())
		}
	} else if messageDF18.ControlField == modeSFields.ControlFieldTISBFineFormat ||
		messageDF18.ControlField == modeSFields.ControlFieldTISBCoarseFormat ||
		messageDF18.ControlField == modeSFields.ControlFieldTISBReservedManagement ||
		messageDF18.ControlField == modeSFields.ControlFieldTISBRelayADSB {

		fmt.Printf(" -- TISB Information --\n")
		fmt.Printf("Not implemented\n")
	}
}

// TODO Add management for ADSBLevel, nicSupplementA and nicSupplementC

func readExampleFile() []string {

	file, err := os.Open("example/example1.txt")
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
