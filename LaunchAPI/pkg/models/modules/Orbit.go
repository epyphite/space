package models

//Orbit model Package fro name of the common orbits
type Orbit struct {
	ID                              string  `json:"id"`
	Name                            string  `json:"name"`
	Description                     string  `json:"description"`
	OrbitPerigee                    float64 `json:"orbitperigee"`
	OrbitApogee                     float64 `json:"orbitapogee"`
	OrbitInclination                float64 `json:"orbitinclination"`
	PerigeeVelocity                 float64 `json:"perigeevelocity"`
	ApogeeVelocity                  float64 `json:"apogeevelocity"`
	OrbitalPeriod                   float64 `json:"orbitalperiod"`
	OrbitalVelocity                 float64 `json:"orbitalvelocity"`
	DeltaV                          float64 `json:"deltav"`
	ExtraVelocity                   float64 `json:"extravelocity"`
	ExtraSpeedForFlightToThePlanets float64 `json:"extraspeedforflighttotheplanets"`
}
