package handler

import (
	"net/http"

	"api/usecase"

	"github.com/labstack/echo/v4"
)

// UserHander User handerのinterface
type UserHandler interface {
	CreateUser() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

// NewUserHandler User handler のコンストラクタ
func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

// req & res json -------------
type ReqCreateUser struct {
	Name string `json:"name"`
}

type DefaultResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// ----------------------------

func (uh *userHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var res DefaultResponse

		postdata := new(ReqCreateUser)
		err := c.Bind(postdata)
		if err != nil {
			res.Status = 400
			res.Message = "parameter error"
			return c.JSON(http.StatusBadRequest, res)
		}

		err = uh.userUsecase.CreateUser(postdata.Name)
		if err != nil {
			res.Status = 400
			res.Message = "create user error"
			return c.JSON(http.StatusBadRequest, res)
		}

		res.Status = 200
		res.Message = "success"

		return c.JSON(http.StatusOK, res)
	}
}
