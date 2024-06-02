// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v2.go at 2024-06-02 22:43:45.7222207 +0300 EEST m=+0.004164501
package messages

import (
	"testing"
)

func TestReadFormat13V2Valid(t *testing.T) {

	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0x68

	msg, err := ReadFormat13V2(false, data)
	if err != nil {
		t.Fatal(err)
	}

	isBDS05V2Valid(t, msg)
}

func TestReadFormat13V2TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS05V2Message()[:6]
	data[0] = data[0] | 0x68

	_, err := ReadFormat13V2(false, data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat13V2BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0x01

	_, err := ReadFormat13V2(false, data)
	if err == nil {
		t.Error(err)
	}
}