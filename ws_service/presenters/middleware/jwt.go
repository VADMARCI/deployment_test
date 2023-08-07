package middleware

import (
	"github.com/labstack/echo/v4"

	// "github.com/pepusz/go_redirect/models"

	"net/http"
	"ws_service/presenters"
)

// Errors
var (
	ErrJWTMissing = echo.NewHTTPError(http.StatusBadRequest, "missing or malformed jwt")
)

func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*presenters.Context)

		return next(cc)
	}
}
