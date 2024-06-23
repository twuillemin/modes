package bitutils

import (
	"testing"
)

func TestGillhamToAltitude(t *testing.T) {

	tests := []struct {
		string
		int32
	}{
		{"00000000010", -1000},
		{"00000001010", -500},
		{"00000011010", 0},
		{"00000011110", 100},
		{"00000010011", 600},
		{"00000110010", 1000},
		{"00001001001", 5800},
		{"00011100100", 10300},
		{"01100011010", 32000},
		{"01110000100", 46300},
		{"10000000011", 126600},
		{"10000000001", 126700},
	}

	for _, test := range tests {
		alt, err := GillhamToAltitude(string2Bits(test.string))
		if err != nil {
			t.Errorf("%v", err)
		}
		if alt != test.int32 {
			t.Errorf("for Altitude: got %v, want %v", alt, test.int32)
		}
	}
}

func string2Bits(s string) (bool, bool, bool, bool, bool, bool, bool, bool, bool, bool, bool, bool) {
	d2 := s[0] == '1'
	d4 := s[1] == '1'
	a1 := s[2] == '1'
	a2 := s[3] == '1'
	a4 := s[4] == '1'
	b1 := s[5] == '1'
	b2 := s[6] == '1'
	b4 := s[7] == '1'
	c1 := s[8] == '1'
	c2 := s[9] == '1'
	c4 := s[10] == '1'

	return false, d2, d4, a1, a2, a4, b1, b2, b4, c1, c2, c4
}
