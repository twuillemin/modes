// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v1.go at 2024-06-02 22:43:45.1317889 +0300 EEST m=+0.005211101
package messages

import (
	"testing"
)

func TestReadFormat18V1Valid(t *testing.T) {

	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0x90

	msg, err := ReadFormat18V1(false, data)
	if err != nil {
		t.Fatal(err)
	}

	isBDS05V1Valid(t, msg)
}

func TestReadFormat18V1TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS05V1Message()[:6]
	data[0] = data[0] | 0x90

	_, err := ReadFormat18V1(false, data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat18V1BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0x01

	_, err := ReadFormat18V1(false, data)
	if err == nil {
		t.Error(err)
	}
}
