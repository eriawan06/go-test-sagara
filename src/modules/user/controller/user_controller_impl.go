package controller

import (
	"net/http"
	"strconv"

	"github.com/eriawan06/go-test-sagara/src/modules/user/service"
	"github.com/eriawan06/go-test-sagara/src/shared"
	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	Service service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		Service: userService,
	}
}

func (controller *UserControllerImpl) FindById() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("userId")
		id, err := strconv.Atoi(userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, shared.ErrorResponse{
				Code:    400,
				Status:  "Bad Request",
				Message: err.Error(),
			})
		}

		userResponse, err := controller.Service.FindById(c.Request().Context(), id)
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

func (controller *UserControllerImpl) FindAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		productResponses, err := controller.Service.FindAll(c.Request().Context())
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
			Data:   productResponses,
		}
		return c.JSON(http.StatusOK, &webResponse)
	}
}
