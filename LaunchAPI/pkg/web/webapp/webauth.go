package webapp

import (
	"epyphite/space/v1/LaunchAPI/pkg/web/constants"
	"log"
	"net/http"
	"time"

	models "epyphite/space/v1/LaunchAPI/pkg/models"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"

	"github.com/dgrijalva/jwt-go"
)

func (a *MainWebAPI) TokenHandler(user models.User) (JResponseToken, error) {

	var response JResponseToken

	// We are happy with the credentials, so build a token. We've given it
	// an expiry of 1 hour.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.Username,
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(constants.APP_KEY))

	if err != nil {

		return response, err
	}
	response.ResponseCode = "200"
	response.Token = tokenString
	return response, err

}

// AuthMiddleware is our middleware to check our token is valid. Returning
// a 401 status to the client if it is not valid.
func (a *MainWebAPI) AuthMiddleware(next http.Handler) http.Handler {
	if len(constants.APP_KEY) == 0 {
		log.Fatal("HTTP server unable to start, expected an APP_KEY for JWT auth")
	}
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return constants.APP_KEY, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	return jwtMiddleware.Handler(next)
}
