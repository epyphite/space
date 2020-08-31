package models

//Rocket Characteristics and Payload
type Rocket struct {
	ID                                   string
	Name                                 string
	Description                          string
	ThrustToWeightRatioOne               float64       `json:"thrusttoweightratioone"`
	ThrustToWeightRatio                  float64       `json:"thrusttoweightratio`
	RocketMass                           float64       `json:"rocketmass"`
	MaxRocketBodyDiameter                float64       `json:"maxrocketbodydiameter"`
	FairingMass                          float64       `json:"fairingmass"`
	FairingJettisonVelocity              float64       `json:"fairingjettisonvelocity"`
	JettisonedBattery                    float64       `json:"jettisonedbattery"`
	AssumedPayloadMass                   float64       `json:"assumedpayloadmass"`
	SecondStageToRocketMassRatio         float64       `json:"secondstagetorocketmassratio"`
	TransferOrbitStageToRocketMassRatio  float64       `json:"transferorbitstagetorocketmassratio"`
	FirstStageDryToWetMassRatio          float64       `json:"firststagedrytowetmassratio"`
	SecondStageDryToWetMassRatio         float64       `json:"secondstagedrytowetmassratio"`
	TransferOrbitStageDryToWetMassRatio  float64       `json:"transferorbitstagedrytowetmassratio"`
	UnusedPropellantOfFirstStage         float64       `json:"unusedpropellantoffirststage"`
	UnusedPropellantOfSecondStage        float64       `json:"unusedpropellantofsecondstage"`
	UnusedPropellantOfTransferOrbitStage float64       `json:"unusedpropellantoftransferorbitstage"`
	FirstStageIspStartAltitude           float64       `json:"firststageispstartaltitude"`
	FirstStageIspVacuum                  float64       `json:"firststageispvacuum"`
	SecondStageIsp                       float64       `json:"secondstageisp"`
	TransferOrbitStageIsp                float64       `json:"transferorbitstageisp"`
	FirstStageFuel                       float64       `json:"firststagefuel"`
	FirstStageCycle                      float64       `json:"firststagecycle"`
	SecondStageFuel                      float64       `json:"secondstagefuel"`
	SecondStageCycle                     float64       `json:"secondstagecycle"`
	ThirdStageFuel                       float64       `json:"thirdstagefuel"`
	ThirdStageCycle                      float64       `json:"thirdstagecycle"`
	SpacePort                            SpacePort     `json:"spaceport"`
	Orbit                                Orbit         `json:"orbit"`
	AirLaunch                            bool          `json:"airlaunch"`
	Engine                               []EngineSpecs `json:"engine"`
}
