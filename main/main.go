package main

import (
	"github.com/fahrurben/gopress/internal/user"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	db, err := sqlx.Connect("mysql", "root:test12345@tcp(localhost:3306)/godoc")
	if err != nil {
		panic(err)
	}

	if err != nil {
		log.Fatalf("failed to set up the validator: %v", err)
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := user.CreateHandler(userService)

	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello world"))
	})
	r.Post("/api/user", userHandler.CreateUserHandler)

	http.ListenAndServe(":3000", r)
}
