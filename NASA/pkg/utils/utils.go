package utils

import (
	"net/url"
)

//GetVarURL will return the value of a URL Parameter
func GetVarURL(uri string, parameter string) string {

	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	return (m[parameter][0])
}
