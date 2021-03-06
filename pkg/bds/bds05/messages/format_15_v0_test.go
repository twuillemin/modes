// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v0.go at 2019-08-15 19:22:02.9880369 +0300 EEST m=+0.008976401
package messages

import (
	"testing"
)

func TestReadFormat15V0Valid(t *testing.T) {

	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0x78

	msg, err := ReadFormat15V0(data)
	if err != nil {
		t.Fatal(err)
	}

	isBDS05V0Valid(t, msg)
}

func TestReadFormat15V0TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS05V0Message()[:6]
	data[0] = data[0] | 0x78

	_, err := ReadFormat15V0(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat15V0BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0x01

	_, err := ReadFormat15V0(data)
	if err == nil {
		t.Error(err)
	}
}
