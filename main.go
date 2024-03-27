package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"internal/database"
	"internal/entity"
	"internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

// GetAPI is a function that returns a JSON response
func GetAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Hello, World!"}`))
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
	r.Post("/users", webUserHandler.CreateUser)

	http.ListenAndServe(":3000", r)
}
