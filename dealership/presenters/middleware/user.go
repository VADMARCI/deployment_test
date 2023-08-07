package middleware

import (
	"dealership/presenters"
	"fmt"

	"github.com/labstack/echo/v4"
	core_models "github.com/pepusz/go_redirect/models"
)

// Errors

func User(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*presenters.Context)
		serviceJWTToken := cc.Request().Header.Get("service-jwt-token")
		fmt.Printf("JWT token got from server.js: %s\n", serviceJWTToken)
		user, err := core_models.GetCoreUserFromToken(serviceJWTToken)
		if err != nil {
			fmt.Printf("WARN: service-jwt-token header is missing or not valid\n")
			return next(cc)
		}
		cc.SetCurrentUser(user)
		return next(cc)
	}
}
