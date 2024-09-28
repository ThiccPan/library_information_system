package config

import (
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/library_information_system/internal/delivery/http"
	"github.com/thiccpan/library_information_system/internal/delivery/http/controller"
	"github.com/thiccpan/library_information_system/internal/repository"
	"github.com/thiccpan/library_information_system/internal/usecase"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB  *gorm.DB
	App *echo.Echo
}

func SetupApp(config BootstrapConfig) {
	userRepository := repository.NewUserRepoImpl(config.DB)
	userUsecase := usecase.NewUserUsecase(config.DB, userRepository)
	userController := controller.NewUserController(userUsecase, NewValidator())

	routerConfig := http.AppConfig{
		App: config.App,
		UserController: userController,
	}

	routerConfig.SetupRoute()
}
