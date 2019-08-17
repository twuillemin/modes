// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v1.go at 2019-08-15 19:22:03.5692791 +0300 EEST m=+0.012962501
package messages

import (
	"testing"
)

func TestReadFormat21V1Valid(t *testing.T) {

	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0xA8

	msg, err := ReadFormat21V1(data)
	if err != nil {
		t.Fatal(err)
	}

	isBDS05V1Valid(t, msg)
}

func TestReadFormat21V1TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS05V1Message()[:6]
	data[0] = data[0] | 0xA8

	_, err := ReadFormat21V1(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat21V1BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0x01

	_, err := ReadFormat21V1(data)
	if err == nil {
		t.Error(err)
	}
}