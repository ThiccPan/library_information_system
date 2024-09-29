package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thiccpan/library_information_system/internal/config"
	"github.com/thiccpan/library_information_system/internal/entity"
)

func CheckAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var user = c.Get("user").(*config.JwtCustomClaims)

			if user.RoleId != entity.ADMIN.Id {
				return c.JSON(http.StatusUnauthorized, map[string]any{
					"message": "unauthorized",
					"error":   "not an admin",
				})
			}

			return next(c)
		}
	}
}
