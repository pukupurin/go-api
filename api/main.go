package main

import (
	"api/interface/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// User関係のDI&ルーティングの初期化
	handler.UserDIRouting(e)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}
