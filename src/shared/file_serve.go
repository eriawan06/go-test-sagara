package shared

import "github.com/labstack/echo/v4"

func FileServe() echo.HandlerFunc {
	return func(c echo.Context) error {
		filename := c.Param("filename")

		return c.File("static/image/" + filename)
	}
}
