package service

import (
	"internal/database"
	"internal/entity"
)

type UserService struct {
	UserDB *database.UserDB
}

func NewUserService(userDB *database.UserDB) *UserService {
	return &UserService{UserDB: userDB}
}

func (u *UserService) CreateUser(name, email, password string) (*entity.User, error) {
	user, err := entity.NewUser(name, email, password)
	if err != nil {
		return nil, err
	}

	user, err = u.UserDB.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}