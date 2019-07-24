package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/twuillemin/modes/internal/pkg/processor"
	"log"
	"net"
	"os"

	"github.com/twuillemin/modes/internal/pkg/plane"
	"github.com/twuillemin/modes/pkg/bds/adsb"
)

func main() {

	// By default, planes are ADSB level 2 compliant (Europe...)
	plane.SetDefaultADSBLevel(adsb.Level2)

	fileName := flag.String("file", "", "the name of the file to be processed")
	airSpyServer := flag.String("adsb_spy_server", "localhost", "the name of the ADSBSpy server (default: localhost)")
	airSpyPort := flag.Int("adsb_spy_port", 47806, "the port of the ADSBSpy server (default: 47806)")

	// If a filename is given, use it and quit
	if len(*fileName) > 0 {
		for _, str := range readExampleFile(*fileName) {
			processor.ProcessSingleLine(str)
		}
		return
	}

	address := fmt.Sprintf("%v:%v", *airSpyServer, *airSpyPort)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		processor.ProcessSingleLine(string(line))
	}

}

func readExampleFile(fileName string) []string {

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
