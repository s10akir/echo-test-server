package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/twinkling-gecko/echo-test-server/pkg/handler"
)

func main() {
	// Echo instance
	app := echo.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// Routes
	app.GET("/", handler.Hello)

	// Start server
	app.Logger.Fatal(app.Start(":3000"))
}
