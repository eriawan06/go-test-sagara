package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	// Create() echo.HandlerFunc
	// Update() echo.HandlerFunc
	// Delete() echo.HandlerFunc
	FindById() echo.HandlerFunc
	FindAll() echo.HandlerFunc
}
