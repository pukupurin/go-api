package handler

import (
	"api/conf"
	"api/infra/mysql"
	"api/usecase"

	"github.com/labstack/echo/v4"
)

func User_DI_Routing(e *echo.Echo) {
	db := conf.NewDB()

	userRepository := infra.NewUserRepository(db)
	userUsecase    := usecase.NewUserUsecase(userRepository)
	userHandler    := NewUserHandler(userUsecase)

	e.POST("/users", userHandler.CreateUser())
}