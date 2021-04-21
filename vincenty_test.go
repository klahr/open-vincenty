package vincenty

import "testing"

func TestInverseProblem_Zero(t *testing.T) {
	azimuth1, azimuth2, distance := InverseProblem(0, 0, 0, 0)
	if azimuth1 != 0 && azimuth2 != 0 && distance != 0 {
		t.Fail()
	}
}

func TestDirectProblem_Zero(t *testing.T) {
	latitude, longitude, azimuth := DirectProblem(0, 0, 0, 0)
	if latitude != 0 && longitude != 0 && azimuth != 0 {
		t.Fail()
	}
}
