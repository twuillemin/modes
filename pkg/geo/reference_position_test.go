package geo

import "testing"

func TestReferencePosition(t *testing.T) {
	// even message: 8D40621D58C382D690C8AC2863A7 => BDS 0,5
	// Encoded Latitude:                  93000
	// Encoded Longitude:                 51372

	lat, lon := GetCPRPositionWithReference(
		93000,
		51372,
		52.258,
		3.918,
		false)

	if lat < 52.257 || lat > 52.258 {
		t.Errorf("latitude: got %v, want ~%v", lat, 52.2572)
	}

	if lon < 3.919 || lon > 3.920 {
		t.Errorf("longitude: got %v, want ~%v", lon, 3.91937)
	}
}
