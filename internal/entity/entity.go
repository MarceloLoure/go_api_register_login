package entity

import (
	"github.com/google/uuid"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func generateCustomUUID() string {
	uuidWithHyphens := uuid.New().String()

	customUUID := uuidWithHyphens[:32]

	return customUUID
}

func NewUser(name, email, password string) *User {
	return &User{
		ID:       generateCustomUUID(),
		Name:     name,
		Email:    email,
		Password: password,
	}
}
