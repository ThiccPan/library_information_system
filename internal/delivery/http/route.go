package http

import (
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/library_information_system/internal/delivery/http/controller"
	"github.com/thiccpan/library_information_system/internal/delivery/http/middleware"
)

type AppConfig struct {
	App              *echo.Echo
	UserController   *controller.UserController
	AuthorController *controller.AuthorController
	TopicController  *controller.TopicController
	BookController  *controller.BookController
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

	authorRoute := route.Group("/authors")
	authorRoute.POST("", e.AuthorController.Create)
	authorRoute.GET("", e.AuthorController.Get)
	authorRoute.GET("/:id", e.AuthorController.GetById)
	authorRoute.POST("/:id", e.AuthorController.Update)
	authorRoute.DELETE("/:id", e.AuthorController.Delete)

	topicRoute := route.Group("/topics")
	topicRoute.POST("", e.TopicController.Create, middleware.JWTUser(), middleware.CheckAdmin())
	topicRoute.GET("", e.TopicController.Get)
	topicRoute.GET("/:id", e.TopicController.GetById)
	topicRoute.POST("/:id", e.TopicController.Update, middleware.JWTUser(), middleware.CheckAdmin())
	topicRoute.DELETE("/:id", e.TopicController.Delete, middleware.JWTUser(), middleware.CheckAdmin())

	bookRoute := route.Group("/books")
	bookRoute.POST("", e.BookController.Create, middleware.JWTUser(), middleware.CheckAdmin())
	bookRoute.GET("", e.BookController.Get)
	bookRoute.GET("/:id", e.BookController.GetById)
	bookRoute.POST("/:id", e.BookController.Update, middleware.JWTUser(), middleware.CheckAdmin())
	bookRoute.DELETE("/:id", e.BookController.Delete, middleware.JWTUser(), middleware.CheckAdmin())
}
