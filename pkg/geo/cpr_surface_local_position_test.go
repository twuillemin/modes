package geo

import "testing"

func TestGetCPRSurfaceLocalPosition1(t *testing.T) {

	// odd message: 8C4841753A9A153237AEF0F275BE => BDS 0,6
	// Encoded Latitude:                  39195
	// Encoded Longitude:                 110320
	// Reference = latitude: 51.990, Longitude: 4.375

	lat, lon := GetCPRSurfaceLocalPosition(
		39195,
		110320,
		51.990,
		4.375,
		true,
	)

	if lat < 52.3205 || lat > 52.3206 {
		t.Errorf("latitude: got %v, want ~%v", lat, 52.32056)
	}

	if lon < 4.7357 || lon > 4.7358 {
		t.Errorf("longitude: got %v, want ~%v", lon, 4.73573)
	}
}

func TestGetCPRSurfaceLocalPosition2(t *testing.T) {

	// odd message: 8FC8200A3AB8F5F893096B000000 => BDS 0,6 // Bad CRC
	// Encoded Latitude:                  64585
	// Encoded Longitude:                 67947
	// Reference = latitude: 51.990, Longitude: 4.375

	lat, lon := GetCPRSurfaceLocalPosition(
		64585,
		67947,
		-43.5,
		172.5,
		true,
	)

	if lat < -43.4857 || lat > -43.4856 {
		t.Errorf("latitude: got %v, want ~%v", lat, -43.48564)
	}

	if lon < 172.5394 || lon > 172.5395 {
		t.Errorf("longitude: got %v, want ~%v", lon, 172.53942)
	}
}
