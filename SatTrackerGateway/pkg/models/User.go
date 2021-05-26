package models

type User struct {
	Username     string
	Password     string
	OutputFormat string
}

type Identity struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}
