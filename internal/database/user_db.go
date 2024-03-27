package database

import (
	"database/sql"

	"entity"
)

type UserDB struct {
	db *sql.DB
}

func NewUserDB(db *sql.DB) *UserDB {
	return &UserDB{
		db: db,
	}
}

func (u *UserDB) CreateUser(user *entity.User) error {
	_, err := u.db.Exec("INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)", user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserDB) GetUserByID(id string) (*entity.User, error) {
	user := &entity.User{}
	err := u.db.QueryRow("SELECT id, name, email, password FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserDB) GetUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	err := u.db.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserDB) UpdateUser(user *entity.User) error {
	_, err := u.db.Exec("UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?", user.Name, user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserDB) GetAllUsers() ([]*entity.User, error) {
	rows, err := u.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*entity.User{}
	for rows.Next() {
		user := &entity.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
