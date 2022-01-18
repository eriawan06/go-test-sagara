package routes

import (
	authRouter "github.com/eriawan06/go-test-sagara/src/modules/auth/router"
	productRouter "github.com/eriawan06/go-test-sagara/src/modules/product/router"
	userRouter "github.com/eriawan06/go-test-sagara/src/modules/user/router"
	"github.com/labstack/echo/v4"
)

func InitRouter(api *echo.Group) {
	productRouter.Init(api)
	userRouter.Init(api)
}

func InitAuthRouter(api *echo.Group) {
	authRouter.Init(api)
}
