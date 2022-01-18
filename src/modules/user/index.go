package user

import (
	"github.com/eriawan06/go-test-sagara/app"
	"github.com/eriawan06/go-test-sagara/src/modules/user/controller"
	"github.com/eriawan06/go-test-sagara/src/modules/user/repository"
	"github.com/eriawan06/go-test-sagara/src/modules/user/service"
	"github.com/go-playground/validator/v10"
)

var (
	productRepository = repository.NewUserUserRepository(app.NewDB())
	userService       = service.NewUserService(productRepository, validator.New())
	userController    = controller.NewUserController(userService)
)

func GetService() service.UserService {
	return userService
}

func GetController() controller.UserController {
	return userController
}
