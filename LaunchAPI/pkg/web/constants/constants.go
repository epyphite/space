package constants

import (
	"github.com/gorilla/sessions"
)

var (
	Store        = sessions.NewCookieStore([]byte("7b24afc8bc80e548d66c4e7ff72171c5"))
	useSSL       = true //Use SSL?
	userAgentKey = ""   // userAgentKey
	err          error
	isPanel      = true
	isNew        = true
	isEnabled    = true
	JWTKey       = []byte("7b24afc8bc80e548d66c4e7ff72171c5#")
	APP_KEY      = []byte("7b24afc8bc80e548d66c4e7ff72171c5")
)
