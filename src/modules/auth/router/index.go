package router

import (
	"github.com/eriawan06/go-test-sagara/src/modules/auth"
	"github.com/labstack/echo/v4"
)

func Init(api *echo.Group) {
	api.POST("/login", auth.GetController().Login())
	api.POST("/register", auth.GetController().Register())
}
