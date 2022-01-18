package main

import (
	"fmt"
	"net/http"

	"github.com/eriawan06/go-test-sagara/src/helper"
	"github.com/eriawan06/go-test-sagara/src/routes"
	"github.com/eriawan06/go-test-sagara/src/shared"
	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	// recover the server if fatal error occur
	e.Use(middleware.Recover())

	//CORS Config
	CORSConfig := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}

	api := e.Group("/api")
	auth := e.Group("/auth")

	api.Use(middleware.CORSWithConfig(CORSConfig))
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(helper.JWT_SECRET),
		SigningMethod: "HS512",
	}))

	routes.InitRouter(api)
	routes.InitAuthRouter(auth)

	api.POST("/refresh-token", func(c echo.Context) (err error) {
		bearer := c.Request().Header.Get("Authorization")

		tokenString := bearer[7:]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			hmacSampleSecret := []byte(helper.JWT_SECRET)
			return hmacSampleSecret, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			//fmt.Println(claims["name"], claims["uid"])
			tokens, _, _ := helper.GenerateJWT(claims["name"].(string))

			return c.JSON(http.StatusOK, struct {
				Code   int             `json:"code"`
				Status string          `json:"status"`
				Data   shared.JwtModel `json:"data"`
			}{
				Code:   200,
				Status: "OK",
				Data:   tokens,
			})
		} else {
			fmt.Println(err)
			return c.String(http.StatusUnauthorized, "unauthorized")
		}

		//parse, _ := jwt.DecodeSegment(c.Request().Header.Get("Authorization")[7:len(bearer)])
		//return c.JSON(http.StatusOK, "")
	})

	// start to listen
	server := &http.Server{
		Addr: "localhost:3000",
	}
	err := e.StartServer(server)
	helper.PanicIfError(err)
}
