package controller

import "github.com/labstack/echo/v4"

type AuthController interface {
	Login() echo.HandlerFunc
	Register() echo.HandlerFunc
}
