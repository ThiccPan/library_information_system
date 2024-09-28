package http

import (
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/library_information_system/internal/delivery/http/controller"
)

type AppConfig struct {
	App            *echo.Echo
	UserController *controller.UserController
}

func (e *AppConfig) SetupRoute() {
	e.App.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(200, map[string]any{"message": "app is online"})
	})
	e.App.POST("/register", e.UserController.RegisterController)
}
