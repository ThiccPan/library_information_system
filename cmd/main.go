package main

import (
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/library_information_system/internal/delivery/http"
)

func main() {
	server := echo.New()
	server.GET("/healthcheck", func(c echo.Context) error { return c.JSON(200, map[string]any{"message": "ok!"}) })
	routerConfig := http.EchoRouteConfig{
		App: server,
	}
	routerConfig.SetupRoute()
	server.Logger.Fatal(server.Start(":8080"))
}
