package models

//EngineSpecs for calculating
type EngineSpecs struct {
	ID                     string
	Name                   string
	Description            string
	FirstStageEngineThrust float64   `json:"firststageenginethrust"`
	FirstStageIspSeaLevel  float64   `json:"firststageispsealevel"`
	FirstStageISPVac       float64   `json:"firststageispvac"`
	FuelType               FuelType  `json:"fueltype"`
	FuelCycle              FuelCycle `json:"fuelcycle"`
	FirstStageEngineNumber float64   `json:"firststageenginenumber"`
}

/*

	var s_engine_data = [["Launcher E-2",		10000.,	290.5, 327., 0, 0, 4,"Launcher 1st stage"],
			     ["Launcher E-2 vac",	12600.,	  0.0, 365., 0, 0, 1,"Launcher 2nd stage"],
			     ["Ursa Ripley",		15900., 262.0, 308., 0, 0, 1,"URSA Major ORSC engine for 1st stage"],
			     ["Ursa Hadley",		2900.,	  0.0, 345., 0, 0, 1,"URSA Major ORSC engine for 2nd stage"],
			     ["RL Rutherford",		1800.,	256.0, 303., 0, 2, 9,"RocketLab Electron 1st stage"],
			     ["RL Rutherford vac", 	2500.,	  0.0, 333.,  0, 2, 1,"Rocketlab Electron 2nd stage"],
			     ["Virgin Newton III",	33600.,	290.0, 308.,  0, 1, 1,"Virgin Orbit Launcherone 1st stage"],
			     ["Virgin Newton IV", 	2200.,	  0.0, 328.,  0, 1, 1,"Virgin Orbit Launcherone 2nd stage"],
			     ["Firefly Reaver",		16900.,	265.0, 295.6, 0, 1, 4,"Firefly Alpha 1st stage"],
			     ["Firefly Lightning",	7100.,    0.0, 322.,  0, 1, 1,"Firefly Alpha 2nd stage"],
			     ["Relativity Aeon-1",	7620.,	267.0, 310.,  1, 1, 9,"Relativity Space Terran-1 1st stage"],
			     ["Relativity Aeon-1 vac",	10200.,	  0.0, 360.,  1, 1, 1,"Relativity Space Terran-1 2nd stage"],
			     ["Isar Aqula SL",		7650.,	256.0, 303.,  0, 1, 9,"Isar Aerospace Spectrum 1st stage"],
			     ["Isar Aqula VAC",		9580.,	  0.0, 333.,  0, 1, 1,"Isar Aerospace Spectrum 2nd stage"],
			     ["SpaceX Merlin 1C",	56700.,	263.0, 304.,  0, 1, 1,"SpaceX Falcone-1 1st stage"],
			     ["SpaceX Kerstel 2",	3100.,	  0.0, 317.,  0, 3, 1,"SpaceX Falcone-1 2nd stage"],
			     ["SpaceX Merlin 1D+",	86200.,	282.0, 311.,  0, 1, 9,"SpaceX Falcone-9FT 1st stage"],
			     ["SpaceX Merlin 1D+v",	95300.,	  0.0, 348.,  0, 1, 1,"SpaceX Falcone-9FT 2nd stage"],
			     ["Energomash RD-170",	740000,	309.5, 337.2, 0, 0, 1,"Yuzhnoye Zenit 1st stage"],
			     ["Energomash RD-120",	93000.,	  0.0, 349.,  0, 0, 1,"Yuzhnoye Zenit 2nd stage"],
			     ["Energomash RD-180",	390000,	311.9, 338.4, 0, 0, 1,"United Launch Alliance Atlas V 1st stage"],
			     ["AR RL-10C",		10400.,	  0.0, 450.,  2, 0, 1,"United Launch Alliance Atlas V 2nd stage"],
			     ["Energomash RD-181",	196000,	311.9, 339.2, 0, 0, 2,"Northrop Grumman Antares 230 1st stage"],
			     ["NG Castor 30XL",		48300.,	  0.0, 304.,  4, 3, 1,"Northrop Grumman Antares 230 2nd stage"],
			     ["AR F-1",			690000,	263.0, 304.,  0, 1, 5,"Saturn V 1st stage"],
			     ["AR J-2",			104000,	  0.0, 421.,  2, 1, 5,"Saturn V 2nd stage"],
			     ["AR RS-68A",		320000,	360.0, 412.,  2, 1, 1,"United Launch Alliance Delta IV 1st stage"],
			     ["AR RL-10B",		11300.,	  0.0, 462.,  2, 0, 1,"United Launch Alliance Delta IV 2nd stage"],
			     ["AR AJ-26",		166000,	297.0, 331.,  0, 0, 2,"Orbital (NG now) Antares 100 1st stage"],
			     ["NG Castor 30A",		30000.,	  0.0, 301.,  4, 3, 1,"Orbital (NG now) Antares 100 2nd stage"],
			     ["Blue Origin BE-4",	250000, 310.0, 335.,  1, 0, 7,"Blue Origin New Glenn 1st stage"],
			     ["BLue Origin BE-4U",	290000,	  0.0, 358.,  1, 0, 1,"Blue Origin New Glenn 2nd stage"],
			     ["SpaceX Raptor",		200000,	330.0, 356.,  1, 0, 37,"SpaceX Starship 1st stage"],
			     ["SpaceX Raptor vac",	230000,	  0.0, 380.,  1, 0, 3,"SpaceX Starship 2nd stage"],
				];
*/
