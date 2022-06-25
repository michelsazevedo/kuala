package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/michelsazevedo/kuala/api"
	"github.com/michelsazevedo/kuala/config"
	"github.com/michelsazevedo/kuala/domain"
	"github.com/michelsazevedo/kuala/repository"
)

func main() {
	conf, _ := config.NewConfig("./config/config.yaml")
	repo, _ := repository.NewPostgresRepository(conf.Database.Host,
		conf.Database.Username, conf.Database.Password, conf.Database.Database)
	service := domain.NewJobService(repo)
	handler := api.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/jobs/{id}", handler.Get)
	r.Get("/jobs", handler.GetAll)
	r.Post("/jobs", handler.Post)
	r.Delete("/jobs/{id}", handler.Delete)
	r.Put("/jobs", handler.Put)
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		healthz := map[string]int{"status": 200}
		json.NewEncoder(w).Encode(healthz)
	})

	log.Fatal(http.ListenAndServe(conf.Server.Port, r))
}
