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
	generateFile("format_09_v0_test.go", "Format09V0", "0x48")
	generateFile("format_10_v0_test.go", "Format10V0", "0x50")
	generateFile("format_11_v0_test.go", "Format11V0", "0x58")
	generateFile("format_12_v0_test.go", "Format12V0", "0x60")
	generateFile("format_13_v0_test.go", "Format13V0", "0x68")
	generateFile("format_14_v0_test.go", "Format14V0", "0x70")
	generateFile("format_15_v0_test.go", "Format15V0", "0x78")
	generateFile("format_16_v0_test.go", "Format16V0", "0x80")
	generateFile("format_17_v0_test.go", "Format17V0", "0x88")
	generateFile("format_18_v0_test.go", "Format18V0", "0x90")
	generateFile("format_20_v0_test.go", "Format20V0", "0xA0")
	generateFile("format_21_v0_test.go", "Format21V0", "0xA8")
	generateFile("format_22_v0_test.go", "Format22V0", "0xB0")
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
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v0.go at {{ .Timestamp }}
package messages

import (
	"testing"
)

func TestRead{{ .Name }}Valid(t *testing.T) {

	data := buildValidBDS05V0Message()
	data[0] = data[0] | {{ .MessageCode }}

	msg, err := Read{{ .Name }}(data)
	if err != nil {
		t.Fatal(err)
	}

	isBDS05V0Valid(t, msg)
}

func TestRead{{ .Name }}TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS05V0Message()[:6]
	data[0] = data[0] | {{ .MessageCode }}

	_, err := Read{{ .Name }}(data)
	if err == nil {
		t.Error(err)
	}
}

func TestRead{{ .Name }}BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0x01

	_, err := Read{{ .Name }}(data)
	if err == nil {
		t.Error(err)
	}
}
`))
