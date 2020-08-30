package models

//SpacePort structure to hold the needed
// information for a eath bound space port to launch
type SpacePort struct {
	ID                    string
	Name                  string
	Description           string
	LaunchPointAltitude   float64
	AdditionalVelocity    float64
	SpaceportLatitude     float64
	LaunchAzimuth         float64
	EarthRotationVelocity float64
}
