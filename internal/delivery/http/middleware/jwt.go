package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/library_information_system/internal/config"
)

func JWTUser() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			SigningKey: []byte(os.Getenv("JWT_SECRET")),
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(config.JwtCustomClaims)
			},
			ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
				return jwt.ParseWithClaims(auth, new(config.JwtCustomClaims), func(token *jwt.Token) (interface{}, error) {
					if token.Claims.(*config.JwtCustomClaims).ExpiresAt.Time.Before(time.Now()) {
						return nil, fmt.Errorf("token expired")
					}
					return []byte(os.Getenv("JWT_SECRET_USER")), nil
				})
			},
			SuccessHandler: func(c echo.Context) {
				data := c.Get("user").(*jwt.Token).Claims.(*config.JwtCustomClaims)
				c.Set("user", data)
			},
			ErrorHandler: func(c echo.Context, err error) error {
				fmt.Println(err.Error())
				if err.Error() == "token expired" {
					return c.JSON(http.StatusUnauthorized, map[string]any{
						"message": "unauthorized",
						"error":   err.Error(),
					})
				}
				return c.JSON(http.StatusUnauthorized, map[string]any{
					"message": "unauthorized",
					"error":   err.Error(),
				})
			},
		})
}
