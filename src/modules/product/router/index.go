package router

import (
	"github.com/eriawan06/go-test-sagara/src/modules/product"
	"github.com/eriawan06/go-test-sagara/src/shared"
	"github.com/labstack/echo/v4"
)

func Init(api *echo.Group) {
	api.GET("/products", product.GetController().FindAll())
	api.GET("/products/:productId", product.GetController().FindById())
	api.POST("/products", product.GetController().Create())
	api.PUT("/products/:productId", product.GetController().Update())
	api.DELETE("/products/:productId", product.GetController().Delete())
	api.POST("/products/image/upload", shared.FileUpload())
	api.GET("/products/image/:filename", shared.FileServe())
}
