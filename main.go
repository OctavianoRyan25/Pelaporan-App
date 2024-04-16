package main

import (
	"github.com/OctavianoRyan25/Pelaporan-App/configs"
	"github.com/OctavianoRyan25/Pelaporan-App/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	configs.CreateDB()
	e := echo.New()

	// Middleware CORS
	e.Use(middleware.CORS())

	// Inisialisasi rute
	routes.Init(e)

	// Mulai server
	e.Logger.Fatal(e.Start(":8080"))
}
