package infra

import (
	"api/domain/model"
	"api/domain/repository"

	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &UserRepository{Conn: conn}
}

// Create Userの新規作成
func (ur *UserRepository) Create(user *model.User) (string, error) {
	if err := ur.Conn.Create(user).Error; err != nil {
		return "", err
	}

	return user.ID, nil
}
