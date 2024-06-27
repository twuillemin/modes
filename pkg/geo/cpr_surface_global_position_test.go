package geo

import "testing"

func TestGetCPRSurfaceGlobalPosition1(t *testing.T) {

	// even message: 8C4841753AAB238733C8CD4020B1 => BDS 0,6
	// Encoded Latitude:                  115609
	// Encoded Longitude:                 116941
	// odd message: 8C4841753A8A35323FAEBDAC702D => BDS 0,6
	// Encoded Latitude:                  39199
	// Encoded Longitude:                 110269
	// Reference = latitude: 51.990, Longitude: 4.375

	lat, lon, err := GetCPRSurfaceGlobalPosition(
		115609,
		116941,
		39199,
		110269,
		false,
		51.990,
		4.375,
	)

	if err != nil {
		t.Errorf("%v", err)
	}

	if lat < 52.3206 || lat > 52.3207 {
		t.Errorf("latitude: got %v, want ~%v", lat, 52.32061)
	}

	if lon < 4.7347 || lon > 4.7348 {
		t.Errorf("longitude: got %v, want ~%v", lon, 4.73473)
	}
}

func TestGetCPRSurfaceGlobalPosition2(t *testing.T) {

	// even message: 8CC8200A3AC8F009BCDEF2000000 => BDS 0,6 / CRC invalid
	// Encoded Latitude:                  1246
	// Encoded Longitude:                 57074
	// odd message: 8FC8200A3AB8F5F893096B000000 => BDS 0,6 / CRC invalid
	// Encoded Latitude:                  64585
	// Encoded Longitude:                 67947
	// Reference = latitude: 51.990, Longitude: 4.375

	lat, lon, err := GetCPRSurfaceGlobalPosition(
		1246,
		57074,
		64585,
		67947,
		false,
		-43.496,
		172.558,
	)

	if err != nil {
		t.Errorf("%v", err)
	}

	if lat < -43.4857 || lat > -43.4856 {
		t.Errorf("latitude: got %v, want ~%v", lat, -43.48564)
	}

	if lon < 172.5394 || lon > 172.5395 {
		t.Errorf("longitude: got %v, want ~%v", lon, 172.53942)
	}
}
