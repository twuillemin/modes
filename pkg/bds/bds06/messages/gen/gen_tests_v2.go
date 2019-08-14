// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates list_converter.go. It can be invoked by running
// go generate

package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

func main() {
	generateFile("format_05_v2_test.go", "Format05V2", "0x28")
	generateFile("format_06_v2_test.go", "Format06V2", "0x30")
	generateFile("format_07_v2_test.go", "Format07V2", "0x38")
	generateFile("format_08_v2_test.go", "Format08V2", "0x40")
}

func generateFile(fileName string, name string, messageCode string) {
	// Open the target file
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// Close at the end
	defer func() {
		closeErr := f.Close()
		if closeErr != nil {
			log.Fatal(err)
		}
	}()

	// Execute the template
	err = builderTemplate.Execute(
		f,
		struct {
			Timestamp   time.Time
			Name        string
			MessageCode string
		}{
			Timestamp:   time.Now(),
			Name:        name,
			MessageCode: messageCode,
		})
	if err != nil {
		log.Fatal(err)
	}
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var builderTemplate = template.Must(template.New("").Parse(`// Package messages holds the definition of the messages
package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"testing"
)

func TestRead{{ .Name }}Valid(t *testing.T) {

	data := buildValidBDS06Message()
	data[0] = data[0] | {{ .MessageCode }}

	msg, err := read{{ .Name }}(false, false, data)
	if err != nil {
		t.Fatal(err)
	}

	if msg.GetMessageFormat() != adsb.{{ .Name }} {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.{{ .Name }}.ToString(),
			msg.GetMessageFormat().ToString())
	}

	isBDS06Valid(t, msg)
}

func TestRead{{ .Name }}TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS06Message()[:6]
	data[0] = data[0] | {{ .MessageCode }}

	_, err := read{{ .Name }}(false, false, data)
	if err == nil {
		t.Error(err)
	}
}

func TestRead{{ .Name }}BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS06Message()
	data[0] = data[0] | 0x01

	_, err := read{{ .Name }}(false, false, data)
	if err == nil {
		t.Error(err)
	}
}`))
