package repository

import (
	"api/domain/model"
)

type UserRepository interface {
	Create(user *model.User) (string, error)
}
