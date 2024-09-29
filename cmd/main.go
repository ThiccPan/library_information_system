package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/thiccpan/library_information_system/internal/config"
	"github.com/thiccpan/library_information_system/internal/delivery/http"
	"github.com/thiccpan/library_information_system/internal/delivery/http/controller"
	"github.com/thiccpan/library_information_system/internal/repository"
	"github.com/thiccpan/library_information_system/internal/usecase"
)

func main() {
	app := echo.New()
	db := config.SetupDB(config.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	})

	userRepository := repository.NewUserRepoImpl(db)
	userUsecase := usecase.NewUserUsecase(db, userRepository)
	userController := controller.NewUserController(userUsecase, config.NewValidator(), config.NewAuthJWT(os.Getenv("JWT_SECRET")))

	authorRepository := repository.NewAuthorRepoImpl(db)
	authorUsecase := usecase.NewAuthorUsecase(db, authorRepository)
	authorController := controller.NewAuthorController(authorUsecase)

	topicRepository := repository.NewTopicRepoImpl(db)
	topicUsecase := usecase.NewTopicUsecase(db, topicRepository)
	topicController := controller.NewTopicController(topicUsecase)

	routerConfig := http.AppConfig{
		App:            app,
		UserController: userController,
		AuthorController: authorController,
		TopicController: topicController,
	}
	routerConfig.SetupRoute()

	app.Logger.Fatal(app.Start(":8080"))
}
