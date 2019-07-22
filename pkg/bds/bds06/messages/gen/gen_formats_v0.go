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

	type FormatParameter struct {
		Name string
	}

	formatParameters := []FormatParameter{
		{"Format05V0"},
		{"Format06V0"},
		{"Format07V0"},
		{"Format08V0"},
	}

	// Open the target file
	f, err := os.Create("formats_05_to_08_v0.go")
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
			Timestamp        time.Time
			FormatParameters []FormatParameter
		}{
			Timestamp:        time.Now(),
			FormatParameters: formatParameters,
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
// This file was generated by gen_formats_v0.go at {{ .Timestamp }}
package messages

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"github.com/twuillemin/modes/pkg/bds/bds"
	"github.com/twuillemin/modes/pkg/bds/bds06/fields"
)

{{ range .FormatParameters }}
// ------------------------------------------------------------------------------------
//
//                                {{ .Name }}
//
// ------------------------------------------------------------------------------------

// {{ .Name }} is a message at the format BDS 0,6
type {{ .Name }} struct {
	Movement                  fields.Movement
	GroundTrackStatus         fields.GroundTrackStatus
	GroundTrack               fields.GroundTrack
	Time                      fields.Time
	CPRFormat                 fields.CompactPositionReportingFormat
	EncodedLatitude           fields.EncodedLatitude
	EncodedLongitude          fields.EncodedLongitude
	HorizontalProtectionLimit fields.HorizontalProtectionLimit
	ContainmentRadius         fields.ContainmentRadius
}

// GetMessageFormat returns the ADSB format of the message
func (message *{{ .Name }}) GetMessageFormat() adsb.MessageFormat {
	return adsb.{{ .Name }}
}

// GetRegister returns the register of the message
func (message *{{ .Name }}) GetRegister() bds.Register {
	return adsb.{{ .Name }}.GetRegister()
}

// GetMovement returns the Movement
func (message *{{ .Name }}) GetMovement() fields.Movement {
	return message.Movement
}

// GetGroundTrackStatus returns the GroundTrackStatus
func (message *{{ .Name }}) GetGroundTrackStatus() fields.GroundTrackStatus {
	return message.GroundTrackStatus
}

// GetGroundTrack returns the GroundTrack
func (message *{{ .Name }}) GetGroundTrack() fields.GroundTrack {
	return message.GroundTrack
}

// GetTime returns the Time
func (message *{{ .Name }}) GetTime() fields.Time {
	return message.Time
}

// GetCPRFormat returns the CompactPositionReportingFormat
func (message *{{ .Name }}) GetCPRFormat() fields.CompactPositionReportingFormat {
	return message.CPRFormat
}

// GetEncodedLatitude returns the EncodedLatitude
func (message *{{ .Name }}) GetEncodedLatitude() fields.EncodedLatitude {
	return message.EncodedLatitude
}

// GetEncodedLongitude returns the EncodedLongitude
func (message *{{ .Name }}) GetEncodedLongitude() fields.EncodedLongitude {
	return message.EncodedLongitude
}

// GetHorizontalProtectionLimit returns the HorizontalProtectionLimit
func (message *{{ .Name }}) GetHorizontalProtectionLimit() fields.HorizontalProtectionLimit {
	return message.HorizontalProtectionLimit
}

// GetContainmentRadius returns the ContainmentRadius
func (message *{{ .Name }}) GetContainmentRadius() fields.ContainmentRadius {
	return message.ContainmentRadius
}

// ToString returns a basic, but readable, representation of the message
func (message *{{ .Name }}) ToString() string {
	return messageBDS06V0ToString(message)
}

// read{{ .Name }} reads a message at the format BDS 0,6
func read{{ .Name }}(data []byte) (*{{ .Name }}, error) {

	formatTypeCode := (data[0] & 0xF8) >> 3

	if formatTypeCode != adsb.{{ .Name }}.GetTypeCode() {
		return nil, fmt.Errorf("the data are given at format %v and can not be read at the format {{ .Name }}", formatTypeCode)
	}

	return &{{ .Name }}{
		Movement:                  fields.ReadMovement(data),
		GroundTrackStatus:         fields.ReadGroundTrackStatus(data),
		GroundTrack:               fields.ReadGroundTrack(data),
		Time:                      fields.ReadTime(data),
		CPRFormat:                 fields.ReadCompactPositionReportingFormat(data),
		EncodedLatitude:           fields.ReadEncodedLatitude(data),
		EncodedLongitude:          fields.ReadEncodedLongitude(data),
		HorizontalProtectionLimit: hplByFormat[formatTypeCode],
		ContainmentRadius:         crByFormat[formatTypeCode],
	}, nil
}

{{ end }}
`))
