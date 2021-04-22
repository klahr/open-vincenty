package vincenty

import "math"

//https://en.wikipedia.org/wiki/Vincenty%27s_formulae

// Fins the ellipsoidal distance between two coordinates.
func InverseProblem(latitude1, longitude1, latitude2, longitude2 float64) (float64, float64, float64) {
	f := 1.0 / 298.257223563
	a := 6378137.0
	b := (1.0 - f) * a

	phi1 := deg2Rad(latitude1)
	phi2 := deg2Rad(latitude2)
	L1 := deg2Rad(longitude1)
	L2 := deg2Rad(longitude2)

	U1 := math.Atan((1.0 - f) * math.Tan(phi1))
	U2 := math.Atan((1.0 - f) * math.Tan(phi2))
	L := L2 - L1

	lambda := L
	cosSqAlpha := 1.0
	sigma := 0.0
	cos2SigmaM := 0.0

	for i := 0; i < 10; i++ {
		sinSigma := math.Sqrt(math.Pow(math.Cos(U2)*math.Sin(lambda), 2) + math.Pow(math.Cos(U1)*math.Sin(U2)-math.Sin(U1)*math.Cos(U2)*math.Cos(lambda), 2))
		cosSigma := math.Sin(U1)*math.Sin(U2) + math.Cos(U1)*math.Cos(U2)*math.Cos(lambda)
		sigma = math.Atan2(sinSigma, cosSigma)
		sinAlpha := (math.Cos(U1) * math.Cos(U2) * math.Sin(lambda)) / sinSigma
		cosSqAlpha = 1.0 - sinAlpha*sinAlpha
		cos2SigmaM = cosSigma - (2.0*math.Sin(U1)*math.Sin(U2))/cosSqAlpha

		C := (f / 16.0) * cosSqAlpha * (4.0 + f*(4.0-3.0*cosSqAlpha))
		lambda = L + (1.0-C)*f*sinAlpha*(sigma+C*sinSigma*(cos2SigmaM+C*cosSigma*(-1.0+2.0*cos2SigmaM*cos2SigmaM)))
	}

	u2 := cosSqAlpha * ((a*a - b*b) / (b * b))

	A := 1.0 + (u2/16384.0)*(4096.0+u2*(-768.0+u2*(320.0-175.0*u2)))
	B := (u2 / 1024.0) * (256.0 + u2*(-128.0+u2*(74.0-47.0*u2)))

	dSigma := B * math.Sin(sigma) * (cos2SigmaM + 0.25*B*(math.Cos(sigma)*(-1.0+2.0*cos2SigmaM*cos2SigmaM)-(B/6.0)*cos2SigmaM*(-3.0+4.0*math.Sin(sigma)*math.Sin(sigma))*(-3.0+4.0*cos2SigmaM*cos2SigmaM)))

	s := b * A * (sigma - dSigma)
	alpha1 := math.Atan2(math.Cos(U2)*math.Sin(lambda), math.Cos(U1)*math.Sin(U2)-math.Sin(U1)*math.Cos(U2)*math.Cos(lambda))
	alpha2 := math.Atan2(math.Cos(U1)*math.Sin(lambda), -math.Sin(U1)*math.Cos(U2)+math.Cos(U1)*math.Sin(U2)*math.Cos(lambda))

	alpha1Deg := rad2Deg(alpha1)
	for alpha1Deg < 0.0 {
		alpha1Deg += 360.0
	}
	for alpha1Deg > 360.0 {
		alpha1Deg -= 360.0
	}

	alpha2Deg := rad2Deg(alpha2)
	for alpha2Deg < 0.0 {
		alpha2Deg += 360.0
	}
	for alpha2Deg > 360.0 {
		alpha2Deg -= 360.0
	}

	return s, alpha1Deg, alpha2Deg
}

// Finds the position given a start position, azimuth and distance.
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

	alpha2Deg := rad2Deg(alpha2)

	for alpha2Deg < 0.0 {
		alpha2Deg += 360.0
	}
	for alpha2Deg > 360.0 {
		alpha2Deg -= 360.0
	}

	return rad2Deg(phi2), rad2Deg(L2), alpha2Deg
}

func deg2Rad(deg float64) float64 {
	return deg * math.Pi / 180.0
}

func rad2Deg(rad float64) float64 {
	return rad * 180.0 / math.Pi
}
