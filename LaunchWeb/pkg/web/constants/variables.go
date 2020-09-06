package constants

import (
	"github.com/gorilla/sessions"
)

var ()

var (
	Store        = sessions.NewCookieStore([]byte("7b24afc8bc80e548d66c4e7ff72171c5"))
	useSSL       = true //Use SSL?
	userAgentKey = ""   // userAgentKey
	err          error
	isPanel      = true
	isNew        = true
	isEnabled    = true
	maxBotList   = 2
	JWTKey       = []byte("7b24afc8bc80e548d66c4e7ff72171c5#")
)
