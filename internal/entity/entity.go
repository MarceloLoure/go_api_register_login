package entity

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Password string `json:"password"`
}

func NewUser(name, email, password string) (*User, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &User{
		ID:    id,
		Name:  name,
		Email: email,
		Password: password,
	}, nil
}