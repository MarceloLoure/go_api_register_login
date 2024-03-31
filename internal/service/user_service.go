package service

import (
	"time"

	"github.com/MarceloLoure/go_api_register_login/internal/database"
	"github.com/MarceloLoure/go_api_register_login/internal/entity"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserDB *database.UserDB
}

func NewUserService(userDB *database.UserDB) *UserService {
	return &UserService{UserDB: userDB}
}

func (userService *UserService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (u *UserService) CreateUser(name, email, password string) (*entity.User, error) {
	user := entity.NewUser(name, email, password)

	createdUser, err := u.UserDB.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (u *UserService) GetUserByEmail(email string) (*entity.UserResponse, error) {
	userResponse, err := u.UserDB.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return userResponse, nil
}

func createToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *UserService) verifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil

}

func (u *UserService) LoginUser(email, password string) (string, error) {
	user, err := u.UserDB.GetUserByLogin(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := createToken(email)
	if err != nil {
		return "", err
	}

	return token, nil
}
