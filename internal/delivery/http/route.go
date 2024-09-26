package http

import "github.com/labstack/echo/v4"

type EchoRouteConfig struct {
	App *echo.Echo
}

func (e *EchoRouteConfig) SetupRoute() {
	e.App.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(200, map[string]any{ "message": "app is online" })
	})
}