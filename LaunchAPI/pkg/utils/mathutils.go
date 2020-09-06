package utils

import "math"

/*

	function Rad2Deg(radians)
	{
		return radians * (180./Math.PI);
	}
*/

//Rad2Deg transform Radians to Degrees
func Rad2Deg(radians float64) float64 {
	return radians * (180. / math.Pi)
}

//Deg2Rad transform Degrees to Radians
func Deg2Rad(degrees float64) float64 {
	return degrees * (math.Pi / 180.)
}
