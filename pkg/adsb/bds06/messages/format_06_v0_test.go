// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_v0.go at 2024-06-02 22:43:48.6763719 +0300 EEST m=+0.002807101
package messages

import (
	"testing"
)

func TestReadFormat06V0Valid(t *testing.T) {

	data := buildValidBDS06V0Message()
	data[0] = data[0] | 0x30

	msg, err := ReadFormat06V0(data)
	if err != nil {
		t.Fatal(err)
	}

	isBDS06V0Valid(t, msg)
}

func TestReadFormat06V0TooShort(t *testing.T) {

	// Get too short data
	data := buildValidBDS06V0Message()[:6]
	data[0] = data[0] | 0x30

	_, err := ReadFormat06V0(data)
	if err == nil {
		t.Error(err)
	}
}

func TestReadFormat06V0BadCode(t *testing.T) {

	// Message code 1
	data := buildValidBDS06V0Message()
	data[0] = data[0] | 0x01

	_, err := ReadFormat06V0(data)
	if err == nil {
		t.Error(err)
	}
}
