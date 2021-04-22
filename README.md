# Open Vincenty
Implementation of Vincenty's formulae.

## Inverse problem

```
latitudeA  := 59.0
longitudeA := 18.0
latitudeB  := 60.0
longitudeB := 19.0

distance, azimuthA, azimuthB := InverseProblem(latitudeA, longitudeA, latitudeB, longitudeB)
```

## Direct problem
```
latitudeA  := 59.0
longitudeA := 18.0
azimuthA   := 45.0
distance   := 10000

latitudeB, longitudeB, azimuthB := DirectProblem(latitudeA, longitudeA, azimuthA, distance)
```
