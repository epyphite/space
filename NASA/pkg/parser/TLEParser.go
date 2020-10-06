package parser

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"epyphite/space/v1/NASA/pkg/models/modules"
)

/*
parseTLELineOne
1 1 Line Number
3-7 25544 Satellite Catalog Number
8 U Elset Classification
10-17 98067A International Designator
19-32 04236.56031392 Element Set Epoch (UTC)*Note: spaces are acceptable in columns 21 & 22
34-43 .00020137 1st Derivative of the Mean Motion with respect to Time
45-52 00000-0 2nd Derivative of the Mean Motion with respect to Time (decimal point assumed)
54-61 16538-3 B* Drag Term
63 0 Element Set Type
65-68 999 Element Number
69 3 Checksum
*/
func parseTLELineOne(LineString string) (modules.TLEFormatLine1, error) {
	var line1 modules.TLEFormatLine1
	var err error

	Line := []rune(LineString)
	line1.LineNumber, err = strconv.Atoi(string(Line[0]))
	if err != nil {
		log.Errorln(err)
	}
	log.Debugln("Parsing LineNumber :", line1.LineNumber)
	line1.SatelliteNumber, err = strconv.Atoi(string(Line[2:7]))
	if err != nil {
		log.Errorln(err)
	}
	log.Debugln("Parsing SatelliteNumber :", line1.SatelliteNumber)

	line1.ElsetClassification = string(Line[7])
	if err != nil {
		log.Errorln(err)
	}
	log.Debugln("Parsing ElsetClassification :", line1.ElsetClassification)

	line1.InternationalDesingator = string(Line[9:16])
	log.Debugln("Parsing InternationalDesingator :", line1.InternationalDesingator)

	line1.ElementSetEpoch = string(Line[18:31])
	log.Debugln("Parsing ElementSetEpoch :", line1.ElementSetEpoch)

	line1.TimeMotionDerivativeOne = string(Line[33:43])
	log.Debugln("Parsing TimeMotionDerivativeOne :", line1.TimeMotionDerivativeOne)

	line1.TimeMotionDerivativeTwo = string(Line[44:52])
	log.Debugln("Parsing TimeMotionDerivativeTwo :", line1.TimeMotionDerivativeTwo)

	line1.BDragTerm = string(Line[53:61])
	log.Debugln("Parsing BDragTerm :", line1.BDragTerm)

	line1.ElementSetType, err = strconv.Atoi(string(Line[62]))
	if err != nil {
		log.Errorln(err)
	}
	log.Debugln("Parsing ElementSetType : ", line1.ElementSetType, Line[62])

	line1.ElementNumber, err = strconv.Atoi(string(Line[65:68]))

	if err != nil {
		log.Errorln(err)
	}
	log.Debugln("Parsing ElementNumber :", line1.ElementNumber, string(Line[64:68]))
	line1.Checksum, err = strconv.Atoi(string(Line[68]))
	if err != nil {
		log.Errorln(err)
	}
	log.Debugln("Parsing Checksum :", line1.Checksum)
	return line1, err
}

/*
parseTLELineTwo
1 2 Line Number
3-7 25544 Satellite Catalog Number
9-16 51.6335 Orbit Inclination (degrees)
18-25 344.7760 Right Ascension of Ascending Node (degrees)
27-33 0007976 Eccentricity (decimal point assumed)
35-42 126.2523 Argument of Perigee (degrees)
44-51 325.9359 Mean Anomaly (degrees)
53-63 15.70406856 Mean Motion (revolutions/day)
64-68 32890 Revolution Number at Epoch
69 3 Checksum
*/
func parseTLELineTwo(LineString string) (modules.TLEFormatLine2, error) {
	var line2 modules.TLEFormatLine2
	var err error
	Line := []rune(LineString)

	line2.LineNumber, err = strconv.Atoi(string(Line[0]))
	if err != nil {
		log.Errorln(err)
	}
	log.Debugln("Parsing LineNumber :", line2.LineNumber)

	line2.SatelliteNumber, err = strconv.Atoi(string(Line[2:7]))
	if err != nil {
		log.Errorln(err)
	}
	log.Debugln("Parsing SatelliteNumber :", line2.SatelliteNumber)

	line2.OrbitInclination, err = strconv.ParseFloat(string(Line[9:16]), 64)
	if err != nil {
		log.Errorln(err)
	}
	log.Debugln("Parsing OrbitInclination :", line2.OrbitInclination)

	line2.RAAN, err = strconv.ParseFloat(string(Line[17:25]), 64)
	if err != nil {
		log.Errorln(err)
	}
	log.Debugln("Parsing RAAN :", line2.RAAN)

	line2.Eccentricity, err = strconv.ParseFloat(string(Line[26:33]), 64)
	if err != nil {
		log.Errorln(err)
	}
	log.Debugln("Parsing Eccentricity :", line2.Eccentricity)

	line2.ArgumentPerigee, err = strconv.ParseFloat(string(Line[35:42]), 64)
	if err != nil {
		log.Errorln(err)
	}
	log.Debugln("Parsing ArgumentPerigee :", line2.ArgumentPerigee)

	line2.MeanAnomaly, err = strconv.ParseFloat(string(Line[43:51]), 64)
	if err != nil {
		log.Errorln(err)
	}
	log.Debugln("Parsing MeanAnomaly :", line2.MeanAnomaly)

	line2.MeanMotion, err = strconv.ParseFloat(string(Line[52:63]), 64)
	if err != nil {
		log.Errorln(err)
	}
	log.Debugln("Parsing MeanMotion :", line2.MeanMotion, string(Line[52:64]))

	line2.RevolutionNumber, err = strconv.ParseFloat(string(Line[64:68]), 64)
	if err != nil {
		log.Errorln(err)
	}
	log.Debugln("Parsing RevolutionNumber :", line2.RevolutionNumber)

	line2.Checksum, err = strconv.Atoi(string(Line[68]))
	if err != nil {
		log.Errorln(err)
	}
	log.Debugln("Parsing Checksum :", line2.Checksum)

	return line2, err
}

//ParseTLEMember Parse a TLEMember struct lines into usable structures
func ParseTLEMember(member *modules.TLEMember) *modules.TLEMemberFormat {

	var formatted modules.TLEMemberFormat
	var line1 modules.TLEFormatLine1
	var line2 modules.TLEFormatLine2
	var err error
	formatted.Name = member.Name

	line1, err = parseTLELineOne(member.Line1)
	if err != nil {
		log.Errorln("ParseTLEMEmber Line 1", err)

	}

	line2, err = parseTLELineTwo(member.Line2)
	if err != nil {
		log.Println("ParseTLEMEmber Line 2", err)

	}

	formatted.Line1 = line1
	formatted.Line2 = line2

	return &formatted
}
