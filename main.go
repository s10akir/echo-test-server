package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/twinkling-gecko/echo-test-server/pkg/router"
)

func main() {
	// Echo instance
	app := echo.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// Routes
	router.Init(app)

	// Start server
	app.Logger.Fatal(app.Start(":3000"))
}
