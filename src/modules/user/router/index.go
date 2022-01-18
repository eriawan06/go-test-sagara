package router

import (
	"github.com/eriawan06/go-test-sagara/src/modules/user"
	"github.com/labstack/echo/v4"
)

func Init(api *echo.Group) {
	api.GET("/users", user.GetController().FindAll())
	api.GET("/users/:userId", user.GetController().FindById())
	// api.PUT("/products/:productId", product.GetController().Update())
	// api.DELETE("/products/:productId", product.GetController().Delete())
}
