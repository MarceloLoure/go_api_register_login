package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/MarceloLoure/go_api_register_login/internal/service"
	"github.com/MarceloLoure/go_api_register_login/internal/entity"
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
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := w.UserService.CreateUser(user.Name, user.Email, user.Password)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(createdUser)
}
