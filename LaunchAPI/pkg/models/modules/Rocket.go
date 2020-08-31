package models

//Rocket Characteristics and Payload
type Rocket struct {
	ID                                        string
	Name                                      string
	Description                               string
	Stages                                    int
	ThrustToWeightRatio                       []float64
	LiftOffMass                               float64
	RocketMass                                float64
	MaxRocketBodyDiameter                     float64
	FairingMass                               float64
	AssumedPayloadMass                        float64
	SecondStageToRocketMassRatio              float64
	TransferOrbitStageToRocketMassRatio       float64
	FirstStageDrytoWetMassRatio               float64
	SecondStageDrytoWetMassRatio              float64
	TransferOrbitStageDryToWetMassRatio       float64
	UnusedPropellantOf1stStage                float64
	UnusedPropellantOf2ndStage                float64
	UnusedPropellantOfTransferOrbitStage      float64
	FirstStageIspSeaLevelOrAtTheStartAltitude float64
	FirstStageIspVacuum                       float64
	SecondStageIsp                            float64
	TransferOrbitStageIsp                     float64
	SpecificImpulseVariation                  float64
	Engine                                    []EngineSpecs
}

// Name, ThrustToWeightRatioFirst,ThrustToWeightRatioSecond, LiftOffMass,RocketMass,MaxRocketBodyDiameter,FairingMass,AssumedPayloadMass,SecondStageToRocketMassRatio,TransferOrbitStageToRocketMassRatio,FirstStageDrytoWetMassRatio,SecondStageDrytoWetMassRatio,TransferOrbitStageDryToWetMassRatio,UnusedPropellantOf1stStage,UnusedPropellantOf2ndStage,UnusedPropellantOfTransferOrbitStage,FirstStageIspSeaLevelOrAtTheStartAltitude,FirstStageIspVacuum,SecondStageIsp,TransferOrbitStageIsp,SpecificImpulseVariation
//["NASA Saturn V",			1.165,0.78,2909200,10.1, 8000, 3500,  0, 48600,  21.284,  4.228,  5.677,  8.081, 10.976,2.6, 2.0, 2,   263.0, 304,   421, 421, 0,1, 2,1, 2,1,    "Cape Canaveral",	"Moon",0, "Apollo lunar program launcher"],
