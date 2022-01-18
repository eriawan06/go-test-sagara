package helper

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/eriawan06/go-test-sagara/src/shared"
	"github.com/labstack/echo/v4"
)

var JWT_SECRET = "1qaz2wsx3edc4rfv"
var JWT_REFRESH_TOKEN_LIFETIME = 24
var JWT_ACCESS_TOKEN_LIFETIME = 300

// GenerateJWT ...
func GenerateJWT(u string) (shared.JwtModel, jwt.MapClaims, error) {

	// Set refresh Token
	refreshToken := jwt.New(jwt.SigningMethodHS512)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * time.Duration(JWT_REFRESH_TOKEN_LIFETIME)).Unix()
	rt, err := refreshToken.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return shared.JwtModel{}, nil, err
	}

	// Set claims
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = u
	claims["name"] = u
	claims["admin"] = true
	claims["roles"] = []string{"ADMIN", "USER"}
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(JWT_ACCESS_TOKEN_LIFETIME)).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return shared.JwtModel{}, nil, err
	}
	// tokens := map[string]string{
	// 	"access_token":  t,
	// 	"refresh_token": rt,
	// }
	tokens := shared.JwtModel{
		AccessToken:  t,
		RefreshToken: rt,
	}
	return tokens, claims, nil
}

func JWTGetClaimsUID(c echo.Context) (string, error) {

	bearer := c.Request().Header.Get("Authorization")
	tokenString := bearer[7:]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		hmacSampleSecret := []byte(JWT_SECRET)
		return hmacSampleSecret, nil
	})

	if err != nil {
		return "error", err
	}
	//spew.Dump("----- token parse -----", token.Claims.(jwt.MapClaims)["uid"])
	return token.Claims.(jwt.MapClaims)["uid"].(string), nil
}
