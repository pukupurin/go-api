package infra

import (

	"api/domain/repository"
	"api/domain/model"
	
	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &UserRepository{Conn: conn}
}

// Create Userの新規作成
func (ur *UserRepository) Create(user *model.User) (int64, error) {
	if err := ur.Conn.Create(user).Error; err != nil {
		return 0, err
	}

	return user.Id, nil
}