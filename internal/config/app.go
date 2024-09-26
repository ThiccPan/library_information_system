package config

import (
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/library_information_system/internal/delivery/http"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB  *gorm.DB
	App *echo.Echo
}

func SetupApp(config BootstrapConfig) {
	routerConfig := http.EchoRouteConfig{
		App: config.App,
	}
	routerConfig.SetupRoute()
}
