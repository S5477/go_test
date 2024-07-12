package repository

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
	"test/internal/entity"
	"test/pkg/password"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(dataBase *gorm.DB) UserRepository {
	return &userRepository{
		db: dataBase,
	}
}

func (r *userRepository) CreateUser(user *entity.User) error {
	var err error

	validate := validator.New()
	err = validate.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
		name := user.Name
		var exists bool
		err := r.db.Model(user).
			Select("count(*) > 0").
			Where("name = ?", name).
			Find(&exists).
			Error

		if err != nil {
			return false
		}

		return !exists
	})

	if err != nil {
		return err
	}

	user.Password, err = password.HashPassword(user.Password)

	if err != nil {
		return err
	}

	err = validate.Struct(*user)
	if err != nil {
		return err
	}

	result := r.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
