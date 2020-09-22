package models

//User stores information about the user currently used, check sed in storage.go for seed information and usage.
type User struct {
	UserID   string `json:"agent" db:"userid"`
	Username string `json:"username" db:"username"`
	Password []byte `json:"password" db:"password"`
	Email    string `json:"email"`
	Token    string `json:"token" db:"token"`
	Banned   bool   `json:"banned"`
	Role     string `json:"role"`
	Approved bool   `json:"approved"`
	IsAdmin  bool   `json:"isadmin"`
}

//JSONLogin for support login only
type JSONLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//JSONResponseUsers A collection for response with the users
type JSONResponseUsers struct {
	ResponseCode string
	Message      string
	ResponseData []User
}

//JSONCreateUser for creating rest post user creation
type JSONCreateUser struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CellPhone string `json:"cellphone"`
}

//SubscriberUser for the basic functionality implemeted
type SubscriberUser struct {
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	FileNames []string `json:"filename"`
	TimeStamp string   `json:"date"`
}
