package services

import (
	"database/user_db"
	"entity"
)

type UserService struct {
	UserDB database.UserDB
}

func NewUserService(userDB database.UserDB) *UserService {
	return &UserService{UserDB: userDB}
}

func (u *UserService) CreateUser(name, email, password) (*entity.User, error) {
	user := entity.NewUser(name, email, password)
	_, err := u.UserDB.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u, *UserService) GetUserByEmail(email string) (*entity.User, error) {
	user, err := u.UserDB.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
