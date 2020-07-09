package router

import (
	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo) {
	initHello(app)
}
