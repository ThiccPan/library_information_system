package main

import (
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/library_information_system/internal/config"
)

func main() {
	app := echo.New()
	db := config.SetupDB()
	config.SetupApp(config.BootstrapConfig{
		App: app,
		DB:  db,
	})
	app.Logger.Fatal(app.Start(":8080"))
}
