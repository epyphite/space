package constants

import "epyphite/space/v1/ESA/pkg/utils"

var (
	_Kilo  = 1000
	_FKilo = float64(_Kilo)
	//GravitationalConstant m^3 kg^-1 s^-2
	GravitationalConstant = 6.67428e-11

	//SolarMass m^3 s^-2
	SolarMass = 1.32712440041e20

	//SunMass Sun Mass
	SunMass = SolarMass / GravitationalConstant

	//GeoCentricGravitational  Geo Centric Gravitational Constast
	GeoCentricGravitational = 3.986004415e14

	mercuryMass = SunMass / 6.0236e6
	venusMass   = SunMass / 4.08523710e5
	marsMass    = SunMass / 3.09870359e6
	jupiterMass = SunMass / 1.047348644e3
	saturnMass  = SunMass / 3.4979018e3
	uranusMass  = SunMass / 2.290298e4
	neptuneMass = SunMass / 1.941226e4

	sunRadiusEquatorial = 696000 * _Kilo

	_mercuryRadiusEquatorial = 2439.7
	mercuryRadiusEquatorial  = _mercuryRadiusEquatorial * _FKilo

	_venusRadiusEquatorial = 6051.8
	venusRadiusEquatorial  = _venusRadiusEquatorial * _FKilo

	_earthRadiusEquatorial = 6378.1366
	earthRadiusEquatorial  = _earthRadiusEquatorial * _FKilo

	_marsRadiusEquatorial = 3396.19
	marsRadiusEquatorial  = _marsRadiusEquatorial * _FKilo

	jupiterRadiusEquatorial = 71492 * _Kilo
	saturnRadiusEquatorial  = 60268 * _Kilo
	uranusRadiusEquatorial  = 25559 * _Kilo
	neptuneRadiusEquatorial = 24764 * _Kilo

	mercuryMu = mercuryMass * GravitationalConstant
	venusMu   = venusMass * GravitationalConstant
	marsMu    = marsMass * GravitationalConstant
	jupiterMu = jupiterMass * GravitationalConstant
	saturnMu  = saturnMass * GravitationalConstant
	uranusMu  = uranusMass * GravitationalConstant
	neptuneMu = neptuneMass * GravitationalConstant
	//EarthMu Eart Gravitational Constant
	EarthMu = GeoCentricGravitational

	mercuryRadiusPolar = mercuryRadiusEquatorial
	mercuryRadiusMean  = mercuryRadiusEquatorial
	venusRadiusPolar   = venusRadiusEquatorial
	venusRadiusMean    = venusRadiusEquatorial

	// The following constants are not from IAU
	_earthRadiusMean  = 6371.0
	earthRadiusMean   = _earthRadiusMean * _FKilo
	_earthRadiusPolar = 6356.8
	earthRadiusPolar  = _earthRadiusPolar * _FKilo

	_marsRadiusMean  = 3389.5
	marsRadiusMean   = _marsRadiusMean * _FKilo
	_marsRadiusPolar = 3376.2
	marsRadiusPolar  = _marsRadiusPolar * _FKilo

	jupiterRadiusMean  = 69911 * _Kilo
	jupiterRadiusPolar = 66854 * _Kilo

	saturnRadiusMean  = 58232 * _Kilo
	saturnRadiusPolar = 54364 * _Kilo

	uranusRadiusMean  = 25362 * _Kilo
	uranusRadiusPolar = 24973 * _Kilo

	neptuneRadiusMean  = 24622 * _Kilo
	neptuneRadiusPolar = 24341 * _Kilo
	// 4.1 s, 56 minutes, 23 hours
	earthSiderealDay = utils.TimeDelta(23, 56, 4.1)
)
