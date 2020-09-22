package models

//Token Basic Structure for token.
type Token struct {
	TokenID   string `json:"tokenid"`
	IsAdmin   bool   `json:"isadmin"` // This identify the token as admin and will be to create other users within the system with privileges.
	TimeStamp int64  `json:"timestamp"`
}
