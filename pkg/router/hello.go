package router

import (
	"github.com/labstack/echo/v4"

	"github.com/twinkling-gecko/echo-test-server/pkg/handler"
)

func initHello(app *echo.Echo) {
	app.GET("/", handler.Hello)
}
