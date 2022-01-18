package controller

import (
	"net/http"
	"strconv"

	"github.com/eriawan06/go-test-sagara/src/modules/product/model/dto"
	"github.com/eriawan06/go-test-sagara/src/modules/product/service"
	"github.com/eriawan06/go-test-sagara/src/shared"
	"github.com/labstack/echo/v4"
)

type ProductControllerImpl struct {
	Service service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		Service: productService,
	}
}

func (controller *ProductControllerImpl) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		productCreateDto := dto.ProductCreateDto{}
		// helper.ReadFromRequestBody(request, &productCreateDto)
		err := c.Bind(&productCreateDto)
		if err != nil {
			return c.JSON(http.StatusBadRequest, shared.ErrorResponse{
				Code:    400,
				Status:  "Bad Request",
				Message: err.Error(),
			})
		}

		productResponse, err := controller.Service.Create(c.Request().Context(), productCreateDto)
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
			Data:   productResponse,
		}
		return c.JSON(http.StatusOK, webResponse)
	}
}

func (controller *ProductControllerImpl) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		productUpdateDto := dto.ProductUpdateDto{}
		// helper.ReadFromRequestBody(request, &productCreateDto)
		err := c.Bind(&productUpdateDto)
		if err != nil {
			return c.JSON(http.StatusBadRequest, shared.ErrorResponse{
				Code:    400,
				Status:  "Bad Request",
				Message: err.Error(),
			})
		}

		productId := c.Param("productId")
		id, err := strconv.Atoi(productId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, shared.ErrorResponse{
				Code:    400,
				Status:  "Bad Request",
				Message: err.Error(),
			})
		}

		productUpdateDto.Id = id
		productResponse, err := controller.Service.Update(c.Request().Context(), productUpdateDto)
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
			Data:   productResponse,
		}
		return c.JSON(http.StatusOK, webResponse)
	}
}

func (controller *ProductControllerImpl) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		productId := c.Param("productId")
		id, err := strconv.Atoi(productId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, shared.ErrorResponse{
				Code:    400,
				Status:  "Bad Request",
				Message: err.Error(),
			})
		}

		err = controller.Service.Delete(c.Request().Context(), id)
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
		}
		return c.JSON(http.StatusOK, webResponse)
	}
}

func (controller *ProductControllerImpl) FindById() echo.HandlerFunc {
	return func(c echo.Context) error {
		productId := c.Param("productId")
		id, err := strconv.Atoi(productId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, shared.ErrorResponse{
				Code:    400,
				Status:  "Bad Request",
				Message: err.Error(),
			})
		}

		productResponse, err := controller.Service.FindById(c.Request().Context(), id)
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
			Data:   productResponse,
		}
		return c.JSON(http.StatusOK, webResponse)
	}
}

func (controller *ProductControllerImpl) FindAll() echo.HandlerFunc {
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
