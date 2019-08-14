// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v0.go at 2019-08-15 00:36:42.3248858 +0300 EEST m=+0.007978901
package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"testing"
)

func TestReadFormat12V0Valid(t *testing.T) {

	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0x60

	msg, err := readFormat12V0(data)
	if err != nil {
		t.Fatal(err)
	}

	if msg.GetMessageFormat() != adsb.Format12V0 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format12V0.ToString(),
			msg.GetMessageFormat().ToString())
	}

	isBDS05V0Valid(t, msg)
}

func TestReadFormat12V0TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS05V0Message()[:6]
	data[0] = data[0] | 0x60

	_, err := readFormat12V0(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat12V0BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0x01

	_, err := readFormat12V0(data)
	if err == nil {
		t.Error(err)
	}
}
