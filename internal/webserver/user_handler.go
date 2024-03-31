package webserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MarceloLoure/go_api_register_login/internal/entity"
	"github.com/MarceloLoure/go_api_register_login/internal/service"
	"github.com/go-chi/chi/v5"
)

type WebUserHandler struct {
	UserService *service.UserService
}

func NewWebUserHandler(userService *service.UserService) *WebUserHandler {
	return &WebUserHandler{UserService: userService}
}

func (w *WebUserHandler) CreateUser(writer http.ResponseWriter, r *http.Request) {
	var user entity.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	existingUser, err := w.UserService.GetUserByEmail(user.Email)
	if existingUser != nil {
		http.Error(writer, "E-mail já cadastrado", http.StatusBadRequest)
		return
	}

	hashedPassword, err := w.UserService.HashPassword(user.Password)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	createdUser, err := w.UserService.CreateUser(user.Name, user.Email, hashedPassword)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(createdUser)
}

func (w *WebUserHandler) GetUserByEmail(writer http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	if email == "" {
		http.Error(writer, "E-mail não especificado na URL", http.StatusBadRequest)
		return
	}

	user, err := w.UserService.GetUserByEmail(email)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(user)
}

func (w *WebUserHandler) Login(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var user entity.UserLogin

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := w.UserService.LoginUser(user.Email, user.Password)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(map[string]string{"token": token})
}
