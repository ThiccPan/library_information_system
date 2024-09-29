package http

import (
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/library_information_system/internal/delivery/http/controller"
	"github.com/thiccpan/library_information_system/internal/delivery/http/middleware"
)

type AppConfig struct {
	App            *echo.Echo
	UserController *controller.UserController
}

func (e *AppConfig) SetupRoute() {
	route := e.App
	route.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(200, map[string]any{"message": "app is online"})
	})
	route.POST("/register", e.UserController.RegisterController)
	route.POST("/login", e.UserController.LoginController)

	rRoute := route.Group("/users", middleware.JWTUser(), middleware.CheckAdmin())
	rRoute.GET("", e.UserController.GetAllController)
	rRoute.GET("/:id", e.UserController.GetByIdController)

	myRoute := route.Group("/my", middleware.JWTUser())
	myRoute.GET("", e.UserController.GetProfileController)
	myRoute.POST("", e.UserController.UpdateController)
	myRoute.POST("/profile", e.UserController.UpdateProfileController)
	route.Static("/pictures/profiles", "resource/user_profile")

}
