package controller

import (
	"net/http"

	"github.com/eriawan06/go-test-sagara/src/modules/auth/model/dto"
	"github.com/eriawan06/go-test-sagara/src/modules/auth/service"
	ud "github.com/eriawan06/go-test-sagara/src/modules/user/model/dto"
	"github.com/eriawan06/go-test-sagara/src/shared"
	"github.com/labstack/echo/v4"
)

type AuthControllerImpl struct {
	Service service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		Service: authService,
	}
}

func (controller *AuthControllerImpl) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		authLoginDto := dto.AuthLoginDto{}

		err := c.Bind(&authLoginDto)
		if err != nil {
			return c.JSON(http.StatusBadRequest, shared.ErrorResponse{
				Code:    400,
				Status:  "Bad Request",
				Message: err.Error(),
			})
		}
		loginResponse, err := controller.Service.Login(c.Request().Context(), authLoginDto)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, shared.ErrorResponse{
				Code:    500,
				Status:  "Internal Server Error",
				Message: err.Error(),
			})
		}

		webResponse := shared.SuccessResponse{
			Code:   200,
			Status: "OK",
			Data:   loginResponse,
		}
		return c.JSON(http.StatusOK, webResponse)
	}
}

func (controller *AuthControllerImpl) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		userCreateDto := ud.UserCreateDto{}

		err := c.Bind(&userCreateDto)
		if err != nil {
			return c.JSON(http.StatusBadRequest, shared.ErrorResponse{
				Code:    400,
				Status:  "Bad Request",
				Message: err.Error(),
			})
		}
		userResponse, err := controller.Service.Register(c.Request().Context(), userCreateDto)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, shared.ErrorResponse{
				Code:    500,
				Status:  "Internal Server Error",
				Message: err.Error(),
			})
		}

		webResponse := shared.SuccessResponse{
			Code:   200,
			Status: "OK",
			Data:   userResponse,
		}
		return c.JSON(http.StatusOK, webResponse)
	}
}
