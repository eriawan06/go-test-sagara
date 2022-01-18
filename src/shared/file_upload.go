package shared

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type FileUploadResponse struct {
	Url string `json:"url"`
}

func FileUpload() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Read form fields
		filename := c.FormValue("filename")

		//-----------
		// Read file
		//-----------

		// Source
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dir := "static/image/"
		file.Filename = filename + ".jpg"
		dst, err := os.Create(dir + file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		fileUrl := "/api/products/image/" + file.Filename

		return c.JSON(http.StatusOK, SuccessResponse{
			Code:   200,
			Status: "OK",
			Data: FileUploadResponse{
				Url: fileUrl,
			},
		})
	}
}
