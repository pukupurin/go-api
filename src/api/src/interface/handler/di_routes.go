package handler

import (
	"api/conf"
	infra "api/infra/mysql"
	"api/usecase"

	"github.com/labstack/echo/v4"
)

func UserDIRouting(e *echo.Echo) {
	db := conf.NewDB()

	userRepository := infra.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := NewUserHandler(userUsecase)

	e.POST("/users", userHandler.CreateUser())
}
