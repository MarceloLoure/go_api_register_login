package webserver

import (
	"encoding/json"
	"net/http"

	"internal/service"
	"internal/entity"
	"github.com/go-chi/chi/v5"
)

type WebUserHandler struct {
	UserService *service.UserService
}

func NewWebUserHandler(userService *service.UserService) *WebUserHandler {
	return &WebUserHandler{UserService: userService}
}

func (w *WebUserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err = w.UserService.CreateUser(user.Name, user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}