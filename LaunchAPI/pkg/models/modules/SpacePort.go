package models

//SpacePort structure to hold the needed
// information for a eath bound space port to launch
type SpacePort struct {
	ID                                 string  `json:"id"`
	Name                               string  `json:"name"`
	Description                        string  `json:"description"`
	LaunchPointAltitude                float64 `json:"launchpointaltitude"`
	AdditionalVelocity                 float64 `json:"additionalvelocity"`
	SpaceportLatitude                  float64 `json:"spaceportlatitude"`
	SpaceportLongitude                 float64 `json:"spaceportlongitude"`
	LaunchAzimuth                      float64 `json:"launchazimuth"`
	EarthRotationVelocity              float64 `json:"earthrotationvelocity"`
	AuxiliaryAngle                     float64 `json:"auxiliaryangle"`
	LaunchPointAltitudeOrbitalVelocity float64 `json:"launchpointaltitudeorbitalvelocity"`
	AbsoluteOrbitalVelocity            float64 `json:"absoluteorbitalvelocity"`
}
