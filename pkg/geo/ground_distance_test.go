package geo

import "testing"

func TestComputeGroundDistance(t *testing.T) {

	distance := ComputeGroundDistance(51.510357, -0.116773, 38.889931, -77.009003)

	if distance < 5897658 || distance > 5897659 {
		t.Errorf("distance: got %v, want ~%v", distance, 5897658.289)
	}
}
