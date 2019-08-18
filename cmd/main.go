package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"

	resolutionAdvisoryMessage "github.com/twuillemin/modes/pkg/acas/ra/messages"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	adsbReader "github.com/twuillemin/modes/pkg/bds/reader"
	modeSCommon "github.com/twuillemin/modes/pkg/modes/common"
	modeSFields "github.com/twuillemin/modes/pkg/modes/fields"
	modeSMessages "github.com/twuillemin/modes/pkg/modes/messages"
	modeSReader "github.com/twuillemin/modes/pkg/modes/reader"
)

func main() {

	fileName := flag.String("file", "", "the name of the file to be processed")
	flag.Parse()

	// If a filename is given, use it and quit
	if len(*fileName) > 0 {
		for _, str := range readFile(*fileName) {
			processSingleLine(str)
		}
		return
	}

	if len(os.Args) > 1 {
		processSingleLine(os.Args[1])
	} else {
		fmt.Printf("No data provided")
	}
}

func readFile(fileName string) []string {

	file, err := os.Open(fileName)
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

// processSingleLine processes a single line of data in the form of hexadecimal text, like "8D4BAB4558AB031C446849B72535"
//
// Params:
//    - str: the line to process
func processSingleLine(str string) {

	fmt.Printf("Reading: %s\n", str)

	binaryData, err := hex.DecodeString(str)
	if err != nil {
		fmt.Printf("unable to convert the given text to hexadecimal values due to: %v", err)
		return
	}

	// Convert to a message if possible
	messageModeS, err := modeSReader.ReadMessage(binaryData)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check the CRC and get the Address or the Interrogator Identifier
	address, err := modeSReader.CheckCRC(messageModeS, binaryData, nil, nil)
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
		address = modeSCommon.ICAOAddress(message11.AddressAnnounced.Address)
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

func postProcessMessage16(messageDF16 *modeSMessages.MessageDF16) {

	fmt.Printf(" -- ACAS Information --\n")

	// Extract the format
	vds1 := (messageDF16.MessageACAS[0] & 0xF0) >> 4
	vds2 := messageDF16.MessageACAS[0] & 0x0F

	if vds1 == 3 && vds2 == 0 {

		// Read the ACAS content
		messageACAS, errACAS := resolutionAdvisoryMessage.ParseResolutionAdvisory(messageDF16.MessageACAS[1:])
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

	processADSBMessage(messageDF17.MessageExtendedSquitter)
}

func postProcessMessage18(messageDF18 *modeSMessages.MessageDF18) {

	if messageDF18.ControlField == modeSFields.ControlFieldADSB ||
		messageDF18.ControlField == modeSFields.ControlFieldADSBReserved {
		processADSBMessage(messageDF18.MessageExtendedSquitter)
	} else if messageDF18.ControlField == modeSFields.ControlFieldTISBFineFormat ||
		messageDF18.ControlField == modeSFields.ControlFieldTISBCoarseFormat ||
		messageDF18.ControlField == modeSFields.ControlFieldTISBReservedManagement ||
		messageDF18.ControlField == modeSFields.ControlFieldTISBRelayADSB {

		fmt.Printf(" -- TISB Information --\n")
		fmt.Printf("Not implemented\n")
	}
}

func processADSBMessage(data []byte) {

	fmt.Printf(" -- ADSB Information --\n")

	// Get the content
	messageADSB, _, errADSB := adsbReader.ReadADSBMessage(adsb.ReaderLevel2, false, false, data)
	if errADSB != nil {
		fmt.Println(errADSB)
		return
	}

	fmt.Printf("%v\n", messageADSB.ToString())
}
