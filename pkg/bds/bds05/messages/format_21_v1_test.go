// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v1.go at 2019-08-15 00:36:42.8648463 +0300 EEST m=+0.012965001
package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"testing"
)

func TestReadFormat21V1Valid(t *testing.T) {

	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0xA8

	msg, err := readFormat21V1(data)
	if err != nil {
		t.Fatal(err)
	}

	if msg.GetMessageFormat() != adsb.Format21V1 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format21V1.ToString(),
			msg.GetMessageFormat().ToString())
	}

	isBDS05V1Valid(t, msg)
}

func TestReadFormat21V1TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS05V1Message()[:6]
	data[0] = data[0] | 0xA8

	_, err := readFormat21V1(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat21V1BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0x01

	_, err := readFormat21V1(data)
	if err == nil {
		t.Error(err)
	}
}
