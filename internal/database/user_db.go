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

func (u *UserDB) CreateUser(user *entity.User) (*entity.User, error) {
	_, err := u.db.Exec("INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)", user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserDB) GetUserByEmail(email string) (*entity.UserResponse, error) {
	var user entity.UserResponse

	row := u.db.QueryRow("SELECT id, name, email FROM users WHERE email = ?", email)
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserDB) GetUserByLogin(email string) (*entity.User, error) {
	var user entity.User

	row := u.db.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
