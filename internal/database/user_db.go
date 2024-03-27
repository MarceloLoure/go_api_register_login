package database

import (
	"database/sql"
	"github.com/MarceloLoure/go_api_register_login/internal/entity"
)

type UserDB struct {
	db *sql.DB
}

func NewUserDB(db *sql.DB) *UserDB {
	return &UserDB{db: db}
}

func (u *UserDB) CreateUser(user *entity.User) ( *entity.User, error) {
	_, err := u.db.Exec("INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)", user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}