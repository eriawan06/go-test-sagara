package main

import (
	"net/http"

	"github.com/eriawan06/go-test-sagara/src/helper"
	"github.com/eriawan06/go-test-sagara/src/routes"
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

	// start to listen
	server := &http.Server{
		Addr: "localhost:3000",
	}
	err := e.StartServer(server)
	helper.PanicIfError(err)
}
