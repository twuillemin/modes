package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"

	resolutionAdvisory "github.com/twuillemin/modes/pkg/acas/ra"
	"github.com/twuillemin/modes/pkg/adsb"
	"github.com/twuillemin/modes/pkg/adsb/reader"
	commbReader "github.com/twuillemin/modes/pkg/commb"
	modeSCommon "github.com/twuillemin/modes/pkg/modes/common"
	modeSFields "github.com/twuillemin/modes/pkg/modes/fields"
	modeSMessages "github.com/twuillemin/modes/pkg/modes/messages"
	modeSReader "github.com/twuillemin/modes/pkg/modes/reader"
)

func main() {

	fileName := flag.String("file", "", "the name of the file to be processed")
	adsbReaderLevelParam := flag.Int("adsb_reader_level", 0, "the version of ADSB reader 0, 1 or 2 to use (ADSB V0 by default)")
	flag.Parse()

	adsbReaderLevel := adsb.ADSBV0
	switch *adsbReaderLevelParam {
	case 0:
		adsbReaderLevel = adsb.ADSBV0
	case 1:
		adsbReaderLevel = adsb.ADSBV1
	case 2:
		adsbReaderLevel = adsb.ADSBV2
	default:
		panic(fmt.Sprintf("Unable to use the given ADSB level %v", *adsbReaderLevelParam))
	}

	// If a filename is given, use it and quit
	if len(*fileName) > 0 {
		for _, str := range readFile(*fileName) {
			processSingleLine(str, adsbReaderLevel)
		}
		return
	}

	if len(os.Args) > 1 {
		processSingleLine(os.Args[1], adsbReaderLevel)
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
//   - str: the line to process
func processSingleLine(str string, readerLevel adsb.ADSBVersion) {

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
		postProcessMessage17(messageModeS.(*modeSMessages.MessageDF17), readerLevel)
	case 18:
		postProcessMessage18(messageModeS.(*modeSMessages.MessageDF18), readerLevel)
	case 20:
		postProcessMessage20(messageModeS.(*modeSMessages.MessageDF20))
	case 21:
		postProcessMessage21(messageModeS.(*modeSMessages.MessageDF21))
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
		messageACAS, errACAS := resolutionAdvisory.ReadResolutionAdvisory(messageDF16.MessageACAS[1:])
		if errACAS != nil {
			fmt.Println(errACAS)
			return
		}

		fmt.Printf("%v\n", messageACAS.ToString())
	} else {
		fmt.Printf("Unknown message type: [%X:%X]\n", vds1, vds2)
	}
}

func postProcessMessage17(messageDF17 *modeSMessages.MessageDF17, adsbVersion adsb.ADSBVersion) {

	processADSBMessage(messageDF17.MessageExtendedSquitter, adsbVersion)
}

func postProcessMessage18(messageDF18 *modeSMessages.MessageDF18, adsbVersion adsb.ADSBVersion) {

	if messageDF18.ControlField == modeSFields.ControlFieldADSB ||
		messageDF18.ControlField == modeSFields.ControlFieldADSBReserved {
		processADSBMessage(messageDF18.MessageExtendedSquitter, adsbVersion)
	} else if messageDF18.ControlField == modeSFields.ControlFieldTISBFineFormat ||
		messageDF18.ControlField == modeSFields.ControlFieldTISBCoarseFormat ||
		messageDF18.ControlField == modeSFields.ControlFieldTISBReservedManagement ||
		messageDF18.ControlField == modeSFields.ControlFieldTISBRelayADSB {

		fmt.Printf(" -- TISB Information --\n")
		fmt.Printf("Not implemented\n")
	}
}

func processADSBMessage(data []byte, adsbVersion adsb.ADSBVersion) {

	fmt.Printf(" -- ADSB Information --\n")

	// Get the content
	messageADSB, errADSB := reader.ReadADSBMessage(adsbVersion, false, false, data)
	if errADSB != nil {
		fmt.Println(errADSB)
		return
	}

	fmt.Printf("%v\n", messageADSB.ToString())
}

func postProcessMessage20(messageDF20 *modeSMessages.MessageDF20) {

	processCommBMessage(messageDF20.MessageCommB)
}

func postProcessMessage21(messageDF21 *modeSMessages.MessageDF21) {

	processCommBMessage(messageDF21.MessageCommB)
}

func processCommBMessage(data []byte) {

	fmt.Printf(" -- Comm-B Information --\n")

	// Get the content
	message, err := commbReader.ReadCommBMessage(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v\n", message.ToString())
}
