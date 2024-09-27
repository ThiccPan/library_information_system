package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/thiccpan/library_information_system/internal/config"
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
	config.SetupApp(config.BootstrapConfig{
		App: app,
		DB:  db,
	})
	app.Logger.Fatal(app.Start(":8080"))
}
