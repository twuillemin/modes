// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v0.go at 2019-08-15 10:16:58.9712124 +0300 EEST m=+0.010968901
package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"testing"
)

func TestReadFormat18V0Valid(t *testing.T) {

	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0x90

	msg, err := ReadFormat18V0(data)
	if err != nil {
		t.Fatal(err)
	}

	if msg.GetMessageFormat() != adsb.Format18V0 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format18V0.ToString(),
			msg.GetMessageFormat().ToString())
	}

	isBDS05V0Valid(t, msg)
}

func TestReadFormat18V0TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS05V0Message()[:6]
	data[0] = data[0] | 0x90

	_, err := ReadFormat18V0(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat18V0BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0x01

	_, err := ReadFormat18V0(data)
	if err == nil {
		t.Error(err)
	}
}
