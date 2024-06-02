// Package messages holds the definition of the messages
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_tests_bds_0_5.go at 2024-06-02 22:43:43.9487145 +0300 EEST m=+0.002136801
package messages

import (
	"github.com/twuillemin/modes/pkg/adsb"
	"testing"
)

func TestDetectAndReadFormat09V0Valid(t *testing.T) {

	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0x48

	msg, err := ReadBDS05(adsb.ReaderLevel0, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format09V0); !ok {
		t.Errorf("expected a message of type Format09V0, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat09V1Valid(t *testing.T) {

	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0x48

	msg, err := ReadBDS05(adsb.ReaderLevel1, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format09V1); !ok {
		t.Errorf("expected a message of type Format09V1, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat09V2Valid(t *testing.T) {

	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0x48

	msg, err := ReadBDS05(adsb.ReaderLevel2, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format09V2); !ok {
		t.Errorf("expected a message of type Format09V2, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat10V0Valid(t *testing.T) {

	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0x50

	msg, err := ReadBDS05(adsb.ReaderLevel0, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format10V0); !ok {
		t.Errorf("expected a message of type Format10V0, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat10V1Valid(t *testing.T) {

	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0x50

	msg, err := ReadBDS05(adsb.ReaderLevel1, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format10V1); !ok {
		t.Errorf("expected a message of type Format10V1, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat10V2Valid(t *testing.T) {

	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0x50

	msg, err := ReadBDS05(adsb.ReaderLevel2, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format10V2); !ok {
		t.Errorf("expected a message of type Format10V2, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat11V0Valid(t *testing.T) {

	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0x58

	msg, err := ReadBDS05(adsb.ReaderLevel0, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format11V0); !ok {
		t.Errorf("expected a message of type Format11V0, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat11V1Valid(t *testing.T) {

	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0x58

	msg, err := ReadBDS05(adsb.ReaderLevel1, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format11V1); !ok {
		t.Errorf("expected a message of type Format11V1, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat11V2Valid(t *testing.T) {

	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0x58

	msg, err := ReadBDS05(adsb.ReaderLevel2, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format11V2); !ok {
		t.Errorf("expected a message of type Format11V2, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat12V0Valid(t *testing.T) {

	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0x60

	msg, err := ReadBDS05(adsb.ReaderLevel0, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format12V0); !ok {
		t.Errorf("expected a message of type Format12V0, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat12V1Valid(t *testing.T) {

	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0x60

	msg, err := ReadBDS05(adsb.ReaderLevel1, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format12V1); !ok {
		t.Errorf("expected a message of type Format12V1, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat12V2Valid(t *testing.T) {

	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0x60

	msg, err := ReadBDS05(adsb.ReaderLevel2, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format12V2); !ok {
		t.Errorf("expected a message of type Format12V2, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat13V0Valid(t *testing.T) {

	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0x68

	msg, err := ReadBDS05(adsb.ReaderLevel0, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format13V0); !ok {
		t.Errorf("expected a message of type Format13V0, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat13V1Valid(t *testing.T) {

	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0x68

	msg, err := ReadBDS05(adsb.ReaderLevel1, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format13V1); !ok {
		t.Errorf("expected a message of type Format13V1, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat13V2Valid(t *testing.T) {

	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0x68

	msg, err := ReadBDS05(adsb.ReaderLevel2, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format13V2); !ok {
		t.Errorf("expected a message of type Format13V2, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat14V0Valid(t *testing.T) {

	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0x70

	msg, err := ReadBDS05(adsb.ReaderLevel0, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format14V0); !ok {
		t.Errorf("expected a message of type Format14V0, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat14V1Valid(t *testing.T) {

	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0x70

	msg, err := ReadBDS05(adsb.ReaderLevel1, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format14V1); !ok {
		t.Errorf("expected a message of type Format14V1, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat14V2Valid(t *testing.T) {

	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0x70

	msg, err := ReadBDS05(adsb.ReaderLevel2, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format14V2); !ok {
		t.Errorf("expected a message of type Format14V2, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat15V0Valid(t *testing.T) {

	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0x78

	msg, err := ReadBDS05(adsb.ReaderLevel0, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format15V0); !ok {
		t.Errorf("expected a message of type Format15V0, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat15V1Valid(t *testing.T) {

	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0x78

	msg, err := ReadBDS05(adsb.ReaderLevel1, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format15V1); !ok {
		t.Errorf("expected a message of type Format15V1, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat15V2Valid(t *testing.T) {

	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0x78

	msg, err := ReadBDS05(adsb.ReaderLevel2, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format15V2); !ok {
		t.Errorf("expected a message of type Format15V2, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat16V0Valid(t *testing.T) {

	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0x80

	msg, err := ReadBDS05(adsb.ReaderLevel0, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format16V0); !ok {
		t.Errorf("expected a message of type Format16V0, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat16V1Valid(t *testing.T) {

	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0x80

	msg, err := ReadBDS05(adsb.ReaderLevel1, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format16V1); !ok {
		t.Errorf("expected a message of type Format16V1, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat16V2Valid(t *testing.T) {

	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0x80

	msg, err := ReadBDS05(adsb.ReaderLevel2, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format16V2); !ok {
		t.Errorf("expected a message of type Format16V2, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat17V0Valid(t *testing.T) {

	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0x88

	msg, err := ReadBDS05(adsb.ReaderLevel0, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format17V0); !ok {
		t.Errorf("expected a message of type Format17V0, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat17V1Valid(t *testing.T) {

	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0x88

	msg, err := ReadBDS05(adsb.ReaderLevel1, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format17V1); !ok {
		t.Errorf("expected a message of type Format17V1, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat17V2Valid(t *testing.T) {

	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0x88

	msg, err := ReadBDS05(adsb.ReaderLevel2, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format17V2); !ok {
		t.Errorf("expected a message of type Format17V2, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat18V0Valid(t *testing.T) {

	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0x90

	msg, err := ReadBDS05(adsb.ReaderLevel0, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format18V0); !ok {
		t.Errorf("expected a message of type Format18V0, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat18V1Valid(t *testing.T) {

	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0x90

	msg, err := ReadBDS05(adsb.ReaderLevel1, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format18V1); !ok {
		t.Errorf("expected a message of type Format18V1, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat18V2Valid(t *testing.T) {

	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0x90

	msg, err := ReadBDS05(adsb.ReaderLevel2, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format18V2); !ok {
		t.Errorf("expected a message of type Format18V2, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat20V0Valid(t *testing.T) {

	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0xA0

	msg, err := ReadBDS05(adsb.ReaderLevel0, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format20V0); !ok {
		t.Errorf("expected a message of type Format20V0, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat20V1Valid(t *testing.T) {

	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0xA0

	msg, err := ReadBDS05(adsb.ReaderLevel1, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format20V1); !ok {
		t.Errorf("expected a message of type Format20V1, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat20V2Valid(t *testing.T) {

	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0xA0

	msg, err := ReadBDS05(adsb.ReaderLevel2, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format20V2); !ok {
		t.Errorf("expected a message of type Format20V2, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat21V0Valid(t *testing.T) {

	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0xA8

	msg, err := ReadBDS05(adsb.ReaderLevel0, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format21V0); !ok {
		t.Errorf("expected a message of type Format21V0, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat21V1Valid(t *testing.T) {

	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0xA8

	msg, err := ReadBDS05(adsb.ReaderLevel1, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format21V1); !ok {
		t.Errorf("expected a message of type Format21V1, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat21V2Valid(t *testing.T) {

	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0xA8

	msg, err := ReadBDS05(adsb.ReaderLevel2, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format21V2); !ok {
		t.Errorf("expected a message of type Format21V2, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat22V0Valid(t *testing.T) {

	data := buildValidBDS05V0Message()
	data[0] = data[0] | 0xB0

	msg, err := ReadBDS05(adsb.ReaderLevel0, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format22V0); !ok {
		t.Errorf("expected a message of type Format22V0, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat22V1Valid(t *testing.T) {

	data := buildValidBDS05V1Message()
	data[0] = data[0] | 0xB0

	msg, err := ReadBDS05(adsb.ReaderLevel1, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format22V1); !ok {
		t.Errorf("expected a message of type Format22V1, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectAndReadFormat22V2Valid(t *testing.T) {

	data := buildValidBDS05V2Message()
	data[0] = data[0] | 0xB0

	msg, err := ReadBDS05(adsb.ReaderLevel2, false, data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := msg.(*Format22V2); !ok {
		t.Errorf("expected a message of type Format22V2, but got %v instead", msg.GetMessageFormat().ToString())
	}
}

func TestDetectBadFormat(t *testing.T) {

	// make a Format 01 message
	data := buildValidBDS05V0Message()
	data[0] = (data[0] & 0x07) | 0x08

	_, err := ReadBDS05(adsb.ReaderLevel0, false, data)
	if err == nil {
		t.Fatal("Expected an error while reading a message with format 01, but message was read")
	}
}