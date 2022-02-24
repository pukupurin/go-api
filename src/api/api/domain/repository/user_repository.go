package repository

import (
	"api/domain/model"
)

type UserRepository interface {
	Create(user *model.User) (int64, error)
}
