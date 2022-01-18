package auth

import (
	"github.com/eriawan06/go-test-sagara/src/modules/auth/controller"
	"github.com/eriawan06/go-test-sagara/src/modules/auth/service"
	"github.com/eriawan06/go-test-sagara/src/modules/user"
	"github.com/go-playground/validator/v10"
)

var (
	userService    = user.GetService()
	authService    = service.NewAuthService(userService, validator.New())
	authController = controller.NewAuthController(authService)
)

func GetService() service.AuthService {
	return authService
}

func GetController() controller.AuthController {
	return authController
}
