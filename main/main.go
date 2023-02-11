package main

import (
	"github.com/fahrurben/gopress/internal/user"
	"github.com/fahrurben/gopress/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/http"
	"os"
)

func main() {
	r := chi.NewRouter()

	db, err := sqlx.Connect("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		panic(err)
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := user.CreateHandler(userService)

	r.Route("/api", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Use(middleware.JwtAuthMiddleware)
			r.Get("/all/{page}/{limit}", userHandler.SelectUserHandler)
			r.Post("/", userHandler.CreateUserHandler)
			r.Get("/{id}", userHandler.GetUserHandler)
			r.Patch("/{id}", userHandler.UpdateUserHandler)
			r.Delete("/{id}", userHandler.DeleteUserHandler)
		})
	})

	http.ListenAndServe(":"+os.Getenv("APP_PORT"), r)
}
