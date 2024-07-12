package user

import (
	"test/internal/entity"
	"test/internal/repository"
)

type UserUsecase interface {
	RegisterUser(user *entity.User) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) RegisterUser(user *entity.User) error {
	return u.userRepository.CreateUser(user)
}
