package service

import (
	"github.com/MarceloLoure/go_api_register_login/internal/database"
	"github.com/MarceloLoure/go_api_register_login/internal/entity"
)

type UserService struct {
	UserDB *database.UserDB
}

func NewUserService(userDB *database.UserDB) *UserService {
	return &UserService{UserDB: userDB}
}

func (u *UserService) CreateUser(name, email, password string) (*entity.User, error) {
	user := entity.NewUser(name, email, password)

	createdUser, err := u.UserDB.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (u *UserService) GetUserByEmail(email string) (*entity.User, error) {
	user, err := u.UserDB.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
