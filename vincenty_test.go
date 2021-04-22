package vincenty

import (
	"testing"
)

func TestInverseProblem(t *testing.T) {
	s, a, b := InverseProblem(-37.951033, 144.424868, -37.652821, 143.926496)
	if s < 54972.0 || s > 54973.0 {
		t.Fail()
	}
	if a < 306.0 || a > 307.0 {
		t.Fail()
	}
	if b < 307.0 || b > 308.0 {
		t.Fail()
	}
}

func TestDirectProblem_Zero(t *testing.T) {
	latitude, longitude, azimuth := DirectProblem(59.349133, 18.022666, 45.0, 10000.0)
	if latitude < 59.0 || latitude > 60.0 {
		t.Fail()
	}
	if longitude < 18.0 || longitude > 19.0 {
		t.Fail()
	}
	if azimuth < 45.0 || azimuth > 46.0 {
		t.Fail()
	}
}
