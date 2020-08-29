package modules

//TLEMemberFormat main format
type TLEMemberFormat struct {
	Name  string         `json:"name"`
	Line1 TLEFormatLine1 `json:"line1"`
	Line2 TLEFormatLine2 `json:"line2"`
}

//TLEFormatLine1 Struct for decoding line 1 of TLE Format
type TLEFormatLine1 struct {
	LineNumber int `json:"linenumber"`
	//Satellite Catalog Number
	SatelliteNumber int `json:"satellitenumber"`
	//Elset Classification
	ElsetClassification string `json:"elsetclassification"`
	//International Designator
	InternationalDesingator string `json:"internationaldesignator"`
	//Element Set Epoch (UTC) *Note: spaces are acceptable in columns 21 & 22
	ElementSetEpoch string `json:"elementsetepoch"`
	//1st Derivative of the Mean Motion with respect to Time
	TimeMotionDerivativeOne string `json:"timemotionderivativeone"`
	//2nd Derivative of the Mean Motion with respect to Time (decimal point assumed)
	TimeMotionDerivativeTwo string `json:"timemotionderivativetwo"`
	//B* Drag Term
	BDragTerm string `json:"bdragterm"`
	//Element Set Type
	ElementSetType int `json:"elementsetype"`
	//Element Number
	ElementNumber int `json:"elementnumber"`
	//Checksum
	Checksum int `json:"checksum"`
}

//TLEFormatLine2 Struct for decoding line 2 of TLE Format
type TLEFormatLine2 struct {
	LineNumber         int    `json:"linenumber"`
	SatelliteNumber    int    `json:"satellitenumber"`
	ElseClassification string `json:"elseclassification"`
	// Orbit Inclination (degrees)
	OrbitInclination float64 `json:"orbitinclination"`
	// Right Ascension of Ascending Node (degrees)
	RAAN float64 `json:"raan"`
	//Eccentricity (decimal point assumed)
	Eccentricity float64 `json:"eccentricity"`
	//Argument of Perigee (degrees)
	ArgumentPerigee float64 `json:"argumentperigee"`
	//Mean Anomaly (degrees)
	MeanAnomaly float64 `json:"meananomaly"`
	//Mean Motion (revolutions/day)
	MeanMotion float64 `json:"meanmotion"`
	//Revolution Number at Epoch
	RevolutionNumber float64 `json:"revolutionnumber"`
	//Checksum
	Checksum int `json:"checksum"`
}
