package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func Error(c echo.Context, status int, err error) error {
	return c.JSON(status, ErrorResponse{Message: err.Error()})
}

func NotFound(c echo.Context, message string) error {
	return c.JSON(http.StatusNotFound, ErrorResponse{Message: message})
}
