// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v0.go at 2024-06-02 22:43:44.5398116 +0300 EEST m=+0.004228801
package messages

import (
	"testing"
)

func TestReadFormat16V0Valid(t *testing.T) {

	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0x80

	msg, err := ReadFormat16V0(data)
	if err != nil {
		t.Fatal(err)
	}

	isBDS05V0Valid(t, msg)
}

func TestReadFormat16V0TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS05V0Message()[:6]
	data[0] = data[0] | 0x80

	_, err := ReadFormat16V0(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat16V0BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0x01

	_, err := ReadFormat16V0(data)
	if err == nil {
		t.Error(err)
	}
}
