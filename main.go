package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
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

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexão com o banco de dados MySQL estabelecida com sucesso!")

	r := chi.NewRouter()
	r.Get("/", GetAPI)

	http.ListenAndServe(":3000", r)
}
