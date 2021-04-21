package vincenty

import "math"

//https://en.wikipedia.org/wiki/Vincenty%27s_formulae
//https://www.movable-type.co.uk/scripts/latlong-vincenty.html

func InverseProblem(latitude1, longitude1, latitude2, longitude2 float64) (float64, float64, float64) {
	// TODO: Implement.

	s := 0.0
	alpha1 := 0.0
	alpha2 := 0.0

	return alpha1, alpha2, s
}

func DirectProblem(latitude, longitude, azimuth, distance float64) (float64, float64, float64) {
	f := 1.0 / 298.257223563
	a := 6378137.0
	b := (1.0 - f) * a

	phi1 := deg2Rad(latitude)
	L1 := deg2Rad(longitude)
	alpha1 := deg2Rad(azimuth)

	U1 := math.Atan((1.0 - f) * math.Tan(phi1))
	sigma1 := math.Atan2(math.Tan(U1), math.Cos(alpha1))
	sinAlpha := math.Cos(U1) * math.Sin(alpha1)
	cos2Alpha := (1.0 - sinAlpha) * (1.0 + sinAlpha)
	u2 := (1.0 - sinAlpha*sinAlpha) * ((a*a - b*b) / (b * b))

	A := 1.0 + (u2/16384.0)*(4096.0+u2*(-768.0+u2*(320.0-175.0*u2)))
	B := (u2 / 1024.0) * (256.0 + u2*(-128.0+u2*(74.0-47.0*u2)))

	sigma := distance / (b * A)

	var twoSigmaM float64
	for i := 0; i < 100; i++ {
		twoSigmaM = 2.0*sigma1 + sigma
		dSigma := B * math.Sin(sigma) * (math.Cos(twoSigmaM) + 0.25*B*(math.Cos(sigma)*(-1.0+2.0*math.Cos(twoSigmaM)*math.Cos(twoSigmaM))-(B/6.0)*math.Cos(twoSigmaM)*(-3.0+4.0*math.Sin(sigma)*math.Sin(sigma))*(-3.0+4.0*math.Cos(twoSigmaM)*math.Cos(twoSigmaM))))
		sigma = distance/(b*A) + dSigma
	}

	phi2 := math.Atan2(math.Sin(U1)*math.Cos(sigma)+math.Cos(U1)*math.Sin(sigma)*math.Cos(alpha1), (1.0-f)*math.Sqrt(sinAlpha*sinAlpha+math.Pow(math.Sin(U1)*math.Sin(sigma)-math.Cos(U1)*math.Cos(sigma)*math.Cos(alpha1), 2)))
	lambda := math.Atan2(math.Sin(sigma)*math.Sin(alpha1), math.Cos(U1)*math.Cos(sigma)-math.Sin(U1)*math.Sin(sigma)*math.Cos(alpha1))
	C := (f / 16.0) * cos2Alpha * (4.0 + f*(4.0-3.0*cos2Alpha))
	L := lambda - (1.0-C)*f*sinAlpha*(sigma+C*math.Sin(sigma)*(math.Cos(twoSigmaM)+C*math.Cos(sigma)*(-1.0+2.0*math.Cos(twoSigmaM)*math.Cos(twoSigmaM))))
	L2 := L + L1
	alpha2 := math.Atan2(sinAlpha, -math.Sin(U1)*math.Sin(sigma)+math.Cos(U1)*math.Cos(sigma)*math.Cos(alpha1))

	return rad2Deg(phi2), rad2Deg(L2), rad2Deg(alpha2)
}

func deg2Rad(deg float64) float64 {
	return deg * math.Pi / 180.0
}

func rad2Deg(rad float64) float64 {
	return rad * 180.0 / math.Pi
}
