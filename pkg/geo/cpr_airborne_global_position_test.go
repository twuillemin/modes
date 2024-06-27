package geo

import "testing"

func TestGetCPRAirborneGlobalPosition1(t *testing.T) {
	// even message: 8D40058B58C901375147EFD09357 => BDS 0,5
	// Encoded Latitude:                  39848
	// Encoded Longitude:                 83951
	// odd message: 8D40058B58C904A87F402D3B8C59 => BDS 0,5
	// Encoded Latitude:                  21567
	// Encoded Longitude:                 81965

	lat, lon, err := GetCPRAirborneGlobalPosition(
		39848,
		83951,
		21567,
		81965,
		false)

	if err != nil {
		t.Errorf("%v", err)
	}

	if lat < 49.817 || lat > 49.818 {
		t.Errorf("latitude: got %v, want ~%v", lat, 49.81755)
	}

	if lon < 6.084 || lon > 6.085 {
		t.Errorf("longitude: got %v, want ~%v", lon, 6.08442)
	}
}

func TestGetCPRAirborneGlobalPosition2(t *testing.T) {
	// even message: 8D40621D58C382D690C8AC2863A7 => BDS 0,5
	// Encoded Latitude:                  93000
	// Encoded Longitude:                 51372
	// odd message: 8D40621D58C386435CC412692AD6 => BDS 0,5
	// Encoded Latitude:                  74158
	// Encoded Longitude:                 50194

	lat, lon, err := GetCPRAirborneGlobalPosition(
		93000,
		51372,
		74158,
		50194,
		true)

	if err != nil {
		t.Errorf("%v", err)
	}

	if lat < 52.257 || lat > 52.258 {
		t.Errorf("latitude: got %v, want ~%v", lat, 52.2572)
	}

	if lon < 3.9193 || lon > 3.9194 {
		t.Errorf("longitude: got %v, want ~%v", lon, 3.91937)
	}
}
