package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/MarceloLoure/go_api_register_login/internal/database"
	"github.com/MarceloLoure/go_api_register_login/internal/service"
	"github.com/MarceloLoure/go_api_register_login/internal/webserver"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

// GetAPI is a function that returns a JSON response
func GetAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Hello, World!"}`))
}

func GetAPI2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Olá usuário!"}`))
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/baseUsers")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userDB := database.NewUserDB(db)
	userService := service.NewUserService(userDB)
	webUserHandler := webserver.NewWebUserHandler(userService)

	fmt.Println("Conexão com o banco de dados MySQL estabelecida com sucesso!")

	r := chi.NewRouter()
	r.Get("/", GetAPI)
	r.Post("/users/register", webUserHandler.CreateUser)
	r.Post("/users/login", webUserHandler.Login)
	r.Get("/users/{email}", webUserHandler.GetUserByEmail)

	fmt.Println("Servidor rodando na porta 3000")
	http.ListenAndServe(":3000", r)
}
