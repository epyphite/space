package srv

import "math"

var (
	secInDay = 24 * 60 * 60
	GM       = 398600.4418
)

//OrbitalElement definition
type OrbitalElement struct {
	Inclination   float64 `json:"inclination"`
	Ascension     float64 `json:"LongitudeAscendingNode"`
	Eccentricity  float64 `json:"eccentricity"`
	Perigee       float64
	Anomaly       float64
	Motion        float64
	SemiMajorAxis float64 `json:"semimajoraxis"`
	Period        float64
	TrueAnomaly   float64
}

//CalculateSemiMajorAxis Will calculate major axis
func (O *OrbitalElement) CalculateSemiMajorAxis() {
	O.SemiMajorAxis = math.Pow(math.Pow(O.Period/(2*math.Pi), 2*GM), (1. / 3))
}

//CalculateTrueAnomaly will get the true anomaly
func (O *OrbitalElement) CalculateTrueAnomaly() {
	O.Period = float64(secInDay) * 1. / O.Motion
	O.TrueAnomaly = O.eccentricityAnomaly(O.Anomaly*math.Pi/180, O.Eccentricity, O.Anomaly*math.Pi/180)
}

func (O *OrbitalElement) eccentricityAnomaly(anomaly float64, eccentricity float64, init float64) float64 {
	iter := 500
	accuracy := 0.0001
	var true float64
	for i := 0; i < iter; i++ {
		true = init - (init-eccentricity*math.Sin(init)-anomaly)/(1.0-eccentricity*math.Cos(init))
		if math.Abs(true-init) < accuracy {
			break
		}
	}
	return true
}
