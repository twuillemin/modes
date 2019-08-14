// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v2.go at 2019-08-15 00:36:43.4276713 +0300 EEST m=+0.013930401
package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"testing"
)

func TestReadFormat16V2Valid(t *testing.T) {

	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0x80

	msg, err := readFormat16V2(false, data)
	if err != nil {
		t.Fatal(err)
	}

	if msg.GetMessageFormat() != adsb.Format16V2 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format16V2.ToString(),
			msg.GetMessageFormat().ToString())
	}

	isBDS05V2Valid(t, msg)
}

func TestReadFormat16V2TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS05V2Message()[:6]
	data[0] = data[0] | 0x80

	_, err := readFormat16V2(false, data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat16V2BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0x01

	_, err := readFormat16V2(false, data)
	if err == nil {
		t.Error(err)
	}
}
