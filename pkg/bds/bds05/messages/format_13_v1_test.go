// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v1.go at 2019-08-15 10:16:59.4810499 +0300 EEST m=+0.009964401
package messages

import (
	"github.com/twuillemin/modes/pkg/bds/adsb"
	"testing"
)

func TestReadFormat13V1Valid(t *testing.T) {

	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0x68

	msg, err := ReadFormat13V1(false, data)
	if err != nil {
		t.Fatal(err)
	}

	if msg.GetMessageFormat() != adsb.Format13V1 {
		t.Errorf("Expected Format \"%v\", got \"%v\"",
			adsb.Format13V1.ToString(),
			msg.GetMessageFormat().ToString())
	}

	isBDS05V1Valid(t, msg)
}

func TestReadFormat13V1TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS05V1Message()[:6]
	data[0] = data[0] | 0x68

	_, err := ReadFormat13V1(false, data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat13V1BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0x01

	_, err := ReadFormat13V1(false, data)
	if err == nil {
		t.Error(err)
	}
}
