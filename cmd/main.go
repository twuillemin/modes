package main

import (
	"bufio"
	"fmt"
	acasReader "github.com/twuillemin/modes/pkg/acas/reader"
	adsbReader "github.com/twuillemin/modes/pkg/adsb/reader"
	"github.com/twuillemin/modes/pkg/adsbspy"
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

	// Read the ACAS content
	messageACAS, errACAS := acasReader.ReadMessage(messageDF16.MessageACAS)
	if errACAS != nil {
		fmt.Println(errACAS)
		return
	}

	fmt.Printf("%v\n", messageACAS.ToString())
}

func postProcessMessage17(messageDF17 *modeSMessages.MessageDF17) {

	fmt.Printf(" -- ADSB Information --\n")

	// Get the content
	messageADSB, errADSB := adsbReader.ReadMessage(messageDF17.MessageExtendedSquitter.Data)
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
		messageADSB, errADSB := adsbReader.ReadMessage(messageDF18.MessageExtendedSquitter.Data)
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

func readExampleFile() []string {

	file, err := os.Open("example/example1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	result := make([]string, 0, 1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
