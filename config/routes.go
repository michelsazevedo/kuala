package config

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/michelsazevedo/kuala/api"
)

func Routes(handler api.JobHandler) chi.Router {
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

	return r
}
