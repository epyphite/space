package models

type Orbit struct {
	ID               string
	Name             string
	Description      string
	OrbitPerigee     float64
	OrbitApogee      float64
	OrbitInclination float64
	PerigeeVelocity  float64
	ApogeeVelocity   float64
	OrbitalPeriod    float64
	OrbitalVelocity  float64
	DeltaV           float64
	ExtraVelocity    float64
}
