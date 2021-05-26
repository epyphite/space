package webapp

import (
	"epyphite/space/v1/LaunchAPI/pkg/web/constants"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	models "epyphite/space/v1/LaunchAPI/pkg/models"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	uuid "github.com/satori/go.uuid"

	"github.com/dgrijalva/jwt-go"
)

//AccessDetails mapping strucure for Redis Usage
type AccessDetails struct {
	AccessUUID string
	UserID     string
}

//CreateToken will create a new access token based on the User.
func (a *MainWebAPI) CreateToken(user models.User) (*TokenDetails, error) {
	var err error
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUUID = uuid.NewV4().String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.NewV4().String()
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = user.UserID
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = user.UserID
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil

}

//CreateAuth Saves the token details into redis.
func (a *MainWebAPI) CreateAuth(user models.User, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := a.RedisClient.Set(td.AccessUUID, user.UserID, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := a.RedisClient.Set(td.RefreshUUID, user.UserID, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
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

//TokenAuthMiddleware function to intercept unauth
func (a *MainWebAPI) TokenAuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := a.TokenValid(r)
		if err != nil {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

//ExtractToken will read the token from the authorization
func (a *MainWebAPI) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

//VerifyToken will matcht the signature
func (a *MainWebAPI) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := a.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//TokenValid validity of the token
func (a *MainWebAPI) TokenValid(r *http.Request) error {
	token, err := a.VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

//ExtractTokenMetadata for the information
func (a *MainWebAPI) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := a.VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userID := fmt.Sprintf("%.f", claims["user_id"])

		return &AccessDetails{
			AccessUUID: accessUUID,
			UserID:     userID,
		}, nil
	}
	return nil, err
}

//FetchAuth will fetch information from Redis
func (a *MainWebAPI) FetchAuth(authD *AccessDetails) (string, error) {
	userid, err := a.RedisClient.Get(authD.AccessUUID).Result()
	if err != nil {
		return "", err
	}
	return userid, nil
}

//DeleteAuth will delete the token
func (a *MainWebAPI) DeleteAuth(givenUUID string) (int64, error) {
	deleted, err := a.RedisClient.Del(givenUUID).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
