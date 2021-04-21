package vincenty

import "testing"

func TestDirectProblem_Zero(t *testing.T) {
	latitude, longitude, azimuth := DirectProblem(0, 0, 0, 0)
	if latitude != 0 && longitude != 0 && azimuth != 0 {
		t.Fail()
	}
}
