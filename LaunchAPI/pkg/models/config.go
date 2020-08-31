package models

//Config Main Configuration File Structure
type Config struct {
	WebPort      string `json:"webport"`
	WebAddress   string `json:"webaddress"`
	BaseURL      string `json:"baseurl"`
	APIKey       string `json:"apikey"`
	SaveOnError  bool   `json:"saveonerror"`
	MaxPages     int    `json:"maxpages"`
	DatabaseName string `json:"databasename"`
}
