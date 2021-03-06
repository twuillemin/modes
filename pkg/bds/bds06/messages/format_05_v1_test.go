// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v1.go at 2019-08-15 19:24:55.4433936 +0300 EEST m=+0.005969801
package messages

import (
	"testing"
)

func TestReadFormat05V1Valid(t *testing.T) {

	data := buildValidBDS06V1Message()
	data[0] = data[0] | 0x28

	msg, err := ReadFormat05V1(false, data)
	if err != nil {
		t.Fatal(err)
	}

	isBDS06V1Valid(t, msg)
}

func TestReadFormat05V1TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS06V1Message()[:6]
	data[0] = data[0] | 0x28

	_, err := ReadFormat05V1(false, data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat05V1BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS06V1Message()
	data[0] = data[0] | 0x01

	_, err := ReadFormat05V1(false, data)
	if err == nil {
		t.Error(err)
	}
}
