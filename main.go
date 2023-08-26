package main

import (
	"database/sql"
	"net/http"

	"github.com/FRSiqueiraBR/golang-url-shortener/internal/handlers"
	"github.com/FRSiqueiraBR/golang-url-shortener/internal/core/services"
	"github.com/FRSiqueiraBR/golang-url-shortener/internal/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "infra/database/UrlShortener.db")
	if err != nil {
		panic(err)
	}

	defer db.Close() //espera tudo rodar depois executa o close

	repo := repositories.NewUrlShortRepository(db)
	hash := services.NewHashService()
	service := services.NewUrlShortService(repo, hash)
	controller := handlers.NewUrlShortHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/url-short", controller.FindAll)
	r.Get("/url-short/{hash}", controller.FindByHash)
	r.Post("/url-short", controller.Create)

	http.ListenAndServe(":8080", r)
}
