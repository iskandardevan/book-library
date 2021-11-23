package middlewares

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

func Log(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
	}))
}