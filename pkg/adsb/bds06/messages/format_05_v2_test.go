// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v2.go at 2024-06-02 22:43:49.8495596 +0300 EEST m=+0.002145801
package messages

import (
	"testing"
)

func TestReadFormat05V2Valid(t *testing.T) {

	data := buildValidBDS06V2Message()
	data[0] = data[0] | 0x28

	msg, err := ReadFormat05V2(false, false, data)
	if err != nil {
		t.Fatal(err)
	}

	isBDS06V2Valid(t, msg)
}

func TestReadFormat05V2TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS06V2Message()[:6]
	data[0] = data[0] | 0x28

	_, err := ReadFormat05V2(false, false, data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat05V2BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS06V2Message()
	data[0] = data[0] | 0x01

	_, err := ReadFormat05V2(false, false, data)
	if err == nil {
		t.Error(err)
	}
}
