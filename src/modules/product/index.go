package product

import (
	"github.com/eriawan06/go-test-sagara/app"
	"github.com/eriawan06/go-test-sagara/src/modules/product/controller"
	"github.com/eriawan06/go-test-sagara/src/modules/product/repository"
	"github.com/eriawan06/go-test-sagara/src/modules/product/service"
	"github.com/go-playground/validator/v10"
)

var (
	productRepository = repository.NewProductRepository(app.NewDB())
	productService    = service.NewProductService(productRepository, validator.New())
	productController = controller.NewProductController(productService)
)

func GetService() service.ProductService {
	return productService
}

func GetController() controller.ProductController {
	return productController
}
