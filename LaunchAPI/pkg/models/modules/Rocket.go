package models

//Rocket Characteristics and Payload
type Rocket struct {
	ID                                   string
	Name                                 string
	Description                          string
	ThrustToWeightRatioOne               float64
	ThrustToWeightRatio                  float64
	RocketMass                           float64
	MaxRocketBodyDiameter                float64
	FairingMass                          float64
	FairingJettisonVelocity              float64
	JettisonedBattery                    float64
	AssumedPayloadMass                   float64
	SecondStageToRocketMassRatio         float64
	TransferOrbitStageToRocketMassRatio  float64
	FirstStageDryToWetMassRatio          float64
	SecondStageDryToWetMassRatio         float64
	TransferOrbitStageDryToWetMassRatio  float64
	UnusedPropellantOfFirstStage         float64
	UnusedPropellantOfSeconStage         float64
	UnusedPropellantOfTransferOrbitStage float64
	FirstStageIspStartAltitude           float64
	FirstStageIspVacuum                  float64
	SecondStageIsp                       float64
	TransferOrbitStageIsp                float64
	FirstStageFuel                       float64
	FirstStageCycle                      float64
	SecondStageFuel                      float64
	SecondStageCycle                     float64
	ThirdStageFuel                       float64
	ThirdStageCycle                      float64
	SpacePort                            SpacePort
	Orbit                                Orbit
	AirLaunch                            bool
	Engine                               []EngineSpecs
}
