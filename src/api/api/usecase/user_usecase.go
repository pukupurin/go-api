package usecase

import (

	"api/domain/model"
	"api/domain/repository"

)

// UserUsecase User 関係のusecaseのinterface
type UserUsecase interface {
	CreateUser(name string) (error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}


// NewUserUsecase User usecaseのコンストラクタ
func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo : userRepo,
	}
}

// CreateUser User を追加する
func (uu *userUsecase) CreateUser(name string) (error) {

	user, err := model.NewUser(name)
	if err != nil {
		return err
	}

	_, err = uu.userRepo.Create(user)
	if err != nil {
		return err
	}

	return nil
}